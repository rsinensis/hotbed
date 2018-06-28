package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"hotbed/models"
	"hotbed/modules/env"
	"hotbed/modules/recovery"
	"hotbed/routers"
	"hotbed/tasks"
	"hotbed/tools/id"
	"hotbed/tools/record"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"golang.org/x/crypto/acme/autocert"
	macaron "gopkg.in/macaron.v1"
)

var (
	// BuildVersion from git tag
	BuildVersion string
	// BuildTime from make time
	BuildTime string
	// BuildMode from make mode
	BuildMode string
)

func redirectToHttps(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see @andreiavrammsd comment: often 307 > 301
		http.StatusTemporaryRedirect)
}

func getHandler() *macaron.Macaron {

	m := macaron.New()

	switch BuildMode {
	case "prod":
		macaron.Env = macaron.PROD
		macaron.ColorLog = false
	case "test":
		macaron.Env = macaron.TEST
	case "dev":
		macaron.Env = macaron.DEV
		m.Use(macaron.Logger())
	}

	m.Use(recovery.Recovery())
	//m.Use(macaron.Recovery())
	//m.Use(gzip.Gziper())

	m.Use(cache.Cacher())

	m.SetDefaultCookieSecret(macaron.Config().Section("cookie").Key("cookie_secret").String())

	m.Use(session.Sessioner(session.Options{
		// Provider:       macaron.Config().Section("sql").Key("sql_type").String(),
		// ProviderConfig: models.GetOrmUrl(),
		Gclifetime:     macaron.Config().Section("session").Key("session_time").MustInt64(),
		CookiePath:     macaron.Config().Section("cookie").Key("cookie_path").String(),
		CookieName:     macaron.Config().Section("cookie").Key("cookie_name").String(),
		CookieLifeTime: macaron.Config().Section("cookie").Key("cookie_time").MustInt(),
	}))

	m.Use(csrf.Csrfer(csrf.Options{
		Secret: macaron.Config().Section("csrf").Key("csrf_secret").String(),
	}))

	m.Use(macaron.Static(macaron.Config().Section("static").Key("static_path").String(),
		macaron.StaticOptions{
			Prefix:      macaron.Config().Section("static").Key("static_prefix").String(),
			SkipLogging: macaron.Config().Section("static").Key("static_skip_log").MustBool(),
		}))

	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory:       macaron.Config().Section("template").Key("template_path").String(),
		Extensions:      []string{".tmpl", ".html"},
		Charset:         "UTF-8",
		IndentJSON:      macaron.Env != macaron.PROD,
		IndentXML:       true,
		HTMLContentType: "text/html",
	}))

	if macaron.Env == macaron.DEV {
		m.Use(toolbox.Toolboxer(m))
	}

	m.Use(env.Enver())

	models.ModelInit()

	routers.RouterInit(m)

	tasks.TaskInit()

	return m
}

func main() {

	var v bool
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.Parse()

	if v {
		log.Println(fmt.Sprintf("\nBuildVersion: %v\n   BuildTime: %v\n   BuildMode: %v", BuildVersion, BuildTime, BuildMode))
		os.Exit(0)
	}

	if len(BuildMode) == 0 {
		BuildMode = "dev"
	}

	configPath := filepath.Join(macaron.Root, "configs", fmt.Sprintf("config_%v.ini", BuildMode))
	if _, err := macaron.SetConfig(configPath); err != nil {
		log.Println(fmt.Sprintf("App load config:%v error:%v", configPath, err))
		os.Exit(1)
	}

	if err := id.NewId(1, 1, id.GetIdTwepoch()); err != nil {
		log.Println(fmt.Sprintf("App new id worker error:%v", err))
		os.Exit(1)
	}

	path := macaron.Config().Section("log").Key("path").String()
	name := macaron.Config().Section("log").Key("name").String()
	level := macaron.Config().Section("log").Key("level").String()
	num := macaron.Config().Section("log").Key("num").MustInt(100)
	more := macaron.Config().Section("log").Key("more").MustBool(false)

	switch BuildMode {
	case "dev":
		record.NewConsoleRecord(record.GetRecordLevel(level), num, more)
	case "test":
	case "prod":
		err := record.NewFileRecord(record.GetRecordLevel(level), num, more, path, name)
		if err != nil {
			log.Println(fmt.Sprintf("NewFileRecord error:%v", err))
			os.Exit(1)
		}
	}

	model := macaron.Config().Section("server").Key("Model").String()
	host := macaron.Config().Section("server").Key("Host").String()
	port := macaron.Config().Section("server").Key("Port").MustInt(80)
	isAcme := macaron.Config().Section("server").Key("IsAcme").MustBool(false)
	acmePath := macaron.Config().Section("server").Key("AcmePath").String()
	isRedirect := macaron.Config().Section("server").Key("IsRedirect").MustBool(false)
	redirectPort := macaron.Config().Section("server").Key("RedirectPort").MustInt(80)
	certFile := macaron.Config().Section("server").Key("CertFile").String()
	keyFile := macaron.Config().Section("server").Key("KeyFile").String()
	readTimeout := macaron.Config().Section("server").Key("ReadTimeout").MustInt(10)
	writeTimeout := macaron.Config().Section("server").Key("WriteTimeout").MustInt(10)
	maxHeaderBytes := macaron.Config().Section("server").Key("MaxHeaderBytes").MustInt(1)

	server := &http.Server{
		Addr:           fmt.Sprintf("%v:%v", host, port),
		Handler:        getHandler(),
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes << 20,
	}

	switch model {
	case "http":
		go func() {
			fmt.Println(server.ListenAndServe())
		}()
	case "https":
		if _, err := tls.LoadX509KeyPair(certFile, keyFile); err != nil {
			log.Println(fmt.Sprintf("https certificate error:%v", err))
			os.Exit(1)
		}

		if isAcme {
			hostPolicy := func(ctx context.Context, checkHost string) error {
				// Note: change to your real host
				if host == checkHost {
					return nil
				}
				return fmt.Errorf("acme/autocert: only %s host is allowed", host)
			}

			certPath := filepath.Join(macaron.Root, acmePath)
			os.MkdirAll(certPath, os.ModePerm)

			m := &autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				HostPolicy: hostPolicy,
				Cache:      autocert.DirCache(certPath),
			}
			server.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
			server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

			go func() {
				log.Println(server.ListenAndServeTLS("", ""))
			}()
		} else {
			cfg := &tls.Config{
				MinVersion:               tls.VersionTLS12,
				CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				PreferServerCipherSuites: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
			}
			server.TLSConfig = cfg
			server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

			go func() {
				log.Println(server.ListenAndServeTLS(certFile, keyFile))
			}()
		}

		// redirect every http request to https
		if isRedirect {
			go func() {
				log.Println(http.ListenAndServe(fmt.Sprintf("%v:%v", host, redirectPort), http.HandlerFunc(redirectToHttps)))
			}()
		}
	}

	log.Println("Server running on:", fmt.Sprintf("%v://%v:%v", model, host, port))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

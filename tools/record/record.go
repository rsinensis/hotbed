package record

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	STACK = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var recorder *Record

type Record struct {
	FileMode   bool
	FilePath   string
	FileName   string
	More       bool
	Level      int
	Daily      int
	ChanNum    int
	MaxDays    int
	Writer     *os.File
	Logger     *log.Logger
	RecordChan chan string
}

func GetRecordLevel(val string) (level int) {
	switch val {
	case "stack":
		level = STACK
	case "fatal":
		level = FATAL
	case "debug":
		level = DEBUG
	case "info":
		level = INFO
	case "warn":
		level = WARN
	case "error":
		level = ERROR
	default:
		level = INFO
	}

	return level
}

func GetRecorder() *Record {
	return recorder
}

func NewConsoleRecord(level, chanNum int, more bool) {

	logger := log.New(os.Stdout, "[App] ", log.Ldate|log.Lmicroseconds)

	r := &Record{FileMode: false, FilePath: "", FileName: "", More: more, Level: level, Daily: time.Now().Day(), Writer: nil, Logger: logger, RecordChan: make(chan string, chanNum)}

	go r.Print()

	recorder = r
}

func NewFileRecord(level, chanNum int, more bool, filePath, fileName string) error {

	os.MkdirAll(filePath, os.ModePerm)

	fd, err := os.OpenFile(filepath.Join(filePath, fileName), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}

	logger := log.New(fd, "[App] ", log.Ldate|log.Lmicroseconds)

	r := &Record{FileMode: true, FilePath: filePath, FileName: fileName, More: more, Level: level, Daily: time.Now().Day(), Writer: nil, Logger: logger, RecordChan: make(chan string, chanNum)}

	go r.Print()

	recorder = r

	return nil
}

func (r *Record) Close() {
	close(r.RecordChan)

	if r.Writer != nil {
		r.Writer.Close()
	}
}

func (r *Record) Detele() {

	filepath.Walk(r.FilePath, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.ModTime().Unix() < (time.Now().Unix()-60*60*24*int64(r.MaxDays)) {
			os.Remove(file)
		}

		return nil
	})

}

func (r *Record) Print() {

	select {
	case <-time.After(5 * time.Second):

	case msg := <-r.RecordChan:
		if r.FileMode {
			day := time.Now().Day()

			if day != r.Daily {

				r.Writer.Close()

				fp := filepath.Join(r.FilePath, r.FileName)

				old := time.Now().AddDate(0, 0, -1).Format("20060102")

				os.Rename(fp, fmt.Sprintf("%s.%s", fp, old))

				fd, err := os.OpenFile(fp, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
				if err != nil {
					log.Println(err)
					os.Exit(1)
				}

				logger := log.New(fd, "", log.Ldate|log.Lmicroseconds)

				r.Writer = fd
				r.Logger = logger
				r.Daily = day

				go r.Detele()
			}
		}

		r.Logger.Println(msg)
	}
}

func getRecordPosition() (string, int) {
	_, file, line, ok := runtime.Caller(2)

	if !ok {
		return "", 0
	}

	n := strings.LastIndex(file, "/")
	if n >= 0 {
		file = file[n+1:]
	}

	return file, line
}

func (r *Record) Debug(msg interface{}) {
	if r.Level <= DEBUG {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[DEBUG] (%s:%d) %v", file, line, msg)
		} else {
			r.RecordChan <- fmt.Sprintf("[DEBUG] %v", msg)

		}

	}
}

func (r *Record) Debugf(format string, v ...interface{}) {
	if r.Level <= DEBUG {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[DEBUG] (%s:%d) %v", file, line, fmt.Sprintf(format, v...))
		} else {
			r.RecordChan <- fmt.Sprintf("[DEBUG] %v", fmt.Sprintf(format, v...))

		}

	}
}

func (r *Record) Info(msg interface{}) {
	if r.Level <= INFO {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[INFO] (%s:%d) %v", file, line, msg)
		} else {
			r.RecordChan <- fmt.Sprintf("[INFO] %v", msg)

		}

	}
}

func (r *Record) Infof(format string, v ...interface{}) {
	if r.Level <= INFO {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[INFO] (%s:%d) %v", file, line, fmt.Sprintf(format, v...))
		} else {
			r.RecordChan <- fmt.Sprintf("[INFO] %v", fmt.Sprintf(format, v...))

		}

	}
}

func (r *Record) Warn(msg interface{}) {
	if r.Level <= WARN {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[WARN] (%s:%d) %v", file, line, msg)
		} else {
			r.RecordChan <- fmt.Sprintf("[WARN] %v", msg)

		}

	}
}

func (r *Record) Warnf(format string, v ...interface{}) {
	if r.Level <= WARN {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[WARN] (%s:%d) %v", file, line, fmt.Sprintf(format, v...))
		} else {
			r.RecordChan <- fmt.Sprintf("[WARN] %v", fmt.Sprintf(format, v...))

		}

	}
}

func (r *Record) Error(msg interface{}) {
	if r.Level <= ERROR {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[ERROR] (%s:%d) %v", file, line, msg)
		} else {
			r.RecordChan <- fmt.Sprintf("[ERROR] %v", msg)

		}

	}
}

func (r *Record) Errorf(format string, v ...interface{}) {
	if r.Level <= ERROR {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[ERROR] (%s:%d) %v", file, line, fmt.Sprintf(format, v...))
		} else {
			r.RecordChan <- fmt.Sprintf("[ERROR] %v", fmt.Sprintf(format, v...))

		}

	}
}

func (r *Record) Fatal(msg interface{}) {
	if r.Level <= FATAL {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[FATAL] (%s:%d) %v", file, line, msg)
		} else {
			r.RecordChan <- fmt.Sprintf("[FATAL] %v", msg)
		}

	}
}

func (r *Record) Fatalf(format string, v ...interface{}) {
	if r.Level <= FATAL {

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[FATAL] (%s:%d) %v", file, line, fmt.Sprintf(format, v...))
		} else {
			r.RecordChan <- fmt.Sprintf("[FATAL] %v", fmt.Sprintf(format, v...))
		}

	}
}

func (r *Record) Stack(msg interface{}) {

	if r.Level <= STACK {

		s := fmt.Sprint(msg)
		s += "\n"
		buf := make([]byte, 1024*1024)
		n := runtime.Stack(buf, true)
		s += string(buf[:n])
		s += "\n"

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[Stack] (%s:%d) %v", file, line, s)
		} else {
			r.RecordChan <- fmt.Sprintf("[Stack] %v", s)
		}

	}

}

func (r *Record) Stackf(format string, v ...interface{}) {

	if r.Level <= STACK {

		s := fmt.Sprintf(format, v...)
		s += "\n"
		buf := make([]byte, 1024*1024)
		n := runtime.Stack(buf, true)
		s += string(buf[:n])
		s += "\n"

		if r.More {
			file, line := getRecordPosition()
			r.RecordChan <- fmt.Sprintf("[Stack] (%s:%d) %v", file, line, s)
		} else {
			r.RecordChan <- fmt.Sprintf("[Stack] %v", s)
		}
	}

}

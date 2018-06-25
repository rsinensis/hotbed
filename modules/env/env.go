package env

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

type Env struct {
	*macaron.Context
	Cache   cache.Cache
	Csrf    csrf.CSRF
	Flash   *session.Flash
	Session session.Store
}

func Enver() macaron.Handler {

	return func(ctx *macaron.Context, cache cache.Cache, sess session.Store, flash *session.Flash, x csrf.CSRF) {

		c := &Env{
			Context: ctx,
			Cache:   cache,
			Flash:   flash,
			Csrf:    x,
			Session: sess,
		}

		ctx.Map(c)
	}
}

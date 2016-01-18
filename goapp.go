package main

import (
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

// GoApp is application templete.
type GoApp struct {
	mux      *http.ServeMux
	store    *sessions.CookieStore
	handlers map[string]http.HandlerFunc
}

// NewGoApp is construct GoApp.
func NewGoApp(name ...string) (app *GoApp) {
	// load settings
	appConf := LoadAppCredentials()
	anaconda.SetConsumerKey(appConf.ConsumerKey)
	anaconda.SetConsumerSecret(appConf.ConsumerSecret)

	app = &GoApp{
		mux:   http.NewServeMux(),
		store: sessions.NewCookieStore([]byte(appConf.CookieSecret)),
	}

	// set default handlers
	app.handlers = map[string]http.HandlerFunc{
		"/":         app.homeHandler,
		"/login":    app.loginHandler,
		"/logout":   app.logoutHandler,
		"/callback": app.callbackHandler,
	}
	return
}

// ListenAndServe listen and serve a server.
func (app *GoApp) ListenAndServe(addr string) error {
	app.registerHandlers()
	// wrap mux for prevent memory leak.
	return http.ListenAndServe(addr, context.ClearHandler(app.mux))
}

// registerHandlers register handlers.
func (app *GoApp) registerHandlers() {
	for k, v := range app.handlers {
		app.mux.HandleFunc(k, v)
	}
}

// AddHandler add a handler
func (app *GoApp) AddHandler(path string, handler http.HandlerFunc, force bool) bool {
	if _, ok := app.handlers[path]; !force && ok {
		return false
	}
	app.handlers[path] = handler
	return true
}

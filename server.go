package main

import (
	"net/http"

	"github.com/gorilla/context"
)

// GoApp is application templete.
type GoApp struct {
	mux *http.ServeMux
}

// NewGoApp is construct GoApp.
func NewGoApp(name ...string) (app *GoApp) {
	return &GoApp{mux: http.NewServeMux()}
}

// ListenAndServe listen and serve a server.
func (a *GoApp) ListenAndServe(addr string) error {
	// wrap mux for prevent memory leak.
	return http.ListenAndServe(addr, context.ClearHandler(a.mux))
}

// RegisterHandlers register handlers.
func (a *GoApp) RegisterHandlers() {
	a.mux.HandleFunc("/", homeHandler)
	a.mux.HandleFunc("/login", loginHandler)
	a.mux.HandleFunc("/logout", logoutHandler)
	a.mux.HandleFunc("/afterlogin", afterloginHandler)
}

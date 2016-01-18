package main

import (
	"net/http"
)

func (app *GoApp) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := app.store.Get(r, sessionName)
	session.Values = make(map[interface{}]interface{})
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

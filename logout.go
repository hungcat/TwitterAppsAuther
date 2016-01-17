package main

import (
	//"fmt"
	"net/http"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionName)
	session.Values = make(map[interface{}]interface{})
	session.Save(r, w)
	//fmt.Println(session.Values)

	http.Redirect(w, r, "/", http.StatusFound)

}

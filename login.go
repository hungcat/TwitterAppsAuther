package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
)

func (app *GoApp) loginHandler(w http.ResponseWriter, r *http.Request) {

	// Twitter認可申請
	url, cred, err := anaconda.AuthorizationURL("")
	if err != nil {
		http.Error(w, "Couldn't connect to twitter.", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// 貰った情報をセッションにストアして
	session, err := app.store.Get(r, sessionName)
	marchaled, _ := json.Marshal(cred)
	session.Values["oauth_credentials"] = string(marchaled)
	session.Save(r, w)
	//fmt.Println(session.Values)

	// 指示されたURLへ
	http.Redirect(w, r, url, http.StatusFound)
}

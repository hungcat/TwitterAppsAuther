package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
)

// LoginTemp contains data to be embeded to login template.
type LoginTemp struct {
	AuthURL string
	Text    string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// ここらへんでTwitter認証開始、リクエストトークン情報をsessionへ

	SetConsumerToAnaconda()
	url, cred, err := anaconda.AuthorizationURL("")
	if err != nil {
		http.Error(w, "Couldn't connect to twitter.", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	session, err := store.Get(r, sessionName)
	marchaled, _ := json.Marshal(cred)
	session.Values["oauth_credentials"] = string(marchaled)
	session.Save(r, w)
	//fmt.Println(session.Values)

	http.Redirect(w, r, url, http.StatusFound)
}

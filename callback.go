package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
)

func (app *GoApp) callbackHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := app.store.Get(r, sessionName)

	// アクセストークンを取得したらファイルに書き込みつつログイン情報をsessionへ

	// セッション中の情報をパース
	var cred oauth.Credentials
	err := json.Unmarshal([]byte(session.Values["oauth_credentials"].(string)), &cred)
	if err != nil {
		http.Error(w, "Session has broken.", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	// 取得した証明書を元にアクセストークンを取得
	_, values, err := anaconda.GetCredentials(&cred, r.FormValue("oauth_verifier"))
	if err != nil {
		http.Error(w, "Couldn't authorize.", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// セッションクリア
	session.Values = make(map[interface{}]interface{})

	// ユーザー情報記録
	// oauth_token, oauth_token_secret, user_id, screen_name, x_auth_expires
	for k, v := range values {
		session.Values[k] = v[0]
	}

	session.Save(r, w)

	// 認可後はhomeへ
	http.Redirect(w, r, "/", http.StatusFound)
}

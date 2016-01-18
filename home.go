package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// HomeTemp is hogehoge.
type HomeTemp struct {
	UserName  string
	ButtonVal string
	ButtonURL string
}

func (app *GoApp) homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		fmt.Fprintf(w, "Under construction.")
		return
	}

	var hTemp HomeTemp
	session, _ := app.store.Get(r, sessionName)

	if _, ok := session.Values["user_id"]; ok {
		hTemp = HomeTemp{
			UserName:  session.Values["screen_name"].(string),
			ButtonVal: "Logout",
			ButtonURL: "/logout",
		}
	} else {
		hTemp = HomeTemp{
			UserName:  "new one",
			ButtonVal: "Login",
			ButtonURL: "/login",
		}
	}

	if err := t.Execute(w, hTemp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
	}
}

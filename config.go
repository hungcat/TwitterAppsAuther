package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gorilla/sessions"
)

const (
	sessionName  = "testapp"
	consumerPath = "appconf.json"
)

var store = sessions.NewCookieStore([]byte("hogecat"))

// AppCredentials is consumerkey and consumersecret.
type AppCredentials struct {
	ConsumerKey    string `json:"key"`
	ConsumerSecret string `json:"secret"`
}

// SetConsumerToAnaconda sets consumer params to anaconda.
func SetConsumerToAnaconda() {
	data, err := ioutil.ReadFile(consumerPath)
	if err != nil {
		return
	}

	var appConf AppCredentials
	if json.Unmarshal(data, &appConf) != nil {
		return
	}

	anaconda.SetConsumerKey(appConf.ConsumerKey)
	anaconda.SetConsumerSecret(appConf.ConsumerSecret)
}

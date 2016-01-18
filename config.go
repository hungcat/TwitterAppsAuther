package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

const (
	sessionName  = "testapp"
	consumerPath = "appconf.json"
)

// AppCredentials is consumerkey and consumersecret.
type AppCredentials struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	CookieSecret   string `json:"cookie_secret"`
}

// LoadAppCredentials loads credentials of application.
func LoadAppCredentials() (appConf *AppCredentials) {
	appConf = &AppCredentials{}

	data, err := ioutil.ReadFile(consumerPath)
	if err != nil {
		return
	}

	if json.Unmarshal(data, appConf) != nil {
		return
	}

	return
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

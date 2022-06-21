package main

import (
	"fmt"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigUser struct {
    Username	string `env:"USERNAME"`
    Password 	string `env:"PASSWORD"`
}

var cfg ConfigUser

func Auth(w http.ResponseWriter, r *http.Request) bool {
	_ = cleanenv.ReadConfig(".env", &cfg)
	username, password, ok := r.BasicAuth()
	if !ok {
		fmt.Println(`Something wrong`)
		w.WriteHeader(401)
		return false
	}

	isValid := (username == cfg.Username) && (password == cfg.Password)
	if !isValid {
		fmt.Println(`wrong username/password`)
		w.WriteHeader(401)
		return false
	}
	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool  {
	if r.Method != "GET" {
		fmt.Println("Only GET is Allowed")
		w.WriteHeader(400)
		return false
	}
	return true
}
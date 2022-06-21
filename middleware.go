package main

import "net/http"

const USERNAME = "username"
const PASSWORD = "password"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`Something wrong`))
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}
	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool  {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is Allowed"))
		return false
	}
	return true
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const USERNAME = "username"
const PASSWORD = "password"

var students = []*Student{}
type Student struct {
	Id    string
	Name  string
	Grade int32
}


func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8080"

	fmt.Println("server started at localhost:8080")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request)  {
	if !Auth(w, r)			{ return }
	if !AllowOnlyGET(w, r)	{ return }

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(0)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

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

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{Id: "s1", Name: "Budi", Grade: 5})
	students = append(students, &Student{Id: "s2", Name: "Yudi", Grade: 1})
	students = append(students, &Student{Id: "s2", Name: "Dudi", Grade: 3})
}
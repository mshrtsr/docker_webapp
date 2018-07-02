package main

import (
	"./CRUD"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var tb_name = os.Getenv("DB_TABLE")

const contentType = "application/json"

type Msg struct {
	Message string `json:"message"`
}

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func message2json(str string) []byte {
	var msg = Msg{
		Message: str,
	}

	jsonData, _ := json.Marshal(msg)
	return jsonData
}

func user2json(user CRUD.User) []byte {
	var response = Response{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Created_at: user.Created_at.Format(time.RFC3339Nano),
		Updated_at: user.Updated_at.Format(time.RFC3339Nano),
	}

	jsonData, _ := json.Marshal(response)
	return jsonData
}

func users2json(users []CRUD.User) []byte {
	var response []Response
	for i := range users {
		response = append(response, Response{
			Id:         users[i].Id,
			Name:       users[i].Name,
			Email:      users[i].Email,
			Created_at: users[i].Created_at.Format(time.RFC3339Nano),
			Updated_at: users[i].Updated_at.Format(time.RFC3339Nano),
		})
	}

	jsonData, _ := json.Marshal(response)
	return jsonData
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Check Content-Type
	if r.Header.Get("Content-Type") != contentType {
		w.WriteHeader(404)
		return
	}

	// Parse URL
	var isState_users bool = false
	var isState_id bool = false
	var id int

	slicedPath := strings.Split(r.URL.Path, "/")
	if len(slicedPath) >= 2 {
		switch slicedPath[1] {
		case "users":
			isState_users = true
		default:
		}
	}
	if len(slicedPath) == 3 {
		id, _ = strconv.Atoi(slicedPath[2])
		if id != 0 {
			isState_id = true
		}
	}

	// Docode json input data
	var rtn int64
	var requests Request

	bufbody := new(bytes.Buffer)
	rtn, _ = bufbody.ReadFrom(r.Body)
	//fmt.Println(rtn)
	if rtn != 0 {
		body := bufbody.Bytes()
		if err := json.Unmarshal(body, &requests); err != nil {
			log.Fatal(err)
		}
	}

	name := requests.Name
	email := requests.Email

	// CRUD
	var user CRUD.User
	var users []CRUD.User
	var jsonData []byte

	//fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		switch {
		case isState_id:
			fmt.Println("Read")
			user = CRUD.ReadData(id, tb_name)
			jsonData = user2json(user)
			w.WriteHeader(200)
		case isState_users:
			fmt.Println("ReadAll")
			users = CRUD.ReadDataAll(tb_name)
			jsonData = users2json(users)
			w.WriteHeader(200)
		default:
			fmt.Println("Hello World!!")
			jsonData = message2json("Hello World!!")
			w.WriteHeader(200)
		}
	case "POST":
		switch {
		case isState_id:
			w.WriteHeader(404)
			return
		case isState_users:
			fmt.Println("Create")
			user = CRUD.CreateData(name, email, tb_name)
			jsonData = user2json(user)
			w.WriteHeader(201)
		default:
			w.WriteHeader(404)
			return
		}
	case "PUT":
		switch {
		case isState_id:
			fmt.Println("Update")
			user = CRUD.UpdateData(id, name, email, tb_name)
			jsonData = user2json(user)
			w.WriteHeader(200)
		default:
			w.WriteHeader(404)
			return
		}
	case "DELETE":
		switch {
		case isState_id:
			fmt.Println("Delete")
			CRUD.DeleteData(id, tb_name)
			w.WriteHeader(204)
		default:
			w.WriteHeader(404)
			return
		}
	default:
		w.WriteHeader(404)
		return
	}

	jsonFormedData := new(bytes.Buffer)
	json.Indent(jsonFormedData, jsonData, "", "    ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, jsonFormedData.String())

}

func main() {
	CRUD.WaitDB()
	CRUD.DropTable(tb_name)
	CRUD.CreateTable(tb_name)
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

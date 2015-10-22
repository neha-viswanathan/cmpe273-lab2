/*
* CMPE 273 Lab 2 - A simple “Hello World” REST API in Go
* Neha Viswanathan
* 010029097
*/

package main

//import statements
import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)


//struct declarations
type Username struct {
	Name string `json:"name"`
}

type Message struct {
	Greet string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

//postHandler function to handle POST requests
func postHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var usr Username
	json.NewDecoder(req.Body).Decode(&usr)
	var msg Message
	msg.Greet = "Hello, " + usr.Name +"!"
	j, _ := json.Marshal(msg)

	rw.Header().Set("Content-Type", "application/json")
	//rw.WriteHeader(http.StatusCreated)
	fmt.Fprintf(rw, "%s", j)
}

//Main method
func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello", postHandler)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

/*
GET and POST REST API calls in golang
Neha Viswanathan
010029097
*/
package main

//import statement
import (
         "encoding/json"
         "io/ioutil"
         "net/http"
         "github.com/gorilla/mux"
         "fmt"
 )

//struct definitions
type User struct {
   Name string
 }

type Greeting struct {
    Greet string
 }

//function to handle GET requests
func getHandler(respW http.ResponseWriter, req *http.Request) {
  userName:=mux.Vars(req)["name"]
  fmt.Fprintf(respW, "Hello , %s!\n", userName)
 }

//function to handle POST requests
func postHandler(respW http.ResponseWriter, req *http.Request) {
  respW.Header().Set("Content-Type", "application/json")
  var user User
  b, _ := ioutil.ReadAll(req.Body)
  json.Unmarshal(b, &user)
  var g Greeting
  g.Greet = "Hello, " + user.Name + "!"
  j, _ := json.Marshal(g)
  respW.Write(j)
}

func main() {
  route := mux.NewRouter()
  route.HandleFunc("/user/{name}", getHandler).Methods("GET")
  route.HandleFunc("/user", postHandler).Methods("POST")

  http.Handle("/", route)
  http.ListenAndServe(":8080", nil)
}

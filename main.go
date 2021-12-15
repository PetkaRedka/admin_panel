package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s", u.name)
}

func main() {

	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8090", nil)

}

func homePage(w http.ResponseWriter, r *http.Request) {

	addCorsHeader(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// else {
	// 	h.APIHandler.ServeHTTP(w, r)
	// }
	// bob := User{"13", 2, 2, 2.3, 3.3}

	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyBytes, err := ioutil.ReadAll(strings.NewReader(`{"connections":[{"surname": "Frozen Yogurt4", "class": 8080, "port": 8089, "status": "Online", "session": "asdas21e909"}, {"surname": "Frozen Yogurt4", "class": 8080, "port": 8089, "status": "Online", "session": "asdas21e909"}]}`))
	//fmt.Fprintf(w, "Go is working! "+bob.getAllInfo())

	//resp := make(map[string]string)
	//resp["message"] = "Status Created"
	//jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(bodyBytes)

}

func addCorsHeader(r http.ResponseWriter) {
	headers := r.Header()
	headers.Add("Content-Type", "application/json")
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct{}

func (h handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	dict := make(map[string]string)
	dict["host"] = (*req).Host
	dict["user_agent"] = (*req).UserAgent()
	dict["request_uri"] = (*req).RequestURI
	var headers string
	for key, value := range (*req).Header {
		headers += fmt.Sprintf("%s : %s\n", key, value)
	}
	dict["headers"] = headers

	response, _ := json.Marshal(dict)
	fmt.Fprintf(resp, string(response))
}

func main() {
	fmt.Println("Server is listening...")
	var h handler
	http.ListenAndServe(":3535", h)
}

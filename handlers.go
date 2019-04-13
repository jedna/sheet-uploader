package main

import (
	"fmt"
	"net/http"
	"strings"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	if r.Method != "POST" {
		http.Error(w, "Bad request", 400)
		return
	}

	//requestDump, err := httputil.DumpRequest(r, true)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(requestDump))

	r.ParseForm()

	Write(strings.Join(r.Form["email"], ""), strings.Join(r.Form["name"], ""), strings.Join(r.Form["rsvp"], "") == "1")

	fmt.Fprintln(w, "Done.")
}

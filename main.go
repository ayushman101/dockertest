package main

import (
	
	"fmt"
	"log"
	"net/http"
	"html"
)


func main()
{

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello there %q,", html.EscapeString(r.URL.Path))
	})


	log.Fatal(http.ListenAndServe(":8081",nil))
}

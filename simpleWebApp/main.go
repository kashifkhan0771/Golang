package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":8080", nil); err != nil{
		_ = fmt.Errorf("cannot start a server due to %+v", err)
	}
}

func homePage(w http.ResponseWriter, r * http.Request){
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil{
		panic(err)
	}
}

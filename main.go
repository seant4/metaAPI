package main

import(
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Meta struct{
	CHARACTER	string `json:"Character on the rise"`
	PLAYER	    string `json:"Player"`
}

var metaData Meta

func getMeta(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metaData)
}

func handleSetChar(){
	var c string
	var p string
	fmt.Println("What would you like to set CHARACTER too?")
	fmt.Scanln(&c)
	fmt.Println("Setting CHARACTER too")
	fmt.Println(c)
	fmt.Println("What would you like to set PLAYER too?")
	fmt.Scanln(&p)
	fmt.Println("Setting PLAYER too")
	fmt.Println(c);
	metaData = Meta{CHARACTER: c, PLAYER: p}
	fmt.Println(metaData)
	fmt.Println("Thank you")
	welcome();
}

func welcome(){
	var input string
	fmt.Println("Welcome to the meta API :) Created by Sean Theisen")
	fmt.Println("Would you like to set data? y/n")
	fmt.Scanln(&input);
	if(input == "y"){
		handleSetChar();
	}else if(input == "n"){
		fmt.Println("Thank you")
	}else{
		welcome();
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/meta", getMeta).Methods("GET");
	go welcome();
	log.Fatal(http.ListenAndServe(":8000", r))
}
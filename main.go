package main

import(
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"crypto/tls"
	"crypto/x509"
)

const (
	localCertFile = "104.131.86.238.crt"
)

type Meta struct{
	CHARACTER	string `json:"Trending Ultimate Character: "`
	PLAYER	    string `json:"Trending Ultimate Player: "`
	CHARACTERM  string `json:"Trending Melee Character: "`
	PLAYERM	    string `json:"Trending Melee Player: "`
}

var metaData Meta

func getMeta(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(metaData)
}

func handleSetChar(){
	var c string
	var p string
	var cm string
	var pm string
	fmt.Println("What would you like to set ULTIMATE CHARACTER too?")
	fmt.Scanln(&c)
	fmt.Println("Setting ULTIMATE CHARACTER too")
	fmt.Println(c)
	fmt.Println("What would you like to set ULTIMATE PLAYER too?")
	fmt.Scanln(&p)
	fmt.Println("Setting ULTIMATE PLAYER too")
	fmt.Println(c)
	fmt.Println("What would you like to set MELEE CHARACTER too?")
	fmt.Scanln(&cm)
	fmt.Println("Setting MELEE CHARACTER too")
	fmt.Println(cm)
	fmt.Println("What would you like to set MELEE PLAYER too?")
	fmt.Scanln(&pm)
	fmt.Println("Setting MELEE PLAYER too")
	fmt.Println(pm)
	metaData = Meta{CHARACTER: c, PLAYER: p, CHARACTERM: cm, PLAYERM: pm}
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
	insecure := flag.Bool("insecure-ssl", false, "Accept/Ignore all server SSL certificates")
	flag.Parse()

	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	certs, err := ioutil.ReadFile(localCertFile)
	if err != nil {
		log.Fatalf("Failed to append %q to RootCAs: %v", localCertFile, err)
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		InsecureSkipVerify: *insecure,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	// Uses local self-signed cert
	req := http.NewRequest(http.MethodGet, "https://localhost/api/version", nil)
	resp, err := client.Do(req)
	// Handle resp and err

	// Still works with host-trusted CAs!
	req = http.NewRequest(http.MethodGet, "https://example.com/", nil)
	resp, err = client.Do(req)
	r := mux.NewRouter()
	r.HandleFunc("/api/meta", getMeta).Methods("GET");
	r.Use(mux.CORSMethodMiddleware(r))
	go welcome();
	log.Fatal(http.ListenAndServeTLS(":8000", localCertFile, "104.131.86.238.key", nil ))
}
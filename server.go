package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

//HelloHandler handles requests for the `/hello` resource
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, Web!\n"))
}

func PayloadHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Reading payload!\n"))
    // read the body of request
    var bodyBytes []byte
    if r.Body != nil {
        bodyBytes, _ = ioutil.ReadAll(r.Body)
    }
    // write bodyBytes to file
    err := ioutil.WriteFile("payload.txt", bodyBytes, 0644)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {

    //get the value of the ADDR environment variable
    addr := goDotEnvVariable("ADDRESS")

    //if it's blank, default to ":80", which means
    //listen port 80 for requests addressed to any host
    if len(addr) == 0 {
        addr = ":80"
    }

    //create a new mux (router)
    //the mux calls different functions for
    //different resource paths
    mux := http.NewServeMux()

    //tell it to call the HelloHandler() function
    //when someone requests the resource path `/hello`
    mux.HandleFunc("/", HelloHandler)
    mux.HandleFunc("/payload", PayloadHandler)

    //start the web server using the mux as the root handler,
    //and report any errors that occur.
    //the ListenAndServe() function will block so
    //this program will continue to run until killed
    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}
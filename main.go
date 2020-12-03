package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// generateJSON creates the spire info and certificate bundle that is returned
// by the web server
func generateJSON() (string, error) {
	// Bundle contains the struct used to generate the spire info json
	type Bundle struct {
		Domain     string
		Server     string
		CertBundle string
	}

	certBundle, err := ioutil.ReadFile(os.Getenv("BUNDLE"))
	if err != nil {
		return "", err
	}

	s := Bundle{os.Getenv("DOMAIN"), os.Getenv("SERVER"), string(certBundle)}

	j, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(j), nil
}

// handler returns the generated json spire info
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	j, err := generateJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, j)
}

// main starts the http server
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

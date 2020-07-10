package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/shvetsiya/distribkv/db"
)

var (
	dbLocation = flag.String("db-location", "", "The path to bold db")
	httpAddr = flag.String("http-addr", "127.0.0.1:8080", "http host and port")
)

func parseFlags() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatal("must provide db location")
	}
}

func main() {
	parseFlags()

	db, close, err := db.NewDatabase(*dbLocation)
	if err != nil {
		log.Fatalf("NewDatabase(%q): v", *dbLocation, err)
	}
	defer close()

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.Form.Get("key")
		value, err := db.GetKey(key)
		fmt.Fprintf(w, "Value = %q, err = %v\n", value, err)
	})


	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.Form.Get("key")
		value := r.Form.Get("value")
		err := db.SetKey(key, []byte(value))
		fmt.Fprintf(w, "Error = %v\n", err)
	})

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
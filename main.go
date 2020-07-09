package main

import (
	"flag"
	"log"
)

var (
	dbLocation = flag.String("db-location", "", "The path to bold db")
)

func parseFlags() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatal("must provide db location")
	}
}

func main() {
	parseFlags()

}

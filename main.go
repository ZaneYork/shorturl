package main

//go:generate rice embed-go

import (
	"fmt"
	"log"
	"os"

	"github.com/asdine/storm"
	"github.com/namsral/flag"
)

var (
	cfg Config
	db  *storm.DB
)

func main() {
	var (
		version   bool
		config    string
		urlLength int
		dbpath    string
		baseurl   string
		bind      string
	)

	flag.BoolVar(&version, "v", false, "display version information")
	flag.StringVar(&config, "config", "", "config file")
	flag.IntVar(&urlLength, "urllen", 5, "short url length")
	flag.StringVar(&dbpath, "dbpath", "urls.db", "Database path")
	flag.StringVar(&baseurl, "baseurl", "", "Base URL for display purposes")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.Parse()

	if version {
		fmt.Printf("shorturl v%s", FullVersion())
		os.Exit(0)
	}

	var err error
	db, err = storm.Open(dbpath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO: Abstract the Config and Handlers better
	cfg.baseURL = baseurl
	cfg.urlLength = urlLength

	NewServer(bind, cfg).ListenAndServe()
}

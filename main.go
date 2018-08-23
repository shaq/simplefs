package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	file := flag.String("f", "", "files to share")
	port := flag.String("p", "8089", "port to serve on")
	flag.Parse()

	if err := shareFile(*file); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	log.Printf("listening on address :%v", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func shareFile(file string) error {
	if len(file) == 0 {
		return fmt.Errorf("please give a file name")
	} else if _, err := os.Stat(file); os.IsNotExist(err) {
		return fmt.Errorf("%v", err)
	}

	log.Printf("got file '%v'", file)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, file)
		}
	})
	return nil
}

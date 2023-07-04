package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa "github.com/tusharr-patil/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA")
	fileName := flag.String("file", "gopher.json", "the json file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *fileName)

	f, err := os.Open(*fileName)

	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

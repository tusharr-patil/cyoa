package main

import (
	"flag"
	"fmt"
	"os"

	cyoa "github.com/tusharr-patil/cyoa"
)

func main() {
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

	fmt.Printf("%+v\n", story["debate"])
}

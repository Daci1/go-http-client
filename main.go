package main

import (
	"log"
	"os"

	"github.com/Daci1/go-http-client/internal"
)

func main() {
	args := os.Args[1:]
	reqConfig, err := internal.ParseFlags(args)

	if err != nil {
		log.Fatalln(err)
	}

	res, err := internal.MakeCall(reqConfig)

	if err != nil {
		log.Fatalln(err)
	}

	if err := res.Print(); err != nil {
		log.Fatalln(err)
	}
}

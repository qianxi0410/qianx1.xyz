package main

import (
	"flag"
	"log"
)

var op = flag.String("op", "", "which operation to run")

func init() {
	flag.Parse()

	if *op != "create" && *op != "delete" && *op != "update" {
		log.Fatalf("unexpected operation: %v", *op)
	}
}

func main() {

	switch *op {
	case "create":
		createPost()
	case "delete":
		deletePost()
	case "update":
		updatePost()
	default:
		log.Fatalf("unexpected operation: %v", *op)
	}

}

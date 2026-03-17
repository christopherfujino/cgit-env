package main

import (
	"fmt"
	"net/http"
	"christopherfujino.com/distributed-compute-monorepo/notes"
)

func main() {
	notes.Create("./notes", "/notes").Register()
	fmt.Println("Listening on 127.0.0.1:8080")

	var err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}

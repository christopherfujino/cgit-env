package main

import (
	"christopherfujino.com/ros/ros-open/notes"
	"christopherfujino.com/ros/ros-open/service"

	"fmt"
	"net/http"
)

var services = []service.T{
	notes.Create("./notes", "/notes"),
}

func main() {
	for _, service := range services {
		service.Register()
	}
	fmt.Println("Listening on 127.0.0.1:8080")

	var err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}

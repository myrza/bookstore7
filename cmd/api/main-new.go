package main

import (
	"fmt"
	"bookstore7/internal/server"
)

func mainnew() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

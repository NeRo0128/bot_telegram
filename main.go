package main

import (
	"bot_telegram/server"
	"fmt"
)

func main() {
	err := server.LoadServer()
	if err != nil {
		fmt.Println(err)
	}
}

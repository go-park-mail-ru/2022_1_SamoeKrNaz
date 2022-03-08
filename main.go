package main

import "fmt"

func main() {

	router := initRouter()

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

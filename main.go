package main

import "fmt"

func main() {

	router := initRouter()
	if router == nil {
		_ = fmt.Errorf("Db error")
		return
	}

	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

package main

import "fmt"

func main() {

	router, err := initRouter()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

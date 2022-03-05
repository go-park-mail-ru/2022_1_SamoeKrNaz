package main

func main() {

	router := initRouter()

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

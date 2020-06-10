package main

func main() {
	router := router.Init()
	router.Logger.Fatal(router.Start(":8080"))
}

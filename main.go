package main

import "github.com/tensaitensai/TimeUS-api/route"

func main() {
	router := route.Init()
	route.Logger.Fatal(route.Start(":8080"))
}

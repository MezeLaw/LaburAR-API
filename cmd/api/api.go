package main

import (
	"LaburAR/cmd/api/routes"
	"LaburAR/cmd/database"
	"fmt"
)

func main() {
	fmt.Println("<<< Starting App >>>")

	database.InitDatabase()
	routes.InitRoutes().Run()

	fmt.Println("<<< App started >>>")
}

// POST https://betterprogramming.pub/hands-on-with-jwt-in-golang-8c986d1bb4c0

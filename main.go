package main

import (
	"fmt"
	"cms-comment/src/infrastructure/database"
	"cms-comment/src/router"
)


func main() {
	fmt.Println("Hello world !!!")
	// Init db mongodb
	database.Init()
	
	// Init router
	r := router.Init()
	fmt.Println(r.Run("0.0.0.0:8888"))
}

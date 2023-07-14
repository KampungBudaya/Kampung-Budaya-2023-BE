package main

import "github.com/KampungBudaya/Kampung-Budaya-2023-BE/app"

//	@title			Kampung Budaya's API
//	@version		1.0
//	@description	APIs that are used on Kampung Budaya 2023's Festival

// @license.name	Apache 2.0
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

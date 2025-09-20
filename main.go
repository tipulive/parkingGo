package main

import (
	"parking/config"
	"parking/db"
	"parking/route"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)
	route.ApiRoute()

}

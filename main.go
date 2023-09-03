package main

import (
	"github.com/MarcoVitoC/pbi-btpns/router"
	"github.com/MarcoVitoC/pbi-btpns/database"
)

func main() {
	database.DB()
	r := router.SetRouter()
	r.Run()
}
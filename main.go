package main

import (
	"GiveGift/db"
	"GiveGift/webService"
)

func main() {
	db.InitDB()
	webService.StartWebService()
	//db.GetRank()
}

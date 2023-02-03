package main

import (
	"log"

	"github.com/mk1589/Week-Two-GL-CipherSchools/database"
	"github.com/mk1589/Week-Two-GL-CipherSchools/routers"
)

func main() {
	database.Setup()                    //establish the database connection
	engine := routers.Router()          //get the customized engien
	err := engine.Run("127.0.0.1:8080") //strat the engien
	if err != nil {
		log.Fatal(err)
	}

}

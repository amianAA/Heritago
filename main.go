package main

import (
	log "heritage/backend/logger"
	"heritage/backend/orm"
	"heritage/backend/server"
)

func main() {
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm)
}

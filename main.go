package main

import (
	log "heritago/backend/logger"
	"heritago/backend/orm"
	"heritago/backend/server"
)

func main() {
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm)
}

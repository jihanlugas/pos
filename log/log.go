package log

import (
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/utils"
	"time"
)

func AddLog(log model.Log) {

	conn, closeConn := db.GetConnection()
	defer closeConn()

	log.ID = utils.GetUniqueID()
	log.CreateDt = time.Now()

	conn.Create(&log)
}

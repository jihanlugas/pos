package log

import (
	"github.com/jihanlugas/pos/db"
	"github.com/jihanlugas/pos/model"
	"github.com/jihanlugas/pos/utils"
	"time"
)

func AddLog(path, loginuser, meessage, request, response string) {

	conn, closeConn := db.GetConnection()
	defer closeConn()

	log := model.Log{
		ID:        utils.GetUniqueID(),
		Path:      path,
		Loginuser: loginuser,
		Message:   meessage,
		Request:   request,
		Response:  response,
		CreateDt:  time.Now(),
	}

	conn.Save(&log)
}

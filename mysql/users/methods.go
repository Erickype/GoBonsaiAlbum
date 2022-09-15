package users

import (
	"database/sql"
	"github.com/Erickype/GoBonsaiAlbum/gRPC/models"
	"github.com/Erickype/GoBonsaiAlbum/mysql"
	"time"
)

func CreateUser(user *models.User) {
	db, err := mysql.GetMysqlConnection()
	if err != nil {
		panic(err)
	}

	q := `INSERT INTO users.user(userName, userLastname, userNickname, createdAt) VALUES(?, ?, ?, ?)`
	stmtIns, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			panic(err)
		}
	}(stmtIns)

	date := time.Now().Format(time.RFC3339)
	_, err = stmtIns.Exec(user.UserName, user.UserLastname, user.UserNickname, date)

	err = db.Close()
	if err != nil {
		panic(err)
	}
}

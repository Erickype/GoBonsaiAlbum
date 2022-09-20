package users

import (
	"database/sql"
	"github.com/Erickype/GoBonsaiAlbum/models"
	"github.com/Erickype/GoBonsaiAlbum/mysql"
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

func CreateUser(user *models.User) (int64, error) {
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

	createdAt := time.Now().Format(time.RFC3339)
	result, err := stmtIns.Exec(user.UserName, user.UserLastname, user.UserNickname, createdAt)

	err = db.Close()
	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func GetUsers() []*models.User {

	var users []*models.User

	db, err := mysql.GetMysqlConnection()
	if err != nil {
		panic(err)
	}
	/*WHERE id > ?*/
	stmtOut, err := db.Prepare("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}
	defer func(stmtOut *sql.Stmt) {
		err := stmtOut.Close()
		if err != nil {
			panic(err)
		}
	}(stmtOut)

	rows, err := stmtOut.Query()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user = models.User{}
		var createdAt time.Time
		err = rows.Scan(&user.Id, &user.UserName, &user.UserLastname, &user.UserNickname, &createdAt)
		if err != nil {
			panic(err)
		}
		user.CreatedAt = date.Date{
			Year:  int32(createdAt.Year()),
			Month: int32(createdAt.Month()),
			Day:   int32(createdAt.Day()),
		}
		users = append(users, &user)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}

	return users
}

func UpdateUser(user *models.User) (int64, error) {

	db, err := mysql.GetMysqlConnection()
	if err != nil {
		panic(err)
	}

	q := `UPDATE users.user 
		  SET userName = ?, userLastname = ?, userNickname = ? 
		  WHERE id = ?`
	stmUpd, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			panic(err)
		}
	}(stmUpd)

	result, err := stmUpd.Exec(user.UserName, user.UserLastname, user.UserNickname, user.Id)

	err = db.Close()
	if err != nil {
		panic(err)
	}

	rows, err := result.RowsAffected()

	return rows, err
}

func DeleteUser(id int32) (int64, error) {
	db, err := mysql.GetMysqlConnection()
	if err != nil {
		panic(err)
	}

	q := `DELETE FROM users.user 
		  WHERE id = ?`
	stmDel, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	defer func(stmDel *sql.Stmt) {
		err := stmDel.Close()
		if err != nil {
			panic(err)
		}
	}(stmDel)

	result, err := stmDel.Exec(id)

	err = db.Close()
	if err != nil {
		panic(err)
	}

	rows, err := result.RowsAffected()

	return rows, err
}

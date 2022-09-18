package models

import "google.golang.org/genproto/googleapis/type/date"

type User struct {
	Id           int32
	UserName     string
	UserLastname string
	UserNickname string
	CreatedAt    date.Date
}

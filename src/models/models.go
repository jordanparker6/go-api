package models

var All = []interface{}{&User{}}

type User struct {
	Id        string `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

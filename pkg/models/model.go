package models

type User struct {
	ID 			string `json:"id" gorm:"primary_key"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	Dob 		string `json:"dob"`
}

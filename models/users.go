package models 


type Users struct{
	
	ID int64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
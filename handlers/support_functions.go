package handlers
import (
	"fmt"
	"golang_test/models"
)

func(app *MyApp) ExistingUser(uname string) (bool){
	var emp *models.Users

	app.DB.First(&emp,"username = ?",uname)
	// fmt.Println("uname: ",uname)
	// fmt.Println("emp.Username: ",emp.Username)
	
	if(uname != ""){
		if uname == emp.Username {
			fmt.Println("already exists")
			return true
		}else{
			return false
		}
	}else{
		return true
	}
}
package handlers
import (
	"github.com/gin-gonic/gin"
	"golang_test/models"
)

func (app *MyApp) ping(c *gin.Context){
	c.JSON(200,"success")
}

func (app *MyApp) get_employees(c *gin.Context){
	var employee []models.Users

	app.DB.Find(&employee)
	c.JSON(200,&employee)
}

func (app *MyApp) post_employees(c *gin.Context){
	var employee models.Users
	var ID int64
	c.Bind(&employee)

	if app.ExistingUser(employee.Username){
		c.JSON(400,gin.H{"message":"not inserted"})
	}else{
	
		app.DB.Raw(insertNewUser,employee.Username,employee.Password).Scan(&ID)

		employee.ID = ID
		c.JSON(201,&employee)
	}
	
	
}


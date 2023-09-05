package handlers

import(
	"github.com/gin-gonic/gin"
)

func(app *MyApp) Routes(r *gin.Engine){
	r.GET("/",app.ping)
	r.GET("/getemployees",app.get_employees)
	r.POST("/postemployees",app.post_employees)

}


package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seifguerbouj/bills/controllers"
	_ "github.com/seifguerbouj/bills/docs" // import swaggo docs
	"github.com/seifguerbouj/bills/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bills API
// @version 1.0
// @description This is a sample Bills API.
// @termsOfService https://example.com/terms/
// @contact.name API Support
// @contact.url https://www.example.com/support
// @contact.email seifguerbouj@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Alive"})
	})
	r.GET("/bills", controllers.FindBills)
	r.POST("/bill", controllers.CreateBill)
	r.PUT("/bill/:id", controllers.UpdateBill)
	r.DELETE("/bill/:id", controllers.DeleteBill)
	r.GET("/expenses", controllers.AllExpenses)
	r.GET("/bill/:date", controllers.GetMonthBills)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

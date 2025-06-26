package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"golang-restraunt-management/database"
	"golang-restraunt-mangement/routes"
	"golang-restraunt-mangement/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection =database.OpenCollection(database.CLient,"food")

func main(){
	port:= os.Getenv("PORT")
	if port ==""{
		port ="8000"
	}

	router:=gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrerRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":"+port)
}
package routes
import ( "github.com/gin-gonic/gin"
	controller "golang-restraunt-management/controllers")

func OrderItemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderItems",controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id",controller.GetOrderItem())
	//route to get all order items for a specific order
	incomingRoutes.GET("/orders/:orderItems-order/:order_id",controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems",controller.CreateOrderItem())
	incomingRoutes.PATCH("orderItems/:orderItem_id",controller.updateOrderItem())
}
package main
import (
  "time"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "example.com/todo/models"
"example.com/todo/controller"
"example.com/todo/config"
"github.com/joho/godotenv"
"log"
  )


func main() {
    db := config.ConnectDatabase()
    models.AutoMigrate(db)
    controller.SetDB(db)
err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
    router := gin.Default()
    
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"*"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
    todoRouter := router.Group("/v1/todo")
    {
    todoRouter.POST("/AddTodo", controller.CreateTodo)
    todoRouter.GET("/GetAllTodos", controller.GetAllTodos)
    todoRouter.GET("/ShowTodoById/:id", controller.ShowTodoById)
    todoRouter.DELETE("/DeleteTodo/:id", controller.DeleteTodo)
    todoRouter.PATCH("/UpdateTodo/:id", controller.UpdateTodo)
    }
    router.Run(":8080")
}


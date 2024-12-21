package routes
 
import (
    controllers "Flashcards/app/controllers/flashcard"
    "Flashcards/app/models"
    services "Flashcards/app/services/flashcard"
 
    "github.com/gin-gonic/gin"
)
 
func SetupFlashcardRoutes(router *gin.Engine) {
    flashcardService := &services.FlashcardService{
        Flashcards: make(map[string]models.Flashcard),
    }
    flashcardController := &controllers.FlashcardController{
        Service: flashcardService,
    }
 
    flashcardGroup := router.Group("/flashcards")
    {
        flashcardGroup.POST("/", flashcardController.CreateFlashcard)  
        flashcardGroup.POST("/:id", flashcardController.UpdateFlashcard) 
        flashcardGroup.GET("/search", flashcardController.SearchFlashcards)
        flashcardGroup.GET("/:id", flashcardController.GetFlashcardByID) 
    }
}
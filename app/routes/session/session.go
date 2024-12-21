package routes
 
import (
    controllers "Flashcards/app/controllers/session"
    sessionService "Flashcards/app/services/session"
 
    "github.com/gin-gonic/gin"
)
 
func SetupSessionRoutes(r *gin.Engine, sessionService *sessionService.SessionService) {
    sessionController := controllers.NewSessionController(sessionService)
 
    sessionGroup := r.Group("/sessions")
    {
        sessionGroup.POST("/", sessionController.Create)              
        sessionGroup.GET("/:id", sessionController.GetState)         
        sessionGroup.POST("/:id/answer", sessionController.SubmitAnswer) 
    }
}
package controllers
 
import (
    "Flashcards/app/models"
    sessionService "Flashcards/app/services/session"
    "net/http"
 
    "github.com/gin-gonic/gin"
)
 
type SessionController struct {
    Service *sessionService.SessionService 
}
 
func NewSessionController(service *sessionService.SessionService) *SessionController {
    return &SessionController{Service: service}
}
 
func (sc *SessionController) Create(ctx *gin.Context) {
    var body struct {
        StudentID string `json:"studentID"`
        Category  string `json:"category"`
    }
 
    if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
 
    session, err := sc.Service.CreateSession(body.StudentID, body.Category)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session", "details": err.Error()})
        return
    }
 
    ctx.JSON(http.StatusCreated, session)
}
 
func (sc *SessionController) GetState(ctx *gin.Context) {
    sessionID := ctx.Param("id")
 
    state, err := sc.Service.GetSessionState(sessionID)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
 
    ctx.JSON(http.StatusOK, state)
}
 
func (sc *SessionController) SubmitAnswer(ctx *gin.Context) {
    sessionID := ctx.Param("id")
 
    var body models.QuestionResponseBody
    if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
 
    state, err := sc.Service.AnswerQuestion(sessionID, body)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
 
    ctx.JSON(http.StatusOK, state)
}
 
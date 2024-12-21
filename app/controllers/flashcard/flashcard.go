package controllers
 
import (
    "Flashcards/app/models"
    services "Flashcards/app/services/flashcard"
    "net/http"
 
    "github.com/gin-gonic/gin"
)
 
type FlashcardController struct {
    Service *services.FlashcardService
}
 
func (c *FlashcardController) CreateFlashcard(ctx *gin.Context) {
    var flashcard models.Flashcard
    if err := ctx.ShouldBindJSON(&flashcard); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
        return
    }
 
    createdFlashcard, err := c.Service.CreateFlashcard(flashcard)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
 
    ctx.JSON(http.StatusOK, createdFlashcard)
}
 
func (c *FlashcardController) UpdateFlashcard(ctx *gin.Context) {
    var updatedFlashcard models.Flashcard
    if err := ctx.ShouldBindJSON(&updatedFlashcard); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
        return
    }
 
    id := ctx.Param("id")
    flashcard, err := c.Service.UpdateFlashcard(id, updatedFlashcard)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
 
    ctx.JSON(http.StatusOK, flashcard)
}
 
func (c *FlashcardController) SearchFlashcards(ctx *gin.Context) {
    tag := ctx.DefaultQuery("tag", "") // Récupère le tag depuis les paramètres de la requête
    if tag == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tag is required"})
        return
    }
 
    flashcards, err := c.Service.SearchFlashcards(tag)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
 
    ctx.JSON(http.StatusOK, flashcards)
}
func (c *FlashcardController) GetFlashcardByID(ctx *gin.Context) {
    id := ctx.Param("id")
    flashcard, err := c.Service.GetFlashcardByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
 
    ctx.JSON(http.StatusOK, flashcard)
}
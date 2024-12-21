// services/flashcardService.go
 
package services
 
import (
    "Flashcards/app/functions"
    "Flashcards/app/models"
    "Flashcards/app/server"
    "context"
    "fmt"
 
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)
 
type FlashcardService struct {
    Flashcards map[string]models.Flashcard 
}
 
func NewFlashcardService() *FlashcardService {
    return &FlashcardService{
        Flashcards: make(map[string]models.Flashcard),
    }
}
 
func (s *FlashcardService) CreateFlashcard(flashcard models.Flashcard) (models.Flashcard, error) {
    flashcard.ID = functions.NewUUID() 
   
    srv := server.GetServer()
    collection := srv.Database.Collection("flashcards") 
 
    _, err := collection.InsertOne(context.TODO(), flashcard)
    if err != nil {
        return models.Flashcard{}, fmt.Errorf("Erreur lors de l'ajout de la flashcard à la base de données: %v", err)
    }
 
    s.Flashcards[flashcard.ID] = flashcard
 
    return flashcard, nil
}
 
func (s *FlashcardService) UpdateFlashcard(id string, updatedFlashcard models.Flashcard) (models.Flashcard, error) {
    var filter bson.M
    var objectID primitive.ObjectID
 
    if objID, err := primitive.ObjectIDFromHex(id); err == nil {
        objectID = objID
        filter = bson.M{"_id": objectID} 
    } else {
        filter = bson.M{"id": id} 
    }
 
    srv := server.GetServer()
    collection := srv.Database.Collection("flashcards") 
 
    var existingFlashcard models.Flashcard
    err := collection.FindOne(context.TODO(), filter).Decode(&existingFlashcard)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return models.Flashcard{}, fmt.Errorf("Flashcard avec l'ID %s non trouvée", id)
        }
        return models.Flashcard{}, fmt.Errorf("Erreur lors de la recherche de la flashcard: %v", err)
    }
 
    update := bson.M{
        "$set": bson.M{
            "question":         updatedFlashcard.Question,
            "responses":        updatedFlashcard.Responses,
            "numRightResponse": updatedFlashcard.NumRightResponse,
            "tags":             updatedFlashcard.Tags,
        },
    }
 
    result, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return models.Flashcard{}, fmt.Errorf("Erreur lors de la mise à jour de la flashcard dans la base de données: %v", err)
    }
 
    if result.MatchedCount == 0 {
        return models.Flashcard{}, fmt.Errorf("Flashcard avec l'ID %s non trouvée ou déjà mise à jour", id)
    }
 
    var flashcard models.Flashcard
    err = collection.FindOne(context.TODO(), filter).Decode(&flashcard)
    if err != nil {
        return models.Flashcard{}, fmt.Errorf("Erreur lors de la récupération de la flashcard mise à jour: %v", err)
    }
 
    return flashcard, nil
}
 
func (s *FlashcardService) SearchFlashcards(tag string) ([]models.Flashcard, error) {
    var result []models.Flashcard
 
    srv := server.GetServer()
    collection := srv.Database.Collection("flashcards")
 
    filter := bson.M{"tags": tag}
 
    cursor, err := collection.Find(context.TODO(), filter)
    if err != nil {
        return nil, fmt.Errorf("Erreur lors de la recherche dans la base de données: %v", err)
    }
    defer cursor.Close(context.TODO())
 
    for cursor.Next(context.TODO()) {
        var flashcard models.Flashcard
        if err := cursor.Decode(&flashcard); err != nil {
            return nil, fmt.Errorf("Erreur lors du décodage des flashcards: %v", err)
        }
        result = append(result, flashcard)
    }
 
    if err := cursor.Err(); err != nil {
        return nil, fmt.Errorf("Erreur de curseur: %v", err)
    }
 
    return result, nil
}
 
func (s *FlashcardService) GetFlashcardByID(id string) (models.Flashcard, error) {
    srv := server.GetServer()
    collection := srv.Database.Collection("flashcards")
 
    var flashcard models.Flashcard
    err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&flashcard)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return models.Flashcard{}, fmt.Errorf("Flashcard with ID %s not found", id)
        }
        return models.Flashcard{}, fmt.Errorf("Erreur lors de la recherche de la flashcard: %v", err)
    }
 
    return flashcard, nil
}
 
 
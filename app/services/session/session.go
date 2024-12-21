package services
 
import (
    "Flashcards/app/functions"
    "Flashcards/app/models"
    "context"
    "errors"
 
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)
 
type SessionService struct {
    db *mongo.Database
}
 
func NewSessionService(database *mongo.Database) *SessionService {
    return &SessionService{
        db: database,
    }
}
 
func (s *SessionService) CreateSession(studentID string, category string) (*models.Session, error) {
    session := models.Session{
        ID:            functions.NewUUID(),
        StudentID:     studentID,
        SessionID:     functions.NewUUID(),
        Category:      category,
        FlashcardList: []string{}, 
        ProposalList:  []int{},
        IsFinished:    false,
        Score:         0,
    }
 
    collection := s.db.Collection("sessions")
    _, err := collection.InsertOne(context.Background(), session)
    if err != nil {
        return nil, err
    }
 
    return &session, nil
}
 
func (s *SessionService) GetSessionState(sessionID string) (*models.SessionState, error) {
    collection := s.db.Collection("sessions")
    var session models.Session
 
    filter := bson.M{"_id": sessionID}
    err := collection.FindOne(context.Background(), filter).Decode(&session)
    if err == mongo.ErrNoDocuments {
        return nil, errors.New("session not found")
    } else if err != nil {
        return nil, err
    }
 
    nextCardID := ""
    if len(session.FlashcardList) > 0 {
        nextCardID = session.FlashcardList[0]
    }
 
    return &models.SessionState{
        NextCardID: nextCardID,
        Score:      session.Score,
        IsFinished: session.IsFinished,
    }, nil
}
 
func (s *SessionService) AnswerQuestion(sessionID string, questionResponse models.QuestionResponseBody) (*models.SessionState, error) {
    collection := s.db.Collection("sessions")
 
    var session models.Session
    filter := bson.M{"_id": sessionID}
    err := collection.FindOne(context.Background(), filter).Decode(&session)
    if err == mongo.ErrNoDocuments {
        return nil, errors.New("session not found")
    } else if err != nil {
        return nil, err
    }
 
    if len(session.FlashcardList) == 0 {
        return nil, errors.New("no more flashcards in the session")
    }
 
    if questionResponse.NumeroResponse == 1 {
        session.Score++
    }
 
    session.FlashcardList = session.FlashcardList[1:]
 
    session.IsFinished = len(session.FlashcardList) == 0
 
    update := bson.M{
        "$set": bson.M{
            "flashcardList": session.FlashcardList,
            "score":         session.Score,
            "isFinished":    session.IsFinished,
        },
    }
    _, err = collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return nil, err
    }
 
    nextCardID := ""
    if len(session.FlashcardList) > 0 {
        nextCardID = session.FlashcardList[0]
    }
 
    return &models.SessionState{
        NextCardID: nextCardID,
        Score:      session.Score,
        IsFinished: session.IsFinished,
    }, nil
}
 
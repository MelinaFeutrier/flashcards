package models

type Session struct {
	ID             string   `json:"id" bson:"_id,omitempty"`
	StudentID      string   `json:"studentID" bson:"studentID"`
	SessionID      string   `json:"sessionID" bson:"sessionID"`
	Score          int      `json:"score" bson:"score"`
	Category       string   `json:"category" bson:"category"`
	FlashcardList  []string `json:"flashcardList" bson:"flashcardList"` 
	ProposalList   []int    `json:"proposalList" bson:"proposalList"`  
	IsFinished     bool     `json:"isFinished" bson:"isFinished"`
}

type SessionState struct {
	NextCardID string `json:"nextCardId"`
	Score      int    `json:"score"`
	IsFinished bool   `json:"isFinished"`
}

type QuestionResponseBody struct {
	FlashcardID   string `json:"flashcardId"`
	NumeroResponse int   `json:"numeroResponse"`
}

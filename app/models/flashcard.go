package models
 
type ResponseCard struct {
    ID       int    `json:"id" bson:"id"`
    Proposal string `json:"proposal" bson:"proposal"`
}
 
type Flashcard struct {
    ID               string         `json:"id" bson:"id"`
    Question         string         `json:"question" bson:"question"`
    Responses        []ResponseCard `json:"responses" bson:"responses"`
    NumRightResponse int            `json:"numRightResponse" bson:"numRightResponse"`
    Tags             []string       `json:"tags" bson:"tags"`
}
# Flashcards

## Project Description
The Flashcards application is a fun and interactive platform designed to help learners memorize information through the use of flashcards. Users can create, manage, and play tailored learning sessions to enhance their knowledge efficiently.

---

## Functional Rules

### Managing Flashcards
**Adding a flashcard:**
- A flashcard includes:
  - A question.
  - 4 possible answers.
  - A correct answer.
  - Tags to categorize the flashcard (e.g., math, history).

### Managing Game Sessions
**Launch a game session for a student:**
- Select a flashcard category.
- Generate a session containing 5 flashcards chosen at random from the selected category.
- The session has no time limit.

**Respond to a list of flashcards:**
- Flashcards are presented one by one.
- Each response is recorded in the session as correct or incorrect.

### Saving Data
The application records:
- The student’s score.
- The list of answers provided by the student for each session.
- The session status, including:
  - The next question.
  - Whether the session is completed or not.

---

## Models

### Flashcard
```json
{
  "answer": "string",
  "responses": [
    {
      "id": "int",
      "proposal": "string"
    }
  ],
  "numRightResponse": "int",
  "tags": ["string"]
}
```

### ResponseCard
```json
{
  "id": "int (1 to 4)",
  "proposal": "string"
}
```

### Session
```json
{
  "studentID": "string",
  "sessionID": "string",
  "score": "int",
  "category": "string",
  "flashcardList": ["Flashcard"],
  "proposalList": ["ResponseCard"],
  "isFinished": "boolean"
}
```

### SessionState
```json
{
  "nextCardId": "int",
  "score": "int",
  "isFinished": "boolean"
}
```

### QuestionResponseBody
```json
{
  "flashcardId": "int",
  "numeroResponse": "int"
}
```

---

## Resources

### Flashcards
- **`POST /flashcards`**
  - Create a new flashcard.
- **`PUT /flashcards`**
  - Update an existing flashcard.
- **`GET /flashcards/search`**
  - Search for flashcards by category or tag.
- **`GET /flashcards/:id`**
  - Retrieve a flashcard by its ID.

### Sessions
- **`POST /sessions`**
  - Create a new session for a specific student and category.
  - **Body:** `{ "studentID": "string", "category": "string" }`

- **`GET /sessions/:id/state`**
  - Retrieve the state of a session.
  - **Params:** `idSession`
  - **Response:** `{ "idProchaineCarte": "int", "score": "int", "isFinished": "boolean" }`

- **`POST /sessions/:id/answer`**
  - Submit an answer for a specific session.
  - **Params:** `idSession`
  - **Body:** `{ "idCard": "int", "numeroReponse": "int" }`
  - **Response:** `{ "idProchaineCarte": "int", "score": "int", "isFinished": "boolean" }`

---

This README outlines the foundational aspects of the Flashcards application, covering its functionality, data models, and API endpoints. It is designed to serve as a reference for developers and collaborators working on the project.


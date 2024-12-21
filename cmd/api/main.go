package main
 
import (
    "Flashcards/app/mongodb"
    flashcardRoutes "Flashcards/app/routes/flashcard"
    sessionRoutes "Flashcards/app/routes/session"
    studentRoutes "Flashcards/app/routes/student"
    "Flashcards/app/server"
    sessionService "Flashcards/app/services/session"
    "os"
 
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)
 
func main() {
    if err := newFlashcardsServer(); err != nil {
        log.Fatal().Err(err).Msg("Unable to create new server")
        os.Exit(51)
    }
    log.Debug().Msg("API launched with human readable log")
 
    srv := server.GetServer()
    srv.ListenAndServe()
}
 
func newFlashcardsServer() error {
    if os.Getenv("MODE") == "" {
        if err := godotenv.Load(); err != nil {
            log.Fatal().Err(err).Msg("Failed to load environment variables")
            return err
        }
    }
 
    srv := &server.Flashcards{}
    srv.ParseParameters()
 
    if srv.Router == nil {
        srv.Router = gin.Default()
    }
 
    switch srv.LogFormat {
    case "HUMAN":
        log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
    default:
        log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true})
    }
 
    client, err := mongodb.OpenMongoDB(srv.DBHost)
    if err != nil {
        log.Fatal().Err(err).Msg("Unable to connect to MongoDB")
        return err
    }
    srv.Database = client.Database("flashcards")
 
    sessionSvc := sessionService.NewSessionService(srv.Database) 
 
    sessionRoutes.SetupSessionRoutes(srv.Router, sessionSvc)
    flashcardRoutes.SetupFlashcardRoutes(srv.Router)
    studentRoutes.SetupRouter(srv.Router)
 
    server.SetServer(srv)
 
    return nil
}
 
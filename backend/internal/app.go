package internal

import (
	"log"

	"context"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"quiz.com/quiz/internal/collection"
	"quiz.com/quiz/internal/controller"
	"quiz.com/quiz/internal/service"
)

type App struct {
	httpServer *fiber.App
	database   *mongo.Database

	quizService *service.QuizService
	netService  *service.NetService
}

func (a App) Stack() {
	panic("unimplemented")
}

func (a *App) Init() {
	a.setupDb()
	a.setupServices()
	a.setupHttp()
	// app.Listen("127.0.0.1:8080")
	log.Fatal(a.httpServer.Listen("0.0.0.0:3000"))
}

func (a *App) setupHttp() {
	app := fiber.New()
	app.Use(cors.New())

	quizController := controller.Quiz(a.quizService)

	app.Post("/api/quiz", quizController.CreateQuiz)
app.Delete("/api/quiz/:quizId", quizController.DeleteQuizById)

	app.Get("/api/quizzes", quizController.GetQuizzes)
	app.Get("/api/quizzes/:quizId", quizController.GetQuizById)
	app.Put("/api/quizzes/:quizId", quizController.UpdateQuizById)

	wsController := controller.Ws(a.netService)
	app.Get("/ws", websocket.New(wsController.Ws))

	a.httpServer = app
}

func (a *App) setupServices() {
	a.quizService = service.Quiz(collection.Quiz(a.database.Collection("quizzes")))
	a.netService = service.Net(a.quizService)
}

func (a *App) setupDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		// options.Client().ApplyURI("mongodb://localhost:27017"))
		options.Client().ApplyURI("mongodb://0.0.0.0:27017"))
	if err != nil {
		panic(err)
	}

	a.database = client.Database("quiz")
}

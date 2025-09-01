package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quiz.com/quiz/internal/entity"
	"quiz.com/quiz/internal/service"
)



type QuizController struct {
	quizService *service.QuizService
}
type CreateQuizRequest struct {
	Name      string                `json:"name"`
	Questions []entity.QuizQuestion `json:"questions"`
}
func (c QuizController) DeleteQuizById(ctx *fiber.Ctx) error {
	quizIdStr := ctx.Params("quizId")
	quizId, err := primitive.ObjectIDFromHex(quizIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid quiz ID")
	}

	err = c.quizService.DeleteQuizById(quizId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to delete quiz")
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c QuizController) CreateQuiz(ctx *fiber.Ctx) error {
	var req CreateQuizRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	
	err := c.quizService.CreateQuiz(req.Name, req.Questions)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create quiz",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
    "message": "Quiz created successfully",
	"id": req.Name, // Assuming the ID is the name for simplicity, adjust as needed;
})

}

func Quiz(quizService *service.QuizService) QuizController {
	return QuizController{
		quizService: quizService,
	}
}

func (c QuizController) GetQuizById(ctx *fiber.Ctx) error {
	quizIdStr := ctx.Params("quizId")
	quizId, err := primitive.ObjectIDFromHex(quizIdStr)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	quiz, err := c.quizService.GetQuizById(quizId)
	if err != nil {
		return err
	}

	if quiz == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(quiz)
}

type UpdateQuizRequest struct {
	Name      string                `json:"name"`
	Questions []entity.QuizQuestion `json:"questions"`
}

func (c QuizController) UpdateQuizById(ctx *fiber.Ctx) error {

	quizIdStr := ctx.Params("quizId")
	quizId, err := primitive.ObjectIDFromHex(quizIdStr)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	var req UpdateQuizRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if err := c.quizService.UpdateQuiz(quizId, req.Name, req.Questions); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := c.quizService.GetQuizzes()
	if err != nil {
		return err
	}

	return ctx.JSON(quizzes)
}

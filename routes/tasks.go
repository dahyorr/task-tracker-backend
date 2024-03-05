package routes

import (
	"strconv"

	"github.com/dahyorr/task-tracker-backend/models"
	"github.com/dahyorr/task-tracker-backend/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func getTask(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return shared.ErrFailedToParseId
	}
	task, err := models.GetTaskById(id)
	if err != nil {
		return shared.ErrUserNotFound
	}
	return c.JSON(task)
}

func createTask(c *fiber.Ctx) error {
	taskForm := models.TaskForm{}
	c.BodyParser(&taskForm)
	task := models.Task{
		Name:        taskForm.Name,
		Description: taskForm.Description,
		Status:      taskForm.Status,
		WorkspaceId: taskForm.WorkspaceId,
		DueDate:     taskForm.DueDate,
	}
	err := task.Create()
	if err != nil {
		return shared.ErrSomethingWentWrong
	}
	return c.JSON(task)
}

func deleteTask(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return shared.ErrFailedToParseId
	}
	task := models.Task{
		Id: id,
	}
	err = task.Delete()
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func updateStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	var formStatus struct {
		Status string
	}
	c.BodyParser(&formStatus)

	if err != nil {
		return shared.ErrFailedToParseId
	}
	task, err := models.GetTaskById(id)
	if err != nil {
		log.Error(err)
		return shared.ErrTaskNotFound
	}
	err = task.UpdateStatus(formStatus.Status)
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.JSON(task)
}

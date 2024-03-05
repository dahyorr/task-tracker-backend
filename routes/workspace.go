package routes

import (
	"strconv"

	"github.com/dahyorr/task-tracker-backend/models"
	"github.com/dahyorr/task-tracker-backend/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func getWorkspaceDetails(c *fiber.Ctx) error {
	workspaceId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return shared.ErrInvalidWorkspaceId
	}
	workspace, err := models.GetWorkspaceById(workspaceId)
	if err != nil {
		return shared.ErrWorkspaceNotFound
	}
	return c.JSON(workspace)
}

func getWorkspaces(c *fiber.Ctx) error {
	userId := c.Locals("session").(*models.Session).UserId
	workspaces, err := models.GetWorkspacesByUserId(userId)
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.JSON(workspaces)
}

func createWorkspace(c *fiber.Ctx) error {
	userId := c.Locals("session").(*models.Session).UserId
	var workspace models.Workspace
	var data models.WorkspaceFormData
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	workspace.OwnerId = userId
	workspace.Name = data.Name
	err := workspace.Create()
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.JSON(workspace)
}

func deleteWorkspace(c *fiber.Ctx) error {
	workspaceId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return shared.ErrInvalidWorkspaceId
	}
	workspace := models.Workspace{
		Id: workspaceId,
	}
	err = workspace.Delete()
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func addUserToWorkspace(c *fiber.Ctx) error {
	workspaceId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return shared.ErrInvalidWorkspaceId
	}
	userId, err := strconv.ParseInt(c.Params("userId"), 10, 64)
	if err != nil {
		return shared.ErrUserNotFound
	}
	workspace, err := models.GetWorkspaceById(workspaceId)
	if err != nil {
		log.Error(err)
		return shared.ErrWorkspaceNotFound
	}
	err = workspace.AddUser(userId)
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func removeUserFromWorkspace(c *fiber.Ctx) error {
	workspaceId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return shared.ErrInvalidWorkspaceId
	}
	userId, err := strconv.ParseInt(c.Params("userId"), 10, 64)
	if err != nil {
		return shared.ErrUserNotFound
	}
	workspace, err := models.GetWorkspaceById(workspaceId)
	if err != nil {
		log.Error(err)
		return shared.ErrWorkspaceNotFound
	}
	err = workspace.RemoveUser(userId)
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func getWorkspaceMembers(c *fiber.Ctx) error {
	workspaceId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return shared.ErrInvalidWorkspaceId
	}
	users, err := models.GetWorkspaceMembers(workspaceId)
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.JSON(users)
}


func getWorkspaceTasks(c *fiber.Ctx) error {
	workspaceId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return shared.ErrInvalidWorkspaceId
	}
	tasks, err := models.GetTasksByWorkspaceId(workspaceId)
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.JSON(tasks)
}
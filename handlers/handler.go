package handlers

import (
	"GoRsyncManager/configs"
	"GoRsyncManager/models"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// @Summary Added a new node.
// @Description Added a new node.
// @Accept json
// @Produce json
// @Param node body models.Node true "Node Details"
// @Success 200 {string} status "ok"
// @Router /add_node [post]
func AddNode(c *fiber.Ctx) error {
	var node models.Node
	err := c.BodyParser(&node)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": err.Error()})
		slog.Error(err.Error())
		return err
	}
	result := configs.DB.Create(&node)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": result.Error})
		slog.Error(result.Error.Error())
		return result.Error
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "node has been added"})
	return nil
}

// @Summary Delete a node.
// @Description Delete a node.
// @Accept json
// @Produce json
// @Param id path integer true "Node Id"
// @Success 200 {string} status "ok"
// @Router /del_node/{id} [delete]
func DelNode(c *fiber.Ctx) error {
	id := c.Params("id")
	var node models.Node
	result := configs.DB.Delete(&node, id)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": result.Error})
		slog.Error(result.Error.Error())
		return result.Error
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "node delete successfully"})
	return nil
}

// @Summary Get list of node.
// @Description Get list of node.
// @Accept json
// @Produce json
// @Success 200 {array} models.Node
// @Router /get_nodes [get]
func GetNodes(c *fiber.Ctx) error {
	var nodes []models.Node
	result := configs.DB.Find(&nodes)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": result.Error})
		slog.Error(result.Error.Error())
		return result.Error
	}
	c.Status(http.StatusOK).JSON(&nodes)
	return nil
}

// @Summary Get a node details by id.
// @Description Get a node details by id.
// @Accept json
// @Produce json
// @Param id path integer true "Node Id"
// @Success 200 {object} models.Node
// @Router /get_node/{id} [get]
func GetNodeById(c *fiber.Ctx) error {
	id := c.Params("id")
	var node models.Node
	result := configs.DB.Find(&node, "id=?", id)
	if result.Error != nil {
		c.Status(http.StatusNotFound).JSON(&fiber.Map{"message": result.Error})
		slog.Error(result.Error.Error())
		return result.Error
	}
	c.Status(http.StatusOK).JSON(&node)
	return nil
}

package handlers

import (
	"GoRsyncManager/configs"
	"GoRsyncManager/models"
	"GoRsyncManager/utility"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// @Summary Added a new host.
// @Description Added a new host.
// @Accept json
// @Produce json
// @Param host body models.Host true "Host Details"
// @Success 200 {string} status "ok"
// @Router /add_host [post]
func AddHost(c *fiber.Ctx) error {
	var host models.Host
	err := c.BodyParser(&host)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": err.Error()})
		slog.Error(err.Error())
		return err
	}
	result := configs.DB.Create(&host)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": result.Error})
		slog.Error(result.Error.Error())
		return result.Error
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "node has been added"})
	return nil
}

// @Summary Delete a host.
// @Description Delete a host.
// @Accept json
// @Produce json
// @Param id path integer true "Host Id"
// @Success 200 {string} status "ok"
// @Router /del_host/{id} [delete]
func DelHost(c *fiber.Ctx) error {
	id := c.Params("id")
	var host models.Host
	result := configs.DB.Delete(&host, id)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": result.Error})
		slog.Error(result.Error.Error())
		return result.Error
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "node delete successfully"})
	return nil
}

// @Summary Get list of host.
// @Description Get list of host.
// @Accept json
// @Produce json
// @Success 200 {array} models.Host
// @Router /get_hosts [get]
func GetHosts(c *fiber.Ctx) error {
	var hosts []models.Host
	result := configs.DB.Omit("key", "password").Find(&hosts)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": result.Error})
		slog.Error(result.Error.Error())
		return result.Error
	}
	c.Status(http.StatusOK).JSON(&hosts)
	return nil
}

// @Summary Get file list from a host
// @Description Get file list from a host
// @Accept json
// @Produce json
// @Param index body models.Index true "Filter Data"
// @Success 200 {array} models.Content
// @Router /get_files [post]
func GetFilesFromHost(c *fiber.Ctx) error {
	var index models.Index
	var content models.Content

	err := c.BodyParser(&index)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": err.Error()})
		slog.Error(err.Error())
		return err
	}

	cmdStr := utility.GetFileFilterCmd(&index)
	err = utility.GetRemoteFileList(index.Id, cmdStr, &content)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": err.Error()})
		slog.Error(err.Error())
		return err
	}
	c.Status(http.StatusOK).JSON(&content)
	return nil
}

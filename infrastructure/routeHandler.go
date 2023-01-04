package infrastructure

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/application"
	"github.com/shuxbot/shux-api/domain"
)

type routeHandler struct {
	userApp *application.UserApp
	channelApp *application.ChannelApp
}

func bodyToUserStruct(c *fiber.Ctx) domain.User {
	var u domain.User
	json.Unmarshal(c.Body(), &u)
	u.UserId = c.Params("user_id")

	return u
}

func result(success bool, err error) map[string]interface{} {
	status := make(map[string]interface{})
	status["error"] = err
	status["success"] = success

	return status
}

func (h *routeHandler) GetUser(c *fiber.Ctx) error {
	u, err := h.userApp.Get(c.Params("user_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err))
	}
	return c.JSON(u)
}

func (h *routeHandler) DeleteUser(c *fiber.Ctx) error {
	err := h.userApp.Delete(c.Params("user_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err))
	}
	return c.JSON(result(true, nil))
}

func (h *routeHandler) UpdateUser(c *fiber.Ctx) error {
	u := bodyToUserStruct(c)
	err := h.userApp.Update(&u, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err))
	}
	return c.JSON(result(true, nil))

}

func (h *routeHandler) CreateUser(c *fiber.Ctx) error {
	u := bodyToUserStruct(c)
	err := h.userApp.Create(&u, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err))
	}
	return c.JSON(result(true, nil))
}


func (h *routeHandler) ListChannels(c *fiber.Ctx) error {
	cMap := make(map[string]interface{})
	cArr, err := h.channelApp.List(c.Params("server_id"))
	cMap["channels"] = cArr

	if err != nil {
		return c.Status(404).JSON(result(false, err))
	}

	return c.JSON(cMap)
}

func NewRouteHandler(userApp *application.UserApp, channelApp *application.ChannelApp) *routeHandler {
	return &routeHandler{userApp: userApp, channelApp: channelApp}
}

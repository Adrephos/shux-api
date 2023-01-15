package infrastructure

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/application"
	"github.com/shuxbot/shux-api/domain"
)

type routeHandler struct {
	userApp    *application.UserApp
	channelApp *application.ChannelApp
	roleApp    *application.RoleApp
	serverApp  *application.ServerApp
}

func bodyToUserStruct(c *fiber.Ctx) domain.User {
	var u domain.User
	json.Unmarshal(c.Body(), &u)
	u.UserId = c.Params("user_id")

	return u
}

func bodyToChannelStruct(c *fiber.Ctx) domain.Channel {
	var ch domain.Channel
	json.Unmarshal(c.Body(), &ch)
	ch.ChannelId = c.Params("channel_id")

	return ch
}

func bodyToRoleStruct(c *fiber.Ctx) domain.Role {
	var rl domain.Role
	json.Unmarshal(c.Body(), &rl)
	rl.RoleId = c.Params("role_id")

	return rl
}

func result(success bool, err error, data interface{}) map[string]interface{} {
	status := make(map[string]interface{})
	if err != nil {
		status["error"] = err.Error()
	} else {
		status["data"] = data
	}
	status["success"] = success

	return status
}

func (h *routeHandler) GetUser(c *fiber.Ctx) error {
	u, err := h.userApp.Get(c.Params("user_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, err, u))
}

func (h *routeHandler) DeleteUser(c *fiber.Ctx) error {
	err := h.userApp.Delete(c.Params("user_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, c.Params("user_id")))
}

func (h *routeHandler) UpdateUser(c *fiber.Ctx) error {
	u := bodyToUserStruct(c)
	err := h.userApp.Update(&u, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, u))

}

func (h *routeHandler) ReplaceUser(c *fiber.Ctx) error {
	u := bodyToUserStruct(c)
	err := h.userApp.Replace(&u, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, u))
}

func (h *routeHandler) CreateUser(c *fiber.Ctx) error {
	u := bodyToUserStruct(c)
	err := h.userApp.Create(&u, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, u))
}

func (h *routeHandler) ListChannels(c *fiber.Ctx) error {
	cMap := make(map[string]interface{})
	cArr, err := h.channelApp.List(c.Params("server_id"))
	cMap["channels"] = cArr

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}

	return c.JSON(result(true, nil, cMap))
}

func (h *routeHandler) GetChannel(c *fiber.Ctx) error {
	ch, err := h.channelApp.Get(c.Params("channel_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ch))
}

func (h *routeHandler) DeleteChannel(c *fiber.Ctx) error {
	err := h.channelApp.Delete(c.Params("channel_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, c.Params("channel_id")))
}

func (h *routeHandler) UpdateChannel(c *fiber.Ctx) error {
	ch := bodyToChannelStruct(c)
	err := h.channelApp.Update(&ch, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ch))

}

func (h *routeHandler) ReplaceChannel(c *fiber.Ctx) error {
	ch := bodyToChannelStruct(c)
	err := h.channelApp.Replace(&ch, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ch))
}

func (h *routeHandler) CreateChannel(c *fiber.Ctx) error {
	ch := bodyToChannelStruct(c)
	err := h.channelApp.Create(&ch, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ch))
}

func (h *routeHandler) ListRoles(c *fiber.Ctx) error {
	rlMap := make(map[string]interface{})
	rlArr, err := h.roleApp.List(c.Params("server_id"))
	rlMap["roles"] = rlArr

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}

	return c.JSON(result(true, nil, rlMap))
}

func (h *routeHandler) GetRole(c *fiber.Ctx) error {
	rl, err := h.roleApp.Get(c.Params("role_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, rl))
}

func (h *routeHandler) DeleteRole(c *fiber.Ctx) error {
	err := h.roleApp.Delete(c.Params("role_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, c.Params("role_id")))
}

func (h *routeHandler) UpdateRole(c *fiber.Ctx) error {
	rl := bodyToRoleStruct(c)
	err := h.roleApp.Update(&rl, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, rl))

}

func (h *routeHandler) ReplaceRole(c *fiber.Ctx) error {
	rl := bodyToRoleStruct(c)
	err := h.roleApp.Replace(&rl, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, rl))
}

func (h *routeHandler) CreateRole(c *fiber.Ctx) error {
	rl := bodyToRoleStruct(c)
	err := h.roleApp.Create(&rl, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, rl))
}

func (h *routeHandler) ListServers(c *fiber.Ctx) error {
	mapId := make(map[string]interface{})
	idArr,err := h.serverApp.List()
	mapId["servers_id"] = idArr

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, mapId))
}

func (h *routeHandler) ServerRanking(c *fiber.Ctx) error{
	ranking := make(map[string]interface{})
	serverRanking, err := h.serverApp.GetRanking(c.Params("server_id"))
	ranking["ranking"] = serverRanking

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ranking))

}

func NewRouteHandler(userApp *application.UserApp, channelApp *application.ChannelApp, roleApp *application.RoleApp, serverApp *application.ServerApp) *routeHandler {
	return &routeHandler{userApp: userApp, channelApp: channelApp, roleApp: roleApp, serverApp: serverApp}
}

package routing

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/application"
	"github.com/shuxbot/shux-api/domain"
)

type RouteHandler struct {
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

// User endpoints
func (h *RouteHandler) GetUser(c *fiber.Ctx) error {
	u, err := h.userApp.Get(c.Params("user_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, err, u))
}

func (h *RouteHandler) DeleteUser(c *fiber.Ctx) error {
	err := h.userApp.Delete(c.Params("user_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, c.Params("user_id")))
}

func (h *RouteHandler) UpdateUser(c *fiber.Ctx) error {
	u := bodyToUserStruct(c)
	err := h.userApp.Update(&u, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, u))

}

func (h *RouteHandler) ReplaceUser(c *fiber.Ctx) error {
	u := bodyToUserStruct(c)
	err := h.userApp.Replace(&u, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, u))
}

func (h *RouteHandler) CreateUser(c *fiber.Ctx) error {
	u := bodyToUserStruct(c)
	err := h.userApp.Create(&u, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, u))
}

// Channel endpoints
func (h *RouteHandler) ListChannels(c *fiber.Ctx) error {
	cMap := make(map[string]interface{})
	cArr, err := h.channelApp.List(c.Params("server_id"))
	cMap["channels"] = cArr

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}

	return c.JSON(result(true, nil, cMap))
}

func (h *RouteHandler) GetChannel(c *fiber.Ctx) error {
	ch, err := h.channelApp.Get(c.Params("channel_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ch))
}

func (h *RouteHandler) DeleteChannel(c *fiber.Ctx) error {
	err := h.channelApp.Delete(c.Params("channel_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, c.Params("channel_id")))
}

func (h *RouteHandler) UpdateChannel(c *fiber.Ctx) error {
	ch := bodyToChannelStruct(c)
	err := h.channelApp.Update(&ch, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ch))

}

func (h *RouteHandler) ReplaceChannel(c *fiber.Ctx) error {
	ch := bodyToChannelStruct(c)
	err := h.channelApp.Replace(&ch, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ch))
}

func (h *RouteHandler) CreateChannel(c *fiber.Ctx) error {
	ch := bodyToChannelStruct(c)
	err := h.channelApp.Create(&ch, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ch))
}

// Role endpoints
func (h *RouteHandler) ListRoles(c *fiber.Ctx) error {
	roleMap := make(map[string]interface{})
	roleArr, err := h.roleApp.List(c.Params("server_id"))
	roleMap["roles"] = roleArr

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}

	return c.JSON(result(true, nil, roleMap))
}

func (h *RouteHandler) GetRole(c *fiber.Ctx) error {
	rl, err := h.roleApp.Get(c.Params("role_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, rl))
}

func (h *RouteHandler) DeleteRole(c *fiber.Ctx) error {
	err := h.roleApp.Delete(c.Params("role_id"), c.Params("server_id"))
	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, c.Params("role_id")))
}

func (h *RouteHandler) UpdateRole(c *fiber.Ctx) error {
	rl := bodyToRoleStruct(c)
	err := h.roleApp.Update(&rl, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, rl))

}

func (h *RouteHandler) ReplaceRole(c *fiber.Ctx) error {
	rl := bodyToRoleStruct(c)
	err := h.roleApp.Replace(&rl, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, rl))
}

func (h *RouteHandler) CreateRole(c *fiber.Ctx) error {
	rl := bodyToRoleStruct(c)
	err := h.roleApp.Create(&rl, c.Params("server_id"))

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, rl))
}

// Server endpoints
func (h *RouteHandler) ListServers(c *fiber.Ctx) error {
	mapId := make(map[string]interface{})
	idArr,err := h.serverApp.List()
	mapId["servers_id"] = idArr

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, mapId))
}

func (h *RouteHandler) ServerRanking(c *fiber.Ctx) error{
	ranking := make(map[string]interface{})
	serverRanking, err := h.serverApp.GetLeaderboard(c.Params("server_id"))
	ranking["ranking"] = serverRanking

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ranking))

}

func (h *RouteHandler) ServerUserRanking(c *fiber.Ctx) error{
	ranking := make(map[string]interface{})
	userRank, err := h.serverApp.GetUserRank(c.Params("server_id"), c.Params("user_id"))
	ranking["user"] = userRank

	if err != nil {
		return c.Status(404).JSON(result(false, err, nil))
	}
	return c.JSON(result(true, nil, ranking))
}

func NewRouteHandler(userApp *application.UserApp, channelApp *application.ChannelApp, roleApp *application.RoleApp, serverApp *application.ServerApp) *RouteHandler {
	return &RouteHandler{userApp: userApp, channelApp: channelApp, roleApp: roleApp, serverApp: serverApp}
}

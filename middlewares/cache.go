// Based on https://github.com/codemicro/fiber-cache
package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	gc "github.com/patrickmn/go-cache"
)

type CacheEntry struct {
	Body        []byte
	StatusCode  int
	ContentType []byte
}

var cache *gc.Cache

func init() {
	cache = gc.New(30 * time.Minute, time.Minute)
}

func CacheAdd(ttl time.Duration) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		key := utils.CopyString(c.Path())
		key = strings.TrimRight(key, "/")
		if c.Method() != fiber.MethodGet {
			cache.Delete(key)
			cache.Delete(key + "/rank")
			path := strings.Split(key, "/")
			key = strings.Join(path[:len(path)-1], "/")
			cache.Delete(key)
			return c.Next()
		}
		val, found := cache.Get(key)

		if found {
			entry := val.(CacheEntry)
			c.Response().SetBody(entry.Body)
			c.Response().SetStatusCode(entry.StatusCode)
			c.Response().Header.SetContentTypeBytes(entry.ContentType)
			return nil
		}

		err := c.Next()

		if err == nil {
			cache.Set(key, CacheEntry{
				Body:        c.Response().Body(),
				StatusCode:  c.Response().StatusCode(),
				ContentType: c.Response().Header.ContentType(),
			}, ttl)
		}

		return err
	}
}

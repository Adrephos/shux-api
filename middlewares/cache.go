// Based on https://github.com/codemicro/fiber-cache
package middlewares

import (
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
	cache = gc.New(time.Hour, time.Hour)
}

func CacheAdd(ttl time.Duration) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		key := utils.CopyString(c.Path())
		val, found := cache.Get(key)

		if c.Method() != fiber.MethodGet {
			cache.Delete(key)
			err := c.Next()
			return err
		}

		if found {
			entry := val.(CacheEntry)
			c.Response().SetBody(entry.Body)
			c.Response().SetStatusCode(entry.StatusCode)
			c.Response().Header.SetContentTypeBytes(entry.ContentType)
			return nil
		}
		c.Locals("cacheKey", key)

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

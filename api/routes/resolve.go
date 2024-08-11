package routes

import (
	"github.com/roh4nyh/shorten_url/database"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()

	// checking if data present in the redis database...
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in the database",
		})

		// if data not found in redis db, might be error while connecting db...
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "can't connect to DB",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	// return c.Status(301).Redirect().To(value)

	// redirect to original URL
	return c.Redirect(value, 301)
}

package utils

import (
	"empty/internal/config"
	"empty/internal/model"

	"github.com/gin-gonic/gin"
)

func GinResponse(c *gin.Context, data *model.Response) {

	switch data.Status {
	case 0:
		c.JSON(200, data.Data)
		return

	case 200:
		c.JSON(200, data.Data)
		return

	case 201:
		c.JSON(201, data.Data)
		return

	case 400:
		config.Log.Error(data.Error)
		c.JSON(400, model.InvalidInput)
		return

	case 401:
		config.Log.Error(data.Error)
		c.JSON(401, model.Unauthorized)
		return

	case 402:
		config.Log.Error(data.Error)
		c.JSON(402, model.PaymentRequired)
		return

	case 403:
		config.Log.Error(data.Error)
		c.JSON(403, model.Forbidden)
		return

	case 404:
		config.Log.Error(data.Error)
		c.JSON(404, model.NotFound)
		return

	case 409:
		config.Log.Error(data.Error)
		c.JSON(409, model.Conflict)
		return

	default:
		config.Log.Error(data.Error)
		c.JSON(500, model.InternalServerError)
		return
	}

}

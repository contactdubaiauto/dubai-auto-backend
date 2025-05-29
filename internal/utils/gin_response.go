package utils

import (
	"dubai-auto/internal/config"
	"dubai-auto/internal/model"

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
		c.JSON(400, model.ResultMessage{
			Tk: model.InvalidInput.Message.Tk,
			Ru: model.InvalidInput.Message.Ru,
			En: data.Error.Error(),
		})
		return

	case 401:
		config.Log.Error(data.Error)
		c.JSON(401, model.ResultMessage{
			Tk: model.Unauthorized.Message.Tk,
			Ru: model.Unauthorized.Message.Ru,
			En: data.Error.Error(),
		})
		return

	case 402:
		config.Log.Error(data.Error)
		c.JSON(402, model.ResultMessage{
			Tk: model.PaymentRequired.Message.Tk,
			Ru: model.PaymentRequired.Message.Ru,
			En: data.Error.Error(),
		})
		return

	case 403:
		config.Log.Error(data.Error)
		c.JSON(403, model.ResultMessage{
			Tk: model.Forbidden.Message.Tk,
			Ru: model.Forbidden.Message.Ru,
			En: data.Error.Error(),
		})
		return

	case 404:
		config.Log.Error(data.Error)
		c.JSON(404, model.ResultMessage{
			Tk: model.NotFound.Message.Tk,
			Ru: model.NotFound.Message.Ru,
			En: data.Error.Error(),
		})
		return

	case 409:
		config.Log.Error(data.Error)
		c.JSON(409, model.ResultMessage{
			Tk: model.Conflict.Message.Tk,
			Ru: model.Conflict.Message.Ru,
			En: data.Error.Error(),
		})
		return

	default:
		config.Log.Error(data.Error)
		c.JSON(500, model.ServerError{
			Message: model.ResultMessage{
				Tk: model.InternalServerError.Message.Tk,
				Ru: model.InternalServerError.Message.Ru,
				En: data.Error.Error(),
			},
		})
		return
	}

}

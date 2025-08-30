package utils

import (
	"github.com/gofiber/fiber/v2"

	"dubai-auto/internal/model"
	"dubai-auto/pkg/logger"
)

func FiberResponse(c *fiber.Ctx, data *model.Response) error {

	switch data.Status {
	case 0:
		c.Status(200).JSON(data.Data)
		return nil

	case 200:
		c.Status(200).JSON(data.Data)
		return nil

	case 201:
		c.Status(201).JSON(data.Data)
		return nil

	case 400:
		logger.Log.Err(data.Error)
		c.Status(400).JSON(model.ResultMessage{
			Tk: model.InvalidInput.Message.Tk,
			Ru: model.InvalidInput.Message.Ru,
			En: data.Error.Error(),
		})
		return data.Error

	case 401:
		logger.Log.Err(data.Error)
		c.Status(401).JSON(model.ResultMessage{
			Tk: model.Unauthorized.Message.Tk,
			Ru: model.Unauthorized.Message.Ru,
			En: data.Error.Error(),
		})
		return data.Error

	case 402:
		logger.Log.Err(data.Error)
		c.Status(402).JSON(model.ResultMessage{
			Tk: model.PaymentRequired.Message.Tk,
			Ru: model.PaymentRequired.Message.Ru,
			En: data.Error.Error(),
		})
		return data.Error

	case 403:
		logger.Log.Err(data.Error)
		c.Status(403).JSON(model.ResultMessage{
			Tk: model.Forbidden.Message.Tk,
			Ru: model.Forbidden.Message.Ru,
			En: data.Error.Error(),
		})
		return data.Error

	case 404:
		logger.Log.Err(data.Error)
		c.Status(404).JSON(model.ResultMessage{
			Tk: model.NotFound.Message.Tk,
			Ru: model.NotFound.Message.Ru,
			En: data.Error.Error(),
		})
		return data.Error

	case 409:
		logger.Log.Err(data.Error)
		c.Status(409).JSON(model.ResultMessage{
			Tk: model.Conflict.Message.Tk,
			Ru: model.Conflict.Message.Ru,
			En: data.Error.Error(),
		})
		return data.Error

	default:
		logger.Log.Err(data.Error)
		c.Status(500).JSON(model.ServerError{
			Message: model.ResultMessage{
				Tk: model.InternalServerError.Message.Tk,
				Ru: model.InternalServerError.Message.Ru,
				En: data.Error.Error(),
			},
		})
		return data.Error
	}

}

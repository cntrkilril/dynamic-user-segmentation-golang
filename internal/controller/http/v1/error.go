package v1

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
)

var errCodeMap = map[entity.ErrCode]int{
	entity.ErrCodeBadRequest: fiber.StatusBadRequest,
	entity.ErrCodeInternal:   fiber.StatusUnauthorized,
	entity.ErrCodeNotFound:   fiber.StatusNotFound,
}

func HandleErrors(ctx *fiber.Ctx, errorArray []error) error {
	appErr := &entity.Error{}
	if errors.As(errorArray[0], &appErr) {
		c, ok := errCodeMap[appErr.Code()]
		if !ok {
			c = fiber.StatusInternalServerError
		}

		return ctx.Status(c).JSON(newResp(nil, errorArray))
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(newResp(nil, errorArray))
}

func HandleRespWithErrors(ctx *fiber.Ctx, data any, errorArray []error) error {
	appErr := &entity.Error{}
	if errors.As(errorArray[0], &appErr) {
		c, ok := errCodeMap[appErr.Code()]
		if !ok {
			c = fiber.StatusInternalServerError
		}

		return ctx.Status(c).JSON(newResp(nil, errorArray))
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(newResp(data, errorArray))
}

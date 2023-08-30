package v1

import (
	"github.com/gofiber/fiber/v2"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/controller"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"github/cntrkilril/dynamic-user-segmentation-golang/pkg/govalidator"
)

type UsersSegmentsHistoryHandler struct {
	val                         *govalidator.Validator
	usersSegmentsHistoryService controller.UsersSegmentsHistoryService
}

func (h *UsersSegmentsHistoryHandler) getHistoryByUserIDAndYearMonth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.GetCSVHistoryByUserIDAndYearMonthDTO
		if err := h.val.ValidateQueryParams(c, &p); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		res, err := h.usersSegmentsHistoryService.GetCSVHistoryByUserID(c.Context(), p)
		if err != nil {
			return HandleErrors(c, []error{err})
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"url": res.Url,
			},
			[]error{},
		))
	}
}

func (h *UsersSegmentsHistoryHandler) Register(r fiber.Router) {
	r.Get("get-by-user-id-and-year-month",
		h.getHistoryByUserIDAndYearMonth())
}

func NewUsersSegmentsHistoryHandler(
	usersSegmentsHistoryService controller.UsersSegmentsHistoryService,
	val *govalidator.Validator,
) *UsersSegmentsHistoryHandler {
	return &UsersSegmentsHistoryHandler{
		usersSegmentsHistoryService: usersSegmentsHistoryService,
		val:                         val,
	}
}

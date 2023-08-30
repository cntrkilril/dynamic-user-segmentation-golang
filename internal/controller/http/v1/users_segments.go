package v1

import (
	"github.com/gofiber/fiber/v2"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/controller"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"github/cntrkilril/dynamic-user-segmentation-golang/pkg/govalidator"
)

type UsersSegmentsHandler struct {
	val                  *govalidator.Validator
	usersSegmentsService controller.UsersSegmentsService
}

func (h *UsersSegmentsHandler) create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.SegmentsByUserIDDTO
		if err := h.val.ValidateRequestBody(c, &p); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		res, errArray := h.usersSegmentsService.Create(c.Context(), p)
		if len(errArray) != 0 {
			if res.UserID == 0 || len(res.Segments) == 0 {
				return HandleRespWithErrors(c, nil, errArray)
			}
			return HandleRespWithErrors(c, res, errArray)
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"userID":   res.UserID,
				"segments": res.Segments,
			},
			[]error{},
		))
	}
}

func (h *UsersSegmentsHandler) getSegmentsByUserID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.GetSegmentsByUserIDDTO
		if err := h.val.ValidateParams(c, &p); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		res, err := h.usersSegmentsService.GetSegmentsByUserID(c.Context(), p)
		if err != nil {
			return HandleErrors(c, []error{err})
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"userID":   res.UserID,
				"segments": res.Segments,
			},
			[]error{},
		))
	}
}

func (h *UsersSegmentsHandler) delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.SegmentsByUserIDDTO
		if err := h.val.ValidateRequestBody(c, &p); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		res, errArray := h.usersSegmentsService.Delete(c.Context(), p)
		if len(errArray) != 0 {
			if res.UserID == 0 || len(res.Segments) == 0 {
				return HandleRespWithErrors(c, nil, errArray)
			}
			return HandleRespWithErrors(c, res, errArray)
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			nil,
			[]error{},
		))
	}
}

func (h *UsersSegmentsHandler) Register(r fiber.Router) {
	r.Post("add-segments-to-user",
		h.create())
	r.Delete("delete-segments-to-user",
		h.delete())
	r.Get("get-segments-by-user-id/:userID",
		h.getSegmentsByUserID())
}

func NewUsersSegmentsHandler(
	usersSegmentsService controller.UsersSegmentsService,
	val *govalidator.Validator,
) *UsersSegmentsHandler {
	return &UsersSegmentsHandler{
		usersSegmentsService: usersSegmentsService,
		val:                  val,
	}
}

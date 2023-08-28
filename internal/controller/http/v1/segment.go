package v1

import (
	"github.com/gofiber/fiber/v2"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/controller"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"github/cntrkilril/dynamic-user-segmentation-golang/pkg/govalidator"
)

type SegmentHandler struct {
	val            *govalidator.Validator
	segmentService controller.SegmentService
}

func (h *SegmentHandler) create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.Segment
		if err := h.val.ValidateRequestBody(c, &p); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		res, err := h.segmentService.Create(c.Context(), p)
		if err != nil {
			return HandleErrors(c, []error{err})
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"segment": res.Slug,
			},
			[]error{},
		))
	}
}

func (h *SegmentHandler) delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.Segment
		if err := h.val.ValidateRequestBody(c, &p); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err := h.segmentService.Delete(c.Context(), p)
		if err != nil {
			return HandleErrors(c, []error{err})
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{},
			[]error{},
		))
	}
}

func (h *SegmentHandler) Register(r fiber.Router) {
	r.Post("create",
		h.create())
	r.Post("delete",
		h.delete())
}

func NewSegmentHandler(
	segmentService controller.SegmentService,
	val *govalidator.Validator,
) *SegmentHandler {
	return &SegmentHandler{
		segmentService: segmentService,
		val:            val,
	}
}

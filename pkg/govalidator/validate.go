package govalidator

import (
	"context"
	gov "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator struct {
	val *gov.Validate
}

func (v *Validator) Validate(ctx context.Context, s any) error {
	return v.val.StructCtx(ctx, s)
}

func (v *Validator) ValidateRequestBody(ctx *fiber.Ctx, s interface{}) error {
	if err := ctx.BodyParser(s); err != nil {
		return err
	}

	return v.val.StructCtx(ctx.Context(), s)
}

func (v *Validator) ValidateQueryParams(ctx *fiber.Ctx, s interface{}) error {
	if err := ctx.QueryParser(s); err != nil {
		return err
	}

	return v.val.StructCtx(ctx.Context(), s)
}

func (v *Validator) ValidateParams(ctx *fiber.Ctx, s interface{}) error {
	if err := ctx.ParamsParser(s); err != nil {
		return err
	}

	return v.val.StructCtx(ctx.Context(), s)
}

func New() *Validator {
	return &Validator{gov.New()}
}

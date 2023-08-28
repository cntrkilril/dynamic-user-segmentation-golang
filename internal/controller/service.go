package controller

import (
	"context"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
)

type (
	SegmentService interface {
		Create(ctx context.Context, dto entity.Segment) (result entity.Segment, err error)
		Delete(ctx context.Context, dto entity.Segment) (err error)
	}

	UsersSegmentsService interface {
		Create(ctx context.Context, dto entity.SegmentsByUserID) (result entity.SegmentsByUserID, err []error)
		Delete(ctx context.Context, dto entity.SegmentsByUserID) (err []error)
		GetSegmentsByUserID(ctx context.Context, dto entity.GetSegmentsByUserIDDTO) (result entity.SegmentsByUserID, err error)
	}
)

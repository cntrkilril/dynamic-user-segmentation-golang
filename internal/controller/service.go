package controller

import (
	"context"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
)

type (
	SegmentService interface {
		Create(ctx context.Context, dto entity.CreateSegmentDTO) (result entity.Segment, err error)
		Delete(ctx context.Context, dto entity.Segment) (err error)
	}

	UsersSegmentsService interface {
		Create(ctx context.Context, dto entity.SegmentsByUserIDDTO) (result entity.SegmentsByUserIDDTO, err []error)
		Delete(ctx context.Context, dto entity.SegmentsByUserIDDTO) (err []error)
		GetSegmentsByUserID(ctx context.Context, dto entity.GetSegmentsByUserIDDTO) (result entity.SegmentsByUserIDDTO, err error)
	}

	UsersSegmentsHistoryService interface {
		GetCSVHistoryByUserID(ctx context.Context, dto entity.GetCSVHistoryByUserIDAndYearMonthDTO) (url entity.CSVUrl, err error)
	}
)

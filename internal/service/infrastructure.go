package service

import (
	"context"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
)

type (
	SegmentGateway interface {
		Save(ctx context.Context, slug string) (entity.Segment, error)
		FindBySlug(ctx context.Context, slug string) (entity.Segment, error)
		Delete(ctx context.Context, slug string) error
	}

	UsersSegmentsGateway interface {
		Save(ctx context.Context, params entity.UsersSegments) (entity.UsersSegments, error)
		Delete(ctx context.Context, params entity.UsersSegments) error
		FindSegmentsByUserID(ctx context.Context, params int64) ([]entity.UsersSegments, error)
	}
)

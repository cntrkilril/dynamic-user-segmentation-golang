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
		DeleteBySegmentSlug(ctx context.Context, slug string) error
		FindSegmentsByUserID(ctx context.Context, userID int64) ([]entity.UsersSegments, error)
		FindRandomUniqueUserID(ctx context.Context, limit int64) ([]entity.UserID, error)
		CountUniqueUser(ctx context.Context) (int64, error)
	}

	EntityManager interface {
		Segment() SegmentGateway
		UsersSegments() UsersSegmentsGateway
	}

	Registry interface {
		EntityManager
		WithTx(context.Context, func(EntityManager) error) error
	}
)

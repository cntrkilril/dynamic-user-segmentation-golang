package service

import (
	"context"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
)

type (
	SegmentGateway interface {
		Save(ctx context.Context, slug string) (entity.Segment, error)
		FindBySegmentSlug(ctx context.Context, slug string) (entity.Segment, error)
		Delete(ctx context.Context, slug string) error
	}

	UsersSegmentsGateway interface {
		Save(ctx context.Context, params entity.UsersSegments) (entity.UsersSegments, error)
		Delete(ctx context.Context, params entity.UsersSegments) error
		DeleteBySegmentSlug(ctx context.Context, slug string) error
		FindSegmentSlugsByUserID(ctx context.Context, userID int64) ([]entity.UsersSegments, error)
		FindUserIDsBySegmentSlug(ctx context.Context, slug string) ([]entity.UserID, error)
		FindRandomUniqueUserIDs(ctx context.Context, limit int64) ([]entity.UserID, error)
		CountUniqueUsers(ctx context.Context) (int64, error)
	}

	UsersSegmentsHistoryGateway interface {
		Save(ctx context.Context, params entity.SaveUsersSegmentHistoryParams) (entity.UsersSegmentHistory, error)
		FindByUserIDAndYearMonth(ctx context.Context, params entity.GetCSVHistoryByUserIDAndYearMonthDTO) ([]entity.UsersSegmentHistory, error)
	}

	EntityManager interface {
		Segment() SegmentGateway
		UsersSegments() UsersSegmentsGateway
		UsersSegmentsHistory() UsersSegmentsHistoryGateway
	}

	Registry interface {
		EntityManager
		WithTx(context.Context, func(EntityManager) error) error
	}
)

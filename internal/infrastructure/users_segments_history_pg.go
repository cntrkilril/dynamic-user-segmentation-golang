package infrastructure

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/service"
)

type UsersSegmentsHistoryRepository struct {
	db queryRunner
}

func (r *UsersSegmentsHistoryRepository) FindByUserIDAndYearMonth(ctx context.Context, params entity.GetCSVHistoryByUserIDAndYearMonthDTO) (result []entity.UsersSegmentHistory, err error) {
	q := `
		SELECT * FROM users_segments_history
		WHERE user_id=$1 and datetime >= make_date($2, $3, 1) and datetime < make_date($2, $3 + 1, 1)
		`

	err = r.db.SelectContext(ctx, &result, q, params.UserID, params.Year, params.Month)
	if err != nil {
		return []entity.UsersSegmentHistory{}, err
	}

	return result, nil
}

func (r *UsersSegmentsHistoryRepository) Save(ctx context.Context, params entity.SaveUsersSegmentHistoryParams) (result entity.UsersSegmentHistory, err error) {
	q := `
		INSERT INTO users_segments_history (user_id, segment_slug, operation, datetime)
		VALUES ($1,$2,$3, now())
		RETURNING user_id, segment_slug, operation, datetime;
		`

	err = r.db.GetContext(ctx, &result, q, params.UserID, params.SegmentSlug, params.Operation.String())
	if err != nil {
		return entity.UsersSegmentHistory{}, err
	}

	return result, nil
}

var _ service.UsersSegmentsHistoryGateway = (*UsersSegmentsHistoryRepository)(nil)

func NewUsersSegmentsHistoryRepository(db sqlx.DB) *UsersSegmentsHistoryRepository {
	return &UsersSegmentsHistoryRepository{
		db: &db,
	}
}

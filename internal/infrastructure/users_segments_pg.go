package infrastructure

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/service"
)

type UsersSegmentsRepository struct {
	db *sqlx.DB
}

func (r *UsersSegmentsRepository) Save(ctx context.Context, params entity.UsersSegments) (result entity.UsersSegments, err error) {
	q := `
		INSERT INTO users_segments (user_id, segment_slug)
		VALUES ($1,$2)
		RETURNING user_id, segment_slug;
		`

	err = r.db.GetContext(ctx, &result, q, params.UserID, params.SegmentSlug)
	if err != nil {
		return entity.UsersSegments{}, err
	}

	return result, nil
}

func (r *UsersSegmentsRepository) FindSegmentsByUserID(ctx context.Context, userID int64) (result []entity.UsersSegments, err error) {
	q := `
		SELECT user_id, segment_slug FROM users_segments
		WHERE user_id=$1;
		`

	err = r.db.SelectContext(ctx, &result, q, userID)
	if err != nil {
		return []entity.UsersSegments{}, err
	}

	return result, nil
}

func (r *UsersSegmentsRepository) Delete(ctx context.Context, params entity.UsersSegments) error {
	q := `
		DELETE
		FROM users_segments
		WHERE user_id=$1 and segment_slug=$2;
		`
	_, err := r.db.ExecContext(ctx, q, params.UserID, params.SegmentSlug)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

var _ service.UsersSegmentsGateway = (*UsersSegmentsRepository)(nil)

func NewUsersSegmentsRepository(db sqlx.DB) *UsersSegmentsRepository {
	return &UsersSegmentsRepository{
		db: &db,
	}
}

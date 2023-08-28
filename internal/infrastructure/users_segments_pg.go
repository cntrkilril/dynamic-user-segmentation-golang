package infrastructure

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/service"
)

type UsersSegmentsRepository struct {
	db queryRunner
}

func (r *UsersSegmentsRepository) DeleteBySegmentSlug(ctx context.Context, slug string) (err error) {

	q := `
		DELETE
		FROM users_segments
		WHERE segment_slug=$1;
		`
	_, err = r.db.ExecContext(ctx, q, slug)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (r *UsersSegmentsRepository) FindRandomUniqueUserID(ctx context.Context, limit int64) (result []entity.UserID, err error) {
	q := `
		SELECT * FROM (SELECT DISTINCT (user_id) from users_segments) users_segments  ORDER BY random() LIMIT $1;
		`

	err = r.db.SelectContext(ctx, &result, q, limit)
	if err != nil {
		return []entity.UserID{}, err
	}

	return result, nil
}

func (r *UsersSegmentsRepository) CountUniqueUser(ctx context.Context) (result int64, err error) {
	q := `
		SELECT COUNT(DISTINCT (user_id)) AS count FROM users_segments;
	`

	err = r.db.GetContext(ctx, &result, q)
	if err != nil {
		return 0, err
	}

	return result, nil
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

func (r *UsersSegmentsRepository) Delete(ctx context.Context, params entity.UsersSegments) (err error) {

	var q string
	if params.SegmentSlug.Valid {
		q = `
		DELETE
		FROM users_segments
		WHERE user_id=$1 and segment_slug=$2;
		`
		_, err = r.db.ExecContext(ctx, q, params.UserID, params.SegmentSlug)
	} else {
		q = `
		DELETE
		FROM users_segments
		WHERE user_id=$1 and segment_slug IS NULL;
		`
		_, err = r.db.ExecContext(ctx, q, params.UserID)
	}

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

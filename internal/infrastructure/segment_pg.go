package infrastructure

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/service"
)

type SegmentRepository struct {
	db *sqlx.DB
}

func (r *SegmentRepository) Save(ctx context.Context, slug string) (result entity.Segment, err error) {
	q := `
		INSERT INTO segments (slug)
		VALUES ($1)
		RETURNING slug;
		`

	err = r.db.GetContext(ctx, &result, q, slug)
	if err != nil {
		return entity.Segment{}, err
	}

	return result, nil
}

func (r *SegmentRepository) FindBySlug(ctx context.Context, slug string) (result entity.Segment, err error) {
	q := `
		SELECT * FROM segments
		WHERE slug=$1;
		`

	err = r.db.GetContext(ctx, &result, q, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Segment{}, entity.ErrSegmentNotFound
		}
		return entity.Segment{}, err
	}

	return result, nil
}

func (r *SegmentRepository) Delete(ctx context.Context, slug string) error {
	q := `
		DELETE
		FROM segments
		WHERE slug=$1;
		`
	_, err := r.db.ExecContext(ctx, q, slug)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

var _ service.SegmentGateway = (*SegmentRepository)(nil)

func NewSegmentRepository(db sqlx.DB) *SegmentRepository {
	return &SegmentRepository{
		db: &db,
	}
}

package infrastructure

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/service"
)

type (
	PGRegistry struct {
		db *sqlx.DB
		m  *pgManager
	}
)

func (r *PGRegistry) UsersSegmentsHistory() service.UsersSegmentsHistoryGateway {
	return r.m.usersSegmentsHistory
}

func (r *PGRegistry) Segment() service.SegmentGateway {
	return r.m.segments
}

func (r *PGRegistry) UsersSegments() service.UsersSegmentsGateway {
	return r.m.usersSegments
}

func (r *PGRegistry) WithTx(ctx context.Context, f func(manager service.EntityManager) error) error {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = f(&pgManager{
		segments:             &SegmentRepository{tx},
		usersSegments:        &UsersSegmentsRepository{tx},
		usersSegmentsHistory: &UsersSegmentsHistoryRepository{tx},
	})
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

var _ service.Registry = &PGRegistry{}

func NewPGRegistry(db *sqlx.DB) *PGRegistry {
	return &PGRegistry{
		db: db,
		m: &pgManager{
			segments:             &SegmentRepository{db},
			usersSegments:        &UsersSegmentsRepository{db},
			usersSegmentsHistory: &UsersSegmentsHistoryRepository{db},
		},
	}
}

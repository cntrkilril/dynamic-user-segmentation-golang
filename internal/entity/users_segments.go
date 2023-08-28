package entity

import "gopkg.in/guregu/null.v4"

type (
	UsersSegments struct {
		UserID      int64       `json:"userID" db:"user_id"`
		SegmentSlug null.String `json:"segmentSlug" db:"segment_slug"`
	}

	GetSegmentsByUserIDDTO struct {
		UserID int64 `params:"userID" db:"user_id" validate:"required,gte=1"`
	}

	SegmentsByUserID struct {
		UserID   int64    `json:"userID" validate:"required,gte=1"`
		Segments []string `json:"segments" validate:"required"`
	}

	UserID struct {
		UserID int64 `db:"user_id"`
	}
)

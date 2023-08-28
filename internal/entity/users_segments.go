package entity

type (
	UsersSegments struct {
		UserID      int64  `json:"userID" db:"user_id"`
		SegmentSlug string `json:"segmentSlug" db:"segment_slug"`
	}

	GetSegmentsByUserIDDTO struct {
		UserID int64 `params:"userID" db:"user_id" validate:"required,gte=1"`
	}

	SegmentsByUserID struct {
		UserID   int64    `json:"userID" validate:"required,gte=1"`
		Segments []string `json:"segments" validate:"required,gte=1"`
	}
)

package entity

type (
	Segment struct {
		Slug string `json:"slug" db:"slug" validate:"required,gte=1"`
	}

	CreateSegmentDTO struct {
		Slug                 string `json:"slug" db:"slug" validate:"required,gte=1"`
		AutoAddToUserPercent int64  `json:"autoAddToUserPercent"  validate:"gte=0,lte=100"`
	}

	FindUniqueUserIDParams struct {
		Limit int64
	}
)

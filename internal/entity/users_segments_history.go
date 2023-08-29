package entity

import "time"

type (
	Operation int64

	UsersSegmentHistory struct {
		UserID      int64     `db:"user_id"`
		SegmentSlug string    `db:"segment_slug"`
		Operation   string    `db:"operation"`
		DateTime    time.Time `db:"datetime"`
	}

	SaveUsersSegmentHistoryParams struct {
		UserID      int64     `db:"user_id"`
		SegmentSlug string    `db:"segment_slug"`
		Operation   Operation `db:"operation"`
	}

	GetCSVHistoryByUserIDAndYearMonthDTO struct {
		UserID int64 `json:"userID" validate:"required,gte=1"`
		Year   int64 `json:"year" validate:"required,gte=2000,lte=3000"`
		Month  int64 `json:"month" validate:"required,gte=1,lte=12"`
	}

	CSVUrl struct {
		Url string `json:"url"`
	}
)

const (
	OperationAdd Operation = iota
	OperationDelete
)

func (s Operation) String() string {
	switch s {
	case OperationAdd:
		return "add"
	case OperationDelete:
		return "delete"
	}
	return "unknown"
}

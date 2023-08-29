package service

import (
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
)

func HandleServiceError(err error) error {
	switch err {
	case entity.ErrSegmentAlreadyExist:
		return entity.NewError(entity.ErrSegmentAlreadyExist.Error(), entity.ErrCodeBadRequest)
	case entity.ErrSegmentNotFound:
		return entity.NewError(entity.ErrSegmentNotFound.Error(), entity.ErrCodeNotFound)
	case entity.ErrUsersSegmentsNotFound:
		return entity.NewError(entity.ErrUsersSegmentsNotFound.Error(), entity.ErrCodeNotFound)
	case entity.ErrUsersSegmentsHistoryNotFound:
		return entity.NewError(entity.ErrUsersSegmentsHistoryNotFound.Error(), entity.ErrCodeNotFound)
	case entity.ErrValidationError:
		return entity.NewError(entity.ErrValidationError.Error(), entity.ErrCodeBadRequest)
	default:
		return entity.NewError(entity.ErrUnknown.Error(), entity.ErrCodeInternal)
	}
}

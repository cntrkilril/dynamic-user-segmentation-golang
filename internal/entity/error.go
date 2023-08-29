package entity

import "errors"

type (
	Error struct {
		msg  string
		code ErrCode
	}
	ErrCode int64
)

const (
	_ = ErrCode(iota)
	ErrCodeBadRequest
	ErrCodeInternal
	ErrCodeNotFound
)

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() ErrCode {
	return e.code
}

var _ error = &Error{}

func NewError(msg string, code ErrCode) *Error {
	return &Error{msg, code}
}

func NewErrSegmentNotFound(segmentSlug string) *Error {
	return &Error{
		msg:  "сегмент " + segmentSlug + " не найден",
		code: ErrCodeNotFound,
	}
}

func NewErrUsersSegmentsIsAlreadyExist(segmentSlug string) *Error {
	return &Error{
		msg:  "сегмент " + segmentSlug + " уже добавлен",
		code: ErrCodeBadRequest,
	}
}

var (
	ErrUnknown                      = errors.New("что-то пошло не так")
	ErrValidationError              = errors.New("невалидные данные")
	ErrSegmentNotFound              = errors.New("сегмент не найден")
	ErrUsersSegmentsNotFound        = errors.New("пользователь и его сегменты не найдены")
	ErrUsersSegmentsHistoryNotFound = errors.New("пользователь или его история не найдены")
	ErrSegmentAlreadyExist          = errors.New("сегмент уже существует")
)

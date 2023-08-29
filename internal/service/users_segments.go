package service

import (
	"context"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/controller"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"gopkg.in/guregu/null.v4"
)

type UsersSegmentsService struct {
	repos Registry
}

func (s *UsersSegmentsService) Create(ctx context.Context, dto entity.SegmentsByUserIDDTO) (result entity.SegmentsByUserIDDTO, errorArray []error) {

	currentSegments, err := s.repos.UsersSegments().FindSegmentSlugsByUserID(ctx, dto.UserID)
	if err != nil {
		return entity.SegmentsByUserIDDTO{}, []error{HandleServiceError(err)}
	}

	currentSegmentsMap := make(map[string]bool)
	for _, v := range currentSegments {
		currentSegmentsMap[v.SegmentSlug.ValueOrZero()] = true
	}

	errorArray = make([]error, 0, len(dto.Segments))

	if len(dto.Segments) == 0 {
		if len(currentSegments) == 0 {
			_, err := s.repos.UsersSegments().Save(ctx, entity.UsersSegments{
				UserID:      dto.UserID,
				SegmentSlug: null.NewString("", false),
			})
			if err != nil {
				return entity.SegmentsByUserIDDTO{}, []error{HandleServiceError(err)}
			}
		}
	} else {
		for _, v := range dto.Segments {
			_, err = s.repos.Segment().FindBySegmentSlug(ctx, v)
			if err != nil {
				if err == entity.ErrSegmentNotFound {
					errorArray = append(errorArray, entity.NewErrSegmentNotFound(v))
					continue
				} else {
					errorArray = append(errorArray, err)
					continue
				}
			}

			if _, ok := currentSegmentsMap[v]; ok {
				errorArray = append(errorArray, entity.NewErrUsersSegmentsIsAlreadyExist(v))
				continue
			}

			var res entity.UsersSegments
			err = s.repos.WithTx(ctx, func(m EntityManager) (err error) {
				res, err = s.repos.UsersSegments().Save(ctx, entity.UsersSegments{
					UserID:      dto.UserID,
					SegmentSlug: null.NewString(v, true),
				})
				if err != nil {
					return err
				}

				_, err = m.UsersSegmentsHistory().Save(ctx, entity.SaveUsersSegmentHistoryParams{UserID: dto.UserID, SegmentSlug: v, Operation: entity.OperationAdd})
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				errorArray = append(errorArray, HandleServiceError(err))
				continue
			}

			result.Segments = append(result.Segments, res.SegmentSlug.ValueOrZero())
		}
	}

	result.UserID = dto.UserID

	return result, errorArray
}

func (s *UsersSegmentsService) Delete(ctx context.Context, dto entity.SegmentsByUserIDDTO) (errorArray []error) {

	errorArray = make([]error, 0, len(dto.Segments))

	if len(dto.Segments) == 0 {
		err := s.repos.UsersSegments().Delete(ctx, entity.UsersSegments{
			UserID:      dto.UserID,
			SegmentSlug: null.NewString("", false),
		})
		if err != nil {
			errorArray = append(errorArray, HandleServiceError(err))
		}
	}

	for _, v := range dto.Segments {
		_, err := s.repos.Segment().FindBySegmentSlug(ctx, v)
		if err != nil {
			if err == entity.ErrSegmentNotFound {
				errorArray = append(errorArray, entity.NewErrSegmentNotFound(v))
				continue
			} else {
				errorArray = append(errorArray, err)
				continue
			}
		}

		err = s.repos.WithTx(ctx, func(m EntityManager) (err error) {
			err = s.repos.UsersSegments().Delete(ctx, entity.UsersSegments{
				UserID:      dto.UserID,
				SegmentSlug: null.NewString(v, true),
			})
			if err != nil {
				return err
			}

			_, err = m.UsersSegmentsHistory().Save(ctx, entity.SaveUsersSegmentHistoryParams{UserID: dto.UserID, SegmentSlug: v, Operation: entity.OperationDelete})
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			errorArray = append(errorArray, HandleServiceError(err))
		}
	}

	return errorArray
}

func (s *UsersSegmentsService) GetSegmentsByUserID(ctx context.Context, dto entity.GetSegmentsByUserIDDTO) (result entity.SegmentsByUserIDDTO, err error) {
	segments, err := s.repos.UsersSegments().FindSegmentSlugsByUserID(ctx, dto.UserID)
	if err != nil {
		return entity.SegmentsByUserIDDTO{}, HandleServiceError(err)
	}
	if len(segments) == 0 {
		return entity.SegmentsByUserIDDTO{}, HandleServiceError(entity.ErrUsersSegmentsNotFound)
	}

	result.Segments = make([]string, 0, len(segments))
	result.UserID = dto.UserID
	for _, v := range segments {
		if v.SegmentSlug.Valid {
			result.Segments = append(result.Segments, v.SegmentSlug.ValueOrZero())
		}
	}
	return result, nil

}

var _ controller.UsersSegmentsService = (*UsersSegmentsService)(nil)

func NewUsersSegmentsService(repos Registry) *UsersSegmentsService {
	return &UsersSegmentsService{
		repos: repos,
	}
}

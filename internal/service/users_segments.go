package service

import (
	"context"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/controller"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
)

type UsersSegmentsService struct {
	usersSegmentsRepo UsersSegmentsGateway
	segmentRepo       SegmentGateway
}

func (s *UsersSegmentsService) Create(ctx context.Context, dto entity.SegmentsByUserID) (result entity.SegmentsByUserID, errorArray []error) {

	currentSegments, err := s.usersSegmentsRepo.FindSegmentsByUserID(ctx, dto.UserID)
	if err != nil {
		return entity.SegmentsByUserID{}, []error{HandleServiceError(err)}
	}

	currentSegmentsMap := make(map[string]bool)
	for _, v := range currentSegments {
		currentSegmentsMap[v.SegmentSlug] = true
	}

	errorArray = make([]error, 0, len(dto.Segments))

	for _, v := range dto.Segments {
		_, err = s.segmentRepo.FindBySlug(ctx, v)
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

		res, err := s.usersSegmentsRepo.Save(ctx, entity.UsersSegments{
			UserID:      dto.UserID,
			SegmentSlug: v,
		})
		if err != nil {
			errorArray = append(errorArray, HandleServiceError(err))
		}

		result.Segments = append(result.Segments, res.SegmentSlug)
	}

	result.UserID = dto.UserID

	return result, errorArray
}

func (s *UsersSegmentsService) Delete(ctx context.Context, dto entity.SegmentsByUserID) (errorArray []error) {

	errorArray = make([]error, 0, len(dto.Segments))

	for _, v := range dto.Segments {
		_, err := s.segmentRepo.FindBySlug(ctx, v)
		if err != nil {
			if err == entity.ErrSegmentNotFound {
				errorArray = append(errorArray, entity.NewErrSegmentNotFound(v))
				continue
			} else {
				errorArray = append(errorArray, err)
				continue
			}
		}

		err = s.usersSegmentsRepo.Delete(ctx, entity.UsersSegments{
			UserID:      dto.UserID,
			SegmentSlug: v,
		})
		if err != nil {
			errorArray = append(errorArray, HandleServiceError(err))
		}
	}

	return errorArray
}

func (s *UsersSegmentsService) GetSegmentsByUserID(ctx context.Context, dto entity.GetSegmentsByUserIDDTO) (result entity.SegmentsByUserID, err error) {
	segments, err := s.usersSegmentsRepo.FindSegmentsByUserID(ctx, dto.UserID)
	if err != nil {
		return entity.SegmentsByUserID{}, HandleServiceError(err)
	}
	if len(segments) == 0 {
		return entity.SegmentsByUserID{}, HandleServiceError(entity.ErrUsersSegmentsNotFound)
	}

	result.Segments = make([]string, 0, len(segments))
	result.UserID = dto.UserID
	for _, v := range segments {
		result.Segments = append(result.Segments, v.SegmentSlug)
	}
	return result, nil

}

var _ controller.UsersSegmentsService = (*UsersSegmentsService)(nil)

func NewUsersSegmentsService(usersSegmentsRepo UsersSegmentsGateway, segmentRepo SegmentGateway) *UsersSegmentsService {
	return &UsersSegmentsService{
		usersSegmentsRepo: usersSegmentsRepo,
		segmentRepo:       segmentRepo,
	}
}

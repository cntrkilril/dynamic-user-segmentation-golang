package service

import (
	"context"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/controller"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
)

type (
	SegmentService struct {
		segmentRepo SegmentGateway
	}
)

func (s *SegmentService) Create(ctx context.Context, dto entity.Segment) (result entity.Segment, err error) {
	_, err = s.segmentRepo.FindBySlug(ctx, dto.Slug)
	if err != entity.ErrSegmentNotFound {
		if err != nil {
			return entity.Segment{}, HandleServiceError(err)
		}
		return entity.Segment{}, HandleServiceError(entity.ErrSegmentAlreadyExist)
	}

	result, err = s.segmentRepo.Save(ctx, dto.Slug)
	if err != nil {
		return entity.Segment{}, HandleServiceError(err)
	}

	return result, nil
}

func (s *SegmentService) Delete(ctx context.Context, dto entity.Segment) (err error) {
	_, err = s.segmentRepo.FindBySlug(ctx, dto.Slug)
	if err != nil {
		return HandleServiceError(err)
	}

	err = s.segmentRepo.Delete(ctx, dto.Slug)
	if err != nil {
		return HandleServiceError(err)
	}

	return nil
}

var _ controller.SegmentService = (*SegmentService)(nil)

func NewSegmentService(segmentRepo SegmentGateway) *SegmentService {
	return &SegmentService{
		segmentRepo: segmentRepo,
	}
}

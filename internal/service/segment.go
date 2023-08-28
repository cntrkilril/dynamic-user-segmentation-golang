package service

import (
	"context"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/controller"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"gopkg.in/guregu/null.v4"
)

type (
	SegmentService struct {
		repos Registry
	}
)

func (s *SegmentService) Create(ctx context.Context, dto entity.CreateSegmentDTO) (result entity.Segment, err error) {
	_, err = s.repos.Segment().FindBySlug(ctx, dto.Slug)
	if err != entity.ErrSegmentNotFound {
		if err != nil {
			return entity.Segment{}, HandleServiceError(err)
		}
		return entity.Segment{}, HandleServiceError(entity.ErrSegmentAlreadyExist)
	}

	userCount, err := s.repos.UsersSegments().CountUniqueUser(ctx)
	if err != nil {
		return entity.Segment{}, HandleServiceError(err)
	}

	randomUserIDs, err := s.repos.UsersSegments().FindRandomUniqueUserID(ctx, userCount*dto.AutoAddToUserPercent/100)
	if err != nil {
		return entity.Segment{}, HandleServiceError(err)
	}

	err = s.repos.WithTx(ctx, func(m EntityManager) (err error) {
		result, err = m.Segment().Save(ctx, dto.Slug)
		if err != nil {
			return err
		}

		for _, v := range randomUserIDs {
			_, err := s.repos.UsersSegments().Save(ctx, entity.UsersSegments{
				UserID:      v.UserID,
				SegmentSlug: null.NewString(dto.Slug, true),
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return entity.Segment{}, HandleServiceError(err)
	}

	return result, nil
}

func (s *SegmentService) Delete(ctx context.Context, dto entity.Segment) (err error) {
	_, err = s.repos.Segment().FindBySlug(ctx, dto.Slug)
	if err != nil {
		return HandleServiceError(err)
	}

	err = s.repos.Segment().Delete(ctx, dto.Slug)
	if err != nil {
		return HandleServiceError(err)
	}

	err = s.repos.WithTx(ctx, func(m EntityManager) (err error) {
		err = s.repos.Segment().Delete(ctx, dto.Slug)
		if err != nil {
			return err
		}

		err = s.repos.UsersSegments().DeleteBySegmentSlug(ctx, dto.Slug)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}

var _ controller.SegmentService = (*SegmentService)(nil)

func NewSegmentService(repos Registry) *SegmentService {
	return &SegmentService{
		repos: repos,
	}
}

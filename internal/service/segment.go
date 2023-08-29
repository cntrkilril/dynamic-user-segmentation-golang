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
	_, err = s.repos.Segment().FindBySegmentSlug(ctx, dto.Slug)
	if err != entity.ErrSegmentNotFound {
		if err != nil {
			return entity.Segment{}, HandleServiceError(err)
		}
		return entity.Segment{}, HandleServiceError(entity.ErrSegmentAlreadyExist)
	}

	userCount, err := s.repos.UsersSegments().CountUniqueUsers(ctx)
	if err != nil {
		return entity.Segment{}, HandleServiceError(err)
	}

	randomUserIDs, err := s.repos.UsersSegments().FindRandomUniqueUserIDs(ctx, userCount*dto.AutoAddToUserPercent/100)
	if err != nil {
		return entity.Segment{}, HandleServiceError(err)
	}

	err = s.repos.WithTx(ctx, func(m EntityManager) (err error) {
		result, err = m.Segment().Save(ctx, dto.Slug)
		if err != nil {
			return err
		}

		for _, v := range randomUserIDs {
			_, err = m.UsersSegmentsHistory().Save(ctx, entity.SaveUsersSegmentHistoryParams{UserID: v.UserID, SegmentSlug: dto.Slug, Operation: entity.OperationAdd})
			if err != nil {
				return err
			}

			_, err = m.UsersSegments().Save(ctx, entity.UsersSegments{
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
	_, err = s.repos.Segment().FindBySegmentSlug(ctx, dto.Slug)
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
		usersIDs, err := s.repos.UsersSegments().FindUserIDsBySegmentSlug(ctx, dto.Slug)

		for _, v := range usersIDs {
			_, err = m.UsersSegmentsHistory().Save(ctx, entity.SaveUsersSegmentHistoryParams{UserID: v.UserID, SegmentSlug: dto.Slug, Operation: entity.OperationDelete})
			if err != nil {
				return err
			}
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

package service

import (
	"context"
	"encoding/csv"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/controller"
	"github/cntrkilril/dynamic-user-segmentation-golang/internal/entity"
	"os"
	"strconv"
	"strings"
)

type UsersSegmentsHistoryService struct {
	usersSegmentsHistoryRepo UsersSegmentsHistoryGateway
	pathToSaveCsv            string
	baseUrl                  string
}

func (s *UsersSegmentsHistoryService) GetCSVHistoryByUserID(ctx context.Context, dto entity.GetCSVHistoryByUserIDAndYearMonthDTO) (url entity.CSVUrl, err error) {
	historyArray, err := s.usersSegmentsHistoryRepo.FindByUserIDAndYearMonth(ctx, dto)
	if err != nil {
		return entity.CSVUrl{}, HandleServiceError(err)
	}

	if len(historyArray) == 0 {
		return entity.CSVUrl{}, HandleServiceError(entity.ErrUsersSegmentsHistoryNotFound)
	}

	userID := strconv.Itoa(int(dto.UserID))

	var fileName strings.Builder

	fileName.WriteString(s.pathToSaveCsv)
	fileName.WriteString("/")
	fileName.WriteString(userID)
	fileName.WriteString(".csv")

	f, err := os.Create(fileName.String())
	defer f.Close()

	if err != nil {
		return entity.CSVUrl{}, HandleServiceError(err)
	}

	records := make([][]string, 0, len(historyArray)+1)
	records = append(records, []string{"user_id", "segment_slug", "operation", "datetime"})
	for _, v := range historyArray {
		newRecord := []string{strconv.Itoa(int(v.UserID)), v.SegmentSlug, v.Operation, v.DateTime.String()}
		records = append(records, newRecord)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)

	if err != nil {
		return entity.CSVUrl{}, HandleServiceError(err)
	}

	return entity.CSVUrl{
		Url: s.baseUrl + userID + ".csv",
	}, nil
}

var _ controller.UsersSegmentsHistoryService = (*UsersSegmentsHistoryService)(nil)

func NewUsersSegmentsHistoryService(usersSegmentsHistoryRepo UsersSegmentsHistoryGateway, pathToSaveCsv string, baseUrl string) *UsersSegmentsHistoryService {
	return &UsersSegmentsHistoryService{
		usersSegmentsHistoryRepo: usersSegmentsHistoryRepo,
		pathToSaveCsv:            pathToSaveCsv,
		baseUrl:                  baseUrl,
	}
}

package infrastructure

import "github/cntrkilril/dynamic-user-segmentation-golang/internal/service"

type (
	pgManager struct {
		usersSegments *UsersSegmentsRepository
		segments      *SegmentRepository
	}
)

func (m *pgManager) Segment() service.SegmentGateway {
	return m.segments
}

func (m *pgManager) UsersSegments() service.UsersSegmentsGateway {
	return m.usersSegments
}

var _ service.EntityManager = &pgManager{}

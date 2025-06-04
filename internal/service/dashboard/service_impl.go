package dashboard

import (
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/dashboard/dto"
    "pencatatan_keuangan/internal/service/dashboard/impl"
    "pencatatan_keuangan/internal/service/dashboard/mapper"

    "github.com/google/uuid"
)

type DashboardService interface {
    GetDashboard(userID uuid.UUID, filter dto.FilterRequest) (*dto.DashboardResponse, error)
}

type dashboardService struct {
    repository      repository.DashboardRepository
    dashboardMapper *mapper.DashboardMapper
}

func NewDashboardService(repository repository.DashboardRepository) DashboardService {
    return &dashboardService{
        repository:      repository,
        dashboardMapper: mapper.NewDashboardMapper(),
    }
}

func (s *dashboardService) GetDashboard(userID uuid.UUID, filter dto.FilterRequest) (*dto.DashboardResponse, error) {
    return impl.GetDashboard(s.repository, s.dashboardMapper, userID, filter)
}

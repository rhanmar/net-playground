package dummy

import (
	"context"

	"net-playground/internal/domain/dto"
)

type repo interface {
	Save(ctx context.Context, data string) error
	GetInfos(ctx context.Context) ([]*dto.GetDummyInfo, error)
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Save(ctx context.Context, data string) error {
	return s.repo.Save(ctx, data)
}

func (s *Service) GetInfos(ctx context.Context) ([]*dto.GetDummyInfo, error) {
	return s.repo.GetInfos(ctx)
}

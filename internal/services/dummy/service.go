package dummy

import "context"

type repo interface {
	Save(ctx context.Context, data string) error
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

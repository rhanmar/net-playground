package dummy

import (
	"context"
	"errors"
	"testing"

	"net-playground/internal/domain/dto"
	"net-playground/internal/services/dummy/mocks"
)

func TestService_Save(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepoMock(t)
		repo.SaveMock.Expect(context.Background(), "test-data").Return(nil)

		s := NewService(repo)
		err := s.Save(context.Background(), "test-data")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("repo error")
		repo := mocks.NewRepoMock(t)
		repo.SaveMock.Expect(context.Background(), "data").Return(expectedErr)

		s := NewService(repo)
		err := s.Save(context.Background(), "data")
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})
}

func TestService_GetInfos(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expected := []*dto.GetDummyInfo{
			{ID: 1, Data: "first"},
			{ID: 2, Data: "second"},
		}
		repo := mocks.NewRepoMock(t)
		repo.GetInfosMock.Expect(context.Background()).Return(expected, nil)

		s := NewService(repo)
		result, err := s.GetInfos(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != len(expected) {
			t.Fatalf("expected %d items, got %d", len(expected), len(result))
		}
		for i := range expected {
			if result[i].ID != expected[i].ID || result[i].Data != expected[i].Data {
				t.Errorf("item %d: expected %+v, got %+v", i, expected[i], result[i])
			}
		}
	})

	t.Run("empty result", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepoMock(t)
		repo.GetInfosMock.Expect(context.Background()).Return([]*dto.GetDummyInfo{}, nil)

		s := NewService(repo)
		result, err := s.GetInfos(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %d items", len(result))
		}
	})

	t.Run("repo error", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("repo error")
		repo := mocks.NewRepoMock(t)
		repo.GetInfosMock.Expect(context.Background()).Return(nil, expectedErr)

		s := NewService(repo)
		_, err := s.GetInfos(context.Background())
		if !errors.Is(err, expectedErr) {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})
}

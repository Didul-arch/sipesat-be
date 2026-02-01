package subjects

import (
	"context"

	repo "github.com/Didul-arch/sipesat/internal/adapters/postgresql/sqlc"
)

type Subject struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Service interface {
	ListSubjects(ctx context.Context) ([]repo.Subject, error)
}

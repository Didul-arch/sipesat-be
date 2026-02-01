package subjects

import (
	"context"

	repo "github.com/Didul-arch/sipesat/internal/adapters/postgresql/sqlc"
)

type svc struct {
	queries *repo.Queries
	// db      *pgx.Pool
}

func NewService(queries *repo.Queries) Service {
	return &svc{
		queries: queries,
		// db:      db,
	}
}

func (s *svc) ListSubjects(ctx context.Context) ([]repo.Subject, error) {
	return s.queries.ListAllSubject(ctx)
}

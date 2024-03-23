package ports

import (
	"context"

	"github.com/R894/bugtracker/internal/domain/bug/core/domain/aggregate"
)

type BugRepository interface {
	SaveBug(ctx context.Context, bug *aggregate.Bug) error
	GetBugByID(ctx context.Context, bugID string) (*aggregate.Bug, error)
	GetBugsByProjectID(ctx context.Context, projectId string) ([]aggregate.Bug, error)
	GetBugs(ctx context.Context) ([]aggregate.Bug, error)
	AssignBugTo(ctx context.Context, bugId, username string) error
	UpdateBug(ctx context.Context, bug *aggregate.UpdateBugRequest) error
	DeleteBug(ctx context.Context, bugID string) error
}

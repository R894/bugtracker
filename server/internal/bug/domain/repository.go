package domain

import "context"

type BugRepository interface {
	SaveBug(ctx context.Context, bug *Bug) error
	GetBugByID(ctx context.Context, bugID string) (*Bug, error)
	UpdateBug(ctx context.Context, bug *Bug) error
	DeleteBug(ctx context.Context, bugID string) error
}

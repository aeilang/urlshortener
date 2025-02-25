// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repo

import (
	"context"
)

type Querier interface {
	CreateURL(ctx context.Context, arg CreateURLParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteURLByShortCode(ctx context.Context, shortCode string) error
	GetURLsByUserID(ctx context.Context, arg GetURLsByUserIDParams) ([]GetURLsByUserIDRow, error)
	GetUrlByShortCode(ctx context.Context, shortCode string) (GetUrlByShortCodeRow, error)
	GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error)
	IsEmailAvaliable(ctx context.Context, email string) (bool, error)
	IsShortCodeAvailable(ctx context.Context, shortCode string) (bool, error)
	UpdatePasswordByEmail(ctx context.Context, arg UpdatePasswordByEmailParams) (UpdatePasswordByEmailRow, error)
	UpdateURLExpiredByShortCode(ctx context.Context, arg UpdateURLExpiredByShortCodeParams) error
	UpdateViewsByShortCode(ctx context.Context, arg UpdateViewsByShortCodeParams) error
}

var _ Querier = (*Queries)(nil)

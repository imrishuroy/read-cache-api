// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
)

type Querier interface {
	AddTagToCache(ctx context.Context, arg AddTagToCacheParams) (CacheTag, error)
	CreateCache(ctx context.Context, arg CreateCacheParams) (Cache, error)
	CreateTag(ctx context.Context, tagName string) (Tag, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCache(ctx context.Context, id int64) error
	DeleteTagFromCacheTagsTable(ctx context.Context, tagID int32) error
	DeleteTagFromTagsTable(ctx context.Context, tagID int32) error
	GetCache(ctx context.Context, id int64) (Cache, error)
	GetUser(ctx context.Context, id string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListCacheTags(ctx context.Context, cacheID int64) ([]Tag, error)
	ListCaches(ctx context.Context, arg ListCachesParams) ([]Cache, error)
	ListPublicCaches(ctx context.Context, tagIds []int32) ([]Cache, error)
	ListUserSubscriptions(ctx context.Context, userID string) ([]Tag, error)
	SubscribeTag(ctx context.Context, arg SubscribeTagParams) (UserTag, error)
	UnsubscribeTag(ctx context.Context, arg UnsubscribeTagParams) error
	UpdateCache(ctx context.Context, arg UpdateCacheParams) (Cache, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)

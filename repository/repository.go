// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"
	"time"

	"github.com/project-inari/core-business-member/dto"
	"github.com/redis/go-redis/v9"
)

// DatabaseRepository represents the repository layer functions of database repository
type DatabaseRepository interface {
	InviteToJoinBusiness(ctx context.Context, inviterUsername, inviteeUsername, businessName string) error
	AcceptBusinessInvitation(ctx context.Context, inviteeUsername, businessName string) error
	GetBusiness(ctx context.Context, businessName string) (*dto.BusinessEntity, error)
	DeclineBusinessInvitation(ctx context.Context, inviteeUsername, businessName string) error
	GetBusinessJoiningStatus(ctx context.Context, filter dto.BusinessJoiningQueryFilter) ([]*dto.BusinessJoiningEntity, error)
	GetBusinessMembers(ctx context.Context, businessName string) ([]*dto.BusinessMemberEntity, error)
	GetUserJoinedBusinesses(ctx context.Context, username string) ([]*dto.UserBusinessesEntity, error)
}

// CacheRepository represents the repository layer functions of cache repository
type CacheRepository interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd
	UpdateUserCacheNewBusinessJoined(ctx context.Context, username string, business dto.BusinessModel) error
}

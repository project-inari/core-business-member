package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/project-inari/core-business-member/dto"
	"github.com/project-inari/core-business-member/pkg/utils"
)

type cacheRepository struct {
	client                 *redis.Client
	keyUserVerifiedAccount string
}

// CacheRepositoryConfig represents the configuration of the cache repository
type CacheRepositoryConfig struct {
	KeyUserVerifiedAccount string
}

// CacheRepositoryDependencies represents the dependencies of the cache repository
type CacheRepositoryDependencies struct {
	Client *redis.Client
}

// NewCacheRepository creates a new cache repository
func NewCacheRepository(c CacheRepositoryConfig, d CacheRepositoryDependencies) CacheRepository {
	return &cacheRepository{
		client:                 d.Client,
		keyUserVerifiedAccount: c.KeyUserVerifiedAccount,
	}
}

func (r *cacheRepository) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *cacheRepository) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd {
	marshalValue, _ := json.Marshal(value)
	return r.client.Set(ctx, key, marshalValue, ttl)
}

func (r *cacheRepository) UpdateUserCacheNewBusinessJoined(ctx context.Context, username string, business dto.BusinessModel) error {
	key := fmt.Sprintf("%s:%s", r.keyUserVerifiedAccount, username)

	remainingTTL, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		return nil
	}
	if remainingTTL < 0 {
		return nil
	}

	existingCache := r.client.Get(ctx, key)
	if existingCache.Err() != nil {
		return existingCache.Err()
	}
	cacheVal := utils.DecodeJSONfromString[dto.UserVerifiedAccountCache](existingCache.Val())
	if cacheVal == nil {
		return fmt.Errorf("failed to decode cache value")
	}
	cacheVal.Businesses = append(cacheVal.Businesses, business)

	_, err = r.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	marshalValue, _ := json.Marshal(cacheVal)
	return r.client.Set(ctx, key, marshalValue, remainingTTL).Err()
}

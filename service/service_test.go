package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/project-inari/core-business-member/dto"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

type mockDatabaseRepository struct {
	err error
}

func (m *mockDatabaseRepository) InviteToJoinBusiness(_, _, _, _ string) error {
	return m.err
}

type mockCacheRepository struct {
	getRes *redis.StringCmd
	setRes *redis.StatusCmd
}

func (m *mockCacheRepository) Get(_ context.Context, _ string) *redis.StringCmd {
	return m.getRes
}

func (m *mockCacheRepository) Set(_ context.Context, _ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	return m.setRes
}

const (
	mockInviterUsername = "inviter"
	mockInviteeUsername = "invitee"
	mockBusinessName    = "business"
)

func TestInvite(t *testing.T) {
	ctx := context.Background()

	req := dto.InviteReq{
		InviterUsername: mockInviterUsername,
		InviteeUsername: mockInviteeUsername,
		BusinessName:    mockBusinessName,
	}

	expectedRes := &dto.InviteRes{
		InviterUsername: mockInviterUsername,
		InviteeUsername: mockInviteeUsername,
		BusinessName:    mockBusinessName,
		Status:          statusInvitePending,
	}

	t.Run("success", func(t *testing.T) {
		mockCacheRepository := &mockCacheRepository{}
		mockDatabaseRepository := &mockDatabaseRepository{
			err: nil,
		}

		s := New(Dependencies{
			DatabaseRepository: mockDatabaseRepository,
			CacheRepository:    mockCacheRepository,
		})

		res, err := s.Invite(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("error - when database repo returned error", func(t *testing.T) {
		mockCacheRepository := &mockCacheRepository{}
		mockDatabaseRepository := &mockDatabaseRepository{
			err: errors.New("error"),
		}

		s := New(Dependencies{
			DatabaseRepository: mockDatabaseRepository,
			CacheRepository:    mockCacheRepository,
		})

		_, err := s.Invite(ctx, req)

		assert.Error(t, err)
	})
}

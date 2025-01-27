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
	getBusinessRes *dto.BusinessEntity
	err            error
}

func (m *mockDatabaseRepository) InviteToJoinBusiness(_ context.Context, _, _, _ string) error {
	return m.err
}

func (m *mockDatabaseRepository) AcceptBusinessInvitation(_ context.Context, _, _ string) error {
	return m.err
}

func (m *mockDatabaseRepository) GetBusiness(_ context.Context, _ string) (*dto.BusinessEntity, error) {
	return m.getBusinessRes, m.err
}

type mockCacheRepository struct {
	getRes *redis.StringCmd
	setRes *redis.StatusCmd
	err    error
}

func (m *mockCacheRepository) Get(_ context.Context, _ string) *redis.StringCmd {
	return m.getRes
}

func (m *mockCacheRepository) Set(_ context.Context, _ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	return m.setRes
}

func (m *mockCacheRepository) UpdateUserCacheNewBusinessJoined(_ context.Context, _ string, _ dto.BusinessCacheModel) error {
	return m.err
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

func TestAcceptInvite(t *testing.T) {
	ctx := context.Background()

	req := dto.AcceptInviteReq{
		InviteeUsername: mockInviteeUsername,
		BusinessName:    mockBusinessName,
	}

	expectedRes := &dto.AcceptInviteRes{
		InviteeUsername: mockInviteeUsername,
		BusinessName:    mockBusinessName,
		Status:          statusInviteAccepted,
	}

	t.Run("success", func(t *testing.T) {

		mockCacheRepository := &mockCacheRepository{}
		mockDatabaseRepository := &mockDatabaseRepository{
			getBusinessRes: &dto.BusinessEntity{},
			err:            nil,
		}

		s := New(Dependencies{
			DatabaseRepository: mockDatabaseRepository,
			CacheRepository:    mockCacheRepository,
		})

		res, err := s.AcceptInvite(ctx, req)

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

		_, err := s.AcceptInvite(ctx, req)

		assert.Error(t, err)
	})

	t.Run("error - when cache repo returned error", func(t *testing.T) {
		mockCacheRepository := &mockCacheRepository{
			err: errors.New("error"),
		}
		mockDatabaseRepository := &mockDatabaseRepository{
			getBusinessRes: &dto.BusinessEntity{},
			err:            nil,
		}

		s := New(Dependencies{
			DatabaseRepository: mockDatabaseRepository,
			CacheRepository:    mockCacheRepository,
		})

		_, err := s.AcceptInvite(ctx, req)

		assert.Error(t, err)
	})
}

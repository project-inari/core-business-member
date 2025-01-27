// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/project-inari/core-business-member/dto"
	"github.com/project-inari/core-business-member/repository"
)

// Port represents the service layer functions
type Port interface {
	Invite(ctx context.Context, req dto.InviteReq) (*dto.InviteRes, error)
}

type service struct {
	databaseRepository repository.DatabaseRepository
	cacheRepository    repository.CacheRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	DatabaseRepository repository.DatabaseRepository
	CacheRepository    repository.CacheRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		databaseRepository: d.DatabaseRepository,
		cacheRepository:    d.CacheRepository,
	}
}

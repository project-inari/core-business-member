package service

import (
	"context"

	"github.com/project-inari/core-business-member/dto"
	"github.com/project-inari/core-business-member/pkg/utils"
)

const (
	roleMember           = "MEMBER"
	statusInviteAccepted = "INVITE_ACCEPTED"
)

// AcceptInvite accepts a business invitation
func (s *service) AcceptInvite(ctx context.Context, req dto.AcceptInviteReq) (*dto.AcceptInviteRes, error) {
	if err := s.databaseRepository.AcceptBusinessInvitation(ctx, req.InviteeUsername, req.BusinessName); err != nil {
		return nil, err
	}

	business, err := s.databaseRepository.GetBusiness(ctx, req.BusinessName)
	if err != nil {
		return nil, err
	}

	if err := s.cacheRepository.UpdateUserCacheNewBusinessJoined(ctx, req.InviteeUsername, *constructBusinessCacheModel(*business)); err != nil {
		return nil, err
	}

	return &dto.AcceptInviteRes{
		InviteeUsername: req.InviteeUsername,
		BusinessName:    req.BusinessName,
		Status:          statusInviteAccepted,
	}, nil
}

func constructBusinessCacheModel(business dto.BusinessEntity) *dto.BusinessCacheModel {
	return &dto.BusinessCacheModel{
		ID:               business.ID,
		Name:             business.Name,
		IndustryType:     business.IndustryType,
		BusinessType:     business.BusinessType,
		Description:      business.Description,
		PhoneNo:          business.PhoneNo,
		OperatingHours:   *utils.DecodeJSONfromString[dto.OperatingHours](business.OperatingHours),
		Address:          business.Address,
		BusinessImageURL: business.BusinessImageURL,
		CreatedAt:        business.CreatedAt,
		UpdatedAt:        business.UpdatedAt,
		UserRole:         roleMember,
	}
}

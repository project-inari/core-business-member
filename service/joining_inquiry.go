package service

import (
	"context"

	"github.com/project-inari/core-business-member/dto"
)

// JoiningInquiry retrieves the joining status of a user to a business
func (s *service) JoiningInquiry(ctx context.Context, req dto.JoiningInquiryReq) (*dto.JoiningInquiryRes, error) {
	filter := dto.BusinessJoiningQueryFilter{
		Username:     req.Username,
		BusinessName: req.BusinessName,
		Status:       req.Status,
	}

	res, err := s.databaseRepository.GetBusinessJoiningStatus(ctx, filter)
	if err != nil {
		return nil, err
	}

	return constructBusinessJoiningResponse(res), nil
}

func constructBusinessJoiningResponse(businessJoining []*dto.BusinessJoiningEntity) *dto.JoiningInquiryRes {
	result := make([]dto.JoiningInquiryResult, 0)
	for _, joining := range businessJoining {
		result = append(result, dto.JoiningInquiryResult{
			ID:           joining.ID,
			Username:     joining.Username,
			BusinessName: joining.BusinessName,
			Status:       joining.Status,
			ActionedBy:   joining.ActionedBy,
			CreatedAt:    joining.CreatedAt,
			UpdatedAt:    joining.UpdatedAt,
		})
	}

	return &dto.JoiningInquiryRes{
		Result: result,
	}
}

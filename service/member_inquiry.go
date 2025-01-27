package service

import (
	"context"

	"github.com/project-inari/core-business-member/dto"
)

// MemberInquiry retrieves the members of a business
func (s *service) MemberInquiry(ctx context.Context, businessName string) (*dto.MemberInquiryRes, error) {
	memberEntities, err := s.databaseRepository.GetBusinessMembers(ctx, businessName)
	if err != nil {
		return nil, err
	}

	return constructMemberInquiryResponse(businessName, memberEntities), nil
}

func constructMemberInquiryResponse(businessName string, members []*dto.BusinessMemberEntity) *dto.MemberInquiryRes {
	result := make([]dto.MemberInquiryResult, 0)
	for _, member := range members {
		result = append(result, dto.MemberInquiryResult{
			ID:        member.ID,
			Username:  member.Username,
			Role:      member.Role,
			CreatedAt: member.CreatedAt,
		})
	}

	return &dto.MemberInquiryRes{
		BusinessName: businessName,
		Result:       result,
	}
}

package repository

import (
	"context"
	"database/sql"

	"github.com/project-inari/core-business-member/dto"
)

const (
	statusInvitePending  = "INVITE_PENDING"
	statusInviteAccepted = "INVITE_ACCEPTED"
	statusInviteDeclined = "INVITE_DECLINED"
	roleMember           = "MEMBER"
)

type databaseRepository struct {
	database string
	client   *sql.DB
}

// DatabaseRepositoryConfig represents the configuration for wiremock API repository
type DatabaseRepositoryConfig struct {
	Database string
}

// DatabaseRepositoryDependencies represents the dependencies for wiremock API repository
type DatabaseRepositoryDependencies struct {
	Client *sql.DB
}

// NewDatabaseRepository creates a new wiremock API repository
func NewDatabaseRepository(c DatabaseRepositoryConfig, d DatabaseRepositoryDependencies) DatabaseRepository {
	return &databaseRepository{
		database: c.Database,
		client:   d.Client,
	}
}

// InviteToJoinBusiness invites a user to join a business
func (r *databaseRepository) InviteToJoinBusiness(ctx context.Context, inviterUsername, inviteeUsername, businessName string) error {
	_, err := r.client.ExecContext(ctx, "INSERT INTO tbl_business_joinings (business_name, username, status, actioned_by) VALUES (?, ?, ?, ?)", businessName, inviteeUsername, statusInvitePending, inviterUsername)
	if err != nil {
		return err
	}

	return nil
}

func (r *databaseRepository) AcceptBusinessInvitation(ctx context.Context, inviteeUsername, businessName string) error {
	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // nolint: errcheck

	_, err = tx.ExecContext(ctx, "UPDATE tbl_business_joinings SET status = ? WHERE business_name = ? AND username = ? AND status = ?", statusInviteAccepted, businessName, inviteeUsername, statusInvitePending)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO tbl_business_members (business_name, username, role) VALUES (?, ?, ?)", businessName, inviteeUsername, roleMember)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetBusiness retrieves a business by its name
func (r *databaseRepository) GetBusiness(ctx context.Context, businessName string) (*dto.BusinessEntity, error) {
	row := r.client.QueryRowContext(ctx, "SELECT id, name, industry_type, business_type, description, phone_no, operating_hours, address, business_image_url, created_at, updated_at FROM tbl_businesses WHERE name = ?", businessName)

	var entity dto.BusinessEntity
	if err := row.Scan(&entity.ID, &entity.Name, &entity.IndustryType, &entity.BusinessType, &entity.Description, &entity.PhoneNo, &entity.OperatingHours, &entity.Address, &entity.BusinessImageURL, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
		return nil, err
	}

	return &entity, nil
}

// DeclineBusinessInvitation declines a business invitation
func (r *databaseRepository) DeclineBusinessInvitation(ctx context.Context, inviteeUsername, businessName string) error {
	_, err := r.client.ExecContext(ctx, "UPDATE tbl_business_joinings SET status = ? WHERE business_name = ? AND username = ? AND status = ?", statusInviteDeclined, businessName, inviteeUsername, statusInvitePending)
	if err != nil {
		return err
	}

	return nil
}

// GetBusinessJoiningStatus retrieves the business joining status with the given filter
func (r *databaseRepository) GetBusinessJoiningStatus(ctx context.Context, filter dto.BusinessJoiningQueryFilter) ([]*dto.BusinessJoiningEntity, error) {
	query := "SELECT id, business_name, username, status, actioned_by, created_at, updated_at FROM tbl_business_joinings WHERE 1=1"
	args := []interface{}{}

	if filter.BusinessName != "" {
		query += " AND business_name = ?"
		args = append(args, filter.BusinessName)
	}
	if filter.Username != "" {
		query += " AND username = ?"
		args = append(args, filter.Username)
	}
	if filter.Status != "" {
		query += " AND status = ?"
		args = append(args, filter.Status)
	}

	query += " ORDER BY created_at DESC"

	rows, err := r.client.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // nolint: errcheck

	var entities []*dto.BusinessJoiningEntity
	for rows.Next() {
		var entity dto.BusinessJoiningEntity
		if err := rows.Scan(&entity.ID, &entity.BusinessName, &entity.Username, &entity.Status, &entity.ActionedBy, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	return entities, nil
}

// GetBusinessMembers retrieves the business members of a business
func (r *databaseRepository) GetBusinessMembers(ctx context.Context, businessName string) ([]*dto.BusinessMemberEntity, error) {
	rows, err := r.client.QueryContext(ctx, "SELECT id, business_name, username, role, created_at, updated_at FROM tbl_business_members WHERE business_name = ?", businessName)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // nolint: errcheck

	var entities []*dto.BusinessMemberEntity
	for rows.Next() {
		var entity dto.BusinessMemberEntity
		if err := rows.Scan(&entity.ID, &entity.BusinessName, &entity.Username, &entity.Role, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	return entities, nil
}

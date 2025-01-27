package repository

import (
	"database/sql"
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
func (r *databaseRepository) InviteToJoinBusiness(inviterUsername, inviteeUsername, businessName, status string) error {
	_, err := r.client.Exec("INSERT INTO tbl_business_joinings (business_name, username, status, actioned_by) VALUES (?, ?, ?, ?)", businessName, inviteeUsername, status, inviterUsername)
	if err != nil {
		return err
	}

	return nil
}

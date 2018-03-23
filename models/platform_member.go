package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type PlatformMember struct {
	ID                uuid.UUID `json:"id" db:"id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	FirstName         string    `json:"first_name" db:"first_name"`
	LastName          string    `json:"last_name" db:"last_name"`
	Dob               time.Time `json:"dob" db:"dob"`
	BlockchainAddress string    `json:"blockchain_address" db:"blockchain_address"`
	RecipientRef      nulls.Int `json:"recipient_id" db:"recipient_id"`

	// Relationships
	Claims PlatformMemberClaims `has_many:"platform_member_claims"`
	Recipient Recipient `belongs_to:"recipients"`
}

// String is not required by pop and may be deleted
func (p PlatformMember) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// PlatformMembers is not required by pop and may be deleted
type PlatformMembers []PlatformMember

// String is not required by pop and may be deleted
func (p PlatformMembers) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *PlatformMember) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.FirstName, Name: "FirstName"},
		&validators.StringIsPresent{Field: p.LastName, Name: "LastName"},
		&validators.TimeIsPresent{Field: p.Dob, Name: "Dob"},
		&validators.StringIsPresent{Field: p.BlockchainAddress, Name: "BlockchainAddress"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *PlatformMember) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *PlatformMember) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

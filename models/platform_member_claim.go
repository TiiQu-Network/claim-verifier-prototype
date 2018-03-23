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

type PlatformMemberClaim struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	MemberID     int        `json:"member_id" db:"member_id"`
	AddressID    int        `json:"address_id" db:"address_id"`
	RecipientRef string     `json:"institution_reference" db:"institution_reference"`
	IssueID      int        `json:"issue_id" db:"issue_id"`
	IssueLevelID int        `json:"issue_level_id" db:"issue_level_id"`
	IssueDate    time.Time  `json:"issue_date" db:"issue_date"`
	Verified     nulls.Bool `json:"verified" db:"verified"`
	ContractID   nulls.Int  `json:"contract_id" db:"contract_id"`

	// Relationships
	// TODO level
	Issueable Issueable      `belongs_to:"issuables"`
	Member    PlatformMember `belongs_to:"platform_members"`
	Address   Address        `belongs_to:"addresses"`
	Contract  DataHash       `belongs_to:"data_hashes"`
}

// String is not required by pop and may be deleted
func (p PlatformMemberClaim) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// PlatformMemberClaims is not required by pop and may be deleted
type PlatformMemberClaims []PlatformMemberClaim

// String is not required by pop and may be deleted
func (p PlatformMemberClaims) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *PlatformMemberClaim) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.MemberID, Name: "MemberID"},
		&validators.IntIsPresent{Field: p.AddressID, Name: "AddressID"},
		&validators.StringIsPresent{Field: p.RecipientRef, Name: "RecipientRef"},
		&validators.IntIsPresent{Field: p.IssueID, Name: "IssueID"},
		&validators.IntIsPresent{Field: p.IssueLevelID, Name: "IssueLevelID"},
		&validators.TimeIsPresent{Field: p.IssueDate, Name: "IssueDate"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *PlatformMemberClaim) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *PlatformMemberClaim) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

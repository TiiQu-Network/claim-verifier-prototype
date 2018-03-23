package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Issue struct {
	ID               uuid.UUID `json:"id" db:"id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	RecipientID      int       `json:"recipient_id" db:"recipient_id"`
	RecipientRef     string    `json:"recipient_ref" db:"recipient_ref"`
	IssueableID      int       `json:"issueable_id" db:"issueable_id"`
	IssueableLevelID int       `json:"issueable_level_id" db:"issueable_level_id"`
	IssueDate        time.Time `json:"issue_date" db:"issue_date"`
}

// String is not required by pop and may be deleted
func (i Issue) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Issues is not required by pop and may be deleted
type Issues []Issue

// String is not required by pop and may be deleted
func (i Issues) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Issue) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: i.RecipientID, Name: "RecipientID"},
		&validators.StringIsPresent{Field: i.RecipientRef, Name: "RecipientRef"},
		&validators.IntIsPresent{Field: i.IssueableID, Name: "IssueableID"},
		&validators.IntIsPresent{Field: i.IssueableLevelID, Name: "IssueableLevelID"},
		&validators.TimeIsPresent{Field: i.IssueDate, Name: "IssueDate"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Issue) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Issue) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

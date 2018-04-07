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

type Recipient struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	FirstName     string    `json:"first_name" db:"first_name"`
	LastName      string    `json:"last_name" db:"last_name"`
	Dob           time.Time `json:"dob" db:"dob"`
	AddressID     nulls.Int `json:"address_id" db:"address_id"`
	InstitutionID nulls.Int `json:"institution_id" db:"institution_id"`

	// Relationships
	Address     Addresses   `has_many:"recipients"`
	Issues      Issues      `has_many:"issues"`
	Institution Institution `belongs_to:"institution"`
}

// String is not required by pop and may be deleted
func (r Recipient) String() string {
	js, _ := json.Marshal(r)
	return string(js)
}

// Recipients is not required by pop and may be deleted
type Recipients []Recipient

// String is not required by pop and may be deleted
func (r Recipients) String() string {
	js, _ := json.Marshal(r)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Recipient) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: r.FirstName, Name: "FirstName"},
		&validators.StringIsPresent{Field: r.LastName, Name: "LastName"},
		&validators.TimeIsPresent{Field: r.Dob, Name: "dob"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Recipient) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Recipient) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

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

type Address struct {
	ID         uuid.UUID    `json:"id" db:"id"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	FirstLine  string       `json:"first_line" db:"first_line"`
	SecondLine nulls.String `json:"second_line" db:"second_line"`
	Town       nulls.String `json:"town" db:"town"`
	County     nulls.String `json:"county" db:"county"`
	Postcode   string       `json:"postcode" db:"postcode"`
	Country    string       `json:"country" db:"country"`
}

// String is not required by pop and may be deleted
func (a Address) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Addresses is not required by pop and may be deleted
type Addresses []Address

// String is not required by pop and may be deleted
func (a Addresses) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Address) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.FirstLine, Name: "FirstLine"},
		&validators.StringIsPresent{Field: a.Postcode, Name: "Postcode"},
		&validators.StringIsPresent{Field: a.Country, Name: "Country"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Address) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Address) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

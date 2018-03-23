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

type DataHash struct {
	ID                        uuid.UUID    `json:"id" db:"id"`
	CreatedAt                 time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt                 time.Time    `json:"updated_at" db:"updated_at"`
	PersonalData              string       `json:"personal_datum" db:"personal_datum"`
	PersonalConstituents      nulls.String `json:"personal_constituent" db:"personal_constituent"`
	CertificationData         string       `json:"certification_datum" db:"certification_datum"`
	CertificationConstituents nulls.String `json:"certification_constituent" db:"certification_constituent"`
	ClaimantAddress           nulls.String `json:"claimant_address" db:"claimant_address"`
	EstablishmentID           nulls.Int    `json:"establishment_id" db:"establishment_id"`
}

// String is not required by pop and may be deleted
func (d DataHash) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// DataHashes is not required by pop and may be deleted
type DataHashes []DataHash

// String is not required by pop and may be deleted
func (d DataHashes) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *DataHash) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.PersonalData, Name: "PersonalData"},
		&validators.StringIsPresent{Field: d.CertificationData, Name: "CertificationData"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *DataHash) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *DataHash) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

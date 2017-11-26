package models

import (
	"encoding/json"
	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/satori/go.uuid"
	"time"
)

type OrganizationAdmin struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	OrganizationID uuid.UUID `json:"organiation_id" db:"organiation_id"`
}

// String is not required by pop and may be deleted
func (o OrganizationAdmin) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OrganizationAdmins is not required by pop and may be deleted
type OrganizationAdmins []OrganizationAdmin

// String is not required by pop and may be deleted
func (o OrganizationAdmins) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OrganizationAdmin) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OrganizationAdmin) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OrganizationAdmin) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

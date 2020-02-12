package model

import "time"

//UserModel struct
type UserModel struct {
	ID           uint16    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	IsActive     bool      `json:"is_active,omitempty"`
	CreatedBy    string    `json:"created_by,omitempty"`
	ModifiedBy   string    `json:"modified_by,omitempty"`
	CreatedDate  time.Time `json:"created_date,omitempty"`
	ModifiedDate time.Time `json:"modified_date,omitempty"`
}

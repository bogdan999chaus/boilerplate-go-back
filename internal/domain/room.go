package domain

import "time"

type Room struct {
	Id             uint64     `json:"id"`
	OrganizationId uint64     `json:"organization_id"`
	Name           string     `json:"name"`
	Description    *string    `json:"description"`
	CreatedDate    time.Time  `json:"created_date"`
	UpdatedDate    time.Time  `json:"updated_date"`
	DeletedDate    *time.Time `json:"deleted_date"`
}

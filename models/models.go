package models

import (
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID               string         `json:"id" gorm:"primary_key"`
	CreatedDate      *time.Time     `json:"createdDate,omitempty"`
	CreatedBy        *string        `json:"createdBy,omitempty"`
	LastModifiedDate *time.Time     `json:"lastModifiedDate,omitempty"`
	LastModifiedBy   *string        `json:"lastModifiedBy,omitempty"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt,omitempty" sql:"index"`
	DeletedBy        *string        `json:"deletedBy,omitempty"`
}

//BeforeCreate ganerate id
func (model *Model) BeforeCreate(tx *gorm.DB) error {
	if len(strings.Trim(model.ID, " ")) == 0 {
		model.ID = uuid.NewV4().String()
	}
	t := time.Now()
	model.CreatedDate = &t
	model.LastModifiedDate = &t
	return nil
}

func (model *Model) BeforeUpdate(tx *gorm.DB) error {
	t := time.Now()
	model.LastModifiedDate = &t

	return nil
}

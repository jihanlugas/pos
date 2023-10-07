package model

import (
	"github.com/jihanlugas/pos/utils"
	"gorm.io/gorm"
	"time"
)

func (m *Usercompany) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.ID != "" {
		m.ID = utils.GetUniqueID()
	}
	m.CreateDt = now
	m.UpdateDt = now
	return
}

func (m *Usercompany) BeforeUpdate(tx *gorm.DB) (err error) {
	if m.DeleteDt == nil {
		now := time.Now()
		m.UpdateDt = now
	}
	return
}

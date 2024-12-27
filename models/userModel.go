package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint        `gorm:"primaryKey"` // Standard field for the primary key
	Name        string      `gorm:"index"` 		// A regular string field
	Email       string     `gorm:"index"` 	
	ActivatedAt sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt   time.Time    // Automatically managed by GORM for creation time
	UpdatedAt   time.Time    // Automatically managed by GORM for update time
}
package model

import (
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Clusters  []Cluster      `gorm:"one2many:cluster_tenants;"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

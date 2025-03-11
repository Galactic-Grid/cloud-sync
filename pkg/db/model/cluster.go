package model

import (
	"time"

	"gorm.io/gorm"
)

type Cluster struct {
	ID          uint           `gorm:"primarykey"`
	Name        string         `gorm:"type:varchar(255);not null"`
	ClusterURL  string         `gorm:"type:varchar(255);not null"`
	ClusterCert string         `gorm:"type:text"`
	TenantID    uint           `gorm:"not null;index"`
	Tenant      Tenant         `gorm:"foreignKey:TenantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

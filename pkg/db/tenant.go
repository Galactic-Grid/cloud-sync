package db

import (
	"github.com/Galactic-Grid/cloud-sync/pkg/db/model"
	"gorm.io/gorm"
)

type TenantService struct {
	db *gorm.DB
}

func NewTenantService() *TenantService {
	return &TenantService{
		db: GetDB(),
	}
}

func (s *TenantService) Create(tenant *model.Tenant) error {
	return s.db.Create(tenant).Error
}

func (s *TenantService) Update(tenant *model.Tenant) error {
	return s.db.Save(tenant).Error
}

func (s *TenantService) Delete(id uint) error {
	return s.db.Delete(&model.Tenant{}, id).Error
}

func (s *TenantService) Get(id uint) (*model.Tenant, error) {
	var tenant model.Tenant
	err := s.db.First(&tenant, id).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (s *TenantService) List() ([]model.Tenant, error) {
	var tenants []model.Tenant
	err := s.db.Find(&tenants).Error
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

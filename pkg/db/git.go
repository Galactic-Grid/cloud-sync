package db

import (
	"github.com/Galactic-Grid/cloud-sync/pkg/db/model"
	"gorm.io/gorm"
)

type GitConfigService struct {
	db *gorm.DB
}

func NewGitConfigService(db *gorm.DB) *GitConfigService {
	return &GitConfigService{
		db: db,
	}
}

func (s *GitConfigService) Create(config *model.GitConfig) error {
	return s.db.Create(config).Error
}

func (s *GitConfigService) Update(config *model.GitConfig) error {
	return s.db.Save(config).Error
}

func (s *GitConfigService) Delete(id uint) error {
	return s.db.Delete(&model.GitConfig{}, id).Error
}

func (s *GitConfigService) Get(id uint) (*model.GitConfig, error) {
	var config model.GitConfig
	err := s.db.First(&config, id).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *GitConfigService) List() ([]model.GitConfig, error) {
	var configs []model.GitConfig
	err := s.db.Find(&configs).Error
	if err != nil {
		return nil, err
	}
	return configs, nil
}

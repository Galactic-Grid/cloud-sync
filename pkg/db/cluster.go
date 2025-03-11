package db

import (
	"github.com/Galactic-Grid/cloud-sync/pkg/db/model"
	"gorm.io/gorm"
)

type ClusterService struct {
	db *gorm.DB
}

func NewClusterService(db *gorm.DB) *ClusterService {
	return &ClusterService{
		db: db,
	}
}

func (s *ClusterService) Create(cluster *model.Cluster) error {
	return s.db.Create(cluster).Error
}

func (s *ClusterService) Update(cluster *model.Cluster) error {
	return s.db.Save(cluster).Error
}

func (s *ClusterService) Delete(id uint) error {
	return s.db.Delete(&model.Cluster{}, id).Error
}

func (s *ClusterService) Get(id uint) (*model.Cluster, error) {
	var cluster model.Cluster
	return &cluster, s.db.First(&cluster, id).Error
}

func (s *ClusterService) List() ([]model.Cluster, error) {
	var clusters []model.Cluster
	return clusters, s.db.Find(&clusters).Error
}

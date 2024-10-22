package mission

import "github.com/blazeisclone/spaceops-mission-ctrl/domain"

type MissionRepository interface {
	Create(mission *domain.Mission) error
	FindByID(id int) (*domain.Mission, error)
	UpdateByID(id int, mission *domain.Mission) error
	DeleteByID(id int) error
}

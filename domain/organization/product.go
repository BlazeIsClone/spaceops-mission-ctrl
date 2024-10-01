package organization

import (
	"errors"

	"github.com/blazeisclone/spaceops-mission-ctrl/domain"
	"github.com/google/uuid"
)

var (
	ErrMissingValues = errors.New("missing values")
)

type Organization struct {
	item  *domain.Item
	price float64
}

func NewOrganization(name, description string, price float64) (Organization, error) {
	if name == "" || description == "" {
		return Organization{}, ErrMissingValues
	}

	return Organization{
		item: &domain.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price: price,
	}, nil
}

func (p Organization) GetID() uuid.UUID {
	return p.item.ID
}

func (p Organization) GetItem() *domain.Item {
	return p.item
}

func (p Organization) GetPrice() float64 {
	return p.price
}

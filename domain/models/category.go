package models

type CategoryDomain interface {
	GetID() string
	GetName() string
	GetDescription() string
}

type category struct {
	id          string
	name        string
	description string
}

func NewCategoryDomain(id, name, description string) CategoryDomain {
	return &category{
		id:          id,
		name:        name,
		description: description,
	}
}

func (c *category) GetID() string {
	return c.id
}

func (c *category) GetName() string {
	return c.name
}

func (c *category) GetDescription() string {
	return c.description
}

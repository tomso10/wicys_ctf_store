package data

import (
	"context"

	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/item"
)

type item_s struct {
	Client *ent.Client
	Ctx    context.Context
}

func (p *item_s) Create(name, description, image string, price int) (*ent.Item, error) {
	// create item
	return p.Client.Item.Create().
		SetName(name).
		SetDescription(description).
		SetImage(image).
		SetPrice(price).
		Save(p.Ctx)
}

func (p *item_s) Update(name, description, image string, price int) (*ent.Item, error) {
	// get item
	_item, err := p.Client.Item.Query().
		Where(item.Name(name)).
		Only(p.Ctx)
	if err != nil {
		// create item if it doesn't exist
		return p.Create(name, description, image, price)
	}

	// update item
	return _item.Update().
		SetName(name).
		SetDescription(description).
		SetImage(image).
		SetPrice(price).
		Save(p.Ctx)
}

func (p *item_s) Get(name string) (*ent.Item, error) {
	// get item
	return p.Client.Item.Query().
		Where(item.Name(name)).
		Only(p.Ctx)
}

func (p *item_s) GetAll() ([]*ent.Item, error) {
	// get all items
	return p.Client.Item.Query().
		All(p.Ctx)
}

func (p *item_s) Exists(name string) bool {
	// check if item exists
	_, err := p.Get(name)
	return err == nil
}

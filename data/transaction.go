package data

import (
	"context"

	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/item"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/transaction"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/user"
)

type transaction_s struct {
	Client *ent.Client
	Ctx    context.Context
}

func (t *transaction_s) Create(username, itemname, instructions string) (*ent.Transaction, error) {
	// get user
	_user, err := t.Client.User.Query().
		Where(user.Username(username)).
		Only(t.Ctx)
	if err != nil {
		return nil, err
	}

	// get item
	_item, err := t.Client.Item.Query().
		Where(item.Name(itemname)).
		Only(t.Ctx)
	if err != nil {
		return nil, err
	}

	// check if user has enough balance
	if _user.Balance < _item.Price {
		return nil, ErrInsufficientFunds
	}

	// charge user balance
	_, err = t.Client.User.UpdateOne(_user).
		SetBalance(_user.Balance - _item.Price).
		Save(t.Ctx)
	if err != nil {
		return nil, err
	}

	// create transaction
	return t.Client.Transaction.Create().
		SetInstructions(instructions).
		SetPrice(_item.Price).
		SetUser(_user).
		SetItem(_item).
		Save(t.Ctx)
}

func (t *transaction_s) GetAll() ([]*ent.Transaction, error) {
	// get all transactions
	return t.Client.Transaction.Query().
		WithUser().
		WithItem().
		Order(ent.Desc(transaction.FieldTime)).
		All(t.Ctx)
}

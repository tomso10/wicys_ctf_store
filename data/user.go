package data

import (
	"context"
	"fmt"

	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type user_s struct {
	Client *ent.Client
	Ctx    context.Context
}

func (u *user_s) Create(username, password, team string) (*ent.User, error) {
	// check if user exists
	_, err := u.Client.User.Query().
		Where(user.Username(username)).
		Only(u.Ctx)
	if err == nil {
		return nil, ErrUserExists
	}

	// hash password
	hash, err := generateHash(password)
	if err != nil {
		return nil, err
	}
	var perms user.Permissions
	// get permissions
	if team == "BLACK" {
		perms = user.PermissionsBlack
	} else if team == "WHITE" {
		perms = user.PermissionsWhite
	} else if team == "RED" {
		perms = user.PermissionsRed
	} else if team == "BLUE" {
		perms = user.PermissionsBlue
	} else {
		return nil, ErrInvalidTeam
	}

	// create user
	return u.Client.User.Create().
		SetUsername(username).
		SetHash(hash).
		SetPermissions(perms).
		Save(u.Ctx)
}

func (u *user_s) Update(username, password, team string) (*ent.User, error) {
	// get user
	_user, err := u.Get(username)
	if err != nil {
		return u.Create(username, password, team)
	}

	// hash password
	hash, err := generateHash(password)
	if err != nil {
		return nil, err
	}

	var perms user.Permissions

	// get permissions
	if team == "BLACK" {
		perms = user.PermissionsBlack
	} else if team == "WHITE" {
		perms = user.PermissionsWhite
	} else if team == "RED" {
		perms = user.PermissionsRed
	} else if team == "BLUE" {
		perms = user.PermissionsBlue
	} else {
		return nil, ErrInvalidTeam
	}

	// update user
	return u.Client.User.UpdateOne(_user).
		SetHash(hash).
		SetPermissions(perms).
		Save(u.Ctx)
}

func (u *user_s) Get(username string) (*ent.User, error) {
	// get user
	return u.Client.User.Query().
		Where(user.Username(username)).
		Only(u.Ctx)
}

func (u *user_s) Authenticate(username, password string) (*ent.User, error) {
	// get user
	user, err := u.Get(username)
	if err != nil {
		return nil, err
	}

	// check password
	if !checkHash(user.Hash, password) {
		return nil, ErrInvalidPassword
	}

	return user, nil
}

func (u *user_s) SetBalance(username string, balance int) (*ent.User, error) {
	// get user
	user, err := u.Get(username)
	if err != nil {
		return nil, err
	}

	_, err = Transaction.Create(username, "ADMIN ACTION ONLY", fmt.Sprintf("set balance to %d", balance))
	if err != nil {
		return nil, err
	}

	// update balance
	return u.Client.User.UpdateOne(user).
		SetBalance(balance).
		Save(u.Ctx)
}

func (u *user_s) DecrementBalance(username string, amount int) (*ent.User, error) {
	// get user
	user, err := User.Get(username)
	if err != nil {
		return nil, err
	}

	// check balance
	if user.Balance < amount {
		return nil, ErrInsufficientFunds
	}

	_, err = Transaction.Create(username, "ADMIN ACTION ONLY", fmt.Sprintf("decrement balance by %d to %d", amount, user.Balance-amount))
	if err != nil {
		return nil, err
	}

	// update balance
	return u.Client.User.UpdateOne(user).
		SetBalance(user.Balance - amount).
		Save(u.Ctx)
}

func (u *user_s) IncrementBalance(username string, amount int) (*ent.User, error) {
	// get user
	user, err := User.Get(username)
	if err != nil {
		return nil, err
	}

	_, err = Transaction.Create(username, "ADMIN ACTION ONLY", fmt.Sprintf("increment balance by %d to %d", amount, amount+user.Balance))
	if err != nil {
		return nil, err
	}

	// update balance
	return u.Client.User.UpdateOne(user).
		AddBalance(amount).
		Save(u.Ctx)
}

func (u *user_s) GetAll() ([]*ent.User, error) {
	// get all users
	return User.Client.User.Query().
		All(User.Ctx)
}

// helpers

func generateHash(password string) (string, error) {
	// convert to bytes
	passwordBytes := []byte(password)

	// generate hash
	hashBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashBytes), nil
}

func checkHash(passwordHash string, password string) bool {
	// convert to bytes
	passwordHashBytes := []byte(passwordHash)
	passwordBytes := []byte(password)

	// check hash
	return nil == bcrypt.CompareHashAndPassword(passwordHashBytes, passwordBytes)
}

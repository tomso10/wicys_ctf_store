package data

import (
	"context"
	"errors"
	"fmt"

	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/token"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent/user"
)

var (
	ErrInvalidTokenType      error = errors.New("invalid token type")
	ErrTokenAlreadyReedeemed error = errors.New("token already redeemed")
)

type token_s struct {
	Client *ent.Client
	Ctx    context.Context
}

func (t *token_s) Create(value string, _type string) (*ent.Token, error) {
	if _type == "koth" {
		return t.Client.Token.Create().
			SetType(token.TypeKoth).
			SetValue(value).
			Save(t.Ctx)
	} else if _type == "purple" {
		return t.Client.Token.Create().
			SetType(token.TypePurple).
			SetValue(value).
			Save(t.Ctx)
	} else if _type == "sponsor" {
		return t.Client.Token.Create().
			SetType(token.TypeSponsor).
			SetValue(value).
			Save(t.Ctx)
	} else if _type == "ctf" {
		return t.Client.Token.Create().
			SetType(token.TypeCtf).
			SetValue(value).
			Save(t.Ctx)
	} else {
		return nil, ErrInvalidTokenType
	}
}

func (t *token_s) Update(value string, _type string) (*ent.Token, error) {
	// get token
	_token, err := t.Client.Token.Query().
		Where(token.Value(value)).
		Only(t.Ctx)
	if err != nil {
		// create item if it doesn't exist
		return t.Create(value, _type)
	}

	// update item
	if _type == "koth" {
		return _token.Update().
			SetType(token.TypeKoth).
			Save(t.Ctx)
	} else if _type == "sponsor" {
		return _token.Update().
			SetType(token.TypeSponsor).
			Save(t.Ctx)
	} else if _type == "purple" {
		return _token.Update().
			SetType(token.TypePurple).
			Save(t.Ctx)
	} else if _type == "ctf" {
		return _token.Update().
			SetType(token.TypeCtf).
			Save(t.Ctx)
	} else {
		return _token, err
	}

}

func (t *token_s) Get(value string) (*ent.Token, error) {
	return t.Client.Token.Query().
		Where(token.Value(value)).
		WithUser().
		Only(t.Ctx)
}

func (t *token_s) Redeem(username string, value string) (*ent.Token, error) {
	_user, err := User.Get(username)
	if err != nil {
		return nil, err
	}

	_token, err := Token.Get(value)
	if err != nil {
		return _token, err
	}

	if ok, _ := _token.QueryUser().Where(user.Username(_user.Username)).Exist(t.Ctx); ok {
		return _token, ErrTokenAlreadyReedeemed
	}

	_, err = _token.Update().AddUser(_user).Save(t.Ctx)
	if err != nil {
		return _token, err
	}

	if len(_token.Edges.User) == 0 {
		if _token.Type == token.TypeKoth {
			_, err = _user.Update().AddBalance(TokenInfo.KothUnredeemed).Save(t.Ctx)
			if err != nil {
				return _token, err
			}

			_, err = Transaction.Create(
				_user.Username,
				"ADMIN ACTION",
				fmt.Sprintf(
					"Redeemed Unredeemed KoTH Token: %v; increase balance by %v to %v",
					_token.Value,
					TokenInfo.KothUnredeemed,
					_user.Balance+TokenInfo.KothUnredeemed,
				),
			)
			return _token, err
		} else if _token.Type == token.TypePurple {
			_, err = _user.Update().AddBalance(TokenInfo.PurpleUnredeemed).Save(t.Ctx)
			if err != nil {
				return _token, err
			}

			_, err = Transaction.Create(
				_user.Username,
				"ADMIN ACTION",
				fmt.Sprintf(
					"Redeemed Unredeemed Purple Token: %v; increase balance by %v to %v",
					_token.Value,
					TokenInfo.PurpleUnredeemed,
					_user.Balance+TokenInfo.PurpleUnredeemed,
				),
			)
			return _token, err
		} else if _token.Type == token.TypeSponsor {
			_, err = _user.Update().AddBalance(TokenInfo.SponsorUnredeemed).Save(t.Ctx)
			if err != nil {
				return _token, err
			}

			_, err = Transaction.Create(
				_user.Username,
				"ADMIN ACTION",
				fmt.Sprintf(
					"Redeemed Unredeemed Sponsor Token: %v; increase balance by %v to %v",
					_token.Value,
					TokenInfo.SponsorUnredeemed,
					_user.Balance+TokenInfo.SponsorUnredeemed,
				),
			)
			return _token, err
		} else if _token.Type == token.TypeCtf {
			_, err = _user.Update().AddBalance(TokenInfo.CTFUnredeemed).Save(t.Ctx)
			if err != nil {
				return _token, err
			}

			_, err = Transaction.Create(
				_user.Username,
				"ADMIN ACTION",
				fmt.Sprintf(
					"Redeemed Unredeemed CTF Token: %v; increase balance by %v to %v",
					_token.Value,
					TokenInfo.CTFUnredeemed,
					_user.Balance+TokenInfo.CTFUnredeemed,
				),
			)
			return _token, err
		} else {
			return _token, ErrInvalidTokenType
		}
	} else {
		if _token.Type == token.TypeKoth {
			_, err = _user.Update().AddBalance(TokenInfo.KothRedeemed).Save(t.Ctx)
			if err != nil {
				return _token, err
			}

			_, err = Transaction.Create(
				_user.Username,
				"ADMIN ACTION",
				fmt.Sprintf(
					"Redeemed Redeemed KoTH Token: %v; increase balance by %v to %v",
					_token.Value,
					TokenInfo.KothRedeemed,
					_user.Balance+TokenInfo.KothRedeemed,
				),
			)
			return _token, err
		} else if _token.Type == token.TypePurple {
			_, err = _user.Update().AddBalance(TokenInfo.PurpleRedeemed).Save(t.Ctx)
			if err != nil {
				return _token, err
			}

			_, err = Transaction.Create(
				_user.Username,
				"ADMIN ACTION",
				fmt.Sprintf(
					"Redeemed Redeemed Purple Token: %v; increase balance by %v to %v",
					_token.Value,
					TokenInfo.PurpleRedeemed,
					_user.Balance+TokenInfo.PurpleRedeemed,
				),
			)
			return _token, err
		} else if _token.Type == token.TypeSponsor {
			_, err = _user.Update().AddBalance(TokenInfo.SponsorRedeemed).Save(t.Ctx)
			if err != nil {
				return _token, err
			}

			_, err = Transaction.Create(
				_user.Username,
				"ADMIN ACTION",
				fmt.Sprintf(
					"Redeemed Redeemed Sponsor Token: %v; increase balance by %v to %v",
					_token.Value,
					TokenInfo.SponsorRedeemed,
					_user.Balance+TokenInfo.SponsorRedeemed,
				),
			)
			return _token, err
		} else if _token.Type == token.TypeCtf {
			_, err = _user.Update().AddBalance(TokenInfo.CTFRedeemed).Save(t.Ctx)
			if err != nil {
				return _token, err
			}

			_, err = Transaction.Create(
				_user.Username,
				"ADMIN ACTION",
				fmt.Sprintf(
					"Redeemed Redeemed CTF Token: %v; increase balance by %v to %v",
					_token.Value,
					TokenInfo.CTFRedeemed,
					_user.Balance+TokenInfo.CTFRedeemed,
				),
			)
			return _token, err
		} else {
			return _token, ErrInvalidTokenType
		}
	}
}

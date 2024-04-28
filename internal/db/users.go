package db

import (
	"errors"
	"storiesservice/internal/consts"
	"storiesservice/internal/sessions"
	"storiesservice/pkg/logger"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	ById       map[int64]*User  `json:"byId"`
	ByUsername map[string]*User `json:"byUsername"`
}

func NewUsers() *Users {
	return &Users{
		ById:       map[int64]*User{},
		ByUsername: map[string]*User{},
	}
}

func (d *Database) GetUsers() *Users {
	users, ok := d.UsersCache.Modify(func() interface{} {
		rows, err := d.db.Query(`
			SELECT u.id, u.username, u.password
			FROM users u
		`)
		if err != nil {
			logger.Eprintln(err)
			return nil
		}

		users := NewUsers()
		for rows.Next() {
			user := &User{}
			if err := rows.Scan(
				&user.Id,
				&user.Username,
				&user.Password,
			); err != nil {
				logger.Eprintln(err)
				return nil
			}
			users.ById[user.Id] = user
			users.ByUsername[user.Username] = user
		}
		return users
	}).(*Users)

	if !ok || users == nil {
		return NewUsers()
	}

	return users
}

func (d *Database) SignIn(username, password string) (session string, err error) {
	users := d.GetUsers()
	if users == nil {
		err = errors.New(consts.I18nGetUsersFailed)
		return
	}

	user, ok := users.ByUsername[username]
	if !ok {
		err = errors.New(consts.I18nUserNotFound)
		return
	}

	if user.Password != password {
		err = errors.New(consts.I18nWrongPassword)
		return
	}

	session = sessions.New(user.Id)
	return
}

package data

import (
	"farm.e-pedion.com/repo/cheetah/logger"
)

type mongoClient struct {
}

type mysqlClient struct {
}

type cassandraClient struct {
}

type Fetchable interface {
	Fetch(interface{}) error
}

// Client is the CRUD persistence interface
type Client interface {
	Persist(interface{}) error
	Query(string, Fetchable) error
	Remove(interface{}) error
}

func NewUserManager() *UserManager {
	return &UserManager{
		logger: logger.GetLogger(),
	}
}

type UserManager struct {
	logger logger.Logger
}

func (u *UserManager) Save(user User) error {
	u.logger.Info("UserSaved",
		logger.String("id", user.ID),
		logger.String("username", user.Username),
		logger.String("email", user.Email),
		logger.Int64("cadunId", user.CadunID),
	)
	return nil
}

func (u *UserManager) Read(id string) error {
	var user User
	u.logger.Info("UserRead",
		logger.String("providedId", id),
		logger.String("id", user.ID),
		logger.String("username", user.Username),
		logger.String("email", user.Email),
		logger.Int64("cadunId", user.CadunID),
	)
	return nil
}

func (u *UserManager) Delete(id string) error {
	u.logger.Info("UserDeleted",
		logger.String("id", id),
	)
	return nil
}

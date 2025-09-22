package fakedb

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

type FakeDb struct {
	Users       map[int]*models.User
	userCounter int
}

func NewFakeDb() *FakeDb {
	return &FakeDb{
		Users:       map[int]*models.User{},
		userCounter: 0,
	}
}

func (fdb *FakeDb) SaveUser(username, password string) (int, error) {

	for _, user := range fdb.Users {
		if user.Username == username {
			return 0, repositories.ErrorDuplicateUser
		}
	}

	fdb.Users[fdb.userCounter] = &models.User{Username: username, Password: password, Id: fdb.userCounter}
	fdb.userCounter++
	return fdb.userCounter - 1, nil
}

func (fdb *FakeDb) GetUserById(id int) (*models.User, error) {
	if _, find := fdb.Users[id]; !find {
		return nil, repositories.ErrorInvalidID
	}
	return fdb.Users[id], nil
}

func (fdb *FakeDb) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	for _, user := range fdb.Users {
		users = append(users, user)
	}
	
	if len(users) == 0 {
		return nil,repositories.ErrorNoUser
	}

	return users,nil
}

func (fdb *FakeDb) DeleteUser(id int) error {
	if _, find := fdb.Users[id]; !find {
		return repositories.ErrorInvalidID
	}

	delete(fdb.Users, id)
	return nil
}

func (fdb *FakeDb) UpdateUser(new *models.User) error {
	if _, find := fdb.Users[new.Id]; !find {
		return repositories.ErrorInvalidID
	}

	fdb.Users[new.Id] = new
	return nil
}

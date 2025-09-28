package repositories

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
)

type User interface {
	SaveUser(username, password string) (int, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUserName(username string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	DeleteUser(id int) error
	UpdateUser(new *models.User) error
}

type userRepo struct {
	db *fakedb.FakeDb
}

func NewUserRepo(db *fakedb.FakeDb) User {
	return &userRepo{
		db: db,
	}
}

func (fdb *userRepo) SaveUser(username, password string) (int, error) {
	fdb.db.UserMu.Lock()
	defer fdb.db.UserMu.Unlock()
	for _, user := range fdb.db.Users {
		if user.Username == username {
			return 0, ErrorDuplicateUser
		}
	}

	fdb.db.Users[fdb.db.UserCounter] = &models.User{Username: username, Password: password, Id: fdb.db.UserCounter, AutoSave: true, Role: models.RoleUser}
	fdb.db.UserCounter++
	return fdb.db.UserCounter - 1, nil
}

func (fdb *userRepo) GetUserById(id int) (*models.User, error) {
	fdb.db.UserMu.RLock()
	defer fdb.db.UserMu.RUnlock()
	if _, find := fdb.db.Users[id]; !find {
		return nil, ErrorInvalidUID
	}
	return fdb.db.Users[id], nil
}

func (fdb *userRepo) GetUserByUserName(username string) (*models.User, error) {
	fdb.db.UserMu.RLock()
	defer fdb.db.UserMu.RUnlock()
	for _, user := range fdb.db.Users {
		if user.Username == username {
			return user, nil
		}
		return nil, ErrorInvalidUName
	}
	return nil, ErrorNoUser
}

func (fdb *userRepo) GetAllUsers() ([]*models.User, error) {
	fdb.db.UserMu.RLock()
	defer fdb.db.UserMu.RUnlock()
	users := []*models.User{}
	for _, user := range fdb.db.Users {
		users = append(users, user)
	}

	// if len(users) == 0 {
	// 	return nil, ErrorNoUser
	// }

	return users, nil
}

func (fdb *userRepo) DeleteUser(id int) error {
	fdb.db.UserMu.Lock()
	defer fdb.db.UserMu.Unlock()
	if _, find := fdb.db.Users[id]; !find {
		return ErrorInvalidUID
	}

	delete(fdb.db.Users, id)
	return nil
}

func (fdb *userRepo) UpdateUser(new *models.User) error {
	fdb.db.UserMu.Lock()
	defer fdb.db.UserMu.Unlock()
	if _, find := fdb.db.Users[new.Id]; !find {
		return ErrorInvalidUID
	}

	fdb.db.Users[new.Id] = new
	return nil
}

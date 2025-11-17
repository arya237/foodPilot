package memory

import (
	"github.com/arya237/foodPilot/internal/db"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

type userRepo struct {
	db *db.FakeDb
}

func NewUserRepo(db *db.FakeDb) repositories.User {
	return &userRepo{
		db: db,
	}
}

func (fdb *userRepo) Save(newUser *models.User) (*models.User, error) {
	fdb.db.UserMu.Lock()
	defer fdb.db.UserMu.Unlock()
	for _, user := range fdb.db.Users {
		if user.Username == newUser.Username {
			return nil, repositories.ErrorDuplicateUser
		}
	}

	newUser.Id = fdb.db.UserCounter
	fdb.db.Users[fdb.db.UserCounter] = newUser
	fdb.db.UserCounter++

	return newUser, nil
}

func (fdb *userRepo) GetById(id int) (*models.User, error) {
	fdb.db.UserMu.RLock()
	defer fdb.db.UserMu.RUnlock()
	if _, find := fdb.db.Users[id]; !find {
		return nil, repositories.ErrorInvalidUID
	}
	return fdb.db.Users[id], nil
}

func (fdb *userRepo) GetByUserName(username string) (*models.User, error) {
	fdb.db.UserMu.RLock()
	defer fdb.db.UserMu.RUnlock()
	for _, user := range fdb.db.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, repositories.ErrorInvalidUName
}

func (fdb *userRepo) GetAll() ([]*models.User, error) {
	fdb.db.UserMu.RLock()
	defer fdb.db.UserMu.RUnlock()
	users := []*models.User{}
	for _, user := range fdb.db.Users {
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, repositories.ErrorNoUser
	}

	return users, nil
}

func (fdb *userRepo) Delete(id int) error {
	fdb.db.UserMu.Lock()
	defer fdb.db.UserMu.Unlock()
	if _, find := fdb.db.Users[id]; !find {
		return repositories.ErrorInvalidUID
	}

	delete(fdb.db.Users, id)
	return nil
}

func (fdb *userRepo) Update(new *models.User) error {
	fdb.db.UserMu.Lock()
	defer fdb.db.UserMu.Unlock()
	if _, find := fdb.db.Users[new.Id]; !find {
		return repositories.ErrorInvalidUID
	}

	fdb.db.Users[new.Id] = new
	return nil
}

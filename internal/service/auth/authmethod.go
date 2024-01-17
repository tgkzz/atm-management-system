package auth

import (
	"atm/internal/models"
)

func (a AuthService) CreateNewUser(user models.User) (err error) {
	if err = validateUserData(user); err != nil {
		return err
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return err
	}

	return a.repo.CreateUser(user)
}

func (a AuthService) GetUserByUsername(username string) (models.User, error) {
	return a.repo.GetUserByUsername(username)
}

func (a AuthService) DeleteUserByUsername(username string) error {
	return a.repo.DeleteUserByUsername(username)
}

func (a AuthService) CheckUserCreds(creds models.User) error {
	user, err := a.repo.GetUserByUsername(creds.Username)
	if err != nil {
		return err
	}

	if !checkPasswordHash(creds.Password, user.Password) {
		return models.ErrIncorrectPassword
	}

	return nil

}

package security

import (
	"chatnews-api/middleware/model"
	"chatnews-api/middleware/repository"
	"chatnews-api/middleware/util"
)

type AuthValidator struct {
	userRepository repository.UserRepository
}

func NewAuthValidator(userRepository repository.UserRepository) *AuthValidator {
	return &AuthValidator{userRepository: userRepository}
}

func (authValidator *AuthValidator) ValidateCredentials(username, password string) (*model.User, bool) {
	user, err := authValidator.userRepository.FindByEmail(username)
	if err != nil || util.VerifyPassword(user.Password, password) != nil {
		return nil, false
	}
	return user, true
}

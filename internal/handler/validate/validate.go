package validate

import (
	"errors"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

func ValidateUser(user domain.User) error {

	if user.Username == "" {
		return errors.New("username is required")
	}

	return nil
}

package validate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
)

var (
	ErrEmptyField           = errors.New("поле не может быть пустым")
	ErrInvalidPhone         = errors.New("неверный формат телефона")
	ErrInvalidSex           = errors.New("пол должен быть 'М' или 'Ж'")
	ErrInvalidSportCategory = errors.New("неверная категория спорта")
	ErrInvalidPassword      = errors.New("пароль должен содержать минимум 8 символов, включая буквы и цифры")
	ErrInvalidUsername      = errors.New("логин должен содержать только буквы, цифры и знак подчеркивания, длиной от 3 до 20 символов")
)

func ValidateUser(user domain.User) error {
	var errs []error

	// Проверка фамилии
	if strings.TrimSpace(user.Surname) == "" {
		errs = append(errs, fmt.Errorf("фамилия: %w", ErrEmptyField))
	}

	// Проверка имени
	if strings.TrimSpace(user.Name) == "" {
		errs = append(errs, fmt.Errorf("имя: %w", ErrEmptyField))
	}

	// Проверка адреса
	if strings.TrimSpace(user.Address) == "" {
		errs = append(errs, fmt.Errorf("адрес: %w", ErrEmptyField))
	}

	// Проверка телефона
	phoneRegex := regexp.MustCompile(`^\+?[78][-\(]?\d{3}\)?-?\d{3}-?\d{2}-?\d{2}$`)
	if !phoneRegex.MatchString(user.Phone) {
		errs = append(errs, fmt.Errorf("телефон: %w", ErrInvalidPhone))
	}

	// Проверка логина
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	if !usernameRegex.MatchString(user.Username) {
		errs = append(errs, fmt.Errorf("логин: %w", ErrInvalidUsername))
	}
	if strings.TrimSpace(user.Username) == "" {
		errs = append(errs, fmt.Errorf("логин: %w", ErrEmptyField))
	}
	// Проверка пароля
	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(user.Password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(user.Password)
	hasMinLength := len(user.Password) >= 8
	if !(hasLetter && hasDigit && hasMinLength) {
		errs = append(errs, fmt.Errorf("пароль: %w", ErrInvalidPassword))
	}
	if strings.TrimSpace(user.Password) == "" {
		errs = append(errs, fmt.Errorf("пароль: %w", ErrEmptyField))
	}

	if len(errs) > 0 {
		var errMsgs []string
		for _, err := range errs {
			errMsgs = append(errMsgs, err.Error())
		}
		return fmt.Errorf(strings.Join(errMsgs, "\n"))
	}

	return nil
}

package validation

import (
    "errors"
    "regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)

func Email(email string) error {
    if !emailRegex.MatchString(email) {
        return errors.New("invalid email format")
    }
    return nil
}

var phoneRegex = regexp.MustCompile(`^[0-9]+$`)

func Phone(phone string) error {
    if len(phone) != 11 {
        return errors.New("phone must be between 11 digits")
    }

	if !phoneRegex.MatchString(phone) {
        return errors.New("invalid phone format")
    }

    return nil
}
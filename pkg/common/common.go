package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/erik-sostenes/receipt-processor-api/pkg/wrongs"
	"github.com/google/uuid"
)

type Map map[string]any

// GetEnv method that reads the environment variables needed in the project.
//
// Note: if an environment variable is not found, a panic will occur.
func GetEnv(key string) string {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		panic(fmt.Sprintf("missing environment variable '%s'", key))
	}
	return value
}

// GenerateUuID generate a new UuID.
func GenerateUuID() string {
	return uuid.New().String()
}

// ParseUuID validate if the format the values is a UuID
func ParseUuID(value string) (uuid.UUID, error) {
	return uuid.Parse(value)
}

// Identifier receives a value to verify if the format is correct
type Identifier string

// Validate method validates if the value is an Uuid, if incorrect returns an wrongs.StatusUnprocessableEntity
func (i Identifier) Validate() (string, error) {
	u, err := ParseUuID(string(i))
	if err != nil {
		return u.String(), wrongs.StatusBadRequest(fmt.Sprintf("incorrect %s uuid unique identifier, must be a uuid value", string(i)))
	}
	return u.String(), nil
}

// Timestamp receives a value to verify if the format is correct
type Timestamp string

// Validate method validates if the value is a time.Time, if incorrect returns an wrongs.StatusUnprocessableEntity
func (t Timestamp) Validate(layout string) (int64, error) {
	v, err := time.Parse(layout, string(t))
	if err != nil {
		return 0, wrongs.StatusBadRequest(fmt.Sprintf("incorrect %s value format, must be a time value", string(t)))
	}
	return v.Unix(), nil
}

// String receives a value to verify if the format is correct
type String string

// The Validate method validates if the value is a string and is not empty, if incorrect returns an wrongs.StatusUnprocessableEntity
func (s String) Validate(fieldName string) (string, error) {
	if strings.TrimSpace(string(s)) == "" {
		return "", wrongs.StatusBadRequest(fmt.Sprintf("the %s field is missing", fieldName))
	}
	return string(s), nil
}

// Float receives a value to verify if the format is correct
type Float string

// Validate method validates if the value is a float, if incorrect returns an errors.StatusUnprocessableEntity
func (f Float) Validate() (float64, error) {
	v, err := strconv.ParseFloat(string(f), 64)
	if err != nil {
		return v, wrongs.StatusBadRequest(fmt.Sprintf("incorrect %s value format, must be a numeric value", string(f)))
	}
	return v, nil
}

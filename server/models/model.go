package models

import (
    "time"
    "math/rand"

    "gopkg.in/bluesuncorp/validator.v5"

    "github.com/satori/go.uuid"
)

// Model is the base model/struct for any model in the application/system
type Model struct {
    ID uint                     `gorm:"primary_key" json:"-"`
    CreatedAt time.Time         `gorm:"not null" json:"-"`
    UpdatedAt time.Time         `json:"-"`
}

const (
    letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// GenerateRandomString returns a random string with `n` as the length
func GenerateRandomString(n int) string {
    b := make([]byte, n)
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}

func generateUUID() string {
    return uuid.NewV4().String()
}

func validateModel(tagName string, model interface{}) error {
    validate := validator.New(tagName, validator.BakedInValidators)
    err := validate.Struct(model)
    if err != nil {
        return err
    }
    return nil
}

// IsValid checks if a `model` entry is valid, given the `tagName` (scope) for validation
func IsValid(tagName string, model interface{}) bool {
    err := validateModel(tagName, model)
    if err != nil {
        return false
    }
    return true
}

package models

import (
    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
)

// Client is the client application model/struct
type Client struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"id"`
    Key string                  `gorm:"not null;unique;index" json:"-"`
    Secret string               `gorm:"not null" validate:"required" json:"-"`
}

// Authentic checks if a secret is valid for a given Client
func (client *Client) Authentic(secret string) bool {
    validSecret := bcrypt.CompareHashAndPassword([]byte(client.Secret), []byte(secret)) == nil
    return validSecret
}

// UpdateSecret updates an Client's secret
func (client *Client) UpdateSecret(secret string) error {
    crypted, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
    if err == nil {
        client.Secret = string(crypted)
        return nil
    }
    return err
}

// BeforeSave Client model/struct hook
func (client *Client) BeforeSave(scope *gorm.Scope) error {
    return validateModel("validate", client)
}

// BeforeCreate Client model/struct hook
func (client *Client) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("UUID", generateUUID())
    scope.SetColumn("Key", GenerateRandomString(32))
    if crypted, err := bcrypt.GenerateFromPassword([]byte(client.Secret), bcrypt.DefaultCost); err == nil {
        scope.SetColumn("Secret", crypted)
    } else {
        return err
    }
    return nil
}

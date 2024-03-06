package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserID               uint   `json:"user_id"`
	Username             string `json:"username,omitempty"`
	Password             string `json:"password,omitempty"`
	PasswordResetToken   string `json:"password_reset_token,omitempty"`
	PasswordResetExpires string `json:"password_reset_token_expiry,omitempty"`
	Email                string `json:"email,omitempty"`
	FirstName            string `json:"first_name,omitempty"`
	LastName             string `json:"last_name,omitempty"`
	DOB                  string `json:"dob,omitempty"`
	Type                 string `json:"type,omitempty"`
}

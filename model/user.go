package model

import "time"

type User struct {
	ID            uint64     `db:"id" json:"id"`
	UID           string     `db:"uid" json:"uid"`
	CardUID       string     `db:"carduid" json:"carduid"`
	UIDCreator    string     `db:"uidCreator" json:"uidCreator"`
	Subscriber    string     `db:"subscriber" json:"subscriber"`
	PhotoURL      string     `db:"photo_url" json:"photo_url"`
	Weight        int        `db:"weight" json:"weight"`
	Name          string     `db:"name" json:"name"`
	Email         string     `db:"email" json:"email"`
	Password      string     `db:"password" json:"password"`
	Ccode         string     `db:"Ccode" json:"Ccode"`
	Phone         string     `db:"phone" json:"phone"`
	PhoneNumber   string     `db:"PhoneNumber" json:"phone_number"`
	Platform      string     `db:"platform" json:"platform"`
	Status        string     `db:"status" json:"status"`
	Gender        string     `db:"gender" json:"gender"`
	Age           string     `db:"age" json:"age"`
	InitCountry   string     `db:"initCountry" json:"init_country"`
	Country       string     `db:"country" json:"country"`
	MaritalStatus string     `db:"marital_status" json:"marital_status"`
	RememberToken *string    `db:"remember_token" json:"remember_token,omitempty"`
	CreatedAt     *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt     *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	UserUpdatedAt *time.Time `db:"userUpdated_at" json:"user_updated_at,omitempty"`
	PackUID       string     `db:"packUid" json:"pack_uid"`
	PackType      string     `db:"packType" json:"pack_type"`
	StatusValid   string     `db:"statusValid" json:"status_valid"`
}

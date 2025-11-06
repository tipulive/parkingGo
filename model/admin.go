package model

import (
	"time"
)

type Admin struct {
	ID              uint64     `db:"id" json:"id" gorm:"primaryKey;autoIncrement"`
	UID             string     `db:"uid" json:"uid"`
	PermissionValue string     `db:"permissionValue" json:"permissionValue"`
	PermissionJSON  string     `db:"permissionJson" json:"permissionJson"` // stores JSON data
	UIDCreator      string     `db:"uidCreator" json:"uidCreator"`
	CompanyName     string     `db:"CompanyName" json:"companyName"`
	PhotoURL        string     `db:"photo_url" json:"photoUrl"`
	Name            string     `db:"name" json:"name"`
	Email           string     `db:"email" json:"email"`
	Password        string     `db:"password" json:"password"`
	Ccode           string     `db:"Ccode" json:"ccode"`
	Phone           string     `db:"phone" json:"phone"`
	PhoneNumber     string     `db:"PhoneNumber" json:"phoneNumber"`
	InitCountry     string     `db:"initCountry" json:"initCountry"`
	Platform        string     `db:"platform" json:"platform"`
	Status          string     `db:"status" json:"status"`
	Subscriber      string     `db:"subscriber" json:"subscriber"`
	Country         string     `db:"country" json:"country"`
	CreatedAt       *time.Time `db:"created_at" json:"createdAt,omitempty"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updatedAt,omitempty"`
}

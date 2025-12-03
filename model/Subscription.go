package model

import "time"

type Subscription struct {
	ID          uint64     `json:"id" db:"id"` // bigint unsigned, primary key
	UID         string     `json:"uid" db:"uid"`
	InputAction string     `json:"input_action"`                         // varchar(191), indexed
	UIDUser     string     `json:"uid_user" db:"uidUser"`                // varchar(191), indexed
	UIDCreator  string     `json:"uid_creator" db:"uidCreator"`          // varchar(191), indexed
	Subscriber  string     `json:"subscriber" db:"subscriber"`           // varchar(191), indexed
	PackName    string     `json:"pack_name" db:"packName"`              // varchar(191)
	PackUID     string     `json:"pack_uid" db:"packUid"`                // varchar(191), indexed
	PackPrice   string     `json:"pack_price" db:"packPrice"`            // varchar(191), default 0
	PackValid   string     `json:"pack_valid" db:"packValid"`            // varchar(191)
	PackType    string     `json:"pack_type" db:"packType"`              // varchar(191)
	Activated   string     `json:"activated" db:"activated"`             // varchar(60), default "on"
	UserStatus  string     `json:"user_status" db:"userStatus"`          // varchar(191), indexed
	PaidStatus  string     `json:"paid_status" db:"paidStatus"`          // varchar(191), indexed
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"created_at"` // timestamp, nullable
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at"` // timestamp, nullable
}

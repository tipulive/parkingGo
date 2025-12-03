package model

import "time"

type Package struct {
	ID                  uint64     `db:"id" json:"id"`                                   // bigint(20) unsigned, primary key
	Name                string     `db:"name" json:"name" validate:"required,min=3"`     // varchar(191), not null
	UID                 string     `db:"uid" json:"uid"`                                 // varchar(191), not null
	PackValid           string     `db:"packValid" json:"packValid"`                     // varchar(191), not null
	PackCover           int        `db:"packCover" json:"packCover" validate:"required"` // int(60), not null
	Status              string     `db:"status" json:"status"`                           // varchar(191), not null
	PackEligible        string     `db:"packEligible" json:"packEligible"`               // varchar(20), not null
	PackType            string     `db:"packType" json:"packType"`                       // varchar(191), not null
	PackDetail          string     `db:"packDetail" json:"packDetail"`                   // varchar(191), not null
	PackJson            *string    `db:"packJson" json:"packJson"`                       // longtext, nullable
	Price               string     `db:"price" json:"price"`                             // varchar(191), not null
	UIDCreator          string     `db:"uidCreator" json:"uidCreator"`                   // varchar(191), not null
	Subscriber          string     `db:"subscriber" json:"subscriber"`                   // varchar(191), not null
	CommentData         *string    `db:"commentData" json:"commentData,omitempty"`       // varchar(191), nullable
	CreatedAt           *time.Time `db:"created_at" json:"created_at,omitempty"`         // timestamp, nullable
	UpdatedAt           *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	CostPerHour         float64    `db:"cost_per_hour"`
	PackagesNeeded      int        `db:"packages_needed"`
	TotalCost           float64    `db:"total_cost"`
	EffectiveHourlyRate float64    `db:"effective_hourly_rate"`
	TotalHoursPurchased float64    `db:"total_hours_purchased"` // Change to float64        // timestamp, nullable
}

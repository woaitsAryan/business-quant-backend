package models

import (
	"time"

	"github.com/google/uuid"
)

type Stock struct {
    ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
    Ticker    string    `gorm:"type:char(10);not null;index" json:"ticker"`
    Date      time.Time `gorm:"type:date;not null" json:"date"`
    Revenue   *int64    `gorm:"type:bigint" json:"revenue"`
    GP        *int64    `gorm:"type:bigint" json:"gp"`
    FCF       *int64    `gorm:"type:bigint" json:"fcf"`
    Capex     *int64    `gorm:"type:bigint" json:"capex"`
}

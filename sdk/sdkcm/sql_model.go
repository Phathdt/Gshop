package sdkcm

import "time"

type SQLModel struct {
	ID        uint32     `json:"id" gorm:"id,PRIMARY_KEY"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

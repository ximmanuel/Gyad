package models

import "time"

// Example model for a Bober
type Bober struct {
	ID        string    `xorm:"pk uuid 'id'"`
	Name      string    `xorm:"varchar(255) 'name'"`
	Age       int       `xorm:"int 'age'"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

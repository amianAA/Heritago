package models

// User defines a user for the app
type User struct {
	BaseModelSoftDelete         // We don't to actually delete the users, maybe audit
	Email               string  `gorm:"not null;unique_index:idx_email"`
	UserID              *string // External user ID
	FirstName           *string
	LastName            *string
	Country             *string
	Password            *string
}

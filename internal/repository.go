// Handles data access operations

package internal

import "gorm.io/gorm"

// Repository provides data access methods for the Message model.
type Repository struct {
	db *gorm.DB // The GORM database connection
}

// NewRepository returns a new instance of Repository with the given DB.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// GetHelloMessage fetches the message with ID 1 from the database.
// Returns the text if found, or an error if not.
func (r *Repository) GetHelloMessage() (string, error) {
	var msg Message
	// Look up the message by primary key ID = 1
	if err := r.db.First(&msg, 1).Error; err != nil {
		return "", err // return error if query fails
	}
	return msg.Text, nil // return the message text if found
}

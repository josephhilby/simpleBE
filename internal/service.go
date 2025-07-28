// Intermediary between handlers and repositories

package internal

import "gorm.io/gorm"

// Service provides business logic using the underlying repository.
type Service struct {
	repo *Repository // Reference to the data access layer
}

// NewService constructs a new Service using a GORM database connection.
// It internally creates a Repository instance.
func NewService(db *gorm.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

// GetHelloMessage retrieves the message text via the repository.
// It acts as a pass-through for now but could include additional logic.
func (s *Service) GetHelloMessage() (string, error) {
	return s.repo.GetHelloMessage()
}

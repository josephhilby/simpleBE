// GORM models, Defines DB tables
// go get gorm.io/gorm
// go get gorm.io/driver/postgres
// go mod tidy

package internal

// Message represents a row in the 'messages' table.
// It includes an ID as the primary key and a text message.
type Message struct {
	ID   uint   `gorm:"primaryKey"` // ID is the primary key (auto-incrementing uint)
	Text string `gorm:"not null"`   // Text is the message content; cannot be null
}

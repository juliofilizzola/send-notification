package campaign

import (
	"time"
)

type Contact struct {
	Email string `json:"email"`
}

type Campaign struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Contact   []Contact `json:"contact"`
}

func NewCampaign(name, content string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}
	return &Campaign{
		ID:        "1",
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contact:   contacts,
	}
}
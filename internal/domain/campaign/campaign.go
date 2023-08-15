package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
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

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(emails) == 0 {
		return nil, errors.New("at least one email has to be provided")
	}
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}
	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contact:   contacts,
	}, nil
}

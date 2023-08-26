package campaign

import (
	"time"

	"github.com/juliofilizzola/send-notification/internal/internalErrors"
	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"min=5,max=1024"`
	CreatedAt time.Time `validate:"required"`
	Contact   []Contact `validate:"min=1,dive"`
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contact:   contacts,
	}
	err := internalErrors.ValidatorStruct(campaign)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}

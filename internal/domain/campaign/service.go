package campaign

import (
	"errors"

	"github.com/juliofilizzola/send-notification/internal/dto"
	"github.com/juliofilizzola/send-notification/internal/internalErrors"
)

type Service struct {
	Repository Repository
}

func (s Service) Create(newCampaign dto.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Contact)
	if err != nil {
		return "", errors.New(err.Error())
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalErrors.ErrInternal
	}
	return campaign.ID, nil
}

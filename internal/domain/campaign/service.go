package campaign

import (
	"errors"

	"github.com/juliofilizzola/send-notification/internal/dto"
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
		return "", err
	}
	return campaign.ID, nil
}

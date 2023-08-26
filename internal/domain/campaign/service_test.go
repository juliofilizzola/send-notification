package campaign

import (
	"testing"

	"github.com/juliofilizzola/send-notification/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

var (
	newCampaign = dto.NewCampaign{
		Content: "this is test service",
		Contact: []string{"email@test.com", "email2@test.com"},
		Name:    "Test Service",
	}
	service = Service{}
)

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func TestService_SaveCampaign(t *testing.T) {
	assertions := assert.New(t)
	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		} else if campaign.Content != newCampaign.Content {
			return false
		} else if len(campaign.Contact) != len(newCampaign.Contact) {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)
	assertions.Nil(err)
	assertions.NotNil(id)

}

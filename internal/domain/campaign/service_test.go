package campaign

import (
	"testing"

	"github.com/juliofilizzola/send-notification/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestService_Create(t *testing.T) {
	assertions := assert.New(t)
	service := Service{}

	newCampaign := dto.NewCampaign{
		Content: "this is test service",
		Contact: []string{"email@test.com", "email2@test.com"},
		Name:    "Test Service",
	}

	id, err := service.Create(newCampaign)
	assertions.Nil(err)
	assertions.NotNil(id)
}

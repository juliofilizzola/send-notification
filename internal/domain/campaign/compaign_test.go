package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "campaign 1"
	content  = "body..."
	contacts = []string{"email@test.com", "email2@test.com"}
)

func Test_NewCampaign_Create(t *testing.T) {
	asserts := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	asserts.Equal(campaign.Name, name)
	asserts.Equal(campaign.Content, content)
	asserts.Equal(len(campaign.Contact), len(contacts))
}

func Test_NewCampaign_IdOnNotNil(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnNow(t *testing.T) {
	assertions := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)
	now := time.Now().Add(-time.Minute)

	assertions.NotNil(campaign.CreatedAt)
	assertions.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assertions := assert.New(t)
	_, err := NewCampaign("", content, contacts)

	assertions.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assertions := assert.New(t)
	_, err := NewCampaign(name, "", contacts)

	assertions.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateEmails(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assertions.Equal("at least one email has to be provided", err.Error())
}

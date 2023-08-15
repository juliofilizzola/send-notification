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

	campaign := NewCampaign(name, content, contacts)

	asserts.Equal(campaign.Name, name)
	asserts.Equal(campaign.Content, content)
	asserts.Equal(len(campaign.Contact), len(contacts))
}

func Test_NewCampaign_IdOnNotNil(t *testing.T) {
	assertions := assert.New(t)

	campaign := NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnNow(t *testing.T) {
	assertions := assert.New(t)
	campaign := NewCampaign(name, content, contacts)
	now := time.Now().Add(-time.Minute)

	assertions.NotNil(campaign.CreatedAt)
	assertions.Greater(campaign.CreatedAt, now)
}

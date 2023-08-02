package campaign

import (
	"testing"
)

func TestNewCampaign(t *testing.T) {
	name := "campaign 1"
	content := "body..."
	contacts := []string{"email@test.com", "email2@test.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("expect 1")
	} else if campaign.Name != name {
		t.Errorf("expect correct name")
	} else if campaign.Content != content {
		t.Errorf("expect correct content")
	} else if len(campaign.Contact) != len(contacts) {
		t.Errorf("error...")
	}
}

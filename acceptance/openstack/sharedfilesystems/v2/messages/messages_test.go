package messages

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-go/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-go/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-go/openstack/sharedfilesystems/v2/messages"
)

const requestID = "req-6f52cd8b-25a1-42cf-b497-7babf70f55f4"
const minimumManilaMessagesMicroVersion = "2.37"

func TestMessageList(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")

	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create a shared file system client: %v", err)
	}
	client.Microversion = minimumManilaMessagesMicroVersion

	allPages, err := messages.List(client, messages.ListOpts{}).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve messages: %v", err)
	}

	allMessages, err := messages.ExtractMessages(allPages)
	if err != nil {
		t.Fatalf("Unable to extract messages: %v", err)
	}

	for _, message := range allMessages {
		tools.PrintResource(t, message)
	}
}

// The test creates 2 messages and verifies that only the one(s) with
// a particular name are being listed
func TestMessageListFiltering(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")

	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create a shared file system client: %v", err)
	}
	client.Microversion = minimumManilaMessagesMicroVersion

	options := messages.ListOpts{
		RequestID: requestID,
	}

	allPages, err := messages.List(client, options).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve messages: %v", err)
	}

	allMessages, err := messages.ExtractMessages(allPages)
	if err != nil {
		t.Fatalf("Unable to extract messages: %v", err)
	}

	for _, listedMessage := range allMessages {
		if listedMessage.RequestID != options.RequestID {
			t.Fatalf("The request id of the message was expected to be %s", options.RequestID)
		}
		tools.PrintResource(t, listedMessage)
	}
}

// Create a message and update the name and description. Get the ity
// service and verify that the name and description have been updated
func TestMessageDelete(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.SkipRelease(t, "stable/newton")
	clients.SkipRelease(t, "stable/ocata")

	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create shared file system client: %v", err)
	}
	client.Microversion = minimumManilaMessagesMicroVersion

	options := messages.ListOpts{
		RequestID: requestID,
	}

	allPages, err := messages.List(client, options).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve messages: %v", err)
	}

	allMessages, err := messages.ExtractMessages(allPages)
	if err != nil {
		t.Fatalf("Unable to extract messages: %v", err)
	}

	if len(allMessages) == 0 {
		t.Skipf("No messages were found")
	}

	var messageID string
	for _, listedMessage := range allMessages {
		if listedMessage.RequestID != options.RequestID {
			t.Fatalf("The request id of the message was expected to be %s", options.RequestID)
		}
		tools.PrintResource(t, listedMessage)
		messageID = listedMessage.ID
	}

	message, err := messages.Get(client, messageID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve the message: %v", err)
	}

	DeleteMessage(t, client, message)
}

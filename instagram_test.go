package facebook

import (
	"log"
	"os"
	"testing"
)

func TestInstagram(t *testing.T) {

	longUserAccessToken := os.Getenv("FACEBOOKACCESSTOKEN")

	// Get facebook user information
	user, err := Me(longUserAccessToken)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	log.Printf("User: %+v", user)

	// Get the facebook accounts associated with the token
	accounts, err := GetAccounts(longUserAccessToken, user.ID)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	log.Printf("Accounts: %+v", accounts)

	// Test getting instagram account info from facebook page
	instagramAccount, instagramAccessToken, err := GetInstagramAccount((*accounts)[0].AccessToken, (*accounts)[0].ID)
	if err != nil || instagramAccessToken == "" {
		t.Fatalf("Error: %s", err)
	}
	log.Printf("InstagramAccount: %+v", instagramAccount)

	// Test getting instagram user
	instagramUser, err := GetInstagramUser((*accounts)[0].AccessToken, instagramAccount.ID)
	if err != nil || instagramAccessToken == "" {
		t.Fatalf("Error: %s", err)
	}
	log.Printf("InstagramUser: %+v", instagramUser)

	// Test uploading of images...
	/*
		container, err := PublishInstagramMedia((*accounts)[0].AccessToken, instagramAccount.ID, "Test", "<IMAGE URL>")
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("InstagramMedia: %+v", container)
		instagramMedia, err := GetInstagramContainer((*accounts)[0].AccessToken, container.ID)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("InstagramMedia: %+v", instagramMedia)
		if instagramMedia.StatusCode == "FINISHED" {
			publish, err := PublishInstagramPost((*accounts)[0].AccessToken, instagramAccount.ID, container.ID)
			if err != nil {
				t.Fatalf("Error: %s", err)
			}
			log.Printf("Publish: %+v", publish)
		}
	*/

	// Test listing of images
	/*
		media, err := GetInstagramMediaPost((*accounts)[0].AccessToken, instagramMediaID)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("Media: %+v", media)
	*/

	/*
		instagramMedias, err := GetInstagramMediaList((*accounts)[0].AccessToken, instagramAccount.ID)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("InstagramMedia: %+v", (*instagramMedias)[0])
	*/

}

package facebook

import (
	"log"
	"os"
	"testing"
)

func TestFacebook(t *testing.T) {
	longUserAccessToken := os.Getenv("FACEBOOKACCESSTOKEN")

	// Get Facebook user information
	user, err := Me(longUserAccessToken)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	log.Printf("User: %+v", user)

	// Get facebook page accounts associated with user
	accounts, err := GetAccounts(longUserAccessToken, user.ID)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	log.Printf("Accounts: %+v", accounts)

	// Testing token
	/*
		debug, err := DebugToken(accessToken, accessToken)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("Debug: %+v", debug)
	*/

	// Testing how to request long lived user and page access token
	// Requires FACEBOOKCLIENTID and FACEBOOKCLIENTSECRET environment variables
	/*
		token, err := GetLongLivedUserAccessToken(accessToken, os.Getenv("FACEBOOKCLIENTID"), os.Getenv("FACEBOOKCLIENTSECRET"))
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("Token: %+v", token)
	*/
	/*
		pages, err := GetLongLivedPagesAccessToken(longUserAccessToken, userID)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("Pages: %+v", pages)
	*/

	// Upload photo
	pageID := (*accounts)[0].ID
	pageAccessToken := (*accounts)[0].AccessToken
	media, err := UploadMedia(pageAccessToken, pageID, "https://mik.earn.se/files/b2c/output/-1645538036.png", "Vi provar med meddelanden.\nOch länk https://www.earn.se/")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	log.Printf("Media: %+v", media)

	// Test posting a request
	/*
		postRequest := PostRequest{Message: "Test", Published: true, Link: "<INSERT LINK>", Picture: "<INSERT PICTURE>"}
		post, err := CreatePost((*accounts)[0].AccessToken, (*accounts)[0].ID, &postRequest)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("Post: %+v", post)
	*/

	// Test listing all facebook posts
	/*
		posts, err := ListPosts(longUserAccessToken, pageID)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("Posts: %+v", posts)

		postStats, err := GetPostStats(longUserAccessToken, postID)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("PostStats: %+v", postStats)
	*/

	// Test publishing a post not previously published
	/*
		status, err := PublishPost(pageAccessToken, postID)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("Status: %+v", status)
	*/

	// Test deleting a post
	/*
		status, err := DeletePost(pageAccessToken, postID)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		log.Printf("Status: %+v", status)
	*/
}

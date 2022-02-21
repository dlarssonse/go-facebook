package facebook

import (
	"fmt"
)

// GetInstagramAccount ...
func GetInstagramAccount(accessToken string, pageID string) (*InstagramAccount, string, error) {
	response := InstagramAccountResponse{}
	request := struct {
		Fields string `json:"fields"`
	}{
		Fields: "instagram_business_account,access_token",
	}
	if err := makeAPIRequest("GET", accessToken, 0, fmt.Sprintf("%s/%s", GRAPH_API_VERSION, pageID), &request, &response); err != nil {
		return nil, "", err
	}

	return &response.InstagramAccount, response.AccessToken, nil
}

// GetInstagramMediaList ...
func GetInstagramMediaList(accessToken string, instagramUserID string) (*[]InstagramMedia, error) {
	response := InstagramMediaResponse{}
	request := struct {
		Fields string `json:"fields"`
	}{
		Fields: "id,caption,comments_count,like_count,is_comment_enabled,media_product_type,media_type,media_url,owner,permalink,shortcode,thumbnail_url,timestamp,username,video_title,children{id,media_type,media_url}",
	}
	if err := makeAPIRequest("GET", accessToken, 0, fmt.Sprintf("%s/%s/media", GRAPH_API_VERSION, instagramUserID), &request, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// UploadInstagramMedia ...
func UploadInstagramMedia(accessToken string, instagramUserID string, caption string, url string) (*InstagramMedia, error) {
	response := InstagramMedia{}
	request := struct {
		ImageURL string `json:"image_url"`
		Caption  string `json:"caption"`
		//MediaType        string `json:"media_type"`
		MediaProductType string `json:"media_product_type"`
		AccessToken      string `json:"access_token"`
	}{
		ImageURL: url,
		Caption:  caption,
		//MediaType:        "CAROUSEL_ALBUM",
		MediaProductType: "FEED",
		AccessToken:      accessToken,
	}
	if err := makeAPIRequest("POST", accessToken, 0, fmt.Sprintf("%s/%s/media", GRAPH_API_VERSION, instagramUserID), &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PublishInstagramMedia
func PublishInstagramMedia(accessToken string, instagramUserID string, containerID string) (*InstagramMedia, error) {
	response := InstagramMedia{}
	request := struct {
		CreationID  string `json:"creation_id"`
		AccessToken string `json:"access_token"`
	}{
		CreationID:  containerID,
		AccessToken: accessToken,
	}
	if err := makeAPIRequest("POST", accessToken, 0, fmt.Sprintf("%s/%s/media_publish", GRAPH_API_VERSION, instagramUserID), &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteInstagramMedia ...
// TODO: This is not yet implemented in the instagram v11.0 API
func DeleteInstagramMedia(accessToken string, instagramMediaID string) (interface{}, error) {
	response := struct {
		Success string `json:"success"`
	}{}
	if err := makeAPIRequest("DELETE", accessToken, 0, fmt.Sprintf("%s/%s", GRAPH_API_VERSION, instagramMediaID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetInstagramContainer ...
func GetInstagramContainer(accessToken string, instagramContainerID string) (*InstagramContainer, error) {
	response := InstagramContainer{}
	request := struct {
		Fields string `json:"fields"`
	}{
		Fields: "id,status,status_code",
	}
	if err := makeAPIRequest("GET", accessToken, 0, fmt.Sprintf("%s/%s", GRAPH_API_VERSION, instagramContainerID), &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetInstagramUser ...
func GetInstagramUser(accessToken string, instagramUserID string) (*InstagramUser, error) {
	response := InstagramUser{}
	request := struct {
		Fields string `json:"fields"`
	}{
		Fields: "biography,id,followers_count,follows_count,media_count,name,profile_picture_url,username,website",
	}
	if err := makeAPIRequest("GET", accessToken, 0, fmt.Sprintf("%s/%s", GRAPH_API_VERSION, instagramUserID), &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

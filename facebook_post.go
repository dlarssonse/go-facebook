package facebook

import "fmt"

// CreatePost ...
func CreatePost(pageAccessToken string, pageID string, request *PostRequest) (*Post, error) {
	response := Post{}
	request.AccessToken = pageAccessToken

	if err := makeAPIRequest("POST", "", 0, fmt.Sprintf("%s/feed", pageID), request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeletePost ...
func DeletePost(pageAccessToken string, postID string) (interface{}, error) {
	response := struct {
		Success bool `json:"success"`
	}{}
	if err := makeAPIRequest("DELETE", pageAccessToken, 0, fmt.Sprintf("%s", postID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPost ...
func ListPosts(pageAccessToken string, pageID string) (*[]Post, error) {
	response := PostResponse{}
	if err := makeAPIRequest("GET", pageAccessToken, 0, fmt.Sprintf("%s/feed", pageID), nil, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// PublishPost ...
func PublishPost(pageAccessToken string, postID string) (interface{}, error) {
	response := struct {
		Success bool `json:"success"`
	}{}
	request := struct {
		AccessToken string `json:"access_token"`
		IsPublished bool   `json:"is_published"`
	}{
		AccessToken: pageAccessToken,
		IsPublished: true,
	}
	if err := makeAPIRequest("POST", "", 0, fmt.Sprintf("%s", postID), &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PublishPost ...
func UploadMedia(pageAccessToken string, pageID string, url string, message string) (*PostMediaResponse, error) {
	response := PostMediaResponse{}
	request := struct {
		AccessToken string `json:"access_token"`
		Url         string `json:"url"`
		Message     string `json:"message"`
	}{
		AccessToken: pageAccessToken,
		Url:         url,
		Message:     message,
	}
	if err := makeAPIRequest("POST", "", 0, fmt.Sprintf("%s/photos", pageID), &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func GetPostStats(pageAccessToken string, postID string) (*PostStats, error) {
	response := PostStats{}
	request := struct {
		AccessToken string `json:"access_token"`
		Fields      string `json:"fields"`
	}{
		AccessToken: pageAccessToken,
		Fields:      "&fields=likes.summary(true),comments.summary(true),shares.summary(true),reactions.summary(true)",
	}
	if err := makeAPIRequest("GET", pageAccessToken, 0, fmt.Sprintf("%s", postID), &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

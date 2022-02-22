package facebook

import "fmt"

const GRAPH_API_VERSION = "v13.0"

// Me ...
func Me(accessToken string) (*User, error) {

	userResponse := User{}
	if err := makeAPIRequest("GET", accessToken, 0, "me", nil, &userResponse); err != nil {
		return nil, err
	}
	return &userResponse, nil
}

// DebugToken ...
func DebugToken(accessToken string, debugToken string) (*TokenInfo, error) {
	response := DebugTokenResponse{}
	request := struct {
		InputToken  string `json:"input_token"`
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
		InputToken:  debugToken,
	}

	if err := makeAPIRequest("GET", "", 0, fmt.Sprintf("%s/debug_token", GRAPH_API_VERSION), &request, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetLongLivedUserAccessToken ...
func GetLongLivedUserAccessToken(accessToken string, clientID string, clientSecret string) (*TokenLongLivedUser, error) {
	response := TokenLongLivedUser{}
	request := TokenLongLivedUserRequest{
		GrantType:       "fb_exchange_token",
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		FBExchangeToken: accessToken,
	}
	if err := makeAPIRequest("GET", "", 0, fmt.Sprintf("%s/oauth/access_token", GRAPH_API_VERSION), &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAccounts ...
func GetAccounts(longUserAccessToken string, userID string) (*[]Account, error) {
	response := AccountResponse{}
	if err := makeAPIRequest("GET", longUserAccessToken, 0, fmt.Sprintf("%s/%s/accounts", GRAPH_API_VERSION, userID), nil, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

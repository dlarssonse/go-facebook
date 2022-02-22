package facebook

// User ...
type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// DebugTokenResponse ...
type DebugTokenResponse struct {
	Data TokenInfo `json:"data"`
}

// TokenInfo ...
type TokenInfo struct {
	AppID               string   `json:"app_id"`
	Type                string   `json:"type"`
	Application         string   `json:"application"`
	DataAccessExpiresAt int      `json:"data_access_expires_at"`
	ExpiresAt           int      `json:"expires_at"`
	IsValid             bool     `json:"is_valid"`
	Scopes              []string `json:"scopes"`
	GranularScopes      []struct {
		Scope     string   `json:"scope"`
		TargetIds []string `json:"target_ids"`
	} `json:"granular_scopes"`
	UserID string `json:"user_id"`
}

// TokenLongLivedUser ...
type TokenLongLivedUser struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"` // The number of seconds until the token expires
}

// TokenLongLivedUserRequest ...
type TokenLongLivedUserRequest struct {
	GrantType       string `json:"grant_type"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	FBExchangeToken string `json:"fb_exchange_token"`
}

// Account ...
type Account struct {
	AccessToken  string `json:"access_token"`
	Category     string `json:"category"`
	CategoryList []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category_list"`
	Name  string
	ID    string   `json:"id"`
	Tasks []string `json:"tasks"`
}

// AccountResponse ...
type AccountResponse struct {
	Data   []Account `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

// PostResponse ...
type PostResponse struct {
	Data   []Post `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

// PostAttachment ...
type PostAttachment struct {
	Description string `json:"description"`
	Link        string `json:"link"`
	ImageHash   string `json:"image_hash"`
	Name        string `json:"name"`
	Picture     string `json:"picture"` // album, animated_image_autoplay, checkin, cover_photo, event, link, multiple, music, note, offer, photo, profile_media, status, video, video_autoplay
}

// PostRequest ...
type PostRequest struct {
	AccessToken          string           `json:"access_token"`
	Message              string           `json:"message"`
	Link                 string           `json:"link,omitempty"` // deprecated for 3.3+
	Published            bool             `json:"published"`
	Picture              string           `json:"picture,omitempty"` // needs domain verification
	ScheduledPublishTime int              `json:"scheduled_publish_time,omitempty"`
	Attachments          []PostAttachment `json:"child_attachments"`
}

// Post ...
type Post struct {
	ID          string `json:"id"`
	Message     string `json:"message"`
	CreatedTime string `json:"created_time"`
}

type PostMediaResponse struct {
	ID     string `json:"id"`
	PostID string `json:"post_id"`
}

type PostStats struct {
	ID    string `json:"id"`
	Likes struct {
		Summary struct {
			TotalCount int `json:"total_count"`
		} `json:"summary"`
	} `json:"likes"`
	Comments struct {
		Summary struct {
			TotalCount int `json:"total_count"`
		}
	} `json:"comments"`
	Reactions struct {
		Summary struct {
			TotalCount     int    `json:"total_count"`
			ViewerReaction string `json:"viewer_reaction"`
		}
	}
}

// InstagramUser ...
type InstagramUser struct {
	ID                string `json:"id"`
	Biography         string `json:"biography"`
	IGID              string `json:"ig_id"`
	FollowersCount    int    `json:"followers_count"`
	FollowsCount      int    `json:"follows_count"`
	Name              string `json:"name"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Username          string `json:"username"`
	Website           string `json:"website"`
}

// InstagramAccount ...
type InstagramAccount struct {
	ID string `json:"id"`
}

// InstagramAccountResponse ...
type InstagramAccountResponse struct {
	ID               string           `json:"id"`
	AccessToken      string           `json:"access_token"`
	InstagramAccount InstagramAccount `json:"instagram_business_account"`
}

// InstagramMedia ...
type InstagramMedia struct {
	ID               string        `json:"id"`
	CommentCount     int           `json:"comment_count"`
	LikeCount        int           `json:"like_count"`
	Caption          string        `json:"caption,omitempty"`
	ImageURL         string        `json:"image_url"`
	MediaProductType string        `json:"media_product_type"`
	MediaType        string        `json:"media_type"`
	MediaURL         string        `json:"media_url"`
	Owner            InstagramUser `json:"owner"`
	Permalink        string        `json:"permalink"`
	Shortcode        string        `json:"shortcode"`
	ThumbnailURL     string        `json:"thumbnail_url"`
	Timestamp        string        `json:"timestamp"`
	Username         string        `json:"username"`
	VideoTitle       string        `json:"video_title"`
	AccessToken      string        `json:"access_token"`
}

// InstagramContainer ...
type InstagramContainer struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
}

// InstagramMediaResponse ...
type InstagramMediaResponse struct {
	Data   []InstagramMedia `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

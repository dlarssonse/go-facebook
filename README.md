# go-facebook

## Usage example

    func main() {
      user, _ := facebook.Me(accessToken)
      log.Printf("Logged in as user: %s", user.Name)
  
      accounts, _ := facebook.GetAccounts(accessToken, user.ID)
      for _, account := range *accounts {
        log.Printf("Page access for: %s", account.Name)
    
        request := facebook.PostRequest{Message: "Testing go-facebook", Published: true, Link: "https://github.com/"}
        post, _ := facebook.CreatePost(account.AccessToken, account.ID, &postRequest)
        log.Printf("Posted message: %+v", post)
      }
    }
    
See more examples in facebook_test.go or instagram_test.go

## Environment variables used

### Required

- export FACEBOOKACCESSTOKEN="INSERT TOKEN FRON FACEBOOK AUTHENTICATION FLOW"

### Optional

- export FACEBOOKCLIENTID="INSERT CLIENT ID"
- export FACEBOOKCLIENTSECRET="INSERT CLIENT SECRET"

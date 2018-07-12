package oauth2device_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rjw57/oauth2device"
	"github.com/rjw57/oauth2device/googledevice"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

// An simple example of using this package for device authorization.
func Example() {
	// The usual OAuth2 configuration
	var clientOAuthConfig = &oauth2.Config{
		ClientID:     "<insert client id here>",
		ClientSecret: "<insert client secret here>",
		Endpoint:     google.Endpoint,

		// for example...
		Scopes: []string{youtube.YoutubeReadonlyScope},
	}

	// Augment OAuth2 configuration with device endpoints.
	var clientDeviceOAuthConfig = &oauth2device.Config{
		Config:         clientOAuthConfig,
		DeviceEndpoint: googledevice.DeviceEndpoint,
	}

	// Use default HTTP client.
	client := http.DefaultClient

	// Get URL and code for user.
	dcr, err := oauth2device.RequestDeviceCode(client, clientDeviceOAuthConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Visit: %v and enter: %v\n", dcr.VerificationURL, dcr.UserCode)

	// Wait for a token. It will be a standard oauth2.Token.
	accessToken, err := oauth2device.WaitForDeviceAuthorization(client,
		clientDeviceOAuthConfig, dcr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Access token: %v\n", accessToken)

	// Now use the token as usual...
}

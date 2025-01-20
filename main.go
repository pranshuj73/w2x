package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)


func main() {
	config, err := loadOrCreateConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	os.Setenv("GOTWI_API_KEY", config.ConsumerKey)
	os.Setenv("GOTWI_API_KEY_SECRET", config.ConsumerSecret)

	clientInput := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           config.AccessToken,
		OAuthTokenSecret:     config.AccessSecret,
	}

	client, err := gotwi.NewClient(clientInput)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("tweet: ")
		tweetText, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		tweetText = strings.TrimSpace(tweetText)

    if tweetText == "" { continue }
		if tweetText == "exit" || tweetText == "quit" || tweetText == "q" {
      goodbyePhrase := goodbyePhrase()
      fmt.Printf("\n%s\n\n", goodbyePhrase)
			break
		}

    if len(tweetText) > 280 {
      fmt.Printf("\nExceeds 280 chars. Trimmed tweet:%s\n\n", tweetText[:280])
      // ask user if they'd like to post the trimmed tweet
      fmt.Print("Post trimmed tweet? [n]: ")
      postTrimmed, err := reader.ReadString('\n')
      if err != nil {
        fmt.Println("Error reading input:", err)
      }

      postTrimmed = strings.TrimSpace(postTrimmed)
      if postTrimmed == "y" || postTrimmed == "yes" {
        tweetText = tweetText[:280]
      } else {
        fmt.Printf("\n\n")
        continue
      }
    }

		tweetInput := &types.CreateInput{
			Text: gotwi.String(tweetText),
		}

		res, err := managetweet.Create(context.Background(), client, tweetInput)
		if err != nil {
			fmt.Println("Error posting tweet:", err)
			continue
		}

		fmt.Printf("Tweet posted successfully! [ID: %s]\n\n",
			gotwi.StringValue(res.Data.ID))
	}
}

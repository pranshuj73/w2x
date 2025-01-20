package main

import (
  "math/rand"
)

func goodbyePhrase() string {
  phrases := []string{
		"Catch you on the TL!",
		"Tweet you later!",
		"Fly high, little bird!",
		"Catch you in the feed!",
		"Retweet you soon!",
		"Scroll ya later!",
		"See you on the timeline, sunshine!",
		"DM you later, gator!",
		"Bye-bye, bluebird!",
		"Unfollowing the day for now!",
	}

	// Pick a random phrase
  randomPhrase := phrases[rand.Intn(len(phrases))]

  return randomPhrase
}

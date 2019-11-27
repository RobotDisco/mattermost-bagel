package main

import (
	"fmt"

	"github.com/RobotDisco/mattermost-bagel/config"
	"github.com/RobotDisco/mattermost-bagel/mattermost"
	"github.com/mattermost/mattermost-server/model"
)

const (
	// DefaultPairingMessage is when we were able to pair two new folks together
	DefaultPairingMessage = "Hello! This week you have been matched up as conversation partners! I hope you meet up and have a great time :)"
	// ConsolingPairingMessage says we tried our best but failed miserably, the lesson learned was to never try!
	ConsolingPairingMessage = "Hello! This week you have been matched up as conversation partners again! Maybe you can continue the conversation from last time? Have a great time :)"
)

func main() {
	bagelConfig := config.CreateBagelConfig()
	persistenceConfig := config.CreatePersistenceConfig()
	matchAndSchedulePairs(bagelConfig, persistenceConfig)
}

func matchAndSchedulePairs(bagelConfig config.BagelConfig, persistenceConfig config.PersistenceConfig) {
	api := mattermost.NewMatterMostClient(bagelConfig.ServerURL, bagelConfig.BotUserName, bagelConfig.BotPassword)

	channelID, members := mattermost.GetActiveChannelMembers(*api, bagelConfig.TeamName, bagelConfig.ChannelName)
	fmt.Printf("There are %d members in channel %s for team %s\n", len(members), bagelConfig.ChannelName, bagelConfig.TeamName)
	bot := mattermost.GetBotUser(*api)
	pairs := mattermost.SplitIntoPairs(channelID, members, bot.Id)
	fmt.Printf("Created %d pair(s)\n", len(pairs))

	verifyResult := persistenceConfig.VerifyPairs(pairs)
	persistenceConfig.LogPairs(verifyResult.Successes)
	mattermost.MessageMembers(*api, verifyResult.Successes, bot, DefaultPairingMessage)
	failingPairCount := len(verifyResult.Failures)

	if failingPairCount == 0 {
		fmt.Println("Scheduling completed with no duplications!")
		return
	}
	if failingPairCount == 1 {
		fmt.Println("1 pair will have to continue the conversation from last time, sorry!")
		persistenceConfig.LogPairs(verifyResult.Failures)
		mattermost.MessageMembers(*api, verifyResult.Failures, bot, ConsolingPairingMessage)
		return
	}
	if persistenceConfig.RetryCount == 0 {
		if failingPairCount > 0 {
			fmt.Printf("%d pair(s) will have to continue the conversation from last time, sorry!\n", len(verifyResult.Failures))
			persistenceConfig.LogPairs(verifyResult.Failures)
			mattermost.MessageMembers(*api, verifyResult.Failures, bot, ConsolingPairingMessage)
		}
		return
	}

	fmt.Printf("%d pairs have been mached previously, retrying.\n", failingPairCount)
	retryScheduling(api, bot, persistenceConfig, verifyResult.Failures)
}

func retryScheduling(api *model.Client4, bot *model.User, persistenceConfig config.PersistenceConfig, failedPairs mattermost.ChannelMemberPairs) {
	persistenceConfig.RetryCount--
	retryPairs := mattermost.RetryPairs(failedPairs)
	retryVerifyResult := persistenceConfig.VerifyPairs(retryPairs)
	persistenceConfig.LogPairs(retryVerifyResult.Successes)
	fmt.Printf("Re-pairing created %d previously unmatched pairs.\n", len(retryVerifyResult.Successes))
	mattermost.MessageMembers(*api, retryVerifyResult.Successes, bot, DefaultPairingMessage)

	retryFailureCount := len(retryVerifyResult.Failures)
	if retryFailureCount == 0 { // Successfully handled all failures!
		return
	}
	// We can't try again if we've reached the limit or have one pair left, log and avoid retrying
	if persistenceConfig.RetryCount == 0 || retryFailureCount == 1 {
		fmt.Printf("%d pair(s) will have to continue the conversation from last time, sorry!\n", retryFailureCount)
		persistenceConfig.LogPairs(retryVerifyResult.Failures)
		mattermost.MessageMembers(*api, retryVerifyResult.Failures, bot, ConsolingPairingMessage)
	} else {
		retryScheduling(api, bot, persistenceConfig, retryVerifyResult.Failures)
	}
}

package mattermost

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mattermost/mattermost-server/model"
)

// ChannelMemberPair represents two channel members that have been grouped
type ChannelMemberPair struct {
	ChannelID string
	First     *model.User
	Second    *model.User
}

// ChannelMemberPairs represents a slice of ChannelMemberPair
type ChannelMemberPairs []ChannelMemberPair

// Identifier is a concatnated, alphabetical string of the two Ids of the pairs
func (pair ChannelMemberPair) Identifier() string {
	if pair.First.Id < pair.Second.Id {
		return pair.First.Id + ":" + pair.Second.Id
	}
	return pair.Second.Id + ":" + pair.First.Id
}

// SplitIntoPairs splits the list of ChannelMembers into randomized list of pairs
func SplitIntoPairs(channelID string, channelMembers model.UserSlice, coffeeBotUserID string) ChannelMemberPairs {
	// Remove coffee bot from our list of prospective matches
	coffeeBotUserIndex := -1
	for index, channelMember := range channelMembers {
		if channelMember.Id == coffeeBotUserID {
			coffeeBotUserIndex = index
			break
		}
	}
	if coffeeBotUserIndex == -1 {
		// How did we access this user list?!?
		return ChannelMemberPairs{}
	}

	// Remove one before we do the swap for less overall math
	memberLength := len(channelMembers) - 1
	channelMembers[coffeeBotUserIndex] = channelMembers[memberLength]
	channelMembers = channelMembers[:memberLength]

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(memberLength, func(i, j int) {
		channelMembers[i], channelMembers[j] = channelMembers[j], channelMembers[i]
	})

	halfMemberLength := memberLength / 2
	pairs := make(ChannelMemberPairs, halfMemberLength)

	for i := 0; i < halfMemberLength; i++ {
		firstMember := channelMembers[i]
		secondMember := channelMembers[memberLength-i-1]

		pair := ChannelMemberPair{channelID, firstMember, secondMember}
		pairs[i] = pair
	}
	return pairs
}

// RetryPairs ... TODO: (TL) documentation
func RetryPairs(pairs ChannelMemberPairs) ChannelMemberPairs {
	// our bot will already be sliced out, so we don't need to remove one
	pairLength := len(pairs)
	halfPairLength := pairLength / 2

	newPairs := make(ChannelMemberPairs, halfPairLength+1) // Allow for odd pair lengths
	if halfPairLength < 1 {
		fmt.Println("Not enough pairs to re-match, aborting.")
		return pairs
	}

	for i := 0; i < halfPairLength; i++ {
		firstPair := pairs[i]
		secondPair := pairs[pairLength-i-1]

		// TODO: (TL) If a list of pairs is does not have a heterogenous mix
		//            of channels, you will be driving for your meetup ...
		newPairs[i] = ChannelMemberPair{
			firstPair.ChannelID,
			firstPair.First,
			secondPair.Second,
		}
		newPairs[pairLength-i-1] = ChannelMemberPair{
			secondPair.ChannelID,
			secondPair.First,
			firstPair.Second,
		}
	}
	// In an uneven scenario, the middle element is left out
	if pairLength%2 != 0 {
		newPairs[halfPairLength] = pairs[halfPairLength]
	}
	return newPairs
}

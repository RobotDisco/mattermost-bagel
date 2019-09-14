package mattermost

import (
	"math/rand"
	"time"

	"github.com/mattermost/mattermost-server/model"
)

// ChannelMemberPair represents two channel members that have been grouped
type ChannelMemberPair struct {
	First  model.ChannelMember
	Second model.ChannelMember
}

// ChannelMemberPairs represents a slice of ChannelMemberPair
type ChannelMemberPairs []ChannelMemberPair

// SplitIntoPairs splits the list of ChannelMembers into randomized list of pairs
func SplitIntoPairs(channelMembers model.ChannelMembers, coffeeBotUserId string) ChannelMemberPairs {
	// Remove coffee bot from our list of prospective matches
	coffeeBotUserIndex := -1
	for index, channelMember := range channelMembers {
		if channelMember.UserId == coffeeBotUserId {
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

		pair := ChannelMemberPair{firstMember, secondMember}
		pairs[i] = pair
	}
	return pairs
}

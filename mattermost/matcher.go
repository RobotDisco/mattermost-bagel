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

// SplitIntoPairs splits the list of ChannelMembers into randomized list of pairs
func SplitIntoPairs(channelMembers model.ChannelMembers) []ChannelMemberPair {
	memberLength := len(channelMembers)
	halfMemberLength := memberLength / 2

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(memberLength, func(i, j int) {
		channelMembers[i], channelMembers[j] = channelMembers[j], channelMembers[i]
	})

	// type ChannelMembers []ChannelMember
	pairs := make([]ChannelMemberPair, halfMemberLength)

	for i := 0; i < halfMemberLength; i++ {
		firstMember := channelMembers[i]
		secondMember := channelMembers[memberLength-i-1]

		pair := ChannelMemberPair{firstMember, secondMember}
		pairs[i] = pair
	}
	return pairs
}

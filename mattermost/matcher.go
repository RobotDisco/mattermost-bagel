package mattermost

import (
	"github.com/mattermost/mattermost-server/model"
)

// ChannelMemberPair represents two channel members that have been grouped
type ChannelMemberPair struct {
	First  model.ChannelMember
	Second model.ChannelMember
}

// SplitIntoPairs Splits the list of ChannelMembers into randomized list of pairs
func SplitIntoPairs(channelMembers model.ChannelMembers) []ChannelMemberPair {
	return []ChannelMemberPair{}
}

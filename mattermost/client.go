package mattermost

import (
	"fmt"
	"os"

	"github.com/mattermost/mattermost-server/model"
)

// ClientV4 is just mattermost's model.Client4
type ClientV4 model.Client4

// NewMatterMostClient returns a NewAPIV4Client after logging in with the provided username and password
func NewMatterMostClient(url string, username string, password string) *model.Client4 {
	api := model.NewAPIv4Client(url)

	api.Login(username, password)

	return api
}

// GetActiveChannelMembers retrieves a list of active members in a given channel for the specified teamName
func GetActiveChannelMembers(m model.Client4, teamName string, channelName string) model.UserSlice {
	team, resp := m.GetTeamByName(teamName, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", team)

	channel, resp := m.GetChannelByName(channelName, team.Id, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", channel)

	members, resp := m.GetUsersInChannel(channel.Id, 0, 100, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}

	slice := model.UserSlice(members)
	return slice.FilterByActive(true)
}

/* GetChannelMembers call result
   [
	   {
			ChannelId:3j4bicr51f8x7coet4shng9kkr
			UserId:3eizzrqatin3te8s3a3yj5r6xc
			Roles:channel_user
			LastViewedAt:1568235154897
			MsgCount:42
			MentionCount:0
			NotifyProps: {
				desktop:default
				email:default
				ignore_channel_mentions:default
				mark_unread:all
				push:default
			}
			LastUpdateAt:1568235154897
			SchemeGuest:false
			SchemeUser:true
			SchemeAdmin:false
			ExplicitRoles
	   },
    ...
   ]

*/

// GetBotUser gets information about the user this program is running as.
func GetBotUser(m model.Client4) *model.User {
	user, resp := m.GetMe("")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}

	return user
}

// MessageMembers sends a message via mattermost to each set of pairs
func MessageMembers(m model.Client4, pairs ChannelMemberPairs, botUser *model.User) {
	for _, p := range pairs {
		uidList := []string{p.First.Id, p.Second.Id, botUser.Id}
		channel, resp := m.CreateGroupChannel(uidList)

		fmt.Printf("Channel: %v", channel)
		fmt.Printf("Received response: %v", resp)

		post := &model.Post{
			ChannelId: channel.Id,
			UserId:    botUser.Id,
			Message:   "Hello! This week you have been matched up as conversation partners! I hope you meet up and have a great time :)",
		}
		_, resp = m.CreatePost(post)
		if resp.Error != nil {
			fmt.Fprintf(os.Stderr, "Error: %+v", resp)
			os.Exit(1)
		}
	}
}

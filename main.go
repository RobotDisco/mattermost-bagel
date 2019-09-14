package main

import (
	"fmt"
	"os"

	"github.com/RobotDisco/mattermost-bagel/mattermost"
)

func main() {
	serverURL := os.Getenv("BAGEL_MATTERMOST_URL")

	botUserName := os.Getenv("BAGEL_USERNAME")
	botPassword := os.Getenv("BAGEL_PASSWORD")

	teamName := os.Getenv("BAGEL_TEAM_NAME")
	channelName := os.Getenv("BAGEL_CHANNEL_NAME")

	api := mattermost.NewMatterMostClient(serverURL, botUserName, botPassword)

	members := mattermost.GetChannelMembers(*api, teamName, channelName)
	fmt.Printf("%+v\n", members)
	fmt.Printf("There are %d members in channel %s for team %s\n", len(*members), channelName, teamName)

	/* This is the result of the GetChannelMembers call
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
}

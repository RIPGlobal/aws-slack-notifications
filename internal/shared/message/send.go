package message

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"os"
	"time"
)

// Initialise Slack API with the Bot Token
//
var api = slack.New(os.Getenv("OAUTH_ACCESS_TOKEN"))

func CreateOrUpdateMessage(channelID string, buildID string, blocks []slack.Block, attachment slack.Attachment){
	slackTS := ""
	slackTS = SlackTSLookup(buildID)

	// Race condition means the write could happen before the read if we don't find a message
	// set a delay based on reported Put Latency of DynamoDB and look again.
	if slackTS == "" {
		time.Sleep(20 * time.Millisecond)
		slackTS = SlackTSLookup(buildID)
	}

	if slackTS == "" {
		_, respTimestamp, err := api.PostMessage(channelID, slack.MsgOptionBlocks(blocks...), slack.MsgOptionAttachments(attachment))
		HandleSlackErrors(err, blocks)
		SaveNewMessageTS(buildID,respTimestamp)
	} else {
		_, _, _, err := api.UpdateMessage(channelID, slackTS, slack.MsgOptionBlocks(blocks...), slack.MsgOptionAttachments(attachment))
		HandleSlackErrors(err, blocks)
	}
}

// Generic Error printer if exists - not much else we can do with them.
//
func HandleSlackErrors(err error, blocks []slack.Block) {
	if err != nil {
		fmt.Println("## Error:")
		fmt.Println(err.Error())
		fmt.Println("## Request:")
		prettyJson, err := json.MarshalIndent(blocks, "", "    ")
		if err != nil {
			fmt.Println("## JSON Indent Error:")
			fmt.Println(err.Error())
		}
		fmt.Printf("\n%s", prettyJson)
	}
}
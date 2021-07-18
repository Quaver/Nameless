package utils

import (
	"fmt"
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/config"
	"github.com/Swan/Nameless/src/db"
	"github.com/andersfylling/snowflake"
	"github.com/nickname32/discordhook"
	log "github.com/sirupsen/logrus"
	"time"
)

var FirstPlace *discordhook.WebhookAPI

func InitializeDiscordWebhooks() {
	var err error
	
	fp := config.Data.DiscordWebhookFirstPlace
	FirstPlace, err = discordhook.NewWebhookAPI(snowflake.Snowflake(fp.Id), fp.Secret, true, nil)
	
	if err != nil {
		panic(err)
	}
	
	log.Info("Successfully initialized First Place Discord Webhook!")
}

// SendFirstPlaceWebhook Sends a first place webhook
func SendFirstPlaceWebhook(user *db.User, score *db.Score, m *db.Map,  oldUser *db.User) error {
	prev := discordhook.EmbedField{
		Name: "Previous #1 Holder",
		Value: "None",
		Inline: true,
	}
	
	if oldUser.Id != 0 {
		prev.Value = fmt.Sprintf("[%v](%v/user/%v)", oldUser.Username, config.Data.WebsiteUrl, oldUser.Id)
	}
	
	t := time.Now().UTC()
	
	_, err := FirstPlace.Execute(nil, &discordhook.WebhookExecuteParams{
		Embeds: []*discordhook.Embed{
			{
				Author: &discordhook.EmbedAuthor{
					Name: user.Username, 
					URL: fmt.Sprintf("%v/user/%v", config.Data.WebsiteUrl, user.Id),
					IconURL: user.AvatarURL,
				},
				Description: "🏆 **Achieved a new first place score!**",
				Fields: []*discordhook.EmbedField {
					{
						Name: "Map",
						Value: fmt.Sprintf("[%v](%v/mapset/map/%v)", m.GetString(), config.Data.WebsiteUrl, m.Id),
						Inline: false,
					},
					{
						Name: "Performance Rating",
						Value: fmt.Sprintf("%.2f", score.PerformanceRating),
						Inline: true,
					},
					{
						Name: "Accuracy",
						Value: fmt.Sprintf("%.2f%%", score.Accuracy),
						Inline: true,
					},
					{
						Name: "Max Combo",
						Value: fmt.Sprintf("%vx", score.MaxCombo),
						Inline: true,
					},
					{
						Name: "Mods",
						Value: common.GetModsString(score.Mods),
						Inline: true,
					},
					&prev,
				},
				Image: &discordhook.EmbedImage{
					URL: fmt.Sprintf("%v/mapsets/%v.jpg", config.Data.CdnUrl, m.MapsetId),
				},
				Timestamp: &t,
				Thumbnail: &discordhook.EmbedThumbnail{URL: config.Data.QuaverAvatar},
				Footer: &discordhook.EmbedFooter{
					Text:         "Quaver",
					IconURL:      config.Data.QuaverAvatar,
				},
				Color: 0x00FF00,
			},
		},
	}, nil, "")
	
	if err != nil {
		return err
	}
	
	return nil
}
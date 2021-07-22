package utils

import (
	"fmt"
	common "github.com/Swan/Nameless/common"
	config "github.com/Swan/Nameless/config"
	db "github.com/Swan/Nameless/db"
	"github.com/andersfylling/snowflake"
	"github.com/nickname32/discordhook"
	log "github.com/sirupsen/logrus"
	"time"
)

var FirstPlace *discordhook.WebhookAPI
var ScoreSubmissionErrors *discordhook.WebhookAPI
var Anticheat *discordhook.WebhookAPI

func InitializeDiscordWebhooks() {
	var err error
	
	// First Places
	fp := config.Data.DiscordWebhookFirstPlace
	FirstPlace, err = discordhook.NewWebhookAPI(snowflake.Snowflake(fp.Id), fp.Secret, true, nil)
	
	if err != nil {
		panic(err)
	}
	
	log.Info("Successfully initialized First Place Discord Webhook!")

	// Score Submission Errors
	score := config.Data.DiscordWebhookSubmissionErrors
	ScoreSubmissionErrors, err = discordhook.NewWebhookAPI(snowflake.Snowflake(score.Id), score.Secret, true, nil)

	if err != nil {
		panic(err)
	}

	log.Info("Successfully initialized Score Submission Error Discord Webhook!")

	// Anti-cheat
	ac := config.Data.DiscordWebhookAnticheat
	Anticheat, err = discordhook.NewWebhookAPI(snowflake.Snowflake(ac.Id), ac.Secret, true, nil)

	if err != nil {
		panic(err)
	}

	log.Info("Successfully initialized Anti-cheat Discord Webhook!")
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
					IconURL: user.GetAvatarURL(),
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
						Name:   "Mods",
						Value:  common.GetModsString(score.Mods),
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
					Text:    "Quaver",
					IconURL: config.Data.QuaverAvatar,
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

// SendScoreSubmissionErrorWebhook Sends a message to Discord that score submission has failed.
func SendScoreSubmissionErrorWebhook(user *db.User, reason string) error {
	t := time.Now().UTC()

	_, err := ScoreSubmissionErrors.Execute(nil, &discordhook.WebhookExecuteParams{
		Embeds: []*discordhook.Embed{
			{
				Author: &discordhook.EmbedAuthor{
					Name: user.Username,
					URL: fmt.Sprintf("%v/user/%v", config.Data.WebsiteUrl, user.Id),
					IconURL: user.GetAvatarURL(),
				},
				Description: "❌ **Score Submission Failed!**",
				Fields: []*discordhook.EmbedField {
					{
						Name: "Reason",
						Value: reason,
						Inline: false,
					},
				},
				Timestamp: &t,
				Thumbnail: &discordhook.EmbedThumbnail{URL: config.Data.QuaverAvatar},
				Footer: &discordhook.EmbedFooter{
					Text:    "Quaver",
					IconURL: config.Data.QuaverAvatar,
				},
				Color: 0xFF0000,
			},
		},
	}, nil, "")

	if err != nil {
		return err
	}
	
	return nil	
}

// SendAnticheatWebhook Sends an anti-cheat related webhook to Discord
func SendAnticheatWebhook(user *db.User, mapData *db.Map, scoreId int, isPersonalBest bool, reason string) error {
	t := time.Now().UTC()

	// Initial Actions that are on every 
	actions := fmt.Sprintf("[View profile](%v/user/%v)", config.Data.WebsiteUrl, user.Id)
	actions += fmt.Sprintf(" | [Ban User](https://a.quavergame.com/ban/%v)", user.Id)
	actions += fmt.Sprintf(" | [Edit User](https://a.quavergame.com/edituser/%v)", user.Id)
	
	if mapData != nil {
		actions += fmt.Sprintf(" | [Download Map](%v/download/mapset/%v)", config.Data.WebsiteUrl, mapData.MapsetId)	
	}
	
	if isPersonalBest {
		actions += fmt.Sprintf(" | [Download Replay](%v/download/replay/%v)", config.Data.WebsiteUrl, scoreId)
	}

	var fields []*discordhook.EmbedField
	
	if mapData != nil {
		fields = append(fields, &discordhook.EmbedField {
			Name: "Map",
			Inline: false,
			Value: fmt.Sprintf("[%v - %v [%v]](%v/mapset/map/%v)",
				mapData.Artist, mapData.Title, mapData.DifficultyName, config.Data.WebsiteUrl, mapData.Id),
		})
	}
	
	fields = append(fields, &discordhook.EmbedField{
		Name: "Detected Issues",
		Value: reason,
		Inline: false,
	})
	
	fields = append(fields, &discordhook.EmbedField{
		Name: "Admin Actions",
		Value: actions,
		Inline: false,
	})

	_, err := Anticheat.Execute(nil, &discordhook.WebhookExecuteParams{
		Embeds: []*discordhook.Embed{
			{
				Author: &discordhook.EmbedAuthor{
					Name: user.Username,
					URL: fmt.Sprintf("%v/user/%v", config.Data.WebsiteUrl, user.Id),
					IconURL: user.GetAvatarURL(),
				},
				Description: "❌ **Anti-cheat Triggered!**",
				Fields: fields,
				Timestamp: &t,
				Thumbnail: &discordhook.EmbedThumbnail{URL: config.Data.QuaverAvatar},
				Footer: &discordhook.EmbedFooter{
					Text:    "Quaver",
					IconURL: config.Data.QuaverAvatar,
				},
				Color: 0xFF0000,
			},
		},
	}, nil, "")

	if err != nil {
		return err
	}

	return nil
}
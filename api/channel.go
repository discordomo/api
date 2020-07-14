package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/discordomo/api/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Channel is the struct for channels
type Channel struct {
	Name string
	Type string
}

// CreateChannel creates a channel
func CreateChannel(c *gin.Context) {
	channel := new(Channel)
	err := c.Bind(channel)
	// Error message did not return in Insomnia
	if err != nil {
		retErr := fmt.Errorf("Unable to parse json body: %w", err)

		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr)

		return
	}

	dg := middleware.RetrieveSession(c)

	cType := parseChannelType(channel.Type)
	guildID := os.Getenv("GUILD_ID")

	_, err = dg.GuildChannelCreate(guildID, channel.Name, cType)
	if err != nil {
		return
	}

	resp := fmt.Sprintf("created %s channel %s", channel.Type, channel.Name)
	c.JSON(http.StatusOK, resp)

	dg.Close()
}

// DeleteChannel deletes a channel
func DeleteChannel(c *gin.Context) {

	channel := c.Param("channel")

	dg := middleware.RetrieveSession(c)

	guildID := os.Getenv("GUILD_ID")

	channels, err := dg.GuildChannels(guildID)
	if err != nil {
		return
	}

	todelete := ""
	for _, sample := range channels {
		if strings.EqualFold(sample.Name, channel) {
			todelete = sample.ID
		}
	}

	_, err = dg.ChannelDelete(todelete)
	if err != nil {
		logrus.Error("Error deleting channel: %w", err)
	}

	resp := fmt.Sprintf("deleted channel %s", channel)
	c.JSON(http.StatusOK, resp)

	return
}

func parseChannelType(cType string) discordgo.ChannelType {
	switch cType {
	case "voice":
		return discordgo.ChannelTypeGuildVoice
	default:
		return discordgo.ChannelTypeGuildText
	}
}

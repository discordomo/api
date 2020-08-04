package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/discordomo/api/router/middleware"
	"github.com/gin-gonic/gin"
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
	if err != nil {
		retErr := fmt.Errorf("Unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())

		return
	}

	dg := middleware.RetrieveSession(c)

	cType := parseChannelType(channel.Type)
	guildID, exists := os.LookupEnv("GUILD_ID")
	if !exists {
		retErr := fmt.Errorf("Guild ID not set")
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, retErr.Error())
		return
	}

	_, err = dg.GuildChannelCreate(guildID, channel.Name, cType)
	if err != nil {
		retErr := fmt.Errorf("Error creating channel: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	resp := fmt.Sprintf("created %s channel %s", channel.Type, channel.Name)
	c.JSON(http.StatusCreated, resp)

	dg.Close()
}

// DeleteChannel deletes a channel
func DeleteChannel(c *gin.Context) {

	channel := c.Param("channel")

	dg := middleware.RetrieveSession(c)

	guildID, exists := os.LookupEnv("GUILD_ID")
	if !exists {
		retErr := fmt.Errorf("Guild ID not set")
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, retErr.Error())
		return
	}

	channels, err := dg.GuildChannels(guildID)
	if err != nil {
		retErr := fmt.Errorf("Unable to return channel list: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())

		return
	}

	todelete := ""
	for _, channelsElement := range channels {
		if strings.EqualFold(channelsElement.Name, channel) {
			todelete = channelsElement.ID
		}
	}

	_, err = dg.ChannelDelete(todelete)
	if err != nil {
		retErr := fmt.Errorf("Error deleting channel: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return

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

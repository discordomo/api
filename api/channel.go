package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/discordomo/api/router/middleware"
	"github.com/gin-gonic/gin"
)

// CreateChannel creates a channel
func CreateChannel(c *gin.Context) {
	channel := new(discordgo.Channel)
	err := c.Bind(channel)
	if err != nil {
		retErr := fmt.Errorf("Unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	dg := middleware.RetrieveSession(c)

	guildID, exists := os.LookupEnv("GUILD_ID")
	if !exists {
		retErr := fmt.Errorf("Guild ID not set")
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, retErr.Error())
		return
	}

	_, err = dg.GuildChannelCreate(guildID, channel.Name, channel.Type)
	if err != nil {
		retErr := fmt.Errorf("Error creating channel: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	resp := fmt.Sprintf("created %v channel %s", channel.Type, channel.Name)
	c.JSON(http.StatusCreated, resp)

	dg.Close()
}

// DeleteChannel deletes a channel
func DeleteChannel(c *gin.Context) {

	// Channel is the channel ID
	channel := c.Param("channel")

	dg := middleware.RetrieveSession(c)

	_, err := dg.ChannelDelete(channel)
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

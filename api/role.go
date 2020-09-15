package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/discordomo/api/router/middleware"
	"github.com/gin-gonic/gin"
)

// CreateRole creates a role
func CreateRole(c *gin.Context) {
	dg := middleware.RetrieveSession(c)

	guildID, exists := os.LookupEnv("GUILD_ID")
	if !exists {
		retErr := fmt.Errorf("Guild ID not set")
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, retErr.Error())
		return
	}

	role, err := dg.GuildRoleCreate(guildID)
	if err != nil {
		retErr := fmt.Errorf("Error creating role: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	modifyRole(c, *role)
	if err != nil {
		retErr := fmt.Errorf("Error modifying role: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	resp := fmt.Sprintf("created new role")
	c.JSON(http.StatusCreated, resp)

	dg.Close()
}

// DeleteRole deletes a role
func DeleteRole(c *gin.Context) {

	// Role is the role ID
	role := c.Param("role")

	dg := middleware.RetrieveSession(c)

	guildID, exists := os.LookupEnv("GUILD_ID")
	if !exists {
		retErr := fmt.Errorf("Guild ID not set")
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, retErr.Error())
		return
	}

	err := dg.GuildRoleDelete(guildID, role)
	if err != nil {
		retErr := fmt.Errorf("Error deleting role: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	resp := fmt.Sprintf("deleted role %s", role)
	c.JSON(http.StatusOK, resp)

	return
}

// modifyRole modifies a role
func modifyRole(c *gin.Context, role discordgo.Role) error {
	newRole := new(discordgo.Role)

	err := c.Bind(newRole)
	if err != nil {
		return fmt.Errorf("Unable to parse json body: %w", err)
	}

	dg := middleware.RetrieveSession(c)

	guildID, exists := os.LookupEnv("GUILD_ID")
	if !exists {
		return fmt.Errorf("Guild ID not set")
	}

	_, err = dg.GuildRoleEdit(guildID, role.ID, newRole.Name, newRole.Color,
		newRole.Hoist, newRole.Permissions, newRole.Mentionable)
	if err != nil {
		return fmt.Errorf("Error updating role: %w", err)
	}

	return nil
}

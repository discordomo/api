package middleware

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

// RetrieveSession retrieves the session
func RetrieveSession(c *gin.Context) *discordgo.Session {
	value := c.Value("session")
	if value == nil {
		return nil
	}

	dg, ok := value.(*discordgo.Session)
	if !ok {
		return nil
	}

	return dg
}

// EstablishSession establishes the session
func EstablishSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, exists := os.LookupEnv("BOT_TOKEN")
		if !exists {
			fmt.Println("Bot token not set")
			return
		}

		dg, err := discordgo.New("Bot " + token)
		if err != nil {
			fmt.Println("error creating Discord session,", err)
			return
		}

		c.Set("session", dg)
	}
}

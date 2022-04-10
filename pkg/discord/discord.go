package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/HalvaPovidlo/discordBotGo/pkg/zap"
)

func OpenSession(token string, logger *zap.Logger) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Errorw("error creating Discord session",
			"err", err)
		return nil, err
	}
	logger.Infow("Bot initialized")

	session.AddHandler(func(s *discordgo.Session, r *discordgo.GuildCreate) {
		logger.Infof("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		guilds := s.State.Guilds
		for _, guild := range guilds {
			fmt.Println(guild.ID, len(guild.VoiceStates), guild.Name)
			for _, state := range guild.VoiceStates {
				fmt.Println(state.UserID, state.ChannelID, state.GuildID)
			}
		}
		fmt.Println("Ready with", len(guilds), "guilds.")
	})

	session.Identify.Intents = discordgo.IntentsAll
	err = session.Open()
	if err != nil {
		logger.Errorw("error opening connection", "err", err)
		return nil, err
	}

	logger.Infow("Bot session opened", "SessionID", session.State.SessionID)
	return session, nil
}
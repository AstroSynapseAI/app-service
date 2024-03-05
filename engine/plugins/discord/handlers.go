package discord

import (
	"context"
	"fmt"
	"github.com/AstroSynapseAI/app-service/engine/chains"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type AvatarHandler struct {
	avatarID   uint
	asaiChain  *chains.AsaiChain
	clientType string
}

func NewAvatarHandler(avatarID uint, chain *chains.AsaiChain) *AvatarHandler {
	return &AvatarHandler{
		avatarID:   avatarID,
		asaiChain:  chain,
		clientType: "Discord",
	}
}

func (handler *AvatarHandler) Messages(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	handler.asaiChain.LoadAvatar(handler.avatarID, msg.Author.ID, handler.clientType)

	if strings.Contains(msg.Content, "@"+session.State.User.ID) {
		response, err := handler.asaiChain.Prompt(context.Background(), msg.Content)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, _ = session.ChannelMessageSend(msg.ChannelID, response)
	}
}

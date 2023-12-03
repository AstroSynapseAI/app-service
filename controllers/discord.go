package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/AstroSynapseAI/app-service/engine/chains"
	"github.com/bwmarrin/discordgo"
)

func DiscordMsgHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	asaiChain, err := chains.NewAsaiChain()
	if err != nil {
		fmt.Println(err)
		return
	}
	if msg.Author.ID == session.State.User.ID {
		return
	}

	sessionID := msg.Author.ID
	userPrompt := msg.Content

	if strings.Contains(msg.Content, "@"+session.State.User.ID) {
		asaiChain.SetSessionID(sessionID)
		response, err := asaiChain.Prompt(context.Background(), userPrompt)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, _ = session.ChannelMessageSend(msg.ChannelID, response)
	}
}

func NewMemberHandler(session *discordgo.Session, addEvent *discordgo.GuildMemberAdd) {
	asaiChain, err := chains.NewAsaiChain()
	if err != nil {
		fmt.Println(err)
		return
	}
	sessionID := addEvent.User.ID
	userName := addEvent.User.Username
	channelID := "1112854836371791944"

	asaiChain.SetSessionID(sessionID)

	userPrompt := fmt.Sprintf("New user, %s (%s), has joined the server. Invoke the onboarding_script.txt and welcome user to the server.", userName, sessionID)

	response, err := asaiChain.Prompt(context.Background(), userPrompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = session.ChannelMessageSend(channelID, response)

}

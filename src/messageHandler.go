package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	content, err := ioutil.ReadFile("money")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	MoneyS := string(content)
	fmt.Println(MoneyS)
	money, err := strconv.Atoi(MoneyS)
	if err != nil {
		return
	}

	args := strings.Fields(m.Content)
	//var r io.ReadCloser
	command := args[0]
	command = strings.ToLower(command)
	for i := 0; i < len(args); i++ {
		args[i] = strings.ToLower(args[i])
		if args[i] == "zvqle" || args[i] == "(no real name you loser)" || args[i] == "864564112126115841" || args[i] == "@zvqle" {
			s.ChannelMessageSend(m.ChannelID, "Is That My Owner I Hear?")
		}
		if args[i] == "@everyone" {
			s.ChannelMessageSend(m.ChannelID, "SHUT THE FUCK UP!")
		}
		if args[i] == "js" || args[i] == "javascript" {
			s.ChannelMessageSend(m.ChannelID, "ew")
		}
		if args[i] == "go" || args[i] == "c" {
			s.ChannelMessageSend(m.ChannelID, "best lanugues out there")
		}

	}
	switch command {

	case ")hello":
		s.ChannelMessageSend(m.ChannelID, "Hello!")
	case ")userid":
		ping := "UserID = " + m.Author.ID + "\nUsername: " + m.Author.Username + "\nAvatar: " + m.Author.AvatarURL("ZivBot")
		s.ChannelMessageSend(m.ChannelID, ping)
	case ")pfp":
		s.ChannelMessageSend(m.ChannelID, m.Author.AvatarURL("ZivBot"))
	case "zvqle":
		s.ChannelMessageSend(m.ChannelID, "My Owner?")
	case "najib":
		s.ChannelMessageSend(m.ChannelID, "My Owner?")
	case ")random":
		s.ChannelMessageSend(m.ChannelID, fmt.Sprint((rand.Int())))
	case ")ask":
		s.ChannelMessageSend(m.ChannelID, ask())
	case ")help":
		s.ChannelMessageSend(m.ChannelID, help())
	case ")risk":
		String, MoneyValue := RiskTaker(money)
		WritetoFile(MoneyValue)
		s.ChannelMessageSend(m.ChannelID, String)
	}

}
func ask() string {
	//var idk string
	num := rand.Int()
	if num%2 == 0 {
		return "yes"
	} else {
		return "no"
	}
	//return "idk"
}

func help() string {
	response := ("ZIVBOT COMMANDS:\n)help: prints all the commands\n)hello: Prints Hello\n)userid: prints your userid (pfp, username, and id)\n)pfp: prints your profile picture\n)random: prints a random number\n)ask: ask a question and itll answer with yes or no")

	return response
}

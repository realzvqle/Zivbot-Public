package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {

	Token := "Insert Your Token Here"
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("failed initilizing token")
		return
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("failed starting bot")
		return
	}
	fmt.Println("Bot Is Now Running!")
	select {}
}

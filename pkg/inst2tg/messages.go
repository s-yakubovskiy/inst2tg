package inst2tg

import (
	"log"

	"github.com/zelenin/go-tdlib/client"
)

func sendMessage(tdlibClient *client.Client, chatID int64, messageText string) {
	// First, check if the chat exists in the chat list
	chats, err := tdlibClient.GetChats(&client.GetChatsRequest{
		Limit: 100,
	})
	if err != nil {
		log.Fatalf("GetChats error: %s", err)
	}

	var chatFound bool
	for _, chat := range chats.ChatIds {
		if chat == chatID {
			chatFound = true
			log.Println("chat found: ", chat)
			break
		}
	}

	if !chatFound {
		log.Fatalf("Chat with ID %d not found in chat list", chatID)
		return
	}

	// Create the input message content
	inputMsgTxt := &client.InputMessageText{
		Text: &client.FormattedText{
			Text:     messageText,
			Entities: []*client.TextEntity{}, // Ensure this is a slice of pointers
		},

		// LinkPreviewOptions: ,
		ClearDraft: false,
	}

	// Create the send message request
	sendMessageRequest := &client.SendMessageRequest{
		ChatId:              chatID,
		InputMessageContent: inputMsgTxt,
	}

	// Send the message
	_, err = tdlibClient.SendMessage(sendMessageRequest)
	if err != nil {
		log.Fatalf("SendMessage error: %s", err)
	} else {
		log.Println("Message sent successfully")
	}
}

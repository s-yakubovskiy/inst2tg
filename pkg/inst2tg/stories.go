package inst2tg

import (
	"log"
	"time"

	"github.com/zelenin/go-tdlib/client"

	"github.com/s-yakubovskiy/inst2vk/pkg/downloader"
)

type SendStoryRequest struct {
	Photo        bool
	Video        bool
	Local        bool
	Path         string
	ActivePeriod int32
}

func (t *TGClient) getPhotoContent(req SendStoryRequest) *client.InputStoryContentPhoto {
	// Create InputFile for the photo
	var inputFile client.InputFile

	if req.Local {
		inputFile = &client.InputFileLocal{Path: req.Path}
	} else {
		path, err := downloader.DownloadFileToTmp(req.Path)
		log.Println("Downloaded file to: ", path)
		if err != nil {
			log.Fatal(err)
		}
		inputFile = &client.InputFileLocal{Path: path}
	}

	// Create and return InputStoryContentPhoto
	return &client.InputStoryContentPhoto{
		Photo: inputFile,
	}
}

func (t *TGClient) getVideoContent(req SendStoryRequest) *client.InputStoryContentVideo {
	// Create InputFile for the photo
	var inputFile client.InputFile

	// NOTE: deprecated because InputFileRemote not working correctly?
	// 	inputFile = &client.InputFileRemote{Id: req.Path}

	if req.Local {
		inputFile = &client.InputFileLocal{Path: req.Path}
	} else {
		path, err := downloader.DownloadFileToTmp(req.Path)
		log.Println("Downloaded file to: ", path)
		if err != nil {
			log.Fatal(err)
		}
		// inputFile = &client.InputFileRemote{Id: path}
		inputFile = &client.InputFileLocal{Path: path}
	}

	// Create and return InputStoryContentPhoto
	return &client.InputStoryContentVideo{
		Video:       inputFile,
		IsAnimation: false,
	}
}

func (t *TGClient) canSendStory(chatID int64) bool {
	_, err := t.Client.CanSendStory(&client.CanSendStoryRequest{ChatId: chatID})
	if err != nil {
		return false
	}
	return true
}

func (t *TGClient) SendStory(chatID int64, req SendStoryRequest) {
	var content client.InputStoryContent

	t.GetChats(chatID)
	if can := t.canSendStory(chatID); !can {
		log.Println("Check Can Send Story [returned false]")
	}

	if req.Video {
		content = t.getVideoContent(req)
	} else {
		content = t.getPhotoContent(req)
	}
	// type StoryPrivacySettingsCloseFriends struct {
	// Create and configure the SendStoryRequest
	storyRequest := &client.SendStoryRequest{
		ChatId:          chatID,
		Content:         content,
		PrivacySettings: &client.StoryPrivacySettingsCloseFriends{},
		// PrivacySettings: &client.StoryPrivacySettingsSelectedUsers{UserIds: []int64{chatID}},
		ActivePeriod: 86400,
	}

	// Send the story
	cStory, err := t.Client.SendStory(storyRequest)
	if err != nil {
		log.Fatalf("SendStory error: %s", err)
	} else {
		log.Println("Story sent successfully")
		time.Sleep(30 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v\n", cStory)
	}
}

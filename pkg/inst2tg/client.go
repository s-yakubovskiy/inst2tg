package inst2tg

import (
	"log"
	"path/filepath"

	"github.com/zelenin/go-tdlib/client"
)

type TGClient struct {
	apiId     int32
	apiHash   string
	verbosity int32
	Client    *client.Client
}

func NewTGClient(apiId int32, apiHash string, verbosity int32) *TGClient {
	return &TGClient{
		apiId:     apiId,
		apiHash:   apiHash,
		verbosity: verbosity,
	}
}

func (tgc *TGClient) Initialize() error {
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     false,
		UseSecretChats:         false,
		ApiId:                  tgc.apiId,
		ApiHash:                tgc.apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: false,
		IgnoreFileNames:        true,
	}

	_, err := client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: tgc.verbosity,
	})
	if err != nil {
		return err
	}

	tgc.Client, err = client.NewClient(authorizer)
	if err != nil {
		return err
	}

	return nil
}

func (tgc *TGClient) GetMe() (*client.User, error) {
	optionValue, err := tgc.Client.GetOption(&client.GetOptionRequest{
		Name: "version",
	})
	if err != nil {
		log.Fatalf("GetOption error: %s", err)
	}

	log.Printf("TDLib version: %s", optionValue.(*client.OptionValueString).Value)

	some, err := tgc.Client.GetAccountTtl()
	log.Printf("Inactivity: %+v\n", some)

	me, err := tgc.Client.GetMe()
	if err != nil {
		log.Fatalf("GetMe error: %s", err)
	}

	log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Usernames.ActiveUsernames[0])
	log.Printf("id: %v\n", me.Id)
	return me, err
}

func (tgc *TGClient) GetChats(chatID int64) {
	chats, err := tgc.Client.GetChats(&client.GetChatsRequest{
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
}

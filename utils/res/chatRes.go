package res

import (
	"peekabook/model/web"
)

func ChatDomaintoChatResponse(chat web.ChatCreateResponse) web.ChatResponse {
	return web.ChatResponse{
		Message:  chat.Message,
		Receiver: chat.Receiver,
		Sender:   chat.Sender,
	}
}

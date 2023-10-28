package req

import (
	"peekabook/model/domain"
	"peekabook/model/web"
)

func ChatCreateRequestToChatDomain(request web.ChatCreateRequest) *domain.Chat {
	return &domain.Chat{
		Message:  request.Message,
		Receiver: request.Receiver,
		Sender:   request.Sender,
	}
}

func ChatUpdateRequestToChatDomain(request web.ChatUpdateRequest) *domain.Chat {
	return &domain.Chat{
		Message:  request.Message,
		Receiver: request.Receiver,
		Sender:   request.Sender,
	}
}

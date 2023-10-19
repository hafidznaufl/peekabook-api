package req

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func ChatDomaintoChatShema(request domain.Chat) *schema.Chat {
	return &schema.Chat{
		Message: request.Message,
		AdminID: request.AdminID,
		UserID:  request.UserID,
		Date:    request.Date,
	}
}

func ChatCreateRequestToChatDomain(request web.ChatCreateRequest) *domain.Chat {
	return &domain.Chat{
		Message: request.Message,
		AdminID: request.AdminID,
		UserID:  request.UserID,
		Date:    request.Date,
	}
}

func ChatUpdateRequestToChatDomain(request web.ChatUpdateRequest) *domain.Chat {
	return &domain.Chat{
		Message: request.Message,
		AdminID: request.AdminID,
		UserID:  request.UserID,
		Date:    request.Date,
	}
}

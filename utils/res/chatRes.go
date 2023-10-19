package res

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func ChatSchematoChatDomain(chat *schema.Chat) *domain.Chat {
	return &domain.Chat{
		ID:      chat.ID,
		Message: chat.Message,
		AdminID: chat.AdminID,
		UserID:  chat.UserID,
		Date:    chat.Date,
	}
}

func ChatDomaintoChatResponse(chat *domain.Chat) web.ChatResponse {
	return web.ChatResponse{
		ID:      chat.ID,
		Message: chat.Message,
		AdminID: chat.AdminID,
		UserID:  chat.UserID,
		Date:    chat.Date,
	}
}

func ConvertChatResponse(chats []domain.Chat) []web.ChatResponse {
	var results []web.ChatResponse
	for _, chat := range chats {
		chatResponse := web.ChatResponse{
			ID:      chat.ID,
			Message: chat.Message,
			AdminID: chat.AdminID,
			UserID:  chat.UserID,
			Date:    chat.Date,
		}
		results = append(results, chatResponse)
	}
	return results
}

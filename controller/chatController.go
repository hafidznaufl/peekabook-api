package controller

import (
	"net/http"
	"peekabook/model/web"
	"peekabook/utils/helper"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
)

type ChatController interface {
	SendMessageController(ctx echo.Context) error
	GetMessagesByReceiverController(ctx echo.Context) error
	GetAllChatsController(ctx echo.Context) error
	UpdateMessageByIDController(ctx echo.Context) error
	DeleteMessageByIDController(ctx echo.Context) error
}

type ChatControllerImpl struct {
	FirebaseApp *firebase.App
}

func NewChatController(Firebase *firebase.App) *ChatControllerImpl {
	return &ChatControllerImpl{FirebaseApp: Firebase}
}

func (c *ChatControllerImpl) SendMessageController(ctx echo.Context) error {
	chatCreateRequest := web.ChatCreateRequest{}
	if err := ctx.Bind(&chatCreateRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	dbClient, err := c.FirebaseApp.Firestore(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Database Connection Error: "+err.Error()))
	}

	chatsCollection := dbClient.Collection("chat")

	chatData := web.ChatCreateResponse{
		Message:  chatCreateRequest.Message,
		Sender:   chatCreateRequest.Sender,
		Receiver: chatCreateRequest.Receiver,
	}

	result, _, err := chatsCollection.Add(ctx.Request().Context(), chatData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Chat Error: "+err.Error()))
	}

	chatID := result.ID

	chatData.ID = chatID

	response := chatData

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Send Message", response))
}

func (c *ChatControllerImpl) GetMessagesByReceiverController(ctx echo.Context) error {
    receiver := ctx.Param("receiver")
    if receiver == "" {
        return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Receiver parameter is required"))
    }

    dbClient, err := c.FirebaseApp.Firestore(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Database Connection Error: " + err.Error()))
    }

    chatsCollection := dbClient.Collection("chat")
    query := chatsCollection.Where("Receiver", "==", receiver)
    docs, err := query.Documents(ctx.Request().Context()).GetAll()
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Messages by Receiver Error: " + err.Error()))
    }

    var messages []web.ChatCreateResponse

    for _, doc := range docs {
        var message web.ChatCreateResponse
        if err := doc.DataTo(&message); err != nil {
            return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error parsing chat data: " + err.Error()))
        }

        // Setel ID dokumen ke ID dalam response.
        message.ID = doc.Ref.ID

        messages = append(messages, message)
    }

    return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Messages by Receiver", messages))
}


func (c *ChatControllerImpl) GetAllChatsController(ctx echo.Context) error {
    dbClient, err := c.FirebaseApp.Firestore(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Database Connection Error: " + err.Error()))
    }

    chatsCollection := dbClient.Collection("chat")
    docs, err := chatsCollection.Documents(ctx.Request().Context()).GetAll()
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Chats Error: " + err.Error()))
    }

    var chats []web.ChatCreateResponse

    for _, doc := range docs {
        var chat web.ChatCreateResponse
        if err := doc.DataTo(&chat); err != nil {
            return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error parsing chat data: " + err.Error()))
        }

        // Setel ID dokumen ke ID dalam response.
        chat.ID = doc.Ref.ID

        chats = append(chats, chat)
    }

    return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Chats", chats))
}


func (c *ChatControllerImpl) UpdateMessageByIDController(ctx echo.Context) error {

	messageID := ctx.Param("id")

	chatUpdateRequest := web.ChatUpdateRequest{}
	if err := ctx.Bind(&chatUpdateRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	dbClient, err := c.FirebaseApp.Firestore(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Database Connection Error: "+err.Error()))
	}

	chatsCollection := dbClient.Collection("chat")

	messageRef := chatsCollection.Doc(messageID)

	_, err = messageRef.Get(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Message not found"))
	}

	_, err = messageRef.Set(ctx.Request().Context(), chatUpdateRequest, firestore.MergeAll)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Message Error: "+err.Error()))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Message Updated Successfully", chatUpdateRequest))
}

func (c *ChatControllerImpl) DeleteMessageByIDController(ctx echo.Context) error {
    // Dapatkan ID pesan yang akan dihapus dari parameter permintaan HTTP.
    messageID := ctx.Param("id")

    // Dapatkan klien Firestore.
    dbClient, err := c.FirebaseApp.Firestore(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Database Connection Error: " + err.Error()))
    }

    // Tentukan koleksi Firestore yang sesuai (misalnya, "chat").
    chatsCollection := dbClient.Collection("chat")

    // Buat referensi dokumen berdasarkan ID pesan.
    messageRef := chatsCollection.Doc(messageID)

    // Hapus pesan berdasarkan ID.
    _, err = messageRef.Delete(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Message Error: " + err.Error()))
    }

    return ctx.JSON(http.StatusOK, helper.SuccessResponse("Message Deleted Successfully", nil))
}

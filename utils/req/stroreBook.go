package req

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
)

func StoreBookDomaintoStoreBookSchema(request domain.Store) *schema.Store {
	return &schema.Store{
		BookID: request.BookID,
		Date: request.Date,
	}
}
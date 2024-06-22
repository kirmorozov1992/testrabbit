package message

type Message struct {
	TransactionID string `json:"transaction_id"`
}

func NewMessage(m string) *Message {
	return &Message{
		TransactionID: m,
	}
}

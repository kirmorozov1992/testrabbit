package message

type Message struct {
	TransactionID int `json:"trans_id"`
}

func NewMessage(m int) *Message {
	return &Message{
		TransactionID: m,
	}
}

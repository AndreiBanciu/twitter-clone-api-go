package dto

type Message struct {
	Message string `json:"message"`
}

func NewMessage(message string) Message {
	m := Message{
		Message: message,
	}

	return m
}

package messages

type MsgPayload struct {
	FromUser uint   `json:"from_user"`
	ToUser   uint   `json:"to_user"`
	Text     string `json:"text"`
}

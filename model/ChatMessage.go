package model

type ChatMessage struct {
	Id          int64   `json:"id"`
	OwerName    string  `json:"owerName"`
	OwerId      string  `json:"owerId"`
	TargetName  string  `json:"targetName"`
	TargetId    int64   `json:"targetId"`
	MessageType int32   `json:"messageType"`
	IsSuccess   bool    `json:"isSuccess"`
	MessageTxt  string  `json:"messageTxt"`
	MessagePic  string  `json:"messagePic"`
	Createat    float64 `json:"creatTime"`
}

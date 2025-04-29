package serializers

import "time"

// HelloSerializer godoc
// @Description Hello information
// HelloSerializer represents a serialized hello
// swagger:model HelloSerializer
type HelloSerializer struct {
	Message   string `json:"message" example:"Joe"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
}

func NewHelloSerializer(message string) *HelloSerializer {
	return &HelloSerializer{
		Message:   message,
		Code:      200,
		Timestamp: time.Now().UnixMilli(),
	}
}

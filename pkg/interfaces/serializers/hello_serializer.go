package serializers

import (
	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/hello"
)

// HelloSerializer godoc
// @Description Hello information
// HelloSerializer represents a serialized hello
// swagger:model HelloSerializer
type HelloSerializer struct {
	Message   string `json:"message" example:"Joe"`
	Timestamp int64  `json:"timestamp"`
}

func NewHelloSerializer(output *dto.HelloOutput) *HelloSerializer {
	return &HelloSerializer{
		Message:   output.Message,
		Timestamp: output.Timestamp,
	}
}

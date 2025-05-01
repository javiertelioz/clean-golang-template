package hello

import "time"

type HelloOption func(*Hello)

func NewHello(options ...HelloOption) *Hello {
	h := &Hello{
		timestamp: time.Now().UnixMilli(),
	}

	for _, option := range options {
		option(h)
	}

	return h
}

func WithName(name string) HelloOption {
	return func(h *Hello) {
		h.name = name
	}
}

func WithTimestamp(timestamp int64) HelloOption {
	return func(h *Hello) {
		h.timestamp = timestamp
	}
}

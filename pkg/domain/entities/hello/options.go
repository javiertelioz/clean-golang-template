package hello

type HelloOption func(*Hello)

func NewHello(options ...HelloOption) *Hello {
	h := &Hello{}

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

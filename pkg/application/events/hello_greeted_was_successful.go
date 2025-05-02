package events

type HelloGreetedWasSuccessful struct {
	Greeting  string
	Timestamp int64
}

func NewHelloGreetedWasSuccessful(greeting string, timestamp int64) HelloGreetedWasSuccessful {
	return HelloGreetedWasSuccessful{Greeting: greeting, Timestamp: timestamp}
}

func (e HelloGreetedWasSuccessful) EventName() string {
	return "HelloGreetedWasSuccessful"
}

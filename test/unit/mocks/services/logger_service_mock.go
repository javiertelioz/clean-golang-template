package services

import (
	"sync"

	"github.com/stretchr/testify/mock"
)

type mockLogMessage struct {
	level   string
	message string
}

type MockLoggerService struct {
	mock.Mock
	DisableOutput bool
	msgChan       chan mockLogMessage
	wg            sync.WaitGroup
	closeOnce     sync.Once
	done          chan struct{}
}

func NewMockLoggerService() *MockLoggerService {
	m := &MockLoggerService{
		DisableOutput: true,
		msgChan:       make(chan mockLogMessage, 1000),
		done:          make(chan struct{}),
	}
	m.startWorker()
	return m
}

func (m *MockLoggerService) startWorker() {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		for {
			select {
			case msg := <-m.msgChan:
				if !m.DisableOutput {
					m.Called(msg.message)
				}
			case <-m.done:
				return
			}
		}
	}()
}

func (m *MockLoggerService) Close() {
	m.closeOnce.Do(func() {
		close(m.done)
		m.wg.Wait()
	})
}

func (m *MockLoggerService) log(level, msg string) {
	select {
	case m.msgChan <- mockLogMessage{level: level, message: msg}:
	default:
		// Channel full, skip in mock
	}
}

func (m *MockLoggerService) Trace(msg string) {
	m.log("trace", msg)
}

func (m *MockLoggerService) Info(msg string) {
	m.log("info", msg)
}

func (m *MockLoggerService) Debug(msg string) {
	m.log("debug", msg)
}

func (m *MockLoggerService) Warn(msg string) {
	m.log("warn", msg)
}

func (m *MockLoggerService) Error(msg string) {
	m.log("error", msg)
}

package pkg

import (
	"errors"
	"sync"
	"time"

	"github.com/arya237/foodPilot/pkg/reservations"
)

type MockRequiredFunctions struct {
	mu                    sync.RWMutex
	getAccessTokenFn      func(studentNumber, password string) (string, error)
	getAccessTokenCounter int

	getFoodProgramFn      func(token string, startDate time.Time) (string, error)
	getFoodProgramCounter int

	reserveFoodFn      func(token string, meal reservations.ReserveModel) (string, error)
	reserveFoodCounter int
}

func NewMockRequiredFunctions() *MockRequiredFunctions {
	return &MockRequiredFunctions{}
}

func (m *MockRequiredFunctions) SetGetAccessToken(fn func(studentNumber, password string) (string, error)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.getAccessTokenFn = fn
}

func (m *MockRequiredFunctions) SetGetFoodProgram(fn func(token string, startDate time.Time) (string, error)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.getFoodProgramFn = fn
}

func (m *MockRequiredFunctions) SetReserveFood(fn func(token string, meal reservations.ReserveModel) (string, error)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.reserveFoodFn = fn
}

func (m *MockRequiredFunctions) GetAccessToken(studentNumber, password string) (string, error) {
	m.mu.RLock()
	fn := m.getAccessTokenFn
	m.mu.RUnlock()

	m.mu.Lock()
	m.getAccessTokenCounter++
	m.mu.Unlock()

	if fn != nil {
		return fn(studentNumber, password)
	}
	return "", errors.New("GetAccessToken not implemented in mock")
}

func (m *MockRequiredFunctions) GetFoodProgram(token string, startDate time.Time) (string, error) {
	m.mu.RLock()
	fn := m.getFoodProgramFn
	m.mu.RUnlock()

	m.mu.Lock()
	m.getFoodProgramCounter++
	m.mu.Unlock()

	if fn != nil {
		return fn(token, startDate)
	}
	return "", errors.New("GetFoodProgram not implemented in mock")
}

func (m *MockRequiredFunctions) ReserveFood(token string, meal reservations.ReserveModel) (string, error) {
	m.mu.RLock()
	fn := m.reserveFoodFn
	m.mu.RUnlock()

	m.mu.Lock()
	m.reserveFoodCounter++
	m.mu.Unlock()

	if fn != nil {
		return fn(token, meal)
	}
	return "", errors.New("ReserveFood not implemented in mock")
}

func (m *MockRequiredFunctions) GetCallCounters() (getAccessToken, getFoodProgram, reserveFood int) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.getAccessTokenCounter, m.getFoodProgramCounter, m.reserveFoodCounter
}

func (m *MockRequiredFunctions) ResetCallCounters() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.getAccessTokenCounter = 0
	m.getFoodProgramCounter = 0
	m.reserveFoodCounter = 0
}

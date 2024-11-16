package sessions

import (
	"errors"
	"github.com/google/uuid"
)

type Manager interface {
	CreateSession() (Session, error)
	GetSession(sessionId string) (Session, error)
	EndSession(sessionId string) error
}

type ManagerImpl struct {
	sessions map[string]Session
}

var _ Manager = (*ManagerImpl)(nil)

var (
	ErrSessionNotFound = errors.New("session not found")
)

func NewManager() *ManagerImpl {
	return &ManagerImpl{
		sessions: make(map[string]Session),
	}
}

func (m *ManagerImpl) CreateSession() (Session, error) {
	id := uuid.NewString()
	m.sessions[id] = Session{
		Id: id,
	}
	return m.sessions[id], nil
}

func (m *ManagerImpl) GetSession(sessionId string) (Session, error) {
	session, ok := m.sessions[sessionId]
	if ok {
		return session, nil
	}
	return Session{}, ErrSessionNotFound
}

func (m *ManagerImpl) EndSession(sessionId string) error {
	_, ok := m.sessions[sessionId]
	if ok {
		delete(m.sessions, sessionId)
	}
	return nil
}

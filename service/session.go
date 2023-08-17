package service

import (
	"money-treez/model"
	"money-treez/repository"
)

type SessionService interface {
	GetSessionByEmail(email string) (model.Session, error)
}

type sessionService struct {
	sessionRepo repository.SessionRepository
}

func NewSessionService(sessionRepo repository.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (c *sessionService) GetSessionByEmail(email string) (model.Session, error) {
	sessionData := model.Session{}
	sessionData, err := c.sessionRepo.SessionAvailEmail(email)
	if err != nil {
		return model.Session{}, err
	}

	return sessionData, nil
}

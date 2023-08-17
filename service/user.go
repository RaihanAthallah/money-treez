package service

import (
	"money-treez/model"
	"money-treez/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	// Create a new user
	CreateUser(user *model.User) (model.User, error)
	// Get a user by id
	GetUser(id int) (model.User, error)
	// Get a user by email
	GetUserByEmail(email string) (model.User, error)
	// login user
	Login(user *model.User) (*string, error)
	// Get all users
}

type userService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

func NewUserService(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) UserService {
	return &userService{userRepo, sessionRepo}
}

func (s *userService) CreateUser(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("user already exists")
	}

	// password hashing
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return *user, err
	}

	user.Password = string(hashedPw)

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) GetUser(id int) (model.User, error) {
	return s.userRepo.GetUser(id)
}

func (s *userService) GetUserByEmail(email string) (model.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if dbUser.Email == "" || dbUser.ID == 0 {
		return nil, errors.New("user not found")
	}

	if user.Password != dbUser.Password {
		return nil, errors.New("wrong email or password")
	}

	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &model.Claims{
		Email: dbUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {
		return nil, err
	}

	session := model.Session{
		Token:  tokenString,
		Email:  user.Email,
		Expiry: expirationTime,
	}

	_, err = s.sessionRepo.SessionAvailEmail(session.Email)
	if err != nil {
		err = s.sessionRepo.AddSessions(session)
	} else {
		err = s.sessionRepo.UpdateSessions(session)
	}

	return &tokenString, nil
}

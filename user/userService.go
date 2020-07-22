package user

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/auth"
	"github.com/well-informed/wellinformed/graph/model"
)

type UserService struct {
	db wellinformed.Persistor
}

func NewUserService(db wellinformed.Persistor) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	existingUser, err := u.db.GetUserByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		log.Printf("error while GetUserByEmail: %v", err)
		return nil, errors.New("email already in used")
	}

	existingUser, err = u.db.GetUserByUsername(input.Username)

	if existingUser != nil {
		return nil, errors.New("username already in used")
	}

	user := &model.User{
		Username:            input.Username,
		Email:               input.Email,
		Firstname:           input.Firstname,
		Lastname:            input.Lastname,
		ActivePreferenceSet: "default",
	}

	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	user.Password = hashedPassword

	//TODO wrap these two statements in a transaction
	createdUser, err := u.db.CreateUser(*user)
	if err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	_, err = u.db.CreatePreferenceSet(&model.PreferenceSet{
		UserID: createdUser.ID,
		Name:   "default",
		Sort:   model.SortTypeChronological,
	})
	if err != nil {
		log.Error("could not create default preference set for user. err: ", err)
	}

	token, err := auth.GenAccessToken(user.ID)
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      &createdUser,
	}, nil
}

func (u *UserService) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	// log.Printf("context: %v", ctx)
	existingUser, err := u.db.GetUserByEmail(input.Email)
	log.Printf("existingUser: %v", existingUser)

	if existingUser == nil || err != nil {
		log.Printf("GetUserByEmail err: %v", err)
		return nil, errors.New("email/password combination don't work 1")
	}

	err = auth.ComparePassword(input.Password, existingUser.Password)
	if err != nil {
		log.Printf("ComparePassword err: %v", err)
		return nil, errors.New("email/password combination don't work 2")
	}

	accessToken, err := auth.GenAccessToken(existingUser.ID)
	// refreshToken, rerr := user.GenRefreshToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: accessToken,
		User:      existingUser,
	}, nil
}

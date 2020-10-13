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

func (u *UserService) userExists(email string, username string) (bool, error) {
	existingUser, err := u.db.GetUserByEmail(email)
	if err != nil {
		return true, err
	}
	if existingUser != nil {
		log.Printf("error while GetUserByEmail: %v", err)
		return true, errors.New("email already in used")
	}

	existingUser, err = u.db.GetUserByUsername(username)
	if err != nil {
		return true, err
	}
	if existingUser != nil {
		return true, errors.New("username already in used")
	}

	return false, nil
}

func (u *UserService) newUser(input model.RegisterInput) (*model.User, error) {
	user := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
	}

	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	user.Password = hashedPassword

	return user, nil
}

func (u *UserService) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	if exists, err := u.userExists(input.Email, input.Username); exists {
		return nil, err
	}

	user, err := u.newUser(input)
	if err != nil {
		return nil, err
	}

	//TODO wrap these two statements in a transaction
	//Transaction starts here

	createdUser, err := u.db.CreateUser(*user)
	if err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	engine, err := u.db.SaveEngine(&model.Engine{
		UserID: createdUser.ID,
		Name:   "default",
		Sort:   model.SortTypeChronological,
	})
	if err != nil {
		log.Error("could not create default engine for user. err: ", err)
		return nil, err
	}

	userFeed, err := u.db.CreateUserFeed(&model.UserFeed{
		UserID:   createdUser.ID,
		EngineID: engine.ID,
		Title:    "default",
		Name:     "default",
	})

	createdUser.ActiveUserFeedID = userFeed.ID
	log.Debugf("user object: %+v", createdUser)
	createdUser, err = u.db.UpdateUser(createdUser)
	if err != nil {
		log.Error("could not save active user feed on new user. err: ", err)
		return nil, err
	}

	token, err := auth.GenAccessToken(createdUser.ID)
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
	log.Debugf("existingUser: %v", existingUser)

	if existingUser == nil || err != nil {
		log.Debugf("GetUserByEmail err: %v", err)
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

func (u *UserService) SaveEngine(ctx context.Context, input *model.EngineInput) (*model.Engine, error) {
	user, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Error("user not logged in. err: ", err)
		return nil, err
	}

	prefSet := &model.Engine{
		UserID:    user.ID,
		Name:      input.Name,
		Sort:      input.Sort,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	}

	updatedPrefSet, err := u.db.SaveEngine(prefSet)
	if err != nil {
		log.Error("couldn't update Engine. err: ", err)
		return nil, err
	}

	return updatedPrefSet, nil
}

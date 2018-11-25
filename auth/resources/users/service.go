package users

import (
	"auth/resources/proto"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type UserService struct {
	Users *UserStorage
}

func NewUserService(users *UserStorage) *UserService {
	return &UserService{
		Users: users,
	}
}

func (au *UserService) BySessionToken(ctx context.Context, token *proto.UUID) (*proto.User, error) {
	tokenUUID, _ := uuid.Parse(token.Value)

	user, err := au.Users.BySessionToken(tokenUUID)

	return &proto.User{
		Uid:      &proto.UUID{Value: user.UID.String()},
		Nickname: *user.Nickname,
		Email:    *user.Email,
		Avatar:   *user.Avatar,
		Points:   int32(*user.Points),
	}, err
}

//func (au *UserService) ByNickname(ctx context.Context, nickname string) (*User, error) {
//	var result User
//	err := us.db.QueryRow(
//		"SELECT nickname, email, points FROM users WHERE nickname = $1",
//		nickname,
//	).Scan(&result.Nickname, &result.Email, &result.Points)
//	if err == nil {
//		return &result, nil
//	}
//
//	return nil, ErrUnknown
//}

func (au *UserService) ByUID(ctx context.Context, userUID *proto.UUID) (*proto.User, error) {
	uid, _ := uuid.Parse(userUID.Value)
	user, err := au.Users.ByUID(uid)
	if err != nil {
		return &proto.User{}, err
	}

	// TODO: rewrite
	if user.Avatar == nil {
		some := ""
		user.Avatar = &some
	}

	return &proto.User{
		Uid:      userUID,
		Nickname: *user.Nickname,
		Email:    *user.Email,
		Avatar:   *user.Avatar,
		Points:   int32(*user.Points),
	}, nil
}

func (au *UserService) SignUp(ctx context.Context, user *proto.UserSignUp) (*proto.Nothing, error) {
	err := au.Users.Add(User{
		Nickname: &user.Nickname,
		Email:    &user.Email,
		Password: &user.Password,
	})

	return &proto.Nothing{}, err
}

func (au *UserService) SignIn(ctx context.Context, user *proto.UserSignIn) (*proto.User, error) {
	result, err := au.Users.Auth(UserAuth{
		Nickname: &user.Nickname,
		Password: &user.Password,
	})

	if err != nil {
		return &proto.User{}, err
	}

	// TODO: rewrite
	if result.Avatar == nil {
		some := ""
		result.Avatar = &some
	}

	return &proto.User{
		Uid:      &proto.UUID{Value: result.UID.String()},
		Nickname: *result.Nickname,
		Email:    *result.Email,
		Avatar:   *result.Avatar,
		Points:   int32(*result.Points),
	}, nil
}

func (au *UserService) AddAvatar(ctx context.Context, request *proto.AddAvatarRequest) (*proto.Nothing, error) {
	// TODO: rewrite

	return &proto.Nothing{}, nil
}

func (au *UserService) Leaders(ctx context.Context, request *proto.LeadersRequest) (*proto.LeadersResponse, error) {
	// TODO: rewrite

	//_, err := au.Users.Leaders(int(request.Limit), int(request.Offset))

	return &proto.LeadersResponse{}, nil
}

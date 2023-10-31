package userservice

import (
	"fmt"
	"shopstoretest/param"
)

func (s Service) Login(req param.UserLoginRequest) (param.UserLoginResponse, error) {
	user, gErr := s.Repository.GetUserByPhoneNumber(req.PhoneNumber)

	if gErr != nil {

		return param.UserLoginResponse{}, fmt.Errorf("unexpected error %w", gErr)
	}

	if req.Password != user.Password {

		return param.UserLoginResponse{}, fmt.Errorf("your id or password is invalid")
	}

	userInfo := param.UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Credit:      user.Credit,
	}

	accessToken, aErr := s.authService.CreateAccessToken(userInfo.ID, user.Role)
	if aErr != nil {

		return param.UserLoginResponse{}, fmt.Errorf("cant create access token %w", aErr)
	}

	refreshToken, rErr := s.authService.CreateRefreshToken(userInfo.ID, user.Role)
	if rErr != nil {

		return param.UserLoginResponse{}, fmt.Errorf("cant create refresh token %w", rErr)
	}

	tokens := param.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	res := param.UserLoginResponse{
		UserInfo: userInfo,
		Tokens:   tokens,
	}

	return res, nil
}

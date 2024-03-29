package services

import (
	"errors"
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/repositories"
)

type UserService struct {
	userRepository     repositories.UserRepository
	followerRepository repositories.FollowerRepository
}

func NewUserService(userRepository repositories.UserRepository, followerRepository repositories.FollowerRepository) UserService {
	return UserService{
		userRepository:     userRepository,
		followerRepository: followerRepository,
	}
}

func (s UserService) MeUpdatePassword(user *models.User, passwordReq *dto.MePasswordRequest) error {
	if !passwordReq.PasswordMatches() {
		return errors.New("new and confirm password not matches")
	}
	valid, err := user.CheckPassword(passwordReq.OldPassword)
	if err != nil || !valid {
		return errors.New("invalid old password")
	}

	user.Password = passwordReq.NewPassword

	err = s.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s UserService) MeUpdate(user *models.User, meReq *dto.MeUpdateRequest) error {
	if meReq.Name != "" {
		user.Name = meReq.Name
	}

	err := s.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) Query(params *dto.UserQueryParams) (any, error) {
	list, pagination, err := s.userRepository.Query(params)
	if err != nil {
		return nil, err
	}

	var users []*dto.User
	for _, user := range *list {
		users = append(users, &dto.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return &dto.UserPagination{
		List:       users,
		Pagination: pagination,
	}, nil
}

func (s UserService) Get(id uint) (*dto.User, error) {
	user, err := s.userRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (s UserService) Follow(user *models.User, followReq *dto.FollowerCreateRequest) error {
	if user.ID == followReq.UserID {
		return errors.New("cannot follow yourself")
	}

	err := s.followerRepository.HasFollowing(user.ID, followReq.UserID)
	if err != nil {
		return err
	}

	var follower models.Follower
	follower.UserID = followReq.UserID
	follower.FollowerID = user.ID

	err = s.followerRepository.Create(&follower)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) UnFollow(user *models.User, followReq *dto.UnFollowerCreateRequest) error {
	err := s.followerRepository.HasFollowing(user.ID, followReq.UserID)
	if err == nil {
		return errors.New("you not follow this user")
	}

	var follower models.Follower
	follower.UserID = followReq.UserID
	follower.FollowerID = user.ID

	err = s.followerRepository.Delete(&follower)
	if err != nil {
		return err
	}

	return nil
}

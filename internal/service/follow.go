package service

import (
	"errors"
	"myapp/internal/model"
	"myapp/internal/repository"
)

type FollowService interface {
	ToggleFollow(followerID, followingID int64) (string, error)
	GetFeedPosts(followerID int64) ([]model.Post, error)
}

type followService struct {
	repo     repository.FollowRepository
	postRepo repository.PostRepository
	userRepo repository.UserRepository
}

func NewFollowService(repo repository.FollowRepository, postRepo repository.PostRepository, userRepo repository.UserRepository) FollowService {
	return &followService{
		repo:     repo,
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

func (s *followService) ToggleFollow(followerID, followingID int64) (string, error) {
	if followerID == followingID {
		return "", errors.New("you can't follow yourself")
	}

	exists, err := s.repo.IsFollowing(followerID, followingID)
	if err != nil {
		return "", err
	}

	if exists {
		if err := s.repo.Unfollow(followerID, followingID); err != nil {
			return "", err
		}
		return "unfollowed", nil
	}

	if err := s.repo.Follow(followerID, followingID); err != nil {
		return "", err
	}
	return "followed", nil
}

func (s *followService) GetFeedPosts(followerID int64) ([]model.Post, error) {
	followingIDs, err := s.repo.GetFollowingIDs(followerID)
	if err != nil {
		return nil, err
	}

	if len(followingIDs) == 0 {
		posts, err := s.postRepo.GetAll()
		if err != nil {
			return nil, err
		}
		return posts, nil
	}

	posts, err := s.postRepo.GetByUserIDs(followingIDs)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

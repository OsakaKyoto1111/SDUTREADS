package service

import (
	"errors"
	"myapp/internal/repository"
)

type LikeService interface {
	ToggleLike(userID, postID int64) (string, error)
	GetLikeCount(postID int64) (int64, error)
}

type likeService struct {
	repo     repository.LikeRepository
	postRepo repository.PostRepository
}

func NewLikeService(r repository.LikeRepository, pr repository.PostRepository) LikeService {
	return &likeService{repo: r, postRepo: pr}
}

func (s *likeService) ToggleLike(userID, postID int64) (string, error) {
	post, err := s.postRepo.GetByID(postID)
	if err != nil || post == nil {
		return "", errors.New("post not found")
	}

	exists, err := s.repo.Exists(userID, postID)
	if err != nil {
		return "", err
	}

	if exists {
		if err := s.repo.Remove(userID, postID); err != nil {
			return "", err
		}
		return "unliked", nil
	}

	if err := s.repo.Add(userID, postID); err != nil {
		return "", err
	}
	return "liked", nil
}

func (s *likeService) GetLikeCount(postID int64) (int64, error) {
	return s.repo.Count(postID)
}

package service

import (
	"errors"
	"myapp/internal/model"
	"myapp/internal/repository"
)

type PostService interface {
	Create(userID int64, content string) error
	GetAll() ([]model.Post, error)
	GetByID(id int64) (*model.Post, error)
	Delete(id int64, userID int64) error
}
type postService struct{ repo repository.PostRepository }

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}

}

func (s *postService) Create(userID int64, content string) error {
	if len(content) == 0 || len(content) > 1000 {
		return errors.New("content must be 1–1000 characters")
	}
	post := &model.Post{UserID: userID, Content: content}
	return s.repo.Create(post)

}
func (s *postService) GetAll() ([]model.Post, error) {
	return s.repo.GetAll()
}

func (s *postService) GetByID(id int64) (*model.Post, error) {
	return s.repo.GetByID(id)
}
func (s *postService) Delete(id int64, userID int64) error {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("post not found")
	}
	if post.UserID != userID {
		return errors.New("forbidden: you can delete only your own posts")
	}
	return s.repo.Delete(id)
}

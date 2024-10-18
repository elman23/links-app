package service

import (
	"links-be/domain"
	"links-be/dto"
)

type ILinkService interface {
	CreateLink(dto.Link) (*dto.Link, error)
	FindLink(string) (*dto.Link, error)
	DeleteLink(string) (*string, error)
	FindLinks() ([]dto.Link, error)
}

type LocalLinkService struct {
	repo domain.ILinkRepo
}

func (lls LocalLinkService) CreateLink(link dto.Link) (*dto.Link, error) {
	return lls.repo.CreateLink(link)
}

func (lls LocalLinkService) FindLink(id string) (*dto.Link, error) {
	return lls.repo.FindLink(id)
}

func (lls LocalLinkService) DeleteLink(id string) (*string, error) {
	return lls.repo.DeleteLink(id)
}

func (lls LocalLinkService) FindLinks() ([]dto.Link, error) {
	return lls.repo.FindLinks()
}

func NewLocalLinkService(repo domain.ILinkRepo) LocalLinkService {
	return LocalLinkService{
		repo: repo,
	}
}

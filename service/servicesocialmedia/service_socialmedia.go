package servicesocialmedia

import (
	"errors"

	"golang-web-service/entity"
	"golang-web-service/model/modelsocialmedia"
	"golang-web-service/repository/repositoryphoto"
	"golang-web-service/repository/repositorysocialmedia"
	"golang-web-service/validation"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ServiceSocialMedia interface {
	Create(data modelsocialmedia.Request) (modelsocialmedia.Response, error)
	GetList() (modelsocialmedia.ResponseListWrapper, error)
	UpdateByID(data modelsocialmedia.Request) (modelsocialmedia.Response, error)
	DeleteByID(id uint) error
}

type service struct {
	repository repositorysocialmedia.RepositorySocialMedia
	repoPhoto  repositoryphoto.RepositoryPhoto
}

func New(repository repositorysocialmedia.RepositorySocialMedia, repoPhoto repositoryphoto.RepositoryPhoto) ServiceSocialMedia {
	return &service{repository: repository, repoPhoto: repoPhoto}
}

func (srv *service) Create(data modelsocialmedia.Request) (modelsocialmedia.Response, error) {
	err := validation.ValidateSocialMediaCreate(data)
	if err != nil {
		return modelsocialmedia.Response{}, err
	}

	entitySocialMedia := new(entity.SocialMedia)
	copier.Copy(&entitySocialMedia, &data)

	createdSocialMedia, err := srv.repository.Create(*entitySocialMedia)
	if err != nil {
		return modelsocialmedia.Response{}, err
	}

	response := modelsocialmedia.Response{}
	copier.Copy(&response, &createdSocialMedia)

	return response, nil
}

func (srv *service) GetList() (modelsocialmedia.ResponseListWrapper, error) {
	listSocialMedia, err := srv.repository.GetList()
	if err != nil {
		return modelsocialmedia.ResponseListWrapper{}, err
	}

	responseList := []modelsocialmedia.ResponseList{}

	for _, socialMedia := range listSocialMedia {
		//get photo by user id
		photo, err := srv.repoPhoto.GetPhotoByUserID(socialMedia.UserID)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return modelsocialmedia.ResponseListWrapper{}, err
		}

		response := new(modelsocialmedia.ResponseList)
		copier.Copy(&response, &socialMedia)
		if photo.PhotoURL != "" {
			response.User.ProfileImageUrl = photo.PhotoURL
		}

		responseList = append(responseList, *response)
	}

	return modelsocialmedia.ResponseListWrapper{SocialMedias: responseList}, nil
}

func (srv *service) UpdateByID(data modelsocialmedia.Request) (modelsocialmedia.Response, error) {
	err := validation.ValidateSocialMediaCreate(data)
	if err != nil {
		return modelsocialmedia.Response{}, err
	}

	entitySocialMedia := new(entity.SocialMedia)
	copier.Copy(&entitySocialMedia, &data)

	updatedSocialMedia, err := srv.repository.UpdateByID(*entitySocialMedia)
	if err != nil {
		return modelsocialmedia.Response{}, err
	}

	response := modelsocialmedia.Response{}
	copier.Copy(&response, &updatedSocialMedia)

	return response, nil
}

func (srv *service) DeleteByID(id uint) error {
	err := srv.repository.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}

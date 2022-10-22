package repositoryphoto

import (
	"golang-web-service/entity"

	"gorm.io/gorm"
)

type RepositoryPhoto interface {
	Create(data entity.Photo) (entity.Photo, error)
	GetPhotos() ([]entity.Photo, error)
	Update(data entity.Photo) (entity.Photo, error)
	Delete(id int) error
	GetPhotoByUserID(id uint) (entity.Photo, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositoryPhoto {
	return &repository{db: db}
}

// Update photo from DB
func (r *repository) Update(data entity.Photo) (entity.Photo, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return entity.Photo{}, err
	}
	return data, nil
}

// Delete photo from DB by photo ID
func (r *repository) Delete(id int) error {
	photo := entity.Photo{}
	photo.ID = uint(id)
	err := r.db.First(&photo).Where("id = ?", id).Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

// Create data Photo
func (r *repository) Create(data entity.Photo) (entity.Photo, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Photo{}, err
	}
	return data, nil
}

// GetPhotos  return slice photo
func (r *repository) GetPhotos() ([]entity.Photo, error) {
	var photo []entity.Photo
	err := r.db.Preload("User").Find(&photo).Error
	if err != nil {
		return []entity.Photo{}, err
	}
	return photo, nil
}

// GetPhoto By User ID
func (r *repository) GetPhotoByUserID(id uint) (entity.Photo, error) {
	var photo entity.Photo
	err := r.db.Preload("User").Where("user_id = ?", id).First(&photo).Error
	if err != nil {
		return entity.Photo{}, err
	}
	return photo, nil
}

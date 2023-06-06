package campaign

import (
	"fmt"
	"mygram/user"

	"gorm.io/gorm"
)

type Repository interface {
	//create User
	FindAll() ([]Campaign, error)
	FindById(ID int) (Campaign, error)
	FindByUserId(userID int) ([]Campaign, error)
	Update(photo Campaign) (Campaign, error)
	Delete(photo Campaign) (Campaign, error)
	CreateImage(photo Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var photo []Campaign

	err := r.db.Preload("User").Find(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *repository) FindById(ID int) (Campaign, error) {
	var photo Campaign

	err := r.db.Preload("User").Where("id = ?", ID).First(&photo).Error

	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *repository) Update(photo Campaign) (Campaign, error) {
	err := r.db.Save(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil

}

func (r *repository) Delete(photo Campaign) (Campaign, error) {
	err := r.db.Delete(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) CreateImage(photo Campaign) (Campaign, error) {
	err := r.db.Create(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) FindByUserId(userID int) ([]Campaign, error) {
	//slice karena campaign dari user_id tertentu mungkin banyak ngepost dll
	var photo []Campaign

	//mencari user_id brapa lalu menggunakan preload yang bertujuan dalam mendapatkan data tertentu
	//dalam
	// err := r.db.Preload("User").Where("id = ?", userID).Find(&photo).Error
	// if err != nil {
	// 	return photo, err
	// }
	// return photo, nil

	err := r.db.Joins("User", r.db.Where(&user.User{ID: userID})).Find(&photo).Error
	fmt.Println("eror", photo)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

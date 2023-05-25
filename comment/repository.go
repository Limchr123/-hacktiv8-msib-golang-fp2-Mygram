package comment

import (
	"mygram/campaign"
	"mygram/user"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Comment, error)
	Create(comment Comment) (Comment, error)
	FindById(ID int) (Comment, error)
	FindByUserId(userID int, campaignID int) ([]Comment, error)
	Update(comment Comment) (Comment, error)
	Delete(comment Comment) (Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Comment, error) {
	var comment []Comment

	// err := r.db.Joins("User").Joins("Campaign").Find(&comment).Error
	err := r.db.Preload("User").Preload("Campaign").Find(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) Create(comment Comment) (Comment, error) {
	err := r.db.Create(&comment).Error

	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) FindByUserId(userID int, campaignID int) ([]Comment, error) {
	//slice karena campaign dari user_id tertentu mungkin banyak ngepost dll
	var comment []Comment

	//mencari user_id brapa lalu menggunakan preload yang bertujuan dalam mendapatkan data tertentu
	//dalam
	// err := r.db.Preload("User").Where("id = ?", userID).Find(&photo).Error
	// if err != nil {
	// 	return photo, err
	// }
	// return photo, nil

	err := r.db.Joins("User", r.db.Where(&user.User{ID: userID})).Joins("Campaign", r.db.Where(&campaign.Campaign{ID: campaignID})).Find(&comment).Error
	// fmt.Println("eror", photo)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) Update(update Comment) (Comment, error) {
	err := r.db.Save(&update).Error
	if err != nil {
		return update, err
	}
	return update, nil
}

func (r *repository) FindById(ID int) (Comment, error) {
	var comment Comment

	err := r.db.Preload("User").Where("id = ?", ID).Find(&comment).Error

	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) Delete(comment Comment) (Comment, error) {
	err := r.db.Delete(&comment).Error

	if err != nil {
		return comment, err
	}
	return comment, nil
}

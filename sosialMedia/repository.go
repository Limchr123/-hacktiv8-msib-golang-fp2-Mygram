package sosialMedia

import (
	"mygram/user"

	"gorm.io/gorm"
)

type Repository interface {
	// FindAll() ([]sosmed, error)
	Create(sosmed SosialMedia) (SosialMedia, error)
	FindById(ID int) (SosialMedia, error)
	FindByUserId(userID int) ([]SosialMedia, error)
	Update(sosmed SosialMedia) (SosialMedia, error)
	Delete(sosmed SosialMedia) (SosialMedia, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) FindAll() ([]SosialMedia, error) {
// 	var sosmed []SosialMedia

// 	// err := r.db.Joins("User").Joins("Campaign").Find(&sosmed).Error
// 	err := r.db.Preload("User").Preload("Campaign").Find(&sosmed).Error
// 	if err != nil {
// 		return sosmed, err
// 	}
// 	return sosmed, nil
// }

func (r *repository) Create(sosmed SosialMedia) (SosialMedia, error) {
	err := r.db.Create(&sosmed).Error

	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (r *repository) FindByUserId(userID int) ([]SosialMedia, error) {
	//slice karena campaign dari user_id tertentu mungkin banyak ngepost dll
	var sosmed []SosialMedia

	//mencari user_id brapa lalu menggunakan preload yang bertujuan dalam mendapatkan data tertentu
	//dalam
	// err := r.db.Preload("User").Where("id = ?", userID).Find(&photo).Error
	// if err != nil {
	// 	return photo, err
	// }
	// return photo, nil

	err := r.db.Joins("User", r.db.Where(&user.User{ID: userID})).Find(&sosmed).Error
	// fmt.Println("eror", photo)
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (r *repository) Update(update SosialMedia) (SosialMedia, error) {
	err := r.db.Save(&update).Error
	if err != nil {
		return update, err
	}
	return update, nil
}

func (r *repository) FindById(ID int) (SosialMedia, error) {
	var sosmed SosialMedia

	err := r.db.Preload("User").Where("id = ?", ID).Find(&sosmed).Error

	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (r *repository) Delete(sosmed SosialMedia) (SosialMedia, error) {
	err := r.db.Delete(&sosmed).Error

	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

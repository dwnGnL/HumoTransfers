package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

func (r Repository) AddUserInfo(user *models.UserInfo) error {

	tx := db.Data.Table("userinfo").Create(user)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *Repository) GetUserInfo(pagination *models.Pagination) ([]*models.UserInfo, error) {
	var UserInfo []*models.UserInfo
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("userinfo").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.UserInfo{}).Find(&UserInfo)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return UserInfo, nil
}

func (r Repository) UpdateUserInfo(userInfo *models.UserInfo) error {

	var UserInfos *models.UserInfo
	query2 := db.Data.Where("id=?", userInfo.ID).Find(&UserInfos)
	if query2.Error != nil {
		log.Println(query2.Error)
		return query2.Error
	}

	log.Println(1, userInfo, 2, UserInfos)
	tx := db.Data.Model(&models.UserInfo{}).Where("id = ?", userInfo.ID).Updates(models.UserInfo{Name: userInfo.Name, Icon: userInfo.Icon, Sort: userInfo.Sort})
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (r Repository) DeleteUserInfo(UserInfo *models.UserInfo) error {
	query := db.Data.Table("userinfo").Where("id =?", UserInfo.ID).Delete(UserInfo)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r Repository) UserInfoStatus(userInfo *models.UserInfo) error {
	tx := db.Data.Where("id", userInfo.ID).Table("userinfo").Scan(&userInfo)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if userInfo.Active == true {
		userInfo.Active = false
	} else {
		userInfo.Active = true
	}
	tx = db.Data.Where("id = ?", userInfo.ID).Table("userinfo").Update("active", userInfo.Active)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) TotalPageUserInfo(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}

	query := db.Data.Table("userinfo").Count(&length)
	if query.Error != nil {
		log.Println(query.Error, "error in TotalPageCountry")
		return 0, query.Error
	}

	totalPage := length / limit
	if length%limit != 0 {
		totalPage++
	}
	return totalPage, nil
}

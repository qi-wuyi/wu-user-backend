package dal

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"wu-user-backend/biz/dal/db"
	"wu-user-backend/utils/pkg/constants"
)

type User struct {
	Id           int64     `gorm:"column:id"`
	UserAccount  string    `gorm:"column:userAccount"`
	UserPassword string    `gorm:"column:userPassword"`
	UserAvatar   string    `gorm:"column:userAvatar"`
	UserProfile  string    `gorm:"column:userProfile"`
	Tags         string    `gorm:"column:tags"`
	UserRole     string    `gorm:"column:userRole"`
	CreatTime    time.Time `gorm:"column:creatTime"`
	UpdateTime   time.Time `gorm:"column:updateTime"`
	IsDeleted    int64     `gorm:"column:isDeleted"`
}

func (User) TableName() string {
	return constants.UserTableName
}

func IsAccountExist(account string) bool {
	u := new(User)
	msg := db.DB.Where("userAccount = ? and isDelete = 0", account).Find(u).Error
	if errors.Is(msg, gorm.ErrRecordNotFound) {
		return false
	}
	if msg != nil {
		return false
	}
	if u.Id == 0 {
		return false
	}
	return true
}

func CreateUser(account, password string) (id int64) {
	u := &User{
		UserAccount:  account,
		UserPassword: password,
	}
	msg := db.DB.Select("userAccount", "userPassword").Create(u).Error
	if msg != nil {
		return 0
	}
	return u.Id
}

func GetUserByAccount(account string) *User {
	u := new(User)
	msg := db.DB.Where("userAccount = ? and isDelete = 0", account).Find(u).Error
	if errors.Is(msg, gorm.ErrRecordNotFound) {
		return nil
	}
	if msg != nil {
		return nil
	}
	if u.Id == 0 {
		return nil
	}
	return u
}

func DeleteUserById(id int64) bool {
	u := new(User)
	msg := db.DB.Where("id =? and isDelete = 0", id).Find(u).Error
	if errors.Is(msg, gorm.ErrRecordNotFound) {
		return false
	}
	if msg != nil {
		return false
	}
	if u.Id == 0 {
		return false
	}
	msg = db.DB.Model(u).Update("isDelete", 1).Error
	if errors.Is(msg, gorm.ErrRecordNotFound) {
		return false
	}
	if msg != nil {
		return false
	}
	return true
}

package pack

import (
	dal "wu-user-backend/biz/dal/user"
	"wu-user-backend/biz/model/user"
)

func User(data *dal.User) *user.User {
	if data == nil {
		return nil
	}
	creatTime := data.CreatTime.String()
	return &user.User{
		UserId:      data.Id,
		UserAccount: data.UserAccount,
		UserAvatar:  data.UserAvatar,
		UserProfile: data.UserProfile,
		Tags:        data.Tags,
		UserRole:    data.UserRole,
		CreatTime:   creatTime,
	}
}

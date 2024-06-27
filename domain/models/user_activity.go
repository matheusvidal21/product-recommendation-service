package models

type UserActivityDomain interface {
	GetUserID() string
	GetProductID() string
	GetAction() string
}

type userActivity struct {
	userID    string
	productID string
	action    Action
}

func NewUserActivity(userID, productID string, action Action) UserActivityDomain {
	return &userActivity{
		userID:    userID,
		productID: productID,
		action:    action,
	}
}

func (u *userActivity) GetUserID() string {
	return u.userID
}

func (u *userActivity) GetProductID() string {
	return u.productID
}

func (u *userActivity) GetAction() string {
	return string(u.action)
}

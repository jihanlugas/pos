package model

import "time"

type UserView struct {
	ID          string     `json:"id"`
	RoleID      string     `json:"role_id"`
	Email       string     `json:"email"`
	Username    string     `json:"username"`
	NoHp        string     `json:"noHp"`
	Fullname    string     `json:"fullname"`
	Passwd      string     `json:"-"`
	PassVersion int        `json:"passVersion"`
	Active      bool       `json:"active"`
	LastLoginDt *time.Time `json:"lastLoginDt"`
	PhotoID     string     `json:"photoId"`
	PhotoUrl    string     `json:"photoUrl"`
	CreateBy    string     `json:"createBy"`
	CreateDt    time.Time  `json:"createDt"`
	UpdateBy    string     `json:"updateBy"`
	UpdateDt    time.Time  `json:"updateDt"`
	DeleteBy    string     `json:"deleteBy"`
	DeleteDt    *time.Time `json:"deleteDt"`
	CreateName  string     `json:"createName"`
	UpdateName  string     `json:"updateName"`
	DeleteName  string     `json:"deleteName"`
}

func (UserView) TableName() string {
	return VIEW_USER
}

package model

import "time"

type UserView struct {
	ID          string     `json:"id"`
	RoleID      string     `json:"roleId"`
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

type CompanyView struct {
	ID          string     `json:"primaryKey"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Address     string     `json:"address"`
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

func (CompanyView) TableName() string {
	return VIEW_COMPANY
}

type UsercompanyView struct {
	ID               string     `json:"id"`
	UserID           string     `json:"userId"`
	CompanyID        string     `json:"companyId"`
	IsDefaultCompany bool       `json:"IsDefaultCompany"`
	IsCreator        bool       `json:"IsCreator"`
	CreateBy         string     `json:"createBy"`
	CreateDt         time.Time  `json:"createDt"`
	UpdateBy         string     `json:"updateBy"`
	UpdateDt         time.Time  `json:"updateDt"`
	DeleteBy         string     `json:"deleteBy"`
	DeleteDt         *time.Time `json:"deleteDt"`
	UserName         string     `json:"userName"`
	CompanyName      string     `json:"companyName"`
	CreateName       string     `json:"createName"`
	UpdateName       string     `json:"updateName"`
	DeleteName       string     `json:"deleteName"`
}

func (UsercompanyView) TableName() string {
	return VIEW_USERCOMPANY
}

type ItemView struct {
	ID          string     `json:"id"`
	CompanyID   string     `json:"companyId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       int64      `json:"price"`
	CreateBy    string     `json:"createBy"`
	CreateDt    time.Time  `json:"createDt"`
	UpdateBy    string     `json:"updateBy"`
	UpdateDt    time.Time  `json:"updateDt"`
	DeleteBy    string     `json:"deleteBy"`
	DeleteDt    *time.Time `json:"deleteDt"`
	CompanyName string     `json:"companyName"`
	CreateName  string     `json:"createName"`
	UpdateName  string     `json:"updateName"`
	DeleteName  string     `json:"deleteName"`
}

func (ItemView) TableName() string {
	return VIEW_ITEM
}

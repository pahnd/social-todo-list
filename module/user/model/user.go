package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"social-todo-list/common"
)

const EntityName = "User"

type UserRole int

const (
	RoleUser UserRole = 1 << iota
	RoleAdmin
	RoleDevelop
	RoleMod
)

// iota trả về value. Hàm này sẽ nhận giá trị số trả về byte string
func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	case RoleDevelop:
		return "develop"
	case RoleMod:
		return "mod"
	default:
		return "user"
	}
}

//hàm thao tác với DB Scan và Value
// Scan = Db đi lên App
// Value = struct đi xuống DB

func (role *UserRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var r UserRole

	roleValue := string(bytes)

	if roleValue == "user" {
		r = RoleUser
	} else if roleValue == "admin" {
		r = RoleAdmin
	} else if roleValue == "develop" {
		r = RoleDevelop
	} else if roleValue == "mod" {
		r = RoleMod
	}

	*role = r
	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}
	return role.String(), nil
}

func (role *UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}

type User struct {
	common.SQLModel
	Username  string   `json:"username" gorm:"column:username;"`
	Password  string   `json:"-" gorm:"column:password;"`
	Salt      string   `json:"-" gorm:"column:salt;"`
	LastName  string   `json:"last_name" gorm:"column:last_name;"`
	FirstName string   `json:"first_name" gorm:"column:first_name;"`
	Phone     string   `json:"phone" gorm:"column:phone;"`
	Role      UserRole `json:"role" gorm:"column:role;"`
	Status    int      `json:"status" gorm:"column:status;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetUserName() string {
	return u.Username
}

func (u *User) GetRole() string {
	return u.Role.String()
}

// Trả về table name DB
func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel
	Username  string  `json:"username" gorm:"column:username;"`
	Password  string  `json:"password" gorm:"column:password;"`
	LastName  string  `json:"last_name" gorm:"column:last_name;"`
	FirstName string  `json:"first_name" gorm:"column:first_name;"`
	Role      *string `json:"role,omitempty" gorm:"column:role;"`
	Salt      string  `json:"-" gorm:"column:salt;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"Username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrUsernameExisted = common.NewCustomError(
		errors.New("username has already existed"),
		"Username has already existed",
		"ErrusernameExisted",
	)
)

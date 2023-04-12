package model

type NcuHomeCreateInput struct {
	Name  string
	Sex   string
	Group string
}

type UserLoginInput struct {
	Password string
	Username string
}

type UserLoginOutput struct {
	Username string
}

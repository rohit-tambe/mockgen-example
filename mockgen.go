//go:build tool

package mockgen_demo

import (
	_ "github.com/golang/mock/mockgen"
)

//go:generate mockgen -package mocks -destination mocks/mockEmployee/emp_mock.go github.com/rohit-tambe/mockgen-example/party Greeter
//go:generate mockgen -package mocks -destination mocks/mockPlayer/player_mock.go github.com/rohit-tambe/mockgen-example/user UserI

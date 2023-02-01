package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	Init()
	// create test:
	// note: 创建用户前需要查询是否存在该用户，保证id的自增性与username的唯一性
	_, err := CreateUser(context.Background(), &(User{Username: "cbcn", Password: "123"}))
	assert.Equal(t, err, nil, "Create User failed")
	_, err = CreateUser(context.Background(), &(User{Username: "cbdn", Password: "123"}))
	assert.Equal(t, err, nil, "Create User failed")
}
func TestQueryUser(t *testing.T) {
	Init()
	// create test:
	// note: 创建用户前需要查询是否存在该用户，保证id的自增性与username的唯一性
	_, err := CreateUser(context.Background(), &(User{Username: "cbcn", Password: "123"}))
	assert.Equal(t, err, nil, "Create User failed")
	_, err = CreateUser(context.Background(), &(User{Username: "cbdn", Password: "123"}))
	assert.Equal(t, err, nil, "Create User failed")
	// query test:
	username := "cbcn"
	_, err = QueryUser(context.Background(), username)
	assert.Equal(t, err, nil, "Not Found User")
	username = "cbcddddn"
	_, err = QueryUser(context.Background(), username)
	assert.Equal(t, err, nil, "Query failed")

}
func TestGetUserInfo(t *testing.T) {
	Init()
	// create test:
	// note: 创建用户前需要查询是否存在该用户，保证id的自增性与username的唯一性
	_, err := CreateUser(context.Background(), &(User{Username: "cbcn", Password: "123"}))
	assert.Equal(t, err, nil, "Create User failed")
	_, err = GetUserInfo(context.Background(), 1)
	assert.Equal(t, err, nil, "Get User Info Failed")

}

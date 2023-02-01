package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLikes(t *testing.T) {
	Init()
	// create test:
	var lk Likes
	lk.Uid = 1
	lk.Vid = 1
	err := CreateLikes(context.Background(), &lk)
	assert.Equal(t, err, nil, "Create likes failed")
	err = DeleteLikes(context.Background(), lk.ID)
	assert.Equal(t, err, nil, "Delete likes failed")
}
func TestGetUserLikes(t *testing.T) {
	Init()
	// Get test:
	var lk Likes
	lk.Uid = 3
	lk.Vid = 2
	err := CreateLikes(context.Background(), &lk)
	assert.Equal(t, err, nil, "Create likes failed")
	var dk Likes
	dk.Uid = 3
	dk.Vid = 8
	err = CreateLikes(context.Background(), &dk)
	assert.Equal(t, err, nil, "Create likes failed")
	res, err := GetUserLikes(context.Background(), 3)
	assert.Equal(t, err, nil, "query failed")
	assert.Equal(t, len(res), 2, "Get likes error")
	DeleteLikes(context.Background(), lk.ID)
	DeleteLikes(context.Background(), dk.ID)

}

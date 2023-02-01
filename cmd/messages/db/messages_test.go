package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMessages(t *testing.T) {
	Init()
	// create test:
	var lk Messages
	lk.Uid = 1
	lk.To_uid = 2
	lk.Content = "cbn"
	err := CreateMessages(context.Background(), &lk)
	assert.Equal(t, err, nil, "Create Messages failed")
	err = DeleteMessages(context.Background(), lk.ID)
	assert.Equal(t, err, nil, "Delete Messages failed")
}
func TestGetUserMessages(t *testing.T) {
	Init()
	// Get test:
	var lk Messages
	lk.Uid = 3
	lk.To_uid = 2
	lk.Content = "cbbb"
	err := CreateMessages(context.Background(), &lk)
	assert.Equal(t, err, nil, "Create likes failed")
	var dk Messages
	dk.Uid = 3
	dk.To_uid = 2
	dk.Content = "cccc"
	err = CreateMessages(context.Background(), &dk)
	assert.Equal(t, err, nil, "Create likes failed")
	res, err := GetUserMessages(context.Background(), 3, 2)
	assert.Equal(t, err, nil, "query failed")
	assert.Equal(t, len(res), 2, "Get likes error")
	DeleteMessages(context.Background(), lk.ID)
	DeleteMessages(context.Background(), dk.ID)

}

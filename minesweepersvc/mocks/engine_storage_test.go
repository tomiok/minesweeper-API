package mocks

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/tomiok/minesweeper-API/minesweepersvc"
	"testing"
	"time"
)

func Test_GetUser(t *testing.T) {
	tomas := "tomas"
	now := time.Now()

	userFound := &minesweepersvc.User{
		Username:  tomas,
		CreatedAt: now,
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockDB := NewMockDB(controller)

	engineStorage := minesweepersvc.NewGameEngineStorage(mockDB)

	mockDB.
		EXPECT().
		Get(tomas).
		Return(userFound, nil)

	user, err := engineStorage.GetUser(tomas)

	if err != nil {
		t.Error(err.Error())
		t.Log("error should be nil")
		t.Fail()
	}

	if user == nil {
		t.Error("user should not be nil")
		t.Fail()
	}

	if user.Username != userFound.Username {
		t.Fail()
	}
}

func Test_GetUser_NotFound(t *testing.T) {
	user1 := "tomas"
	errMsg := "user not found"

	user := &minesweepersvc.User{
		Username:  user1,
		CreatedAt: time.Now(),
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDB := NewMockDB(controller)

	mockDB.EXPECT().Get(user1).Return(nil, errors.New(errMsg))

	engineStorage := minesweepersvc.NewGameEngineStorage(mockDB)

	user, err := engineStorage.GetUser(user1)

	if user != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}

	if err.Error() != errMsg {
		t.Fail()
	}
	fmt.Println("success")
}

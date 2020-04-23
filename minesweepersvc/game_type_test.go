package minesweepersvc

import (
	"github.com/golang/mock/gomock"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"testing"
	"time"
)

func Test_createGame(t *testing.T) {
	logs.InitDefault("test")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	gameStorageMock := NewMockMineSweeperGameStorage(mockCtrl)
	s := MSGameService{gameStorageMock}
	gameStorageMock.EXPECT().GetUser("some").Return(&User{
		Username:  "some",
		CreatedAt: time.Now(),
	}, nil)

	game := &Game{
		Name:         "",
		Rows:         5,
		Cols:         5,
		Mines:        5,
		Status:       nil,
		S:            "",
		Grid:         nil,
		ClickCounter: 0,
		Username:     "some",
		CreatedAt:    time.Time{},
		StartedAt:    time.Time{},
		TimeSpent:    0,
		Points:       0,
	}

	gameStorageMock.EXPECT().CreateGame(game).Return(nil)
	err := s.CreateGame(game)

	if err != nil {
		t.Fail()
	}

	if game.Name == "" {
		t.Fail()
	}
}

func Test_createGameWithName(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	gameStorageMock := NewMockMineSweeperGameStorage(mockCtrl)
	s := MSGameService{gameStorageMock}
	gameStorageMock.EXPECT().GetUser("some").Return(&User{
		Username:  "some",
		CreatedAt: time.Now(),
	}, nil)

	game := &Game{
		Name:         "game1",
		Rows:         5,
		Cols:         5,
		Mines:        5,
		Status:       nil,
		S:            "",
		Grid:         nil,
		ClickCounter: 0,
		Username:     "some",
		CreatedAt:    time.Time{},
		StartedAt:    time.Time{},
		TimeSpent:    0,
		Points:       0,
	}

	gameStorageMock.EXPECT().CreateGame(game).Return(nil)
	err := s.CreateGame(game)

	if err != nil {
		t.Fail()
	}

	if game.Name != "game1" {
		t.Fail()
	}

	if game.S != "new" {
		t.Fail()
	}
}

func Test_createGameFail_WithoutUsername(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	gameStorageMock := NewMockMineSweeperGameStorage(mockCtrl)
	s := MSGameService{gameStorageMock}

	game := &Game{
		Name:         "game1",
		Rows:         5,
		Cols:         5,
		Mines:        5,
		Status:       nil,
		S:            "",
		Grid:         nil,
		ClickCounter: 0,
		Username:     "",
		CreatedAt:    time.Time{},
		StartedAt:    time.Time{},
		TimeSpent:    0,
		Points:       0,
	}

	err := s.CreateGame(game)

	if err != nil {
		t.Log("should fail without username" + err.Error())
		return
	}
	t.Fail()
}

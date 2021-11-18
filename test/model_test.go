package test

import (
	"gRPCServer/model"
	"testing"
)

func TestCheckListInsert(t *testing.T) {
	data := model.Todo{
		FileName: "21IHPA0000A_100101_005332_DDA142164623.mp4",
	}
	data.Insert()
}

func TestCheckListUpdate(t *testing.T) {
	data := model.Todo{
		FileName: "21IHPA0000A_100101_005332_DDA142164623.mp4",
		Project:  true,
		Storage:  true,
	}
	data.Update()
}

func TestCheckListRemove(t *testing.T) {
	data := model.Todo{
		FileName: "21IHPA0000A_100101_005332_DDA142164623.mp4",
	}
	data.Delete()
}

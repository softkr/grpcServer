package test

import (
	"fmt"
	db "gRPCServer/model"
	"testing"
)

func TestConnect(t *testing.T) {
	info := db.DBType{
		Database:   "iot",
		Collection: "watches",
	}
	db.Collection(&info)
}

func TestWatchSocketStatus(t *testing.T) {
	on := db.SocketType{
		Guid:   "21IHPA0000A",
		Status: "on",
		Addr:   "1.244.141.56:36116",
	}
	db.WatchSocketStatus(&on)

	off := db.SocketType{
		Guid:   "21IHPA0000A",
		Status: "off",
		Addr:   "1.244.141.56:36116",
	}
	db.WatchSocketStatus(&off)
}

func TestWatchUpdate(t *testing.T) {
	up := db.SocketType{
		Guid:         "21IHPA0000A",
		Wear:         10,
		TakeMedicine: 10,
	}
	db.WatchUpdate(&up)
}

func TestFileCollection(t *testing.T) {
	db.FileCollection()
}

/******** 프로젝트 *******/
//프로젝트 관리
func TestInserts(t *testing.T) {
	data := db.ProjectType{
		Code:       "ihp007scm",
		Name:       "속초의료원 치매환자 프로젝트",
		Url:        "mongodb://localhost:27017",
		Account:    "iotservertest",
		AccessKey: "ZrbbWFyUZHCUCdy7+rKtpqI5lGmRZdg3AZ4jKgtYIT9c4bCKMEp0uzkz8/BJ9/nLQFJBPFk/dRgaUT0SQCAI9A==",
	}
	prj, err := db.Projects(&data)
	if err != nil {
		fmt.Println(err)
	}
	//입력
	prj.Insert()

}

// 유저 관리
func TestUsers(t *testing.T) {
	data := db.UserType{
		Guid:    "21IHPA0000A",
		Project: "ihp007scm",
	}
	user := db.Users(&data)
	user.Insert()
}

// 프로젝트 정보 찾기
func TestFindProject(t *testing.T) {
	project := db.FindProject("21IHPA00000A")
	fmt.Println(project)
}

// 파일 기록 생성
func TestInsert(t *testing.T) {
	args := []string{
		"a376ceeae4a2b0bee79aba4d6c100080",
		"d31119f325db61d41358722652b66c81",
		"2c0134ae4f33525c367ab70d7849603e",
		"7440ee1217367722141360fbb2c4d196",
		"7e524e73d435eb227a6f19de62caaa44",
		"e50ca16a818567a47f6fc1e4318f9f7d",
		"611e2056f0703ec455aea39fd76b46de",
		"e04ab7b88242b409fe592fea9bb96257",
	}
	data := db.FileInfoType{
		Guid:         "21IHPA02720B",
		FileName:     "100101_005332_DDA142164623.mp4",
		VideoMD5:     "ad4458a45e68ca64e2301ebe975f927f",
		SubFile:      args,
		SubFileCount: 0,
	}
	db.FileInsert(&data)
}

// 파일 존재 여부
func TestFine(t *testing.T) {
	dd := db.FileFind("e04ab7b88242b409fe592fea9bb96257")
	fmt.Println(dd)
}

// 서브파일 카운트 업데이트
func TestFileUpdate(t *testing.T) {
	db.FileUpdate("611e2056f0703ec455aea39fd76b46de")
}

// 서브 파일 부족 개수 카운트
func TestFileCount(t *testing.T) {
	dd := db.FileCount("ad4458a45e68ca64e2301ebe975f927f")
	fmt.Println(dd)
}

// 파일 삭제
func TestFileDelete(t *testing.T) {
	db.FileDeleteOne("ad4458a45e68ca64e2301ebe975f927f")
}

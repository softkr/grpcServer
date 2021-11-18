package main

import (
	"context"
	db "gRPCServer/model"
	pb "gRPCServer/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

type server struct {
	pb.UnimplementedGreeterServer
}

/**** 프로젝트 *******/

func (s *server) TodoListInsert(ctx context.Context, in *pb.CheckList) (*pb.CheckListReply, error) {
	data := db.Todo{
		FileName: in.GetFilename(),
		Project:  in.GetProject(),
		Storage:  in.GetStorage(),
		Api:      in.GetApi(),
		Remove:   in.GetRemove(),
	}
	data.Insert()
	return &pb.CheckListReply{Status: 200, Message: "정상처리 되었습니다."}, nil
}

func (s *server) TodoListUpdate(ctx context.Context, in *pb.CheckList) (*pb.CheckListReply, error) {
	data := db.Todo{
		FileName: in.GetFilename(),
		Project:  in.GetProject(),
		Storage:  in.GetStorage(),
		Api:      in.GetApi(),
		Remove:   in.GetRemove(),
	}
	data.Update()
	return &pb.CheckListReply{Status: 200, Message: "정상처리 되었습니다."}, nil
}
func (s *server) TodoListRemove(ctx context.Context, in *pb.CheckList) (*pb.CheckListReply, error) {
	data := db.Todo{
		FileName: in.GetFilename(),
	}
	data.Delete()
	return &pb.CheckListReply{Status: 200, Message: "정상처리 되었습니다."}, nil
}

//등록

// 찾기
func (s *server) GetProject(ctx context.Context, in *pb.ProjectRequest) (*pb.ProjectReply, error) {
	data := db.Users{
		Guid: in.GetGuid(),
	}
	project := data.Find()
	return &pb.ProjectReply{Status: 200, Project: &project}, nil
}

/*** 워치 ****/

// WatchStatus 워치 온라인 오프라인
func (s *server) WatchStatus(ctx context.Context, in *pb.WatchState) (*pb.WatchStateReply, error) {
	statues := db.Sockets{
		Guid:   in.GetSn(),
		Status: in.GetStatus().String(),
		Addr:   in.GetAddr(),
	}
	statues.OnOff()
	return &pb.WatchStateReply{Status: 200, Message: "성공했습니다."}, nil
}

// 워치 상태 업데이트
func (s *server) WatchUpdate(ctx context.Context, in *pb.WatchUpdates) (*pb.WatchStateReply, error) {
	data := db.Sockets{
		Guid:         in.GetSn(),
		Wear:         in.GetWear(),
		TakeMedicine: in.GetTakeMedicine(),
	}
	data.Update()
	return &pb.WatchStateReply{Status: 200, Message: "성공했습니다."}, nil
}

/*****워치 끝*****/

/********동영상파일***********/

// SetFileInfo 파일 업로드전 파일 상세정보 저장
func (s *server) SetFileInfo(ctx context.Context, in *pb.SetFileInfoRequest) (*pb.SetFileInfoReply, error) {
	subFile := strings.Split(in.SubFile, "|")
	data := db.Files{
		Guid:         in.GetGuid(),
		FileName:     in.FileName,
		VideoMD5:     in.VideoMd5,
		SubFile:      subFile,
		SubFileCount: 0,
	}
	data.Insert()
	return &pb.SetFileInfoReply{Message: in.GetSubFile()}, nil
}

// PutFileInfo 서브파일 업로드 카운트 추가
func (s *server) PutFileInfo(ctx context.Context, in *pb.PutFileInfoRequest) (*pb.PutFileInfoReply, error) {
	data := db.Files{
		Value: in.GetSubFile(),
	}
	result := data.Update()
	return &pb.PutFileInfoReply{SubFileCount: result}, nil
}

// FindSubFile 서브파일 존재여부 확인
func (s *server) FindSubFile(ctx context.Context, in *pb.GetFindFileInfoRequest) (*pb.GetFindFileInReply, error) {
	data := db.Files{
		Value: in.GetSubFile(),
	}
	result := data.Find()
	subFile := strings.Join(result.SubFile, " ")
	return &pb.GetFindFileInReply{Guid: result.Guid, FileName: result.FileName, VideoMd5: result.VideoMD5, SubFile: subFile}, nil
}

// DeleteFileInfo 업로드 파일 끝나면 파일정도 삭제
func (s *server) DeleteFileInfo(ctx context.Context, in *pb.RemoveFileInfoRequest) (*pb.RemoveFileInfoReply, error) {
	data := db.Files{
		Value: in.GetVideoMd5(),
	}
	data.Remove()
	return &pb.RemoveFileInfoReply{
		Message: "ok",
	}, nil
}

// SubFileCount 서브파일 숫자 카운팅하는 함수임
func (s *server) SubFileCount(ctx context.Context, in *pb.SubFileCountRequest) (*pb.SubFileCountReply, error) {
	data := db.Files{
		Value: in.GetVideoMd5(),
	}
	count := data.Count()
	return &pb.SubFileCountReply{
		Message: count,
	}, nil
}

/*******동영상 파일 끝********/

func main() {
	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

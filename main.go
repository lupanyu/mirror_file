package main

import (
	"fmt"
	"mirror_file/task"
)


type Client interface {
	Connect()          //与服务端建立链接
	Md5(filename string) (map[string][]byte)
	AbsFileList(dir string,f *tools.FileList)
	SendFile()         //发送文件
}


type Server interface {
	Listen(string) (error)           //打开监听
	Md5(filename string) ([]byte)    //
	AbsFileList(dir string,f *tools.FileList)
	SaveFile(filename string)
}

func main() {
	var config tools.Conf
	f :=config.Config()

	//hostAddr := f.Host

	var filelist tools.FileList
	tools.ListDir(f.Dir, &filelist)
	filelist.Relative(f.Dir)
	fmt.Println(filelist)
	fmt.Println(len(filelist))
	//根据类型字段 来启动不同类型的服务
/*	if f.ServerType == "server"{
		var run Server
		e := run.Listen(hostAddr)
		if e != nil {
			fmt.Println(e.Error())
		}
	}
	if f.ServerType == "client"{
		var run Client
		run.Connect()
	}
*/
}

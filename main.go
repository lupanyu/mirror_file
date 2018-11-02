package main

import (
	"bufio"
	"fmt"
	"io"
	"mirror_file/task"
)


type Client interface {
	Connect( addr string ) (*bufio.ReadWriter, error)         //与服务端建立链接
	Md5(filename string) (map[string][]byte)
	AbsFileList(dir string,f *task.FileList)
	SendFile()         //发送文件
}



type Server interface {
	Listen() (error)           //打开监听
	Md5(filename string) ([]byte)    //
	AbsFileList(dir string,f *task.FileList)
	SaveFile(filename string)
}

func main() {
	var config task.Conf
	f :=config.Config("conf.yaml")


	if f.ServerType == "server"{
		//var run Server
		con ,e := task.Listen(f.Host)
		if e != nil {
			fmt.Println(e.Error())
		}
		data := make([]byte,1024)
		for {
		n, err2 := con.Read(data)
		if err2 == io.EOF {
			fmt.Println("du wan le ")
			return
		}
		fmt.Println(n,string(data[:n]))}
	}
	if f.ServerType == "client"{
		var run Client
		run.Connect(f.Host)
	}

}

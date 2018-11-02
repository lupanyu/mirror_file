package main

import (
	"bufio"
	"fmt"
	"mirror_file/task"
	"mirror_file_client/task"
)

type FileOption interface {
	Md5Count()
	Relative(dir string )
}
type Clienter interface {
	Connect( addr string ) (*bufio.ReadWriter, error)         //与服务端建立链接
	Md5(filename string) (map[string][]byte)
	AbsFileList(dir string,f *task.FileList)
	SendFile()         //发送文件
}

func main() {
	var config task.Conf
	f :=config.Config("client.yaml")

	//hostAddr := f.Host

	var filelist tools.FileList
	tools.ListDir(f.Dir, &filelist)
	filelist.Relative(f.Dir)
	//fmt.Println(filelist)
	//fmt.Println(len(filelist))
	//根据类型字段 来启动不同类型的服务
	if f.ServerType == "client"{


		m, e := task.Connect(f.Host)
		if e != nil {
			fmt.Println(e.Error(),"\n--------------connect error-------------------\n")
		}
		// a :=task.Pack{PackType:"test",PackData:"testdata"}

		m.WriteString(`{"PackType":"test","PackData":"testdata"}`)
		m.Flush()
	}

}

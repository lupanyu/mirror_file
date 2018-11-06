package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"mirror_file/task"
	"mirror_file_client/task"
	"os"
)

type FileOption interface {
	Md5Count()
	Relative(dir string )
}
type Clienter interface {
	Connect( addr string ) (*bufio.ReadWriter, error)         //与服务端建立链接
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
	file, e := os.Open("/var/log/messages")
	if e != nil {
	}
	byte_file, e := ioutil.ReadAll(file)
	if e != nil {
		log.Panic(e)
	}
	if f.ServerType == "client"{


		m, e := task.Connect(f.Host)
		if e != nil {
			fmt.Println(e.Error(),"\n--------------connect error-------------------\n")
		}
		// a :=task.Pack{PackType:"test",PackData:"testdata"}
		byte_data := make([]byte,1024)
		data := `{"PackType":"file","PackData":"message"}`
		copy(byte_data, data)
		fmt.Println(byte_data)
		//nn, err := m.WriteString(`{"PackType":"file","PackData":"testdata"}`)
		nn , err := m.Write(byte_data)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("写入%v字节",nn)
		m.Flush()

	    nn,err =  m.Write(byte_file)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("写入%v字节",nn)

		m.Flush()
	    }


}

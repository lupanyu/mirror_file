package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"mirror_file/pack"
	"mirror_file/task"
	//	"mirror_file_client/task"
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


func MakePack(file *file,readsize int){
	ReadSize := make([]byte,readsize) //指定每次读取的大小为1024。
	ReadByte := make([]byte,4096,4096) //指定读取到的字节数。
	r := bufio.NewReader(&file)
	ActualSize,err := r.Read(ReadSize)
	ReadByte = append(ReadByte,ReadSize[:ActualSize]...)
	test1 := &pack.File{
		FileName: proto.String(file.FileName),
		RelName:  proto.String(file.RelFile),
		Md5: 	  proto.String(file.Md5),
		FileData: byte_file,
		FileEnd: proto.Bool(true),
	}
}

func main() {
	var config task.Conf
	f :=config.Config("client.yaml")

	//建立链接
	if f.ServerType == "client"{
		m, e := task.Connect(f.Host)
		if e != nil {
			fmt.Println(e.Error(),"\n--------------connect error-------------------\n")
		}
		// a :=task.Pack{PackType:"test",PackData:"testdata"}

	var filelist task.FileList
	task.ListDir(f.Dir, &filelist)
	filelist.Relative(f.Dir)
	//fmt.Println(filelist)
	for _,file := range filelist{
		//filename := file.RelFile
		files, e := os.Open(file.AbsFile)
		if e != nil {
			log.Println(e)
		}
		defer files.Close()
		byte_file, e := ioutil.ReadAll(files)
		test1 := &pack.File{
			FileName: proto.String(file.FileName),
			RelName:  proto.String(file.RelFile),
			Md5: 	  proto.String(file.Md5),
			FileData: byte_file,
			FileEnd: proto.Bool(true),
		}
		//fmt.Println(test1)
		bytedata, i := proto.Marshal(test1)
		if i != nil {
			log.Panic("--------56--------\n",i)
		}
		fmt.Println("60")
		n := len(bytedata)/102400
		for i:=0 ; i < n; i++ {
			fmt.Println(i)
			nn, err := m.Write(bytedata[i*102400:(i+1)*102400])
			if err != nil {
				log.Panic(err)
			}
			m.Flush()
			log.Printf("send %v 字节",nn,file.RelFile)
		}
		fmt.Println(n)
		m.Write(bytedata[n*102400:])
		m.Flush()
		fmt.Println("65")

		for {
			data := make([]byte, 4096)
			n, err := m.Read(data)
			if err != nil {
				log.Println(err, "------------66-----------")
			}
			//去掉 []byte中的空位置
			index := bytes.IndexByte(data, 0)
			rbyf_pn := data[0:index]
			fmt.Println("收到字节数:", n, rbyf_pn)
			if n < 4096 {
				break
			}
			fmt.Println("xia yi ge !!")
		}
		//发送文件
		//err := task.SendFile(filename,m,&byte_file)
		//if err != nil {
		//	log.Println("send file err ",err)
		//}
	}
	}
}

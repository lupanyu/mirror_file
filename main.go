package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"mirror_file/task"
	"time"
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
	f := config.Config("conf.yaml")

	if f.ServerType == "server" {
		//var run Server
		con, e := task.Listen(f.Host)
		if e != nil {
			fmt.Println(e.Error())
		}


		for {
			time.Sleep(time.Second*10)
			data := make([]byte, 1024)
			n, err2 := con.Read(data)
			//去掉 []byte中的空位置
			index := bytes.IndexByte(data, 0)
			rbyf_pn := data[0:index]

			if err2 != nil  {
				fmt.Println(err2)

			}
			if n == 0 {
				fmt.Println("没有收到数据")
				continue
			}
			fmt.Println("收到了%v个字节",n)
			pack := task.Pack{}
			e := pack.DecodeJson(rbyf_pn)
			if e != nil {
				log.Printf("in convert json :%#v",e)
				continue
			}
			if pack.PackType == "file"{
				//var filename string

				filename := f.Dir + "/" + pack.PackData.(string)

				log.Printf("收到客户端请求 写入文件%v",filename)
				e := task.SaveFile(con,data,filename)
				if e != nil {
					fmt.Printf("写入文件%v失败:%v",filename,e)
				}

			}
		}


	}
}
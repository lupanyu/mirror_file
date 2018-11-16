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
	"net"
	"os"
	"strings"
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


func ConRead(conn net.Conn)([]byte){
	var buffer bytes.Buffer
	//var jieshu bool = false
		READCON:
		subdata := make([]byte,409600)
		n, err := conn.Read(subdata)
		fmt.Println(n)
		if n == 0 {
			fmt.Println("du bu dao shuju")
			return buffer.Bytes()
		}
		if err != nil  {
			fmt.Println(err)

		}
		//去掉 []byte中的空位置
		index := bytes.IndexByte(subdata, 0)
		//fmt.Println(index,reflect.TypeOf(index))
		if index == -1 {
			//rbyf_pn := subdata[0:index]
			fmt.Println("ba数据写到了buffer")
			buffer.Write(subdata)
			goto READCON

		}
		if index < 409600{
			rbyf_pn := subdata[0:index]
			buffer.Write(rbyf_pn)
			return buffer.Bytes()
		}
		if index == 0 {
			return buffer.Bytes()
		}
			return buffer.Bytes()
}

func PackSaveFile(basedir string,p *pack.File)error{
 	md5 := p.GetMd5()
 	fmt.Println("md5:[",md5,"]")
 	fileData := p.GetFileData()
 	absFile := basedir  + p.GetRelName()
	absDir := strings.Replace(absFile,p.GetFileName(),"",1)
	//如果目录不存在 新建目录
	_,err := os.Stat(absDir)
	if os.IsNotExist(err){
		os.MkdirAll(absDir,0755)
	}
	//写文件内容
	err = ioutil.WriteFile(absFile, fileData, 0644)
	if err != nil {
		log.Println("---------73-----",p.GetFileName(),p.GetRelName(),p.GetFileMode())
		log.Println("-----------in save file-------------",absFile,err)
		return err
	}
	return nil
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
		NEXT:
		for {

			time.Sleep(time.Second*3)
			fmt.Println(time.Local)
			//一次吧数据读完
			data:= ConRead(con)

			if data == nil {
				fmt.Println("no data")
				continue
			}
			fmt.Println("shoudao shuju :",len(data))
			file := &pack.File{}
			//解码
			err := proto.Unmarshal(data,file)

			fmt.Println(file,"-------------------")

			result := &pack.Result{Ok:proto.Bool(true),Info:proto.String("fdsa"),}
			results ,err := proto.Marshal(result)
			if err != nil {
				log.Println("解码出现意外：",err)
				fmt.Println("fanhui shuju:",results)
				con.Write(results)
				goto NEXT
			}
			err = PackSaveFile(f.Dir,file)
			if err != nil {
				log.Println("保存文件失败",file.GetFileName())
			}
		//	result := &pack.Result{Ok:proto.Bool(true),Info:proto.String("fdsa"),}
		//	results ,err := proto.Marshal(result)
			if err != nil{
				log.Println(err)
			}
			con.Write(results)
		}
		/*for {
//			time.Sleep(time.Second*10)
			data := make([]byte, 40960)
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

*/
	}
}
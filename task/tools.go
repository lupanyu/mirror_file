package tools

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)


type File struct {

	AbsFile 	string
	Md5 		string
	RelFile 	string
}

type FileList []File

//为FileList更新每个文件的RelFile信息
func (f *FileList)Relative(dir string ) {

	for k, file := range *f {
		//fmt.Println(file,dir)
		(*f)[k].RelFile = strings.Replace(file.AbsFile, dir,"",1) //方法一
		//方法2
		//file.RelFile = strings.Replace(file.AbsFile, dir,"",1)
		//(*f)[k] = file
	}

}


//返回文件的md5
func (f *File)Md5Count(){
	//var md5str map[string][]byte

	file, inerr := os.Open(f.AbsFile)
	defer file.Close()
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, file)
		md51  := md5h.Sum([]byte(""))
		f.Md5 = fmt.Sprintf("%x",md51)
	}
}


func ListDir(dir string,d *FileList) {
	//var filelist []string
	files, _ := ioutil.ReadDir(dir)
	for  _,file :=  range files{
		var f File
		if file.IsDir(){
			subDir := dir + "/" + file.Name()
			ListDir(subDir,d)
		}
		//fmt.Println(file.Name())
		f.AbsFile = dir+"/"+file.Name()
		if  F,_ := os.Stat(f.AbsFile) ; ! F.IsDir() {
			f.Md5Count()
			//f.RelFile = file.Name()
			*d = append(*d,f)
		}
	}
}
type Conf struct{
	ServerType  string `yaml:"serverType"`
	Host 		string `yaml:"host"`
	//Port 		string `yaml:"port"`
	Dir 		string `yaml:"dir"`
}
//加载配置文件
func (c *Conf)Config() *Conf{
	confFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil{
		panic(err)
	}
	err = yaml.Unmarshal(confFile,c)
	if err != nil {
		panic(err.Error())
	}
	return c
}

//数据包的格式
type Pack struct {
	PackType interface{}
	PackData interface{}
}

func (p *Pack)DecodeJson(data []byte){
	e := json.Unmarshal(data,p)
	if e != nil {
		log.Fatal(e)
	}
}
///	data := []byte(`{"PackType":"127.0.0.1","PackData":123}`)
//	m := Pack{}
//	m.DecodeJson(data)
//	fmt.Println(m)


func Connect(){

}
//处理客户端发送的数据
func Worker(conn net.Conn){

}
type Tcp struct {

}

func (*Tcp)Listen(host string)error{
	fmt.Println("listen start")
	conn, e := net.Listen("tcp", host)
	if e != nil {
		return errors.New(e.Error() + "TCP服务无法监听在端口"+host)
	}
	fmt.Println("listen ok")
	for {
		conn, err := conn.Accept()
		if err != nil{
			fmt.Println("心请求监听失败!")
			continue
		}
		// 开始处理新链接数据
		//go e.handleMessage(conn)
		Worker(conn)
	}
}
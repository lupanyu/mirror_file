package task

import (
	"bufio"
	"bytes"
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
	FileName    string
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
			f.FileName = file.Name()
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
func (c *Conf)Config(filename string) *Conf{
	confFile, err := ioutil.ReadFile(filename)
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
//解析 json形式的包为Pack类型
func (p *Pack)DecodeJson(data []byte)error{
	e := json.Unmarshal(data,p)
	if e != nil {
		return e
	}
	return nil
}
///	data := []byte(`{"PackType":"127.0.0.1","PackData":123}`)
//	m := Pack{}
//	m.DecodeJson(data)
//	fmt.Println(m)



//保存文件
func SaveFile(conn net.Conn,data []byte,filename string,)error{
	var buffer bytes.Buffer
	for {
		i, err := conn.Read(data)
		buffer.Write(data[:i])
		if err != nil {
			log.Println(err)
			if err == io.EOF{break}
		}
		}
	fmt.Printf("文件长度是:%v",buffer.Len())
	file_err := ioutil.WriteFile(filename, buffer.Bytes(), 0644)
	if file_err != nil {
		log.Println(file_err)
	}
	return nil
}


//发送文件
func SendFile(filename string,m *bufio.ReadWriter,byte_file *[]byte )error{
	byte_data := make([]byte,1024)
	data := fmt.Sprintf(`{"PackType":"file","PackData":"%v"}`,filename)
	copy(byte_data, data)
	fmt.Println(byte_data)
	//nn, err := m.WriteString(`{"PackType":"file","PackData":"testdata"}`)
	nn , err := m.Write(byte_data)
	if err != nil {
		return err
	}
	fmt.Printf("写入%v字节",nn)
	m.Flush()

	nn,err =  m.Write(*byte_file)
	if err != nil {
		return  err
	}
	fmt.Printf("写入%v字节",nn)

	m.Flush()
	return nil
}



//返回一个有超时的TCP链接缓冲readwrite
func Connect(addr string) (*bufio.ReadWriter, error) {
	// Dial the remote process.
	// Note that the local port is chosen on the fly. If the local port
	// must be a specific one, use DialTCP() instead.
	//fmt.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}
//处理客户端发送的数据
func Worker(conn net.Conn){

}
type Tcp struct {

}

func Listen(addr string )(net.Conn,error){
	fmt.Println("listen start")
	conn, e := net.Listen("tcp", addr )
	if e != nil {
		return nil,errors.New(e.Error() + "TCP服务无法监听在端口"+addr )
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
		return conn,nil
	}
}
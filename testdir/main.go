package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type FileList []string

//返回所有文件列表,软连会被认为成文件
func AbsFileList(dir string,f *FileList) {
	//var filelist []string
	files, _ := ioutil.ReadDir(dir)

	for  _,file :=  range files{
		fmt.Println(file)
		if file.IsDir(){
			subDir := dir + "/" + file.Name()
			AbsFileList(subDir,f)
		}
		//fmt.Println(file.Name())
		absfile := dir+"/"+file.Name()
		if  F,_ := os.Stat(absfile) ; ! F.IsDir() {
			*f = append(*f, absfile)
		}
	}
}

//
//为FileList更新每个文件的RelFile信息
func (f *File)Relative(dir string ) {


		//fmt.Println(file,dir)
		(*f).RelFile = strings.Replace(f.AbsFile, dir,"",1) //方法一
		//方法2
		//file.RelFile = strings.Replace(file.AbsFile, dir,"",1)
		//(*f)[k] = file


}

type CompleData struct {
	S string
	Ms map[string]string
	Mn map[string]int
	N int
	B []byte
	C *CompleData
}

func HandleGob(rw *bufio.ReadWriter){
	var data CompleData

	dec := gob.NewDecoder(rw)
	err := dec.Decode(&data)
	if err != nil {
		fmt.Println("无法解析的数据")
		return
	}
	fmt.Println("输出",data,data.C)
}

type conn interface {
	abc() string
}

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
type File struct {

	AbsFile 	string
	Md5 		string
	RelFile 	string
}

type FileOption interface {
	Md5Count()
	Relative(dir string )
}

func (p *Pack)BecodeJson(data []byte)error{
	e := json.Unmarshal(data,p)
	if e != nil {
		return e
	}
	return nil
}

func (f *File)Md5Count(){
	//var md5str map[string][]byte

	file, inerr := os.Open(f.AbsFile)
	defer file.Close()
	if inerr != nil {
		panic(inerr)
	}
	md5h := md5.New()
	io.Copy(md5h, file)
	md51  := md5h.Sum([]byte(""))
	f.Md5 = fmt.Sprintf("%x",md51)
}

func main() {
	//var f FileOption
	//local := File{"/var/log/messages","",""}
	var buf bytes.Buffer
	buf.Write([]byte(`{"PackType":"file","PackData":"testdata"}`))
	fmt.Println(buf)
	type pack Pack
	Pack.BecodeJson(buf)
}

package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
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
func (f FileList)Relative(dir string)[]string{
	var relaiveFileList []string
	for _,file :=  range  f {
		relaiveFile := strings.Trim(file,dir)
		relaiveFileList = append(relaiveFileList, relaiveFile)
	}
	return relaiveFileList

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

func main() {
	//var f FileList
	//AbsFileList("/var/log",&f)
	data := []byte(`{"PackType":"127.0.0.1","PackData":123}`)
	m := Pack{}
	m.DecodeJson(data)
	fmt.Println(m)

}

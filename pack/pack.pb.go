// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test1.proto

package pack


import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//客户端发送
type File struct {
	FileName             *string  `protobuf:"bytes,1,req,name=FileName" json:"FileName,omitempty"`
	RelName              *string  `protobuf:"bytes,2,opt,name=RelName" json:"RelName,omitempty"`
	Md5                  *string  `protobuf:"bytes,3,opt,name=Md5" json:"Md5,omitempty"`
	FileData             []byte   `protobuf:"bytes,4,opt,name=FileData" json:"FileData,omitempty"`
	FileMode             *string  `protobuf:"bytes,5,opt,name=FileMode" json:"FileMode,omitempty"`
	FileBlock            *int32   `protobuf:"varint,6,opt,name=FileBlock" json:"FileBlock,omitempty"`
	FileEnd              *bool    `protobuf:"varint,7,opt,name=FileEnd,def=1" json:"FileEnd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_a038b2eef6c1cfff, []int{0}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

const Default_File_FileEnd bool = true

func (m *File) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *File) GetRelName() string {
	if m != nil && m.RelName != nil {
		return *m.RelName
	}
	return ""
}

func (m *File) GetMd5() string {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return ""
}

func (m *File) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

func (m *File) GetFileMode() string {
	if m != nil && m.FileMode != nil {
		return *m.FileMode
	}
	return ""
}

func (m *File) GetFileBlock() int32 {
	if m != nil && m.FileBlock != nil {
		return *m.FileBlock
	}
	return 0
}

func (m *File) GetFileEnd() bool {
	if m != nil && m.FileEnd != nil {
		return *m.FileEnd
	}
	return Default_File_FileEnd
}

type FileList struct {
	Test1                *File    `protobuf:"bytes,1,opt,name=test1" json:"test1,omitempty"`
	Test2                *File    `protobuf:"bytes,2,opt,name=test2" json:"test2,omitempty"`
	Test3                *File    `protobuf:"bytes,3,opt,name=test3" json:"test3,omitempty"`
	Test4                *File    `protobuf:"bytes,4,opt,name=test4" json:"test4,omitempty"`
	Test5                *File    `protobuf:"bytes,5,opt,name=test5" json:"test5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileList) Reset()         { *m = FileList{} }
func (m *FileList) String() string { return proto.CompactTextString(m) }
func (*FileList) ProtoMessage()    {}
func (*FileList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a038b2eef6c1cfff, []int{1}
}

func (m *FileList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileList.Unmarshal(m, b)
}
func (m *FileList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileList.Marshal(b, m, deterministic)
}
func (m *FileList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileList.Merge(m, src)
}
func (m *FileList) XXX_Size() int {
	return xxx_messageInfo_FileList.Size(m)
}
func (m *FileList) XXX_DiscardUnknown() {
	xxx_messageInfo_FileList.DiscardUnknown(m)
}

var xxx_messageInfo_FileList proto.InternalMessageInfo

func (m *FileList) GetTest1() *File {
	if m != nil {
		return m.Test1
	}
	return nil
}

func (m *FileList) GetTest2() *File {
	if m != nil {
		return m.Test2
	}
	return nil
}

func (m *FileList) GetTest3() *File {
	if m != nil {
		return m.Test3
	}
	return nil
}

func (m *FileList) GetTest4() *File {
	if m != nil {
		return m.Test4
	}
	return nil
}

func (m *FileList) GetTest5() *File {
	if m != nil {
		return m.Test5
	}
	return nil
}

//从服务器返回
type Result struct {
	Fileinfo             *File    `protobuf:"bytes,1,opt,name=fileinfo" json:"fileinfo,omitempty"`
	Ok                   *bool    `protobuf:"varint,2,opt,name=ok,def=0" json:"ok,omitempty"`
	Info                 *string  `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_a038b2eef6c1cfff, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

const Default_Result_Ok bool = false

func (m *Result) GetFileinfo() *File {
	if m != nil {
		return m.Fileinfo
	}
	return nil
}

func (m *Result) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return Default_Result_Ok
}

func (m *Result) GetInfo() string {
	if m != nil && m.Info != nil {
		return *m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*File)(nil), "File")
	proto.RegisterType((*FileList)(nil), "FileList")
	proto.RegisterType((*Result)(nil), "Result")
}

func init() { proto.RegisterFile("test1.proto", fileDescriptor_a038b2eef6c1cfff) }

var fileDescriptor_a038b2eef6c1cfff = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x80, 0x61, 0x66, 0xbb, 0xe9, 0x66, 0x67, 0x0b, 0xda, 0xa0, 0x98, 0x63, 0xd8, 0x53, 0x4e,
	0x05, 0xab, 0xbd, 0xf4, 0x58, 0xd4, 0x93, 0xf5, 0xd0, 0x37, 0x08, 0xee, 0x2c, 0x84, 0xc6, 0x46,
	0x9a, 0x14, 0x7c, 0x00, 0x1f, 0x5c, 0x76, 0x6c, 0x0f, 0xdd, 0x53, 0xc8, 0xc7, 0x84, 0xf9, 0x09,
	0x36, 0x99, 0x52, 0x7e, 0x5c, 0x7c, 0x1f, 0x63, 0x8e, 0xed, 0x2f, 0x60, 0xf9, 0xe6, 0x03, 0xa9,
	0x5b, 0x94, 0xc3, 0xf9, 0xe1, 0xbe, 0x48, 0x83, 0x29, 0x6c, 0xad, 0x6e, 0xb0, 0xda, 0x51, 0x60,
	0x28, 0x0c, 0xd8, 0x5a, 0x35, 0x38, 0xd9, 0x76, 0x2b, 0x3d, 0xe1, 0xcb, 0x79, 0xfe, 0xc5, 0x65,
	0xa7, 0x4b, 0x03, 0x76, 0x76, 0x91, 0x6d, 0xec, 0x48, 0x0b, 0x9e, 0x99, 0x63, 0x3d, 0xc8, 0x26,
	0xc4, 0xcf, 0xbd, 0x9e, 0x1a, 0xb0, 0x42, 0xdd, 0x63, 0x35, 0xd0, 0xeb, 0xa1, 0xd3, 0x95, 0x01,
	0x2b, 0xd7, 0x65, 0x3e, 0x9e, 0xa8, 0xfd, 0xf9, 0x7f, 0xfb, 0xee, 0x53, 0x56, 0x77, 0x28, 0xb8,
	0x50, 0x83, 0x01, 0xdb, 0x2c, 0xc5, 0x82, 0xfb, 0xce, 0xba, 0xe4, 0x96, 0xb1, 0x3e, 0x71, 0xd4,
	0x58, 0x9f, 0x39, 0x6c, 0xac, 0x2b, 0x8e, 0xbb, 0x68, 0xbb, 0xc1, 0xe9, 0x8e, 0xd2, 0x29, 0x64,
	0xf5, 0x80, 0xb2, 0xf7, 0x81, 0xfc, 0xa1, 0x8f, 0xd7, 0xab, 0xe7, 0x58, 0xc4, 0x3d, 0xef, 0x95,
	0x6b, 0xd1, 0xbb, 0x90, 0x48, 0xcd, 0xb0, 0xe4, 0x39, 0xfe, 0x8b, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc7, 0x1c, 0x59, 0xca, 0x52, 0x01, 0x00, 0x00,
}

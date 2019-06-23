// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movie.proto

package movie

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type CreateResponse struct {
	Movie                *Movie   `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetMovie() *Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

type DeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{2}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResponse struct {
	Movie                *Movie   `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{3}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetMovie() *Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

type FindAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllRequest) Reset()         { *m = FindAllRequest{} }
func (m *FindAllRequest) String() string { return proto.CompactTextString(m) }
func (*FindAllRequest) ProtoMessage()    {}
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{4}
}

func (m *FindAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllRequest.Unmarshal(m, b)
}
func (m *FindAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllRequest.Marshal(b, m, deterministic)
}
func (m *FindAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllRequest.Merge(m, src)
}
func (m *FindAllRequest) XXX_Size() int {
	return xxx_messageInfo_FindAllRequest.Size(m)
}
func (m *FindAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllRequest proto.InternalMessageInfo

type FindAllResponse struct {
	Movies               []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllResponse) Reset()         { *m = FindAllResponse{} }
func (m *FindAllResponse) String() string { return proto.CompactTextString(m) }
func (*FindAllResponse) ProtoMessage()    {}
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{5}
}

func (m *FindAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllResponse.Unmarshal(m, b)
}
func (m *FindAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllResponse.Marshal(b, m, deterministic)
}
func (m *FindAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllResponse.Merge(m, src)
}
func (m *FindAllResponse) XXX_Size() int {
	return xxx_messageInfo_FindAllResponse.Size(m)
}
func (m *FindAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllResponse proto.InternalMessageInfo

func (m *FindAllResponse) GetMovies() []*Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

type Movie struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{6}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*FindAllRequest)(nil), "FindAllRequest")
	proto.RegisterType((*FindAllResponse)(nil), "FindAllResponse")
	proto.RegisterType((*Movie)(nil), "Movie")
}

func init() { proto.RegisterFile("movie.proto", fileDescriptor_fde087a4194eda75) }

var fileDescriptor_fde087a4194eda75 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x45, 0x99, 0x29, 0x13, 0xf1, 0xd5, 0x26, 0x25, 0xb8, 0x28, 0x83, 0x68, 0x09, 0x08, 0x16,
	0xf4, 0x81, 0xf5, 0x0b, 0x44, 0x71, 0xe7, 0x26, 0x7e, 0x81, 0x3a, 0x6f, 0x11, 0x88, 0x4d, 0x9d,
	0xc4, 0x7e, 0x88, 0x5f, 0x2c, 0x4d, 0x52, 0xdb, 0xc9, 0xaa, 0xcb, 0x77, 0xb9, 0x1c, 0x72, 0x6e,
	0x60, 0xfc, 0xe5, 0x36, 0x86, 0x70, 0xdd, 0xbb, 0xe0, 0xd4, 0x35, 0x4c, 0x9e, 0x7a, 0x7a, 0x0f,
	0xa4, 0xe9, 0xfb, 0x87, 0x7c, 0x90, 0xe7, 0xd0, 0x04, 0x13, 0x2c, 0xcd, 0xaa, 0x79, 0x75, 0x73,
	0xaa, 0xd3, 0xa1, 0x10, 0xf8, 0xae, 0xe6, 0xd7, 0x6e, 0xe5, 0x49, 0x5e, 0x40, 0x13, 0x39, 0xb1,
	0x37, 0x5e, 0x32, 0x7c, 0xdd, 0x5e, 0x3a, 0x85, 0xea, 0x0a, 0x26, 0xcf, 0x64, 0x69, 0x8f, 0xe5,
	0x50, 0x9b, 0x2e, 0x76, 0x47, 0xba, 0x36, 0xdd, 0x16, 0xb8, 0x2b, 0x1c, 0x05, 0x9c, 0x02, 0x7f,
	0x31, 0xab, 0xee, 0xd1, 0xda, 0x4c, 0x54, 0xf7, 0x20, 0xfe, 0x93, 0x8c, 0xb8, 0x04, 0x16, 0xdb,
	0x7e, 0x56, 0xcd, 0x47, 0x07, 0x8c, 0x9c, 0xaa, 0x3b, 0x68, 0x62, 0x50, 0xbe, 0x66, 0x2f, 0x5d,
	0x1f, 0x48, 0x2f, 0x7f, 0x2b, 0x38, 0x8b, 0xfd, 0x37, 0xea, 0x37, 0xe6, 0x93, 0xe4, 0x02, 0x58,
	0x5a, 0x41, 0x72, 0x1c, 0xac, 0xd6, 0x0a, 0x2c, 0xe6, 0x59, 0x00, 0x4b, 0x7e, 0x92, 0xe3, 0x60,
	0x89, 0x56, 0x60, 0x21, 0x7e, 0x0b, 0x27, 0x59, 0x44, 0x0a, 0x1c, 0x4a, 0xb6, 0x53, 0x2c, 0x1c,
	0x3f, 0x58, 0xfc, 0xb7, 0x87, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x6e, 0x0c, 0xb7, 0xc6,
	0x01, 0x00, 0x00,
}

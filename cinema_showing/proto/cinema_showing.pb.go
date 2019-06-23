// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cinema_showing.proto

package cinema_showing

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

type CreateCinemaShowingRequest struct {
	Movie                int64    `protobuf:"varint,1,opt,name=movie,proto3" json:"movie,omitempty"`
	CinemaHall           int64    `protobuf:"varint,2,opt,name=cinemaHall,proto3" json:"cinemaHall,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCinemaShowingRequest) Reset()         { *m = CreateCinemaShowingRequest{} }
func (m *CreateCinemaShowingRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCinemaShowingRequest) ProtoMessage()    {}
func (*CreateCinemaShowingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50fce743d1c301cd, []int{0}
}

func (m *CreateCinemaShowingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCinemaShowingRequest.Unmarshal(m, b)
}
func (m *CreateCinemaShowingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCinemaShowingRequest.Marshal(b, m, deterministic)
}
func (m *CreateCinemaShowingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCinemaShowingRequest.Merge(m, src)
}
func (m *CreateCinemaShowingRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCinemaShowingRequest.Size(m)
}
func (m *CreateCinemaShowingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCinemaShowingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCinemaShowingRequest proto.InternalMessageInfo

func (m *CreateCinemaShowingRequest) GetMovie() int64 {
	if m != nil {
		return m.Movie
	}
	return 0
}

func (m *CreateCinemaShowingRequest) GetCinemaHall() int64 {
	if m != nil {
		return m.CinemaHall
	}
	return 0
}

type CreateCinemaShowingResponse struct {
	Showing              *CinemaShowing `protobuf:"bytes,1,opt,name=showing,proto3" json:"showing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateCinemaShowingResponse) Reset()         { *m = CreateCinemaShowingResponse{} }
func (m *CreateCinemaShowingResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCinemaShowingResponse) ProtoMessage()    {}
func (*CreateCinemaShowingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50fce743d1c301cd, []int{1}
}

func (m *CreateCinemaShowingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCinemaShowingResponse.Unmarshal(m, b)
}
func (m *CreateCinemaShowingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCinemaShowingResponse.Marshal(b, m, deterministic)
}
func (m *CreateCinemaShowingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCinemaShowingResponse.Merge(m, src)
}
func (m *CreateCinemaShowingResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCinemaShowingResponse.Size(m)
}
func (m *CreateCinemaShowingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCinemaShowingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCinemaShowingResponse proto.InternalMessageInfo

func (m *CreateCinemaShowingResponse) GetShowing() *CinemaShowing {
	if m != nil {
		return m.Showing
	}
	return nil
}

type DeleteCinemaShowingRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCinemaShowingRequest) Reset()         { *m = DeleteCinemaShowingRequest{} }
func (m *DeleteCinemaShowingRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCinemaShowingRequest) ProtoMessage()    {}
func (*DeleteCinemaShowingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50fce743d1c301cd, []int{2}
}

func (m *DeleteCinemaShowingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCinemaShowingRequest.Unmarshal(m, b)
}
func (m *DeleteCinemaShowingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCinemaShowingRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCinemaShowingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCinemaShowingRequest.Merge(m, src)
}
func (m *DeleteCinemaShowingRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCinemaShowingRequest.Size(m)
}
func (m *DeleteCinemaShowingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCinemaShowingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCinemaShowingRequest proto.InternalMessageInfo

func (m *DeleteCinemaShowingRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteCinemaShowingResponse struct {
	Showing              *CinemaShowing `protobuf:"bytes,1,opt,name=showing,proto3" json:"showing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteCinemaShowingResponse) Reset()         { *m = DeleteCinemaShowingResponse{} }
func (m *DeleteCinemaShowingResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCinemaShowingResponse) ProtoMessage()    {}
func (*DeleteCinemaShowingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50fce743d1c301cd, []int{3}
}

func (m *DeleteCinemaShowingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCinemaShowingResponse.Unmarshal(m, b)
}
func (m *DeleteCinemaShowingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCinemaShowingResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCinemaShowingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCinemaShowingResponse.Merge(m, src)
}
func (m *DeleteCinemaShowingResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCinemaShowingResponse.Size(m)
}
func (m *DeleteCinemaShowingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCinemaShowingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCinemaShowingResponse proto.InternalMessageInfo

func (m *DeleteCinemaShowingResponse) GetShowing() *CinemaShowing {
	if m != nil {
		return m.Showing
	}
	return nil
}

type FindAllCinemaShowingsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllCinemaShowingsRequest) Reset()         { *m = FindAllCinemaShowingsRequest{} }
func (m *FindAllCinemaShowingsRequest) String() string { return proto.CompactTextString(m) }
func (*FindAllCinemaShowingsRequest) ProtoMessage()    {}
func (*FindAllCinemaShowingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50fce743d1c301cd, []int{4}
}

func (m *FindAllCinemaShowingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllCinemaShowingsRequest.Unmarshal(m, b)
}
func (m *FindAllCinemaShowingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllCinemaShowingsRequest.Marshal(b, m, deterministic)
}
func (m *FindAllCinemaShowingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllCinemaShowingsRequest.Merge(m, src)
}
func (m *FindAllCinemaShowingsRequest) XXX_Size() int {
	return xxx_messageInfo_FindAllCinemaShowingsRequest.Size(m)
}
func (m *FindAllCinemaShowingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllCinemaShowingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllCinemaShowingsRequest proto.InternalMessageInfo

type FindAllCinemaShowingsResponse struct {
	Showings             []*CinemaShowing `protobuf:"bytes,1,rep,name=showings,proto3" json:"showings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FindAllCinemaShowingsResponse) Reset()         { *m = FindAllCinemaShowingsResponse{} }
func (m *FindAllCinemaShowingsResponse) String() string { return proto.CompactTextString(m) }
func (*FindAllCinemaShowingsResponse) ProtoMessage()    {}
func (*FindAllCinemaShowingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50fce743d1c301cd, []int{5}
}

func (m *FindAllCinemaShowingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllCinemaShowingsResponse.Unmarshal(m, b)
}
func (m *FindAllCinemaShowingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllCinemaShowingsResponse.Marshal(b, m, deterministic)
}
func (m *FindAllCinemaShowingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllCinemaShowingsResponse.Merge(m, src)
}
func (m *FindAllCinemaShowingsResponse) XXX_Size() int {
	return xxx_messageInfo_FindAllCinemaShowingsResponse.Size(m)
}
func (m *FindAllCinemaShowingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllCinemaShowingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllCinemaShowingsResponse proto.InternalMessageInfo

func (m *FindAllCinemaShowingsResponse) GetShowings() []*CinemaShowing {
	if m != nil {
		return m.Showings
	}
	return nil
}

type CinemaShowing struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Movie                int64    `protobuf:"varint,2,opt,name=movie,proto3" json:"movie,omitempty"`
	CinemaHall           int64    `protobuf:"varint,3,opt,name=cinemaHall,proto3" json:"cinemaHall,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CinemaShowing) Reset()         { *m = CinemaShowing{} }
func (m *CinemaShowing) String() string { return proto.CompactTextString(m) }
func (*CinemaShowing) ProtoMessage()    {}
func (*CinemaShowing) Descriptor() ([]byte, []int) {
	return fileDescriptor_50fce743d1c301cd, []int{6}
}

func (m *CinemaShowing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CinemaShowing.Unmarshal(m, b)
}
func (m *CinemaShowing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CinemaShowing.Marshal(b, m, deterministic)
}
func (m *CinemaShowing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CinemaShowing.Merge(m, src)
}
func (m *CinemaShowing) XXX_Size() int {
	return xxx_messageInfo_CinemaShowing.Size(m)
}
func (m *CinemaShowing) XXX_DiscardUnknown() {
	xxx_messageInfo_CinemaShowing.DiscardUnknown(m)
}

var xxx_messageInfo_CinemaShowing proto.InternalMessageInfo

func (m *CinemaShowing) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CinemaShowing) GetMovie() int64 {
	if m != nil {
		return m.Movie
	}
	return 0
}

func (m *CinemaShowing) GetCinemaHall() int64 {
	if m != nil {
		return m.CinemaHall
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateCinemaShowingRequest)(nil), "CreateCinemaShowingRequest")
	proto.RegisterType((*CreateCinemaShowingResponse)(nil), "CreateCinemaShowingResponse")
	proto.RegisterType((*DeleteCinemaShowingRequest)(nil), "DeleteCinemaShowingRequest")
	proto.RegisterType((*DeleteCinemaShowingResponse)(nil), "DeleteCinemaShowingResponse")
	proto.RegisterType((*FindAllCinemaShowingsRequest)(nil), "FindAllCinemaShowingsRequest")
	proto.RegisterType((*FindAllCinemaShowingsResponse)(nil), "FindAllCinemaShowingsResponse")
	proto.RegisterType((*CinemaShowing)(nil), "CinemaShowing")
}

func init() { proto.RegisterFile("cinema_showing.proto", fileDescriptor_50fce743d1c301cd) }

var fileDescriptor_50fce743d1c301cd = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x69, 0x8b, 0x9d, 0x5c, 0x71, 0x0f, 0xa1, 0x0f, 0x25, 0xed, 0xca, 0xc8, 0xd3, 0x10,
	0xc9, 0xc3, 0xfc, 0x04, 0x52, 0xd1, 0x81, 0x6f, 0x1d, 0x3e, 0x4b, 0x5d, 0x2f, 0x1a, 0xc8, 0x9a,
	0xd9, 0xd4, 0xf9, 0x9d, 0xfd, 0x14, 0x62, 0x13, 0xff, 0xb4, 0x34, 0x05, 0x1f, 0x93, 0x73, 0x39,
	0xe7, 0xe6, 0xfc, 0x02, 0xd1, 0x4e, 0xd4, 0xb8, 0x2f, 0x1f, 0xf5, 0x8b, 0x7a, 0x17, 0xf5, 0x33,
	0x3f, 0x34, 0xaa, 0x55, 0xac, 0x00, 0x9a, 0x37, 0x58, 0xb6, 0x98, 0x77, 0xea, 0xd6, 0x88, 0x05,
	0xbe, 0xbe, 0xa1, 0x6e, 0x49, 0x04, 0x27, 0x7b, 0x75, 0x14, 0x18, 0x7b, 0x4b, 0x6f, 0x15, 0x14,
	0xe6, 0x40, 0x32, 0x00, 0xe3, 0xb5, 0x29, 0xa5, 0x8c, 0xfd, 0x4e, 0xfa, 0x73, 0xc3, 0xee, 0x20,
	0x19, 0xf5, 0xd4, 0x07, 0x55, 0x6b, 0x24, 0x2b, 0x98, 0xd9, 0x1d, 0x3a, 0xdb, 0xb3, 0xf5, 0x9c,
	0xf7, 0x07, 0xbf, 0x65, 0x76, 0x09, 0xf4, 0x06, 0x25, 0x3a, 0x96, 0x9b, 0x83, 0x2f, 0x2a, 0xbb,
	0x99, 0x2f, 0xaa, 0xaf, 0xd8, 0xd1, 0xe9, 0x7f, 0xc7, 0x66, 0x90, 0xde, 0x8a, 0xba, 0xba, 0x96,
	0xb2, 0x37, 0xa0, 0x6d, 0x30, 0xbb, 0x87, 0x85, 0x43, 0xb7, 0x51, 0x17, 0x70, 0x6a, 0xbd, 0x74,
	0xec, 0x2d, 0x83, 0x91, 0xac, 0x1f, 0x9d, 0x3d, 0xc0, 0x79, 0x4f, 0x1a, 0x3e, 0xeb, 0x97, 0x81,
	0xef, 0x66, 0x10, 0x0c, 0x19, 0xac, 0x3f, 0x3c, 0x88, 0x7a, 0xbe, 0x5b, 0x6c, 0x8e, 0x62, 0x87,
	0x24, 0x87, 0xd0, 0xc0, 0x21, 0x09, 0x77, 0x93, 0xa7, 0x29, 0x9f, 0x42, 0x98, 0x43, 0x68, 0xaa,
	0x26, 0x09, 0x77, 0x13, 0xa2, 0x29, 0x9f, 0x02, 0xb2, 0x81, 0x99, 0xad, 0x91, 0x2c, 0xf8, 0x54,
	0xe1, 0x34, 0xe3, 0x93, 0x7d, 0x3f, 0x85, 0xdd, 0x5f, 0xbe, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0xab, 0x5f, 0x77, 0xa7, 0xe3, 0x02, 0x00, 0x00,
}
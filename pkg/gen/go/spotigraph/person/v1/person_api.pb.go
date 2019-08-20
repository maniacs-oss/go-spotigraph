// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spotigraph/person/v1/person_api.proto

package personv1

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Principal            string   `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{0}
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

func (m *CreateRequest) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}

func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}

func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}

func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}

func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateRequest struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            *types.StringValue `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             *types.StringValue `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{2}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}

func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}

func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}

func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}

func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetFirstName() *types.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *UpdateRequest) GetLastName() *types.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{3}
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

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchRequest struct {
	Page                 uint32             `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32             `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Sorts                []string           `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Cursor               *types.StringValue `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	PersonId             *types.StringValue `protobuf:"bytes,5,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	Principal            *types.StringValue `protobuf:"bytes,6,opt,name=principal,proto3" json:"principal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{4}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}

func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}

func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}

func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}

func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchRequest) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *SearchRequest) GetSorts() []string {
	if m != nil {
		return m.Sorts
	}
	return nil
}

func (m *SearchRequest) GetCursor() *types.StringValue {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *SearchRequest) GetPersonId() *types.StringValue {
	if m != nil {
		return m.PersonId
	}
	return nil
}

func (m *SearchRequest) GetPrincipal() *types.StringValue {
	if m != nil {
		return m.Principal
	}
	return nil
}

type CreateResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Person   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{5}
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

func (m *CreateResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CreateResponse) GetEntity() *Person {
	if m != nil {
		return m.Entity
	}
	return nil
}

type GetResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Person   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}

func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}

func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}

func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}

func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetResponse) GetEntity() *Person {
	if m != nil {
		return m.Entity
	}
	return nil
}

type UpdateResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Person   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{7}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}

func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}

func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}

func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}

func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *UpdateResponse) GetEntity() *Person {
	if m != nil {
		return m.Entity
	}
	return nil
}

type DeleteResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{8}
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

func (m *DeleteResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type SearchResponse struct {
	Error                *v1.Error          `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Total                uint32             `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	PerPage              uint32             `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Count                uint32             `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	CurrentPage          uint32             `protobuf:"varint,5,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	NextCursor           *types.StringValue `protobuf:"bytes,6,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	Members              []*Person          `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd1ef20e07f3284, []int{9}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}

func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}

func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}

func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}

func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SearchResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *SearchResponse) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *SearchResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SearchResponse) GetCurrentPage() uint32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *SearchResponse) GetNextCursor() *types.StringValue {
	if m != nil {
		return m.NextCursor
	}
	return nil
}

func (m *SearchResponse) GetMembers() []*Person {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "spotigraph.person.v1.CreateRequest")
	proto.RegisterType((*GetRequest)(nil), "spotigraph.person.v1.GetRequest")
	proto.RegisterType((*UpdateRequest)(nil), "spotigraph.person.v1.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "spotigraph.person.v1.DeleteRequest")
	proto.RegisterType((*SearchRequest)(nil), "spotigraph.person.v1.SearchRequest")
	proto.RegisterType((*CreateResponse)(nil), "spotigraph.person.v1.CreateResponse")
	proto.RegisterType((*GetResponse)(nil), "spotigraph.person.v1.GetResponse")
	proto.RegisterType((*UpdateResponse)(nil), "spotigraph.person.v1.UpdateResponse")
	proto.RegisterType((*DeleteResponse)(nil), "spotigraph.person.v1.DeleteResponse")
	proto.RegisterType((*SearchResponse)(nil), "spotigraph.person.v1.SearchResponse")
}

func init() {
	proto.RegisterFile("spotigraph/person/v1/person_api.proto", fileDescriptor_3cd1ef20e07f3284)
}

var fileDescriptor_3cd1ef20e07f3284 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x6f, 0x12, 0x41,
	0x14, 0xc7, 0xc3, 0x6e, 0x81, 0xf2, 0x10, 0x0e, 0x2b, 0x26, 0x88, 0xc6, 0x00, 0xd6, 0x84, 0xc4,
	0xb0, 0x08, 0x9a, 0x26, 0x3d, 0x78, 0x28, 0xb5, 0x36, 0xf5, 0x60, 0xc8, 0x6e, 0x6c, 0x4c, 0xad,
	0x92, 0x01, 0x5e, 0x97, 0x4d, 0x60, 0x67, 0x9c, 0x1d, 0xb0, 0xf4, 0xe4, 0x9f, 0xe1, 0xd9, 0x98,
	0x98, 0xe8, 0x5f, 0xa2, 0x7f, 0x52, 0x4f, 0x66, 0x77, 0x66, 0xa5, 0xdb, 0x42, 0x7f, 0x1d, 0x7a,
	0x9b, 0x7d, 0xf3, 0xfd, 0xbe, 0x99, 0x79, 0xfb, 0x79, 0x0f, 0x9e, 0xf8, 0x8c, 0x0a, 0xd7, 0xe1,
	0x84, 0x0d, 0x1b, 0x0c, 0xb9, 0x4f, 0xbd, 0xc6, 0xb4, 0xa9, 0x56, 0x5d, 0xc2, 0x5c, 0x93, 0x71,
	0x2a, 0xa8, 0x51, 0x98, 0xcb, 0x4c, 0xb9, 0x69, 0x4e, 0x9b, 0xa5, 0x4d, 0xc7, 0x15, 0xc3, 0x49,
	0xcf, 0xec, 0xd3, 0x71, 0x03, 0xbd, 0x29, 0x9d, 0x31, 0x4e, 0x8f, 0x66, 0x8d, 0xd0, 0xd2, 0xaf,
	0x3b, 0xe8, 0xd5, 0xa7, 0x64, 0xe4, 0x0e, 0x88, 0xc0, 0xc6, 0xb9, 0x85, 0x4c, 0x5c, 0x7a, 0xe4,
	0x50, 0xea, 0x8c, 0x50, 0x7a, 0x7a, 0x93, 0xc3, 0xc6, 0x17, 0x4e, 0x58, 0x70, 0x82, 0xda, 0xaf,
	0x5c, 0x70, 0xbf, 0x05, 0x12, 0x7f, 0xe6, 0x0b, 0x1c, 0x07, 0x12, 0xb9, 0x92, 0x92, 0xea, 0x06,
	0xe4, 0xb6, 0x38, 0x12, 0x81, 0x16, 0x7e, 0x9e, 0xa0, 0x2f, 0x8c, 0x1a, 0x64, 0x18, 0x77, 0xbd,
	0xbe, 0xcb, 0xc8, 0xa8, 0x98, 0x28, 0x27, 0x6a, 0x99, 0x36, 0x9c, 0xb4, 0xd3, 0x3c, 0x59, 0xd6,
	0x6b, 0x5f, 0x35, 0x6b, 0xbe, 0x59, 0x5d, 0x07, 0xd8, 0x41, 0x31, 0xf7, 0x69, 0xee, 0x40, 0x19,
	0x8a, 0x27, 0xed, 0x7b, 0xfc, 0x6e, 0x2b, 0xff, 0xe9, 0xc3, 0xb3, 0xfa, 0xc6, 0x66, 0x7d, 0x9f,
	0xd4, 0x8f, 0x3f, 0x3e, 0x5d, 0xfb, 0x96, 0x28, 0x5b, 0x9a, 0x3b, 0xa8, 0xfe, 0x49, 0x40, 0xee,
	0x1d, 0x1b, 0xc4, 0xce, 0xbc, 0xa2, 0xd7, 0xd8, 0x01, 0x38, 0x74, 0xb9, 0x2f, 0xba, 0x1e, 0x19,
	0x63, 0x51, 0x2b, 0x27, 0x6a, 0xd9, 0xd6, 0x43, 0x53, 0x56, 0xca, 0x8c, 0x2a, 0x65, 0xda, 0x82,
	0xbb, 0x9e, 0xb3, 0x47, 0x46, 0x13, 0x8c, 0x5f, 0x3e, 0xf4, 0xbe, 0x25, 0x63, 0x34, 0xb6, 0x21,
	0x33, 0x22, 0x51, 0x1e, 0xfd, 0x9a, 0x79, 0x56, 0x03, 0x6b, 0x90, 0x26, 0x28, 0xdf, 0x2b, 0x1c,
	0xe1, 0x0d, 0x9e, 0x52, 0xfd, 0xa9, 0x41, 0xce, 0x46, 0xc2, 0xfb, 0xc3, 0xc8, 0x6b, 0xc0, 0x0a,
	0x23, 0x0e, 0x86, 0xee, 0x9c, 0x15, 0xae, 0x8d, 0xfb, 0xb0, 0xca, 0x90, 0x77, 0xc3, 0xb8, 0x16,
	0xc6, 0xd3, 0x0c, 0x79, 0x27, 0xd8, 0x2a, 0x40, 0xd2, 0xa7, 0x5c, 0xf8, 0x45, 0xbd, 0xac, 0xd7,
	0x32, 0x96, 0xfc, 0x30, 0x5e, 0x40, 0xaa, 0x3f, 0xe1, 0x3e, 0xe5, 0xc5, 0x95, 0xcb, 0x5f, 0x65,
	0x29, 0xad, 0x61, 0x43, 0x46, 0x91, 0xed, 0x0e, 0x8a, 0xc9, 0x2b, 0x94, 0x63, 0xf9, 0xdb, 0x56,
	0x65, 0xa2, 0xdd, 0x81, 0xf1, 0xfa, 0x34, 0x4a, 0xa9, 0xeb, 0xfe, 0xab, 0x39, 0x68, 0x33, 0xc8,
	0x47, 0x8c, 0xfa, 0x8c, 0x7a, 0x3e, 0x1a, 0x4d, 0x48, 0x22, 0xe7, 0x94, 0x87, 0xa5, 0xca, 0xb6,
	0x1e, 0x98, 0xa7, 0x9a, 0x50, 0xe1, 0x3d, 0x6d, 0x9a, 0xdb, 0x81, 0xc4, 0x92, 0xca, 0xa0, 0x2e,
	0xe8, 0x09, 0x57, 0xcc, 0xfe, 0x53, 0xb3, 0xa8, 0x71, 0xcd, 0x4e, 0xb8, 0xb2, 0x94, 0xb6, 0x3a,
	0x85, 0x6c, 0xc8, 0xf8, 0x6d, 0x9f, 0x3b, 0x83, 0x7c, 0xd4, 0x22, 0xb7, 0x7d, 0xf4, 0x16, 0xe4,
	0x23, 0xa4, 0x6f, 0x7c, 0x74, 0xf5, 0x87, 0x06, 0xf9, 0x08, 0xee, 0x9b, 0x3f, 0xa0, 0x00, 0x49,
	0x41, 0x05, 0x19, 0x29, 0xf2, 0xe5, 0x47, 0xac, 0x25, 0xf4, 0x73, 0x2d, 0xd1, 0xa7, 0x13, 0x4f,
	0x84, 0xec, 0xe7, 0x2c, 0xf9, 0x61, 0x54, 0xe0, 0x4e, 0x7f, 0xc2, 0x39, 0x7a, 0x42, 0x9a, 0x92,
	0xe1, 0x66, 0x56, 0xc5, 0x42, 0xe3, 0x4b, 0xc8, 0x7a, 0x78, 0x24, 0xba, 0xaa, 0x75, 0xae, 0x00,
	0xab, 0x05, 0x81, 0x61, 0x4b, 0xb6, 0xcf, 0x3a, 0xa4, 0xc7, 0x38, 0xee, 0x21, 0xf7, 0x8b, 0xe9,
	0xb2, 0x7e, 0x69, 0xa9, 0x23, 0x71, 0xeb, 0xb7, 0x0e, 0x19, 0x19, 0xdb, 0xec, 0xec, 0x1a, 0x36,
	0xa4, 0x24, 0xe7, 0xc6, 0xe3, 0xc5, 0xf6, 0xd8, 0xa4, 0x2e, 0xad, 0x5d, 0x2c, 0x52, 0x65, 0x7f,
	0x03, 0xfa, 0x0e, 0x0a, 0xa3, 0xbc, 0x58, 0x3c, 0x1f, 0xe0, 0xa5, 0xca, 0x05, 0x0a, 0x95, 0xcb,
	0x86, 0x94, 0xa4, 0x72, 0xd9, 0x05, 0x63, 0x63, 0x7d, 0xd9, 0x05, 0xcf, 0x80, 0x6d, 0x43, 0x4a,
	0xf2, 0xb6, 0x2c, 0x69, 0x6c, 0xc0, 0x2e, 0x4b, 0x7a, 0x06, 0x59, 0x1b, 0x52, 0x12, 0xbf, 0x65,
	0x49, 0x63, 0x93, 0x77, 0x59, 0xd2, 0x38, 0xc1, 0x6d, 0x1f, 0x2a, 0x94, 0x3b, 0xe6, 0x31, 0x7a,
	0xae, 0x18, 0x12, 0xbe, 0xd0, 0xd3, 0xce, 0xab, 0xff, 0xc9, 0xdc, 0x4e, 0x40, 0x4d, 0x27, 0xb1,
	0xaf, 0x06, 0xe2, 0xb4, 0xf9, 0x5d, 0xd3, 0xed, 0xce, 0xfb, 0x5f, 0x5a, 0xc1, 0x9e, 0x1b, 0xa5,
	0xda, 0xdc, 0x6b, 0xfe, 0x3d, 0x1d, 0x3e, 0x90, 0xe1, 0x83, 0xbd, 0x66, 0x2f, 0x15, 0xc2, 0xf7,
	0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xe5, 0xf4, 0xfa, 0x8f, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PersonAPIClient is the client API for PersonAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonAPIClient interface {
	// Create a person.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get a person by id.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Update person attributes.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a person by id.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Search for persons.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type personAPIClient struct {
	cc *grpc.ClientConn
}

func NewPersonAPIClient(cc *grpc.ClientConn) PersonAPIClient {
	return &personAPIClient{cc}
}

func (c *personAPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.person.v1.PersonAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.person.v1.PersonAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.person.v1.PersonAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.person.v1.PersonAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personAPIClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.person.v1.PersonAPI/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonAPIServer is the server API for PersonAPI service.
type PersonAPIServer interface {
	// Create a person.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get a person by id.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Update person attributes.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a person by id.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Search for persons.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterPersonAPIServer(s *grpc.Server, srv PersonAPIServer) {
	s.RegisterService(&_PersonAPI_serviceDesc, srv)
}

func _PersonAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.person.v1.PersonAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonAPIServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.person.v1.PersonAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.person.v1.PersonAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.person.v1.PersonAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonAPI_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonAPIServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.person.v1.PersonAPI/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonAPIServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersonAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spotigraph.person.v1.PersonAPI",
	HandlerType: (*PersonAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PersonAPI_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PersonAPI_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PersonAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PersonAPI_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _PersonAPI_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spotigraph/person/v1/person_api.proto",
}

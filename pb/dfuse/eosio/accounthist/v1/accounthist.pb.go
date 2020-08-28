// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/eosio/accounthist/v1/accounthist.proto

package pbaccounthist

import (
	context "context"
	fmt "fmt"
	v1 "github.com/dfuse-io/dfuse-eosio/pb/dfuse/eosio/codec/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetActionsRequest struct {
	Account              uint64   `protobuf:"varint,1,opt,name=account,proto3" json:"account,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Cursor               *Cursor  `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetActionsRequest) Reset()         { *m = GetActionsRequest{} }
func (m *GetActionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetActionsRequest) ProtoMessage()    {}
func (*GetActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{0}
}

func (m *GetActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActionsRequest.Unmarshal(m, b)
}
func (m *GetActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActionsRequest.Marshal(b, m, deterministic)
}
func (m *GetActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActionsRequest.Merge(m, src)
}
func (m *GetActionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetActionsRequest.Size(m)
}
func (m *GetActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetActionsRequest proto.InternalMessageInfo

func (m *GetActionsRequest) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *GetActionsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetActionsRequest) GetCursor() *Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

type ActionResponse struct {
	Cursor               *Cursor         `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	ActionTrace          *v1.ActionTrace `protobuf:"bytes,2,opt,name=action_trace,json=actionTrace,proto3" json:"action_trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ActionResponse) Reset()         { *m = ActionResponse{} }
func (m *ActionResponse) String() string { return proto.CompactTextString(m) }
func (*ActionResponse) ProtoMessage()    {}
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{1}
}

func (m *ActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionResponse.Unmarshal(m, b)
}
func (m *ActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionResponse.Marshal(b, m, deterministic)
}
func (m *ActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionResponse.Merge(m, src)
}
func (m *ActionResponse) XXX_Size() int {
	return xxx_messageInfo_ActionResponse.Size(m)
}
func (m *ActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActionResponse proto.InternalMessageInfo

func (m *ActionResponse) GetCursor() *Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *ActionResponse) GetActionTrace() *v1.ActionTrace {
	if m != nil {
		return m.ActionTrace
	}
	return nil
}

type ActionRow struct {
	Version              uint32          `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ActionTrace          *v1.ActionTrace `protobuf:"bytes,2,opt,name=action_trace,json=actionTrace,proto3" json:"action_trace,omitempty"`
	LastDeletedSeq       uint64          `protobuf:"varint,3,opt,name=last_deleted_seq,json=lastDeletedSeq,proto3" json:"last_deleted_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ActionRow) Reset()         { *m = ActionRow{} }
func (m *ActionRow) String() string { return proto.CompactTextString(m) }
func (*ActionRow) ProtoMessage()    {}
func (*ActionRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{2}
}

func (m *ActionRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRow.Unmarshal(m, b)
}
func (m *ActionRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRow.Marshal(b, m, deterministic)
}
func (m *ActionRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRow.Merge(m, src)
}
func (m *ActionRow) XXX_Size() int {
	return xxx_messageInfo_ActionRow.Size(m)
}
func (m *ActionRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRow.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRow proto.InternalMessageInfo

func (m *ActionRow) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ActionRow) GetActionTrace() *v1.ActionTrace {
	if m != nil {
		return m.ActionTrace
	}
	return nil
}

func (m *ActionRow) GetLastDeletedSeq() uint64 {
	if m != nil {
		return m.LastDeletedSeq
	}
	return 0
}

type ShardCheckpoint struct {
	InitialStartBlock    uint64   `protobuf:"varint,1,opt,name=initial_start_block,json=initialStartBlock,proto3" json:"initial_start_block,omitempty"`
	TargetStopBlock      uint64   `protobuf:"varint,2,opt,name=target_stop_block,json=targetStopBlock,proto3" json:"target_stop_block,omitempty"`
	LastWrittenBlockNum  uint64   `protobuf:"varint,3,opt,name=last_written_block_num,json=lastWrittenBlockNum,proto3" json:"last_written_block_num,omitempty"`
	LastWrittenBlockId   string   `protobuf:"bytes,4,opt,name=last_written_block_id,json=lastWrittenBlockId,proto3" json:"last_written_block_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardCheckpoint) Reset()         { *m = ShardCheckpoint{} }
func (m *ShardCheckpoint) String() string { return proto.CompactTextString(m) }
func (*ShardCheckpoint) ProtoMessage()    {}
func (*ShardCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{3}
}

func (m *ShardCheckpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardCheckpoint.Unmarshal(m, b)
}
func (m *ShardCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardCheckpoint.Marshal(b, m, deterministic)
}
func (m *ShardCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardCheckpoint.Merge(m, src)
}
func (m *ShardCheckpoint) XXX_Size() int {
	return xxx_messageInfo_ShardCheckpoint.Size(m)
}
func (m *ShardCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ShardCheckpoint proto.InternalMessageInfo

func (m *ShardCheckpoint) GetInitialStartBlock() uint64 {
	if m != nil {
		return m.InitialStartBlock
	}
	return 0
}

func (m *ShardCheckpoint) GetTargetStopBlock() uint64 {
	if m != nil {
		return m.TargetStopBlock
	}
	return 0
}

func (m *ShardCheckpoint) GetLastWrittenBlockNum() uint64 {
	if m != nil {
		return m.LastWrittenBlockNum
	}
	return 0
}

func (m *ShardCheckpoint) GetLastWrittenBlockId() string {
	if m != nil {
		return m.LastWrittenBlockId
	}
	return ""
}

type ActionRowAppend struct {
	LastDeletedSeq       uint64   `protobuf:"varint,3,opt,name=last_deleted_seq,json=lastDeletedSeq,proto3" json:"last_deleted_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionRowAppend) Reset()         { *m = ActionRowAppend{} }
func (m *ActionRowAppend) String() string { return proto.CompactTextString(m) }
func (*ActionRowAppend) ProtoMessage()    {}
func (*ActionRowAppend) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{4}
}

func (m *ActionRowAppend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRowAppend.Unmarshal(m, b)
}
func (m *ActionRowAppend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRowAppend.Marshal(b, m, deterministic)
}
func (m *ActionRowAppend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRowAppend.Merge(m, src)
}
func (m *ActionRowAppend) XXX_Size() int {
	return xxx_messageInfo_ActionRowAppend.Size(m)
}
func (m *ActionRowAppend) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRowAppend.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRowAppend proto.InternalMessageInfo

func (m *ActionRowAppend) GetLastDeletedSeq() uint64 {
	if m != nil {
		return m.LastDeletedSeq
	}
	return 0
}

type Cursor struct {
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Magic                uint32   `protobuf:"varint,2,opt,name=magic,proto3" json:"magic,omitempty"`
	Account              uint64   `protobuf:"varint,3,opt,name=account,proto3" json:"account,omitempty"`
	ShardNum             uint32   `protobuf:"varint,4,opt,name=shard_num,json=shardNum,proto3" json:"shard_num,omitempty"`
	SequenceNumber       uint64   `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cursor) Reset()         { *m = Cursor{} }
func (m *Cursor) String() string { return proto.CompactTextString(m) }
func (*Cursor) ProtoMessage()    {}
func (*Cursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{5}
}

func (m *Cursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cursor.Unmarshal(m, b)
}
func (m *Cursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cursor.Marshal(b, m, deterministic)
}
func (m *Cursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cursor.Merge(m, src)
}
func (m *Cursor) XXX_Size() int {
	return xxx_messageInfo_Cursor.Size(m)
}
func (m *Cursor) XXX_DiscardUnknown() {
	xxx_messageInfo_Cursor.DiscardUnknown(m)
}

var xxx_messageInfo_Cursor proto.InternalMessageInfo

func (m *Cursor) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Cursor) GetMagic() uint32 {
	if m != nil {
		return m.Magic
	}
	return 0
}

func (m *Cursor) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *Cursor) GetShardNum() uint32 {
	if m != nil {
		return m.ShardNum
	}
	return 0
}

func (m *Cursor) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*GetActionsRequest)(nil), "dfuse.eosio.accounthist.v1.GetActionsRequest")
	proto.RegisterType((*ActionResponse)(nil), "dfuse.eosio.accounthist.v1.ActionResponse")
	proto.RegisterType((*ActionRow)(nil), "dfuse.eosio.accounthist.v1.ActionRow")
	proto.RegisterType((*ShardCheckpoint)(nil), "dfuse.eosio.accounthist.v1.ShardCheckpoint")
	proto.RegisterType((*ActionRowAppend)(nil), "dfuse.eosio.accounthist.v1.ActionRowAppend")
	proto.RegisterType((*Cursor)(nil), "dfuse.eosio.accounthist.v1.Cursor")
}

func init() {
	proto.RegisterFile("dfuse/eosio/accounthist/v1/accounthist.proto", fileDescriptor_4c22ddb60199ece6)
}

var fileDescriptor_4c22ddb60199ece6 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xd5, 0xb6, 0x69, 0x20, 0x1b, 0xda, 0x90, 0x6d, 0x41, 0x56, 0xb8, 0x04, 0x5f, 0x88, 0x2a,
	0xea, 0x90, 0xf4, 0x46, 0x4f, 0xfd, 0x90, 0x00, 0x21, 0x7a, 0x70, 0x90, 0x90, 0xb8, 0x58, 0xf6,
	0x7a, 0x48, 0x56, 0x8d, 0xbd, 0xce, 0xee, 0x38, 0x15, 0xe2, 0xc0, 0x7f, 0x40, 0x42, 0xe2, 0x8f,
	0xf1, 0x7f, 0xd0, 0xee, 0x3a, 0xd4, 0x01, 0x1a, 0xa9, 0x12, 0x37, 0xcf, 0xbc, 0x37, 0x6f, 0xdf,
	0x8e, 0x67, 0x96, 0x3e, 0x4f, 0x3f, 0x95, 0x1a, 0x86, 0x20, 0xb5, 0x90, 0xc3, 0x98, 0x73, 0x59,
	0xe6, 0x38, 0x13, 0x1a, 0x87, 0xcb, 0x51, 0x3d, 0x0c, 0x0a, 0x25, 0x51, 0xb2, 0x9e, 0x65, 0x07,
	0x96, 0x1d, 0xd4, 0xe1, 0xe5, 0xa8, 0xd7, 0xaf, 0x2b, 0x71, 0x99, 0x02, 0x37, 0x1a, 0xf6, 0xc3,
	0x55, 0xfb, 0x5f, 0x69, 0xf7, 0x15, 0xe0, 0x29, 0x47, 0x21, 0x73, 0x1d, 0xc2, 0xa2, 0x04, 0x8d,
	0xcc, 0xa3, 0xf7, 0x2a, 0x21, 0x8f, 0xf4, 0xc9, 0xa0, 0x11, 0xae, 0x42, 0x76, 0x40, 0x77, 0xe6,
	0x22, 0x13, 0xe8, 0x6d, 0xf5, 0xc9, 0x60, 0x37, 0x74, 0x01, 0x7b, 0x49, 0x9b, 0xbc, 0x54, 0x5a,
	0x2a, 0x6f, 0xbb, 0x4f, 0x06, 0xed, 0xb1, 0x1f, 0xdc, 0xee, 0x29, 0x38, 0xb7, 0xcc, 0xb0, 0xaa,
	0xf0, 0xbf, 0x11, 0xba, 0xe7, 0x8e, 0x0f, 0x41, 0x17, 0x32, 0xd7, 0x50, 0x93, 0x23, 0x77, 0x95,
	0x63, 0x17, 0xf4, 0x41, 0x6c, 0xd5, 0x22, 0x54, 0x31, 0x07, 0xeb, 0xb3, 0x3d, 0x7e, 0xba, 0xa6,
	0xe0, 0xee, 0xbf, 0x1c, 0x05, 0xee, 0xdc, 0xf7, 0x86, 0x18, 0xb6, 0xe3, 0x9b, 0xc0, 0xff, 0x4e,
	0x68, 0xab, 0x32, 0x25, 0xaf, 0x4d, 0x3b, 0x96, 0xa0, 0xb4, 0x90, 0xb9, 0x35, 0xb4, 0x1b, 0xae,
	0xc2, 0xff, 0x73, 0x1a, 0x1b, 0xd0, 0x87, 0xf3, 0x58, 0x63, 0x94, 0xc2, 0x1c, 0x10, 0xd2, 0x48,
	0xc3, 0xc2, 0x36, 0xb2, 0x11, 0xee, 0x99, 0xfc, 0x85, 0x4b, 0x4f, 0x60, 0xe1, 0xff, 0x24, 0xb4,
	0x33, 0x99, 0xc5, 0x2a, 0x3d, 0x9f, 0x01, 0xbf, 0x2a, 0xa4, 0xc8, 0x91, 0x05, 0x74, 0x5f, 0xe4,
	0x02, 0x45, 0x3c, 0x8f, 0x34, 0xc6, 0x0a, 0xa3, 0x64, 0x2e, 0xf9, 0x55, 0xf5, 0xe3, 0xba, 0x15,
	0x34, 0x31, 0xc8, 0x99, 0x01, 0xd8, 0x21, 0xed, 0x62, 0xac, 0xa6, 0x80, 0x91, 0x46, 0x59, 0x54,
	0xec, 0x2d, 0xcb, 0xee, 0x38, 0x60, 0x82, 0xb2, 0x70, 0xdc, 0x63, 0xfa, 0xd8, 0x3a, 0xbb, 0x56,
	0x02, 0x11, 0x72, 0x47, 0x8e, 0xf2, 0x32, 0xab, 0xfc, 0xed, 0x1b, 0xf4, 0x83, 0x03, 0x6d, 0xc5,
	0x65, 0x99, 0xb1, 0x11, 0x7d, 0xf4, 0x8f, 0x22, 0x91, 0x7a, 0x8d, 0x3e, 0x19, 0xb4, 0x42, 0xf6,
	0x67, 0xcd, 0x9b, 0xd4, 0x3f, 0xa1, 0x9d, 0xdf, 0xed, 0x3e, 0x2d, 0x0a, 0xc8, 0xd3, 0x3b, 0x34,
	0xe5, 0x07, 0xa1, 0x4d, 0x37, 0x05, 0x1b, 0xfe, 0xd4, 0x01, 0xdd, 0xc9, 0xe2, 0xa9, 0xe0, 0xab,
	0xc1, 0xb5, 0x41, 0x7d, 0xd0, 0xb7, 0xd7, 0x07, 0xfd, 0x09, 0x6d, 0x69, 0xd3, 0x68, 0x7b, 0xd9,
	0x86, 0xad, 0xb9, 0x6f, 0x13, 0xe6, 0x86, 0xcf, 0x68, 0x47, 0x9b, 0x55, 0xc9, 0x39, 0x18, 0x3c,
	0x01, 0xe5, 0xed, 0x38, 0x6b, 0xab, 0xf4, 0xa5, 0xcd, 0x8e, 0xbf, 0x98, 0xd9, 0xb6, 0x82, 0xaf,
	0x85, 0x46, 0xa9, 0x3e, 0x33, 0x41, 0xe9, 0xcd, 0xbe, 0xb1, 0xa3, 0x4d, 0x93, 0xfd, 0xd7, 0x5e,
	0xf6, 0x0e, 0x37, 0xd1, 0xd7, 0x97, 0xe8, 0x05, 0x39, 0x7b, 0xf7, 0xf1, 0xed, 0x54, 0xe0, 0xac,
	0x4c, 0x02, 0x2e, 0xb3, 0xa1, 0xad, 0x3c, 0x12, 0xb2, 0xfa, 0x70, 0x4f, 0x42, 0x91, 0x0c, 0x6f,
	0x7f, 0x6b, 0x4e, 0x8a, 0xa4, 0x96, 0x48, 0x9a, 0xf6, 0xc1, 0x38, 0xfe, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0xcc, 0x82, 0x15, 0x9e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountHistoryClient is the client API for AccountHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountHistoryClient interface {
	GetActions(ctx context.Context, in *GetActionsRequest, opts ...grpc.CallOption) (AccountHistory_GetActionsClient, error)
}

type accountHistoryClient struct {
	cc *grpc.ClientConn
}

func NewAccountHistoryClient(cc *grpc.ClientConn) AccountHistoryClient {
	return &accountHistoryClient{cc}
}

func (c *accountHistoryClient) GetActions(ctx context.Context, in *GetActionsRequest, opts ...grpc.CallOption) (AccountHistory_GetActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccountHistory_serviceDesc.Streams[0], "/dfuse.eosio.accounthist.v1.AccountHistory/GetActions", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountHistoryGetActionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountHistory_GetActionsClient interface {
	Recv() (*ActionResponse, error)
	grpc.ClientStream
}

type accountHistoryGetActionsClient struct {
	grpc.ClientStream
}

func (x *accountHistoryGetActionsClient) Recv() (*ActionResponse, error) {
	m := new(ActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountHistoryServer is the server API for AccountHistory service.
type AccountHistoryServer interface {
	GetActions(*GetActionsRequest, AccountHistory_GetActionsServer) error
}

// UnimplementedAccountHistoryServer can be embedded to have forward compatible implementations.
type UnimplementedAccountHistoryServer struct {
}

func (*UnimplementedAccountHistoryServer) GetActions(req *GetActionsRequest, srv AccountHistory_GetActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActions not implemented")
}

func RegisterAccountHistoryServer(s *grpc.Server, srv AccountHistoryServer) {
	s.RegisterService(&_AccountHistory_serviceDesc, srv)
}

func _AccountHistory_GetActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetActionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountHistoryServer).GetActions(m, &accountHistoryGetActionsServer{stream})
}

type AccountHistory_GetActionsServer interface {
	Send(*ActionResponse) error
	grpc.ServerStream
}

type accountHistoryGetActionsServer struct {
	grpc.ServerStream
}

func (x *accountHistoryGetActionsServer) Send(m *ActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AccountHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.eosio.accounthist.v1.AccountHistory",
	HandlerType: (*AccountHistoryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActions",
			Handler:       _AccountHistory_GetActions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/eosio/accounthist/v1/accounthist.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/service_status.proto

package protos

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

// ServiceResult enumeration as defined in service "result" by systemd
type ServiceExitStatus_ServiceResult int32

const (
	ServiceExitStatus_UNUSED          ServiceExitStatus_ServiceResult = 0
	ServiceExitStatus_SUCCESS         ServiceExitStatus_ServiceResult = 1
	ServiceExitStatus_PROTOCOL        ServiceExitStatus_ServiceResult = 2
	ServiceExitStatus_TIMEOUT         ServiceExitStatus_ServiceResult = 3
	ServiceExitStatus_EXIT_CODE       ServiceExitStatus_ServiceResult = 4
	ServiceExitStatus_SIGNAL          ServiceExitStatus_ServiceResult = 5
	ServiceExitStatus_CORE_DUMP       ServiceExitStatus_ServiceResult = 6
	ServiceExitStatus_WATCHDOG        ServiceExitStatus_ServiceResult = 7
	ServiceExitStatus_START_LIMIT_HIT ServiceExitStatus_ServiceResult = 8
	ServiceExitStatus_RESOURCES       ServiceExitStatus_ServiceResult = 9
)

var ServiceExitStatus_ServiceResult_name = map[int32]string{
	0: "UNUSED",
	1: "SUCCESS",
	2: "PROTOCOL",
	3: "TIMEOUT",
	4: "EXIT_CODE",
	5: "SIGNAL",
	6: "CORE_DUMP",
	7: "WATCHDOG",
	8: "START_LIMIT_HIT",
	9: "RESOURCES",
}

var ServiceExitStatus_ServiceResult_value = map[string]int32{
	"UNUSED":          0,
	"SUCCESS":         1,
	"PROTOCOL":        2,
	"TIMEOUT":         3,
	"EXIT_CODE":       4,
	"SIGNAL":          5,
	"CORE_DUMP":       6,
	"WATCHDOG":        7,
	"START_LIMIT_HIT": 8,
	"RESOURCES":       9,
}

func (x ServiceExitStatus_ServiceResult) String() string {
	return proto.EnumName(ServiceExitStatus_ServiceResult_name, int32(x))
}

func (ServiceExitStatus_ServiceResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e31fb6a24e18c33, []int{0, 0}
}

// ExitCode enumeration as defined in service "result" by systemd
type ServiceExitStatus_ExitCode int32

const (
	ServiceExitStatus_UNUSED_EXIT_CODE ServiceExitStatus_ExitCode = 0
	ServiceExitStatus_EXITED           ServiceExitStatus_ExitCode = 1
	ServiceExitStatus_KILLED           ServiceExitStatus_ExitCode = 2
	ServiceExitStatus_DUMPED           ServiceExitStatus_ExitCode = 3
)

var ServiceExitStatus_ExitCode_name = map[int32]string{
	0: "UNUSED_EXIT_CODE",
	1: "EXITED",
	2: "KILLED",
	3: "DUMPED",
}

var ServiceExitStatus_ExitCode_value = map[string]int32{
	"UNUSED_EXIT_CODE": 0,
	"EXITED":           1,
	"KILLED":           2,
	"DUMPED":           3,
}

func (x ServiceExitStatus_ExitCode) String() string {
	return proto.EnumName(ServiceExitStatus_ExitCode_name, int32(x))
}

func (ServiceExitStatus_ExitCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e31fb6a24e18c33, []int{0, 1}
}

type ServiceExitStatus struct {
	LatestServiceResult ServiceExitStatus_ServiceResult `protobuf:"varint,1,opt,name=latest_service_result,json=latestServiceResult,proto3,enum=magma.orc8r.ServiceExitStatus_ServiceResult" json:"latest_service_result,omitempty"`
	LatestExitCode      ServiceExitStatus_ExitCode      `protobuf:"varint,2,opt,name=latest_exit_code,json=latestExitCode,proto3,enum=magma.orc8r.ServiceExitStatus_ExitCode" json:"latest_exit_code,omitempty"`
	// Optional return code returned by the service during exit
	LatestRc uint32 `protobuf:"varint,3,opt,name=latest_rc,json=latestRc,proto3" json:"latest_rc,omitempty"`
	// Clean exit, e.g. SIGNKILL
	NumCleanExits uint32 `protobuf:"varint,4,opt,name=num_clean_exits,json=numCleanExits,proto3" json:"num_clean_exits,omitempty"`
	// Unclean exit e.g. CORE_DUMP or non zero exit code.
	NumFailExits         uint32   `protobuf:"varint,5,opt,name=num_fail_exits,json=numFailExits,proto3" json:"num_fail_exits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceExitStatus) Reset()         { *m = ServiceExitStatus{} }
func (m *ServiceExitStatus) String() string { return proto.CompactTextString(m) }
func (*ServiceExitStatus) ProtoMessage()    {}
func (*ServiceExitStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e31fb6a24e18c33, []int{0}
}

func (m *ServiceExitStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceExitStatus.Unmarshal(m, b)
}
func (m *ServiceExitStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceExitStatus.Marshal(b, m, deterministic)
}
func (m *ServiceExitStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceExitStatus.Merge(m, src)
}
func (m *ServiceExitStatus) XXX_Size() int {
	return xxx_messageInfo_ServiceExitStatus.Size(m)
}
func (m *ServiceExitStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceExitStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceExitStatus proto.InternalMessageInfo

func (m *ServiceExitStatus) GetLatestServiceResult() ServiceExitStatus_ServiceResult {
	if m != nil {
		return m.LatestServiceResult
	}
	return ServiceExitStatus_UNUSED
}

func (m *ServiceExitStatus) GetLatestExitCode() ServiceExitStatus_ExitCode {
	if m != nil {
		return m.LatestExitCode
	}
	return ServiceExitStatus_UNUSED_EXIT_CODE
}

func (m *ServiceExitStatus) GetLatestRc() uint32 {
	if m != nil {
		return m.LatestRc
	}
	return 0
}

func (m *ServiceExitStatus) GetNumCleanExits() uint32 {
	if m != nil {
		return m.NumCleanExits
	}
	return 0
}

func (m *ServiceExitStatus) GetNumFailExits() uint32 {
	if m != nil {
		return m.NumFailExits
	}
	return 0
}

func init() {
	proto.RegisterEnum("magma.orc8r.ServiceExitStatus_ServiceResult", ServiceExitStatus_ServiceResult_name, ServiceExitStatus_ServiceResult_value)
	proto.RegisterEnum("magma.orc8r.ServiceExitStatus_ExitCode", ServiceExitStatus_ExitCode_name, ServiceExitStatus_ExitCode_value)
	proto.RegisterType((*ServiceExitStatus)(nil), "magma.orc8r.ServiceExitStatus")
}

func init() { proto.RegisterFile("orc8r/protos/service_status.proto", fileDescriptor_5e31fb6a24e18c33) }

var fileDescriptor_5e31fb6a24e18c33 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xb6, 0xeb, 0xda, 0xb7, 0xb5, 0x33, 0x1e, 0x48, 0x45, 0xbb, 0x8c, 0x0a, 0xc1,
	0x0e, 0xa8, 0x95, 0xe0, 0xc2, 0xb5, 0x38, 0x66, 0x8b, 0x48, 0xe7, 0x61, 0x3b, 0x02, 0x71, 0x31,
	0x59, 0x66, 0xa6, 0x48, 0x49, 0x83, 0x62, 0x07, 0xed, 0x7f, 0xe1, 0xc0, 0xbf, 0x8a, 0xec, 0xa4,
	0x40, 0xc5, 0x61, 0xa7, 0x3c, 0x7f, 0xef, 0x7b, 0xbf, 0x2f, 0x89, 0x1f, 0x3c, 0xab, 0xea, 0xec,
	0x6d, 0xbd, 0xfc, 0x5e, 0x57, 0xb6, 0x32, 0x4b, 0xa3, 0xeb, 0x1f, 0x79, 0xa6, 0x95, 0xb1, 0xa9,
	0x6d, 0xcc, 0xc2, 0xab, 0xf8, 0xb0, 0x4c, 0xef, 0xca, 0x74, 0xe1, 0x8d, 0xf3, 0x9f, 0x03, 0x78,
	0x24, 0x5a, 0x17, 0xbd, 0xcf, 0xad, 0xf0, 0x46, 0xfc, 0x15, 0x9e, 0x14, 0xa9, 0xd5, 0xc6, 0xaa,
	0x2d, 0xa1, 0xd6, 0xa6, 0x29, 0xec, 0x2c, 0x38, 0x0b, 0xce, 0xa7, 0xaf, 0x5f, 0x2d, 0xfe, 0x41,
	0x2c, 0xfe, 0x1b, 0xdf, 0x2a, 0xdc, 0xcf, 0xf0, 0x93, 0x16, 0xb5, 0x23, 0xe2, 0x8f, 0x80, 0xba,
	0x04, 0x7d, 0x9f, 0x5b, 0x95, 0x55, 0xb7, 0x7a, 0xd6, 0xf3, 0xf0, 0x97, 0x0f, 0xc0, 0x5d, 0x49,
	0xaa, 0x5b, 0xcd, 0xa7, 0x2d, 0x60, 0x7b, 0xc6, 0xa7, 0x30, 0xee, 0x90, 0x75, 0x36, 0xeb, 0x9f,
	0x05, 0xe7, 0x13, 0x3e, 0x6a, 0x05, 0x9e, 0xe1, 0x17, 0x70, 0xbc, 0x69, 0x4a, 0x95, 0x15, 0x3a,
	0xdd, 0xf8, 0x48, 0x33, 0x1b, 0x78, 0xcb, 0x64, 0xd3, 0x94, 0xc4, 0xa9, 0x8e, 0x63, 0xf0, 0x73,
	0x98, 0x3a, 0xdf, 0xb7, 0x34, 0x2f, 0x3a, 0xdb, 0xbe, 0xb7, 0x1d, 0x6d, 0x9a, 0xf2, 0x7d, 0x9a,
	0x17, 0xde, 0x35, 0xff, 0x15, 0xc0, 0x64, 0xf7, 0x7b, 0x00, 0x86, 0xc9, 0x55, 0x22, 0x68, 0x88,
	0xf6, 0xf0, 0x21, 0x1c, 0x88, 0x84, 0x10, 0x2a, 0x04, 0x0a, 0xf0, 0x11, 0x8c, 0xae, 0x39, 0x93,
	0x8c, 0xb0, 0x18, 0xf5, 0x5c, 0x4b, 0x46, 0x6b, 0xca, 0x12, 0x89, 0xfa, 0x78, 0x02, 0x63, 0xfa,
	0x39, 0x92, 0x8a, 0xb0, 0x90, 0xa2, 0x81, 0x43, 0x88, 0xe8, 0xe2, 0x6a, 0x15, 0xa3, 0x7d, 0xd7,
	0x22, 0x8c, 0x53, 0x15, 0x26, 0xeb, 0x6b, 0x34, 0x74, 0x90, 0x4f, 0x2b, 0x49, 0x2e, 0x43, 0x76,
	0x81, 0x0e, 0xf0, 0x09, 0x1c, 0x0b, 0xb9, 0xe2, 0x52, 0xc5, 0xd1, 0x3a, 0x92, 0xea, 0x32, 0x92,
	0x68, 0xe4, 0x26, 0x38, 0x15, 0x2c, 0xe1, 0x84, 0x0a, 0x34, 0x9e, 0x87, 0x30, 0xfa, 0xf3, 0x63,
	0x1e, 0x03, 0x6a, 0xdf, 0x4d, 0xfd, 0x8d, 0xdb, 0x73, 0x71, 0xee, 0x48, 0x43, 0x14, 0xb8, 0xfa,
	0x43, 0x14, 0xc7, 0x34, 0x44, 0x3d, 0x57, 0xbb, 0x54, 0x1a, 0xa2, 0xfe, 0xbb, 0xd3, 0x2f, 0x4f,
	0xfd, 0x65, 0x2c, 0xdb, 0xad, 0x2a, 0xf2, 0x9b, 0xe5, 0x5d, 0xd5, 0x2d, 0xd7, 0xcd, 0xd0, 0x3f,
	0xdf, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x6c, 0x45, 0xeb, 0x73, 0x02, 0x00, 0x00,
}

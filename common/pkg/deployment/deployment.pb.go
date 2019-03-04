// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/deployment.proto

package deployment

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

type GithubDeploymentState int32

const (
	GithubDeploymentState_success     GithubDeploymentState = 0
	GithubDeploymentState_error       GithubDeploymentState = 1
	GithubDeploymentState_failure     GithubDeploymentState = 2
	GithubDeploymentState_inactive    GithubDeploymentState = 3
	GithubDeploymentState_in_progress GithubDeploymentState = 4
	GithubDeploymentState_queued      GithubDeploymentState = 5
	GithubDeploymentState_pending     GithubDeploymentState = 6
)

var GithubDeploymentState_name = map[int32]string{
	0: "success",
	1: "error",
	2: "failure",
	3: "inactive",
	4: "in_progress",
	5: "queued",
	6: "pending",
}

var GithubDeploymentState_value = map[string]int32{
	"success":     0,
	"error":       1,
	"failure":     2,
	"inactive":    3,
	"in_progress": 4,
	"queued":      5,
	"pending":     6,
}

func (x GithubDeploymentState) String() string {
	return proto.EnumName(GithubDeploymentState_name, int32(x))
}

func (GithubDeploymentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f728dae8917e37ee, []int{0}
}

type GithubRepository struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GithubRepository) Reset()         { *m = GithubRepository{} }
func (m *GithubRepository) String() string { return proto.CompactTextString(m) }
func (*GithubRepository) ProtoMessage()    {}
func (*GithubRepository) Descriptor() ([]byte, []int) {
	return fileDescriptor_f728dae8917e37ee, []int{0}
}

func (m *GithubRepository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GithubRepository.Unmarshal(m, b)
}
func (m *GithubRepository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GithubRepository.Marshal(b, m, deterministic)
}
func (m *GithubRepository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GithubRepository.Merge(m, src)
}
func (m *GithubRepository) XXX_Size() int {
	return xxx_messageInfo_GithubRepository.Size(m)
}
func (m *GithubRepository) XXX_DiscardUnknown() {
	xxx_messageInfo_GithubRepository.DiscardUnknown(m)
}

var xxx_messageInfo_GithubRepository proto.InternalMessageInfo

func (m *GithubRepository) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GithubRepository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeploymentSpec struct {
	Repository           *GithubRepository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	DeploymentID         int64             `protobuf:"varint,2,opt,name=deploymentID,proto3" json:"deploymentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeploymentSpec) Reset()         { *m = DeploymentSpec{} }
func (m *DeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*DeploymentSpec) ProtoMessage()    {}
func (*DeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_f728dae8917e37ee, []int{1}
}

func (m *DeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentSpec.Unmarshal(m, b)
}
func (m *DeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentSpec.Marshal(b, m, deterministic)
}
func (m *DeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentSpec.Merge(m, src)
}
func (m *DeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_DeploymentSpec.Size(m)
}
func (m *DeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentSpec proto.InternalMessageInfo

func (m *DeploymentSpec) GetRepository() *GithubRepository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *DeploymentSpec) GetDeploymentID() int64 {
	if m != nil {
		return m.DeploymentID
	}
	return 0
}

type DeploymentRequest struct {
	Deployment           *DeploymentSpec `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Timestamp            int64           `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Deadline             int64           `protobuf:"varint,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Payload              []byte          `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Cluster              string          `protobuf:"bytes,5,opt,name=cluster,proto3" json:"cluster,omitempty"`
	DeliveryID           string          `protobuf:"bytes,6,opt,name=deliveryID,proto3" json:"deliveryID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DeploymentRequest) Reset()         { *m = DeploymentRequest{} }
func (m *DeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*DeploymentRequest) ProtoMessage()    {}
func (*DeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f728dae8917e37ee, []int{2}
}

func (m *DeploymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentRequest.Unmarshal(m, b)
}
func (m *DeploymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentRequest.Marshal(b, m, deterministic)
}
func (m *DeploymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentRequest.Merge(m, src)
}
func (m *DeploymentRequest) XXX_Size() int {
	return xxx_messageInfo_DeploymentRequest.Size(m)
}
func (m *DeploymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentRequest proto.InternalMessageInfo

func (m *DeploymentRequest) GetDeployment() *DeploymentSpec {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *DeploymentRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DeploymentRequest) GetDeadline() int64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

func (m *DeploymentRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *DeploymentRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *DeploymentRequest) GetDeliveryID() string {
	if m != nil {
		return m.DeliveryID
	}
	return ""
}

type DeploymentStatus struct {
	Deployment           *DeploymentSpec       `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	State                GithubDeploymentState `protobuf:"varint,2,opt,name=state,proto3,enum=deployment.GithubDeploymentState" json:"state,omitempty"`
	Description          string                `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DeliveryID           string                `protobuf:"bytes,4,opt,name=deliveryID,proto3" json:"deliveryID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeploymentStatus) Reset()         { *m = DeploymentStatus{} }
func (m *DeploymentStatus) String() string { return proto.CompactTextString(m) }
func (*DeploymentStatus) ProtoMessage()    {}
func (*DeploymentStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f728dae8917e37ee, []int{3}
}

func (m *DeploymentStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentStatus.Unmarshal(m, b)
}
func (m *DeploymentStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentStatus.Marshal(b, m, deterministic)
}
func (m *DeploymentStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentStatus.Merge(m, src)
}
func (m *DeploymentStatus) XXX_Size() int {
	return xxx_messageInfo_DeploymentStatus.Size(m)
}
func (m *DeploymentStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentStatus proto.InternalMessageInfo

func (m *DeploymentStatus) GetDeployment() *DeploymentSpec {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *DeploymentStatus) GetState() GithubDeploymentState {
	if m != nil {
		return m.State
	}
	return GithubDeploymentState_success
}

func (m *DeploymentStatus) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DeploymentStatus) GetDeliveryID() string {
	if m != nil {
		return m.DeliveryID
	}
	return ""
}

func init() {
	proto.RegisterEnum("deployment.GithubDeploymentState", GithubDeploymentState_name, GithubDeploymentState_value)
	proto.RegisterType((*GithubRepository)(nil), "deployment.GithubRepository")
	proto.RegisterType((*DeploymentSpec)(nil), "deployment.DeploymentSpec")
	proto.RegisterType((*DeploymentRequest)(nil), "deployment.DeploymentRequest")
	proto.RegisterType((*DeploymentStatus)(nil), "deployment.DeploymentStatus")
}

func init() { proto.RegisterFile("protobuf/deployment.proto", fileDescriptor_f728dae8917e37ee) }

var fileDescriptor_f728dae8917e37ee = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x49, 0x77, 0xb3, 0x6d, 0x26, 0xab, 0x62, 0x46, 0x20, 0x85, 0xaa, 0x42, 0x4b, 0x4e,
	0x15, 0x87, 0x22, 0x95, 0x03, 0x12, 0xea, 0x71, 0x25, 0xd4, 0xab, 0x79, 0x00, 0xe4, 0x4d, 0xa6,
	0xc5, 0x52, 0x62, 0xbb, 0xfe, 0x53, 0x94, 0x47, 0xe4, 0x21, 0x78, 0x17, 0x14, 0x6f, 0x77, 0xe3,
	0x5d, 0xb8, 0x71, 0xcb, 0x7c, 0x33, 0x9e, 0xdf, 0xf7, 0xc5, 0x86, 0xb7, 0xc6, 0x6a, 0xaf, 0x37,
	0xe1, 0xfe, 0x63, 0x4b, 0xa6, 0xd3, 0x43, 0x4f, 0xca, 0x5f, 0x47, 0x0d, 0x61, 0x52, 0xea, 0x5b,
	0x60, 0x5f, 0xa5, 0xff, 0x11, 0x36, 0x9c, 0x8c, 0x76, 0xd2, 0x6b, 0x3b, 0xe0, 0x6b, 0xc8, 0xf5,
	0x4f, 0x45, 0xb6, 0xca, 0x56, 0xd9, 0x55, 0xc1, 0xb7, 0x05, 0x22, 0xcc, 0x95, 0xe8, 0xa9, 0x3a,
	0x89, 0x62, 0xfc, 0xae, 0x2d, 0x9c, 0xaf, 0xf7, 0xbb, 0xbe, 0x19, 0x6a, 0xf0, 0x16, 0xc0, 0xee,
	0x37, 0xc5, 0x05, 0xe5, 0xcd, 0xe5, 0x75, 0x62, 0xe1, 0x98, 0xc6, 0x93, 0x79, 0xac, 0x61, 0x39,
	0x8d, 0xde, 0xad, 0x23, 0x6b, 0xc6, 0x0f, 0xb4, 0xfa, 0x77, 0x06, 0xaf, 0x26, 0x28, 0xa7, 0xc7,
	0x40, 0xce, 0xe3, 0x17, 0x48, 0x52, 0x3d, 0x73, 0x2f, 0x52, 0xee, 0xa1, 0x4f, 0x9e, 0x4c, 0xe3,
	0x25, 0x14, 0x5e, 0xf6, 0xe4, 0xbc, 0xe8, 0xcd, 0x33, 0x72, 0x12, 0xf0, 0x02, 0xce, 0x5a, 0x12,
	0x6d, 0x27, 0x15, 0x55, 0xb3, 0xd8, 0xdc, 0xd7, 0x58, 0xc1, 0xa9, 0x11, 0x43, 0xa7, 0x45, 0x5b,
	0xcd, 0x57, 0xd9, 0xd5, 0x92, 0xef, 0xca, 0xb1, 0xd3, 0x74, 0xc1, 0x79, 0xb2, 0x55, 0x1e, 0x7f,
	0xd8, 0xae, 0xc4, 0x77, 0xa3, 0xd3, 0x4e, 0x3e, 0x91, 0x1d, 0xee, 0xd6, 0xd5, 0x22, 0x36, 0x13,
	0xa5, 0xfe, 0x95, 0x01, 0x4b, 0xcc, 0x7a, 0xe1, 0x83, 0xfb, 0xaf, 0x78, 0x9f, 0x21, 0x77, 0x5e,
	0xf8, 0xed, 0xcd, 0x9d, 0xdf, 0xbc, 0xff, 0xfb, 0x36, 0x0e, 0x71, 0xc4, 0xb7, 0xf3, 0xb8, 0x82,
	0xb2, 0x25, 0xd7, 0x58, 0x69, 0xbc, 0xd4, 0x2a, 0x86, 0x2f, 0x78, 0x2a, 0x1d, 0x65, 0x99, 0x1f,
	0x67, 0xf9, 0xe0, 0xe1, 0xcd, 0x3f, 0x09, 0x58, 0xc2, 0xa9, 0x0b, 0x4d, 0x43, 0xce, 0xb1, 0x17,
	0x58, 0x40, 0x4e, 0xd6, 0x6a, 0xcb, 0xb2, 0x51, 0xbf, 0x17, 0xb2, 0x0b, 0x96, 0xd8, 0x09, 0x2e,
	0xe1, 0x4c, 0x2a, 0xd1, 0x78, 0xf9, 0x44, 0x6c, 0x86, 0x2f, 0xa1, 0x94, 0xea, 0xbb, 0xb1, 0xfa,
	0xc1, 0x8e, 0xc7, 0xe6, 0x08, 0xb0, 0x78, 0x0c, 0x14, 0xa8, 0x65, 0xf9, 0x78, 0xce, 0x90, 0x6a,
	0xa5, 0x7a, 0x60, 0x8b, 0xcd, 0x22, 0x3e, 0xf3, 0x4f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xff,
	0xc1, 0xa1, 0x0f, 0x03, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/surfacers/stackdriver/config.proto

/*
Package stackdriver is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/surfacers/stackdriver/config.proto

It has these top-level messages:
	SurfacerConf
*/
package stackdriver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SurfacerConf struct {
	Batch            *bool   `protobuf:"varint,1,opt,name=batch,def=1" json:"batch,omitempty"`
	BatchSize        *uint64 `protobuf:"varint,2,opt,name=batch_size,json=batchSize,def=200" json:"batch_size,omitempty"`
	BatchTimerSec    *uint64 `protobuf:"varint,3,opt,name=batch_timer_sec,json=batchTimerSec,def=10" json:"batch_timer_sec,omitempty"`
	MonitoringUrl    *string `protobuf:"bytes,5,opt,name=monitoring_url,json=monitoringUrl,def=custom.googleapis.com/cloudprober/" json:"monitoring_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SurfacerConf) Reset()                    { *m = SurfacerConf{} }
func (m *SurfacerConf) String() string            { return proto.CompactTextString(m) }
func (*SurfacerConf) ProtoMessage()               {}
func (*SurfacerConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_SurfacerConf_Batch bool = true
const Default_SurfacerConf_BatchSize uint64 = 200
const Default_SurfacerConf_BatchTimerSec uint64 = 10
const Default_SurfacerConf_MonitoringUrl string = "custom.googleapis.com/cloudprober/"

func (m *SurfacerConf) GetBatch() bool {
	if m != nil && m.Batch != nil {
		return *m.Batch
	}
	return Default_SurfacerConf_Batch
}

func (m *SurfacerConf) GetBatchSize() uint64 {
	if m != nil && m.BatchSize != nil {
		return *m.BatchSize
	}
	return Default_SurfacerConf_BatchSize
}

func (m *SurfacerConf) GetBatchTimerSec() uint64 {
	if m != nil && m.BatchTimerSec != nil {
		return *m.BatchTimerSec
	}
	return Default_SurfacerConf_BatchTimerSec
}

func (m *SurfacerConf) GetMonitoringUrl() string {
	if m != nil && m.MonitoringUrl != nil {
		return *m.MonitoringUrl
	}
	return Default_SurfacerConf_MonitoringUrl
}

func init() {
	proto.RegisterType((*SurfacerConf)(nil), "cloudprober.surfacer.stackdriver.SurfacerConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/surfacers/stackdriver/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xc4, 0x30,
	0x10, 0x84, 0xe5, 0xfb, 0x91, 0x38, 0x8b, 0x03, 0xc9, 0x55, 0x44, 0x15, 0xa5, 0x8a, 0x28, 0xec,
	0x40, 0x99, 0x0a, 0x89, 0x8a, 0x36, 0x81, 0x3a, 0x4a, 0x7c, 0x1b, 0x9f, 0x45, 0x92, 0x8d, 0xd6,
	0x36, 0xc5, 0xbd, 0x1f, 0xef, 0x85, 0x70, 0x40, 0x97, 0x6e, 0x35, 0xdf, 0xb7, 0xc5, 0x0c, 0x7f,
	0x31, 0xd6, 0x9f, 0x43, 0x27, 0x35, 0x8e, 0xca, 0x20, 0x9a, 0x01, 0x94, 0x1e, 0x30, 0x9c, 0x66,
	0xc2, 0x0e, 0x48, 0xb9, 0x40, 0x7d, 0xab, 0x81, 0x9c, 0x72, 0xbe, 0xd5, 0x9f, 0x27, 0xb2, 0x5f,
	0x40, 0x4a, 0xe3, 0xd4, 0x5b, 0x23, 0x67, 0x42, 0x8f, 0x22, 0x5d, 0xf9, 0xf2, 0xdf, 0x97, 0x2b,
	0x3d, 0xfb, 0x66, 0xfc, 0xb6, 0xfe, 0x03, 0xaf, 0x38, 0xf5, 0xe2, 0x81, 0xef, 0xbb, 0xd6, 0xeb,
	0x73, 0xc2, 0x52, 0x96, 0xdf, 0x94, 0x3b, 0x4f, 0x01, 0xaa, 0x25, 0x12, 0x19, 0xe7, 0xf1, 0x68,
	0x9c, 0xbd, 0x40, 0xb2, 0x49, 0x59, 0xbe, 0x2b, 0xb7, 0xcf, 0x45, 0x51, 0x1d, 0x62, 0x5c, 0xdb,
	0x0b, 0x88, 0x47, 0x7e, 0xbf, 0x38, 0xde, 0x8e, 0x40, 0x8d, 0x03, 0x9d, 0x6c, 0xa3, 0xb8, 0x79,
	0x2a, 0xaa, 0x63, 0x44, 0xef, 0xbf, 0xa4, 0x06, 0x2d, 0xde, 0xf8, 0xdd, 0x88, 0x93, 0xf5, 0x48,
	0x76, 0x32, 0x4d, 0xa0, 0x21, 0xd9, 0xa7, 0x2c, 0x3f, 0x94, 0x99, 0x0e, 0xce, 0xe3, 0x28, 0x97,
	0xd2, 0xed, 0x6c, 0x5d, 0xdc, 0x60, 0x5d, 0xbe, 0x3a, 0x5e, 0x3f, 0x3f, 0x68, 0xf8, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xa6, 0x89, 0x92, 0x5c, 0x2c, 0x01, 0x00, 0x00,
}

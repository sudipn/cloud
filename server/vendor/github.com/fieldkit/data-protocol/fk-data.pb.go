// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fk-data.proto

/*
Package fk_data is a generated protocol buffer package.

It is generated from these files:
	fk-data.proto

It has these top-level messages:
	DeviceLocation
	SensorReading
	LoggedReading
	SensorInfo
	Metadata
	DataRecord
*/
package fk_data

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

type DeviceLocation struct {
	Fix       uint32  `protobuf:"varint,1,opt,name=fix" json:"fix,omitempty"`
	Time      uint64  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Longitude float32 `protobuf:"fixed32,3,opt,name=longitude" json:"longitude,omitempty"`
	Latitude  float32 `protobuf:"fixed32,4,opt,name=latitude" json:"latitude,omitempty"`
	Altitude  float32 `protobuf:"fixed32,5,opt,name=altitude" json:"altitude,omitempty"`
}

func (m *DeviceLocation) Reset()                    { *m = DeviceLocation{} }
func (m *DeviceLocation) String() string            { return proto.CompactTextString(m) }
func (*DeviceLocation) ProtoMessage()               {}
func (*DeviceLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeviceLocation) GetFix() uint32 {
	if m != nil {
		return m.Fix
	}
	return 0
}

func (m *DeviceLocation) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *DeviceLocation) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *DeviceLocation) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *DeviceLocation) GetAltitude() float32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

type SensorReading struct {
	Reading uint32  `protobuf:"varint,1,opt,name=reading" json:"reading,omitempty"`
	Time    uint64  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Sensor  uint32  `protobuf:"varint,3,opt,name=sensor" json:"sensor,omitempty"`
	Value   float32 `protobuf:"fixed32,4,opt,name=value" json:"value,omitempty"`
}

func (m *SensorReading) Reset()                    { *m = SensorReading{} }
func (m *SensorReading) String() string            { return proto.CompactTextString(m) }
func (*SensorReading) ProtoMessage()               {}
func (*SensorReading) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SensorReading) GetReading() uint32 {
	if m != nil {
		return m.Reading
	}
	return 0
}

func (m *SensorReading) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SensorReading) GetSensor() uint32 {
	if m != nil {
		return m.Sensor
	}
	return 0
}

func (m *SensorReading) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type LoggedReading struct {
	Version  uint32          `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Location *DeviceLocation `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	Reading  *SensorReading  `protobuf:"bytes,3,opt,name=reading" json:"reading,omitempty"`
}

func (m *LoggedReading) Reset()                    { *m = LoggedReading{} }
func (m *LoggedReading) String() string            { return proto.CompactTextString(m) }
func (*LoggedReading) ProtoMessage()               {}
func (*LoggedReading) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoggedReading) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *LoggedReading) GetLocation() *DeviceLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *LoggedReading) GetReading() *SensorReading {
	if m != nil {
		return m.Reading
	}
	return nil
}

type SensorInfo struct {
	Sensor        uint32 `protobuf:"varint,1,opt,name=sensor" json:"sensor,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	UnitOfMeasure string `protobuf:"bytes,3,opt,name=unitOfMeasure" json:"unitOfMeasure,omitempty"`
}

func (m *SensorInfo) Reset()                    { *m = SensorInfo{} }
func (m *SensorInfo) String() string            { return proto.CompactTextString(m) }
func (*SensorInfo) ProtoMessage()               {}
func (*SensorInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SensorInfo) GetSensor() uint32 {
	if m != nil {
		return m.Sensor
	}
	return 0
}

func (m *SensorInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SensorInfo) GetUnitOfMeasure() string {
	if m != nil {
		return m.UnitOfMeasure
	}
	return ""
}

type Metadata struct {
	DeviceId   []byte        `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Time       uint64        `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Git        string        `protobuf:"bytes,3,opt,name=git" json:"git,omitempty"`
	ResetCause uint32        `protobuf:"varint,4,opt,name=resetCause" json:"resetCause,omitempty"`
	Sensors    []*SensorInfo `protobuf:"bytes,5,rep,name=sensors" json:"sensors,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Metadata) GetDeviceId() []byte {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *Metadata) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Metadata) GetGit() string {
	if m != nil {
		return m.Git
	}
	return ""
}

func (m *Metadata) GetResetCause() uint32 {
	if m != nil {
		return m.ResetCause
	}
	return 0
}

func (m *Metadata) GetSensors() []*SensorInfo {
	if m != nil {
		return m.Sensors
	}
	return nil
}

type DataRecord struct {
	LoggedReading *LoggedReading `protobuf:"bytes,1,opt,name=loggedReading" json:"loggedReading,omitempty"`
	Metadata      *Metadata      `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *DataRecord) Reset()                    { *m = DataRecord{} }
func (m *DataRecord) String() string            { return proto.CompactTextString(m) }
func (*DataRecord) ProtoMessage()               {}
func (*DataRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DataRecord) GetLoggedReading() *LoggedReading {
	if m != nil {
		return m.LoggedReading
	}
	return nil
}

func (m *DataRecord) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceLocation)(nil), "fk_data.DeviceLocation")
	proto.RegisterType((*SensorReading)(nil), "fk_data.SensorReading")
	proto.RegisterType((*LoggedReading)(nil), "fk_data.LoggedReading")
	proto.RegisterType((*SensorInfo)(nil), "fk_data.SensorInfo")
	proto.RegisterType((*Metadata)(nil), "fk_data.Metadata")
	proto.RegisterType((*DataRecord)(nil), "fk_data.DataRecord")
}

func init() { proto.RegisterFile("fk-data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x95, 0x37, 0xe9, 0x36, 0x9d, 0xc5, 0x68, 0x19, 0xd0, 0x12, 0x21, 0x84, 0xa2, 0x88, 0x43,
	0x2e, 0xad, 0x50, 0x7b, 0xe5, 0x46, 0x2f, 0x95, 0x5a, 0x21, 0x99, 0x3b, 0xc8, 0x34, 0x4e, 0x64,
	0x35, 0xb5, 0x51, 0xe2, 0x54, 0xf0, 0x0f, 0x1c, 0xb8, 0xf3, 0xb3, 0x28, 0xb6, 0x93, 0x26, 0xa8,
	0xb7, 0x79, 0x7e, 0x33, 0x99, 0xf7, 0xde, 0x28, 0x40, 0x8b, 0xd3, 0x32, 0xe7, 0x86, 0xaf, 0x7e,
	0xd4, 0xda, 0x68, 0x9c, 0x17, 0xa7, 0x6f, 0x1d, 0x4c, 0x7f, 0x13, 0x78, 0xbe, 0x15, 0x17, 0x79,
	0x14, 0x7b, 0x7d, 0xe4, 0x46, 0x6a, 0x85, 0x8f, 0x10, 0x14, 0xf2, 0x67, 0x4c, 0x12, 0x92, 0x51,
	0xd6, 0x95, 0x88, 0x10, 0x1a, 0x79, 0x16, 0xf1, 0x5d, 0x42, 0xb2, 0x90, 0xd9, 0x1a, 0xdf, 0xc2,
	0xa2, 0xd2, 0xaa, 0x94, 0xa6, 0xcd, 0x45, 0x1c, 0x24, 0x24, 0xbb, 0x63, 0xd7, 0x07, 0x7c, 0x03,
	0x51, 0xc5, 0x8d, 0x23, 0x43, 0x4b, 0x0e, 0xb8, 0xe3, 0x78, 0xe5, 0xb9, 0x99, 0xe3, 0x7a, 0x9c,
	0x9e, 0x80, 0x7e, 0x11, 0xaa, 0xd1, 0x35, 0x13, 0x3c, 0x97, 0xaa, 0xc4, 0x18, 0xe6, 0xb5, 0x2b,
	0xbd, 0xa0, 0x1e, 0xde, 0x14, 0xf5, 0x04, 0xf7, 0x8d, 0x1d, 0xb7, 0x8a, 0x28, 0xf3, 0x08, 0x5f,
	0xc1, 0xec, 0xc2, 0xab, 0xb6, 0xd7, 0xe2, 0x40, 0xfa, 0x87, 0x00, 0xdd, 0xeb, 0xb2, 0x14, 0xf9,
	0x68, 0xdb, 0x45, 0xd4, 0x8d, 0xd4, 0xaa, 0xdf, 0xe6, 0x21, 0x6e, 0x20, 0xaa, 0x7c, 0x40, 0x76,
	0xe3, 0xc3, 0xfa, 0xf5, 0xca, 0x67, 0xb8, 0x9a, 0xe6, 0xc7, 0x86, 0x46, 0xfc, 0x70, 0x15, 0x1f,
	0xd8, 0x99, 0xa7, 0x61, 0x66, 0xe2, 0x72, 0x30, 0x95, 0x7e, 0x05, 0x70, 0xcc, 0x4e, 0x15, 0x7a,
	0x64, 0x87, 0x4c, 0xec, 0x20, 0x84, 0x8a, 0x7b, 0xeb, 0x0b, 0x66, 0x6b, 0x7c, 0x0f, 0xb4, 0x55,
	0xd2, 0x7c, 0x2e, 0x0e, 0x82, 0x37, 0x6d, 0xed, 0x6e, 0xb2, 0x60, 0xd3, 0xc7, 0xf4, 0x2f, 0x81,
	0xe8, 0x20, 0x0c, 0xef, 0x34, 0x74, 0x87, 0xc8, 0xad, 0xf4, 0x5d, 0x6e, 0x17, 0x3c, 0x63, 0x03,
	0xbe, 0x99, 0xee, 0x23, 0x04, 0xa5, 0x34, 0xfe, 0xc3, 0x5d, 0x89, 0xef, 0x00, 0x6a, 0xd1, 0x08,
	0xf3, 0x89, 0xb7, 0x8d, 0x0b, 0x97, 0xb2, 0xd1, 0x0b, 0x2e, 0x61, 0xee, 0x24, 0x37, 0xf1, 0x2c,
	0x09, 0xb2, 0x87, 0xf5, 0xcb, 0xff, 0x02, 0xe8, 0x6c, 0xb2, 0xbe, 0x27, 0xfd, 0x05, 0xb0, 0xe5,
	0x86, 0x33, 0x71, 0xd4, 0x75, 0x8e, 0x1f, 0x81, 0x56, 0xe3, 0xeb, 0x58, 0x8d, 0xe3, 0x0c, 0x27,
	0xb7, 0x63, 0xd3, 0x66, 0x5c, 0x42, 0x74, 0xf6, 0x46, 0xfd, 0xc1, 0x5e, 0x0c, 0x83, 0x7d, 0x02,
	0x6c, 0x68, 0xf9, 0x7e, 0x6f, 0xff, 0x8b, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x9d,
	0x81, 0x03, 0x28, 0x03, 0x00, 0x00,
}
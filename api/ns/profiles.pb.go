// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profiles.proto

package ns

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

type RatePolicy int32

const (
	// Drop
	RatePolicy_DROP RatePolicy = 0
	// Mark
	RatePolicy_MARK RatePolicy = 1
)

var RatePolicy_name = map[int32]string{
	0: "DROP",
	1: "MARK",
}
var RatePolicy_value = map[string]int32{
	"DROP": 0,
	"MARK": 1,
}

func (x RatePolicy) String() string {
	return proto.EnumName(RatePolicy_name, int32(x))
}
func (RatePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_profiles_c13dad724479dafe, []int{0}
}

type ServiceProfile struct {
	ServiceProfileID       string     `protobuf:"bytes,1,opt,name=serviceProfileID,proto3" json:"serviceProfileID,omitempty"`
	UlRate                 uint32     `protobuf:"varint,2,opt,name=ulRate,proto3" json:"ulRate,omitempty"`
	UlBucketSize           uint32     `protobuf:"varint,3,opt,name=ulBucketSize,proto3" json:"ulBucketSize,omitempty"`
	UlRatePolicy           RatePolicy `protobuf:"varint,4,opt,name=ulRatePolicy,proto3,enum=ns.RatePolicy" json:"ulRatePolicy,omitempty"`
	DlRate                 uint32     `protobuf:"varint,5,opt,name=dlRate,proto3" json:"dlRate,omitempty"`
	DlBucketSize           uint32     `protobuf:"varint,6,opt,name=dlBucketSize,proto3" json:"dlBucketSize,omitempty"`
	DlRatePolicy           RatePolicy `protobuf:"varint,7,opt,name=dlRatePolicy,proto3,enum=ns.RatePolicy" json:"dlRatePolicy,omitempty"`
	AddGWMetadata          bool       `protobuf:"varint,8,opt,name=addGWMetadata,proto3" json:"addGWMetadata,omitempty"`
	DevStatusReqFreq       uint32     `protobuf:"varint,9,opt,name=devStatusReqFreq,proto3" json:"devStatusReqFreq,omitempty"`
	ReportDevStatusBattery bool       `protobuf:"varint,10,opt,name=reportDevStatusBattery,proto3" json:"reportDevStatusBattery,omitempty"`
	ReportDevStatusMargin  bool       `protobuf:"varint,11,opt,name=reportDevStatusMargin,proto3" json:"reportDevStatusMargin,omitempty"`
	DrMin                  uint32     `protobuf:"varint,12,opt,name=drMin,proto3" json:"drMin,omitempty"`
	DrMax                  uint32     `protobuf:"varint,13,opt,name=drMax,proto3" json:"drMax,omitempty"`
	ChannelMask            []byte     `protobuf:"bytes,14,opt,name=channelMask,proto3" json:"channelMask,omitempty"`
	PrAllowed              bool       `protobuf:"varint,15,opt,name=prAllowed,proto3" json:"prAllowed,omitempty"`
	HrAllowed              bool       `protobuf:"varint,16,opt,name=hrAllowed,proto3" json:"hrAllowed,omitempty"`
	RaAllowed              bool       `protobuf:"varint,17,opt,name=raAllowed,proto3" json:"raAllowed,omitempty"`
	NwkGeoLoc              bool       `protobuf:"varint,18,opt,name=nwkGeoLoc,proto3" json:"nwkGeoLoc,omitempty"`
	TargetPER              uint32     `protobuf:"varint,19,opt,name=targetPER,proto3" json:"targetPER,omitempty"`
	MinGWDiversity         uint32     `protobuf:"varint,20,opt,name=minGWDiversity,proto3" json:"minGWDiversity,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}   `json:"-"`
	XXX_unrecognized       []byte     `json:"-"`
	XXX_sizecache          int32      `json:"-"`
}

func (m *ServiceProfile) Reset()         { *m = ServiceProfile{} }
func (m *ServiceProfile) String() string { return proto.CompactTextString(m) }
func (*ServiceProfile) ProtoMessage()    {}
func (*ServiceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_profiles_c13dad724479dafe, []int{0}
}
func (m *ServiceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceProfile.Unmarshal(m, b)
}
func (m *ServiceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceProfile.Marshal(b, m, deterministic)
}
func (dst *ServiceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceProfile.Merge(dst, src)
}
func (m *ServiceProfile) XXX_Size() int {
	return xxx_messageInfo_ServiceProfile.Size(m)
}
func (m *ServiceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceProfile proto.InternalMessageInfo

func (m *ServiceProfile) GetServiceProfileID() string {
	if m != nil {
		return m.ServiceProfileID
	}
	return ""
}

func (m *ServiceProfile) GetUlRate() uint32 {
	if m != nil {
		return m.UlRate
	}
	return 0
}

func (m *ServiceProfile) GetUlBucketSize() uint32 {
	if m != nil {
		return m.UlBucketSize
	}
	return 0
}

func (m *ServiceProfile) GetUlRatePolicy() RatePolicy {
	if m != nil {
		return m.UlRatePolicy
	}
	return RatePolicy_DROP
}

func (m *ServiceProfile) GetDlRate() uint32 {
	if m != nil {
		return m.DlRate
	}
	return 0
}

func (m *ServiceProfile) GetDlBucketSize() uint32 {
	if m != nil {
		return m.DlBucketSize
	}
	return 0
}

func (m *ServiceProfile) GetDlRatePolicy() RatePolicy {
	if m != nil {
		return m.DlRatePolicy
	}
	return RatePolicy_DROP
}

func (m *ServiceProfile) GetAddGWMetadata() bool {
	if m != nil {
		return m.AddGWMetadata
	}
	return false
}

func (m *ServiceProfile) GetDevStatusReqFreq() uint32 {
	if m != nil {
		return m.DevStatusReqFreq
	}
	return 0
}

func (m *ServiceProfile) GetReportDevStatusBattery() bool {
	if m != nil {
		return m.ReportDevStatusBattery
	}
	return false
}

func (m *ServiceProfile) GetReportDevStatusMargin() bool {
	if m != nil {
		return m.ReportDevStatusMargin
	}
	return false
}

func (m *ServiceProfile) GetDrMin() uint32 {
	if m != nil {
		return m.DrMin
	}
	return 0
}

func (m *ServiceProfile) GetDrMax() uint32 {
	if m != nil {
		return m.DrMax
	}
	return 0
}

func (m *ServiceProfile) GetChannelMask() []byte {
	if m != nil {
		return m.ChannelMask
	}
	return nil
}

func (m *ServiceProfile) GetPrAllowed() bool {
	if m != nil {
		return m.PrAllowed
	}
	return false
}

func (m *ServiceProfile) GetHrAllowed() bool {
	if m != nil {
		return m.HrAllowed
	}
	return false
}

func (m *ServiceProfile) GetRaAllowed() bool {
	if m != nil {
		return m.RaAllowed
	}
	return false
}

func (m *ServiceProfile) GetNwkGeoLoc() bool {
	if m != nil {
		return m.NwkGeoLoc
	}
	return false
}

func (m *ServiceProfile) GetTargetPER() uint32 {
	if m != nil {
		return m.TargetPER
	}
	return 0
}

func (m *ServiceProfile) GetMinGWDiversity() uint32 {
	if m != nil {
		return m.MinGWDiversity
	}
	return 0
}

type DeviceProfile struct {
	DeviceProfileID      string   `protobuf:"bytes,1,opt,name=deviceProfileID,proto3" json:"deviceProfileID,omitempty"`
	SupportsClassB       bool     `protobuf:"varint,2,opt,name=supportsClassB,proto3" json:"supportsClassB,omitempty"`
	ClassBTimeout        uint32   `protobuf:"varint,3,opt,name=classBTimeout,proto3" json:"classBTimeout,omitempty"`
	PingSlotPeriod       uint32   `protobuf:"varint,4,opt,name=pingSlotPeriod,proto3" json:"pingSlotPeriod,omitempty"`
	PingSlotDR           uint32   `protobuf:"varint,5,opt,name=pingSlotDR,proto3" json:"pingSlotDR,omitempty"`
	PingSlotFreq         uint32   `protobuf:"varint,6,opt,name=pingSlotFreq,proto3" json:"pingSlotFreq,omitempty"`
	SupportsClassC       bool     `protobuf:"varint,7,opt,name=supportsClassC,proto3" json:"supportsClassC,omitempty"`
	ClassCTimeout        uint32   `protobuf:"varint,8,opt,name=classCTimeout,proto3" json:"classCTimeout,omitempty"`
	MacVersion           string   `protobuf:"bytes,9,opt,name=macVersion,proto3" json:"macVersion,omitempty"`
	RegParamsRevision    string   `protobuf:"bytes,10,opt,name=regParamsRevision,proto3" json:"regParamsRevision,omitempty"`
	RxDelay1             uint32   `protobuf:"varint,11,opt,name=rxDelay1,proto3" json:"rxDelay1,omitempty"`
	RxDROffset1          uint32   `protobuf:"varint,12,opt,name=rxDROffset1,proto3" json:"rxDROffset1,omitempty"`
	RxDataRate2          uint32   `protobuf:"varint,13,opt,name=rxDataRate2,proto3" json:"rxDataRate2,omitempty"`
	RxFreq2              uint32   `protobuf:"varint,14,opt,name=rxFreq2,proto3" json:"rxFreq2,omitempty"`
	FactoryPresetFreqs   []uint32 `protobuf:"varint,15,rep,packed,name=factoryPresetFreqs,proto3" json:"factoryPresetFreqs,omitempty"`
	MaxEIRP              uint32   `protobuf:"varint,16,opt,name=maxEIRP,proto3" json:"maxEIRP,omitempty"`
	MaxDutyCycle         uint32   `protobuf:"varint,17,opt,name=maxDutyCycle,proto3" json:"maxDutyCycle,omitempty"`
	SupportsJoin         bool     `protobuf:"varint,18,opt,name=supportsJoin,proto3" json:"supportsJoin,omitempty"`
	RfRegion             string   `protobuf:"bytes,19,opt,name=rfRegion,proto3" json:"rfRegion,omitempty"`
	Supports32BitFCnt    bool     `protobuf:"varint,20,opt,name=supports32bitFCnt,proto3" json:"supports32bitFCnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProfile) Reset()         { *m = DeviceProfile{} }
func (m *DeviceProfile) String() string { return proto.CompactTextString(m) }
func (*DeviceProfile) ProtoMessage()    {}
func (*DeviceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_profiles_c13dad724479dafe, []int{1}
}
func (m *DeviceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProfile.Unmarshal(m, b)
}
func (m *DeviceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProfile.Marshal(b, m, deterministic)
}
func (dst *DeviceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProfile.Merge(dst, src)
}
func (m *DeviceProfile) XXX_Size() int {
	return xxx_messageInfo_DeviceProfile.Size(m)
}
func (m *DeviceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProfile proto.InternalMessageInfo

func (m *DeviceProfile) GetDeviceProfileID() string {
	if m != nil {
		return m.DeviceProfileID
	}
	return ""
}

func (m *DeviceProfile) GetSupportsClassB() bool {
	if m != nil {
		return m.SupportsClassB
	}
	return false
}

func (m *DeviceProfile) GetClassBTimeout() uint32 {
	if m != nil {
		return m.ClassBTimeout
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotPeriod() uint32 {
	if m != nil {
		return m.PingSlotPeriod
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotDR() uint32 {
	if m != nil {
		return m.PingSlotDR
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotFreq() uint32 {
	if m != nil {
		return m.PingSlotFreq
	}
	return 0
}

func (m *DeviceProfile) GetSupportsClassC() bool {
	if m != nil {
		return m.SupportsClassC
	}
	return false
}

func (m *DeviceProfile) GetClassCTimeout() uint32 {
	if m != nil {
		return m.ClassCTimeout
	}
	return 0
}

func (m *DeviceProfile) GetMacVersion() string {
	if m != nil {
		return m.MacVersion
	}
	return ""
}

func (m *DeviceProfile) GetRegParamsRevision() string {
	if m != nil {
		return m.RegParamsRevision
	}
	return ""
}

func (m *DeviceProfile) GetRxDelay1() uint32 {
	if m != nil {
		return m.RxDelay1
	}
	return 0
}

func (m *DeviceProfile) GetRxDROffset1() uint32 {
	if m != nil {
		return m.RxDROffset1
	}
	return 0
}

func (m *DeviceProfile) GetRxDataRate2() uint32 {
	if m != nil {
		return m.RxDataRate2
	}
	return 0
}

func (m *DeviceProfile) GetRxFreq2() uint32 {
	if m != nil {
		return m.RxFreq2
	}
	return 0
}

func (m *DeviceProfile) GetFactoryPresetFreqs() []uint32 {
	if m != nil {
		return m.FactoryPresetFreqs
	}
	return nil
}

func (m *DeviceProfile) GetMaxEIRP() uint32 {
	if m != nil {
		return m.MaxEIRP
	}
	return 0
}

func (m *DeviceProfile) GetMaxDutyCycle() uint32 {
	if m != nil {
		return m.MaxDutyCycle
	}
	return 0
}

func (m *DeviceProfile) GetSupportsJoin() bool {
	if m != nil {
		return m.SupportsJoin
	}
	return false
}

func (m *DeviceProfile) GetRfRegion() string {
	if m != nil {
		return m.RfRegion
	}
	return ""
}

func (m *DeviceProfile) GetSupports32BitFCnt() bool {
	if m != nil {
		return m.Supports32BitFCnt
	}
	return false
}

func init() {
	proto.RegisterType((*ServiceProfile)(nil), "ns.ServiceProfile")
	proto.RegisterType((*DeviceProfile)(nil), "ns.DeviceProfile")
	proto.RegisterEnum("ns.RatePolicy", RatePolicy_name, RatePolicy_value)
}

func init() { proto.RegisterFile("profiles.proto", fileDescriptor_profiles_c13dad724479dafe) }

var fileDescriptor_profiles_c13dad724479dafe = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x6f, 0x4f, 0xdb, 0x30,
	0x10, 0x87, 0x57, 0xa0, 0xd0, 0x1a, 0x52, 0x8a, 0x61, 0xc8, 0x9a, 0xa6, 0x29, 0x42, 0xd3, 0x54,
	0xa1, 0xa9, 0x12, 0x65, 0xda, 0x7b, 0x68, 0x00, 0xb1, 0xad, 0x22, 0x72, 0xa7, 0xf1, 0xda, 0x24,
	0xd7, 0x62, 0x91, 0xc6, 0xc5, 0x76, 0x4b, 0xbb, 0xef, 0x3b, 0x69, 0x1f, 0x63, 0xb2, 0xf3, 0x87,
	0x24, 0xed, 0xde, 0xe5, 0x9e, 0xdf, 0x91, 0x3b, 0x1d, 0x4f, 0x83, 0x5a, 0x53, 0x29, 0x46, 0x3c,
	0x02, 0xd5, 0x9d, 0x4a, 0xa1, 0x05, 0xde, 0x88, 0xd5, 0xc9, 0xdf, 0x3a, 0x6a, 0x0d, 0x41, 0xce,
	0x79, 0x00, 0x7e, 0x92, 0xe2, 0x53, 0xd4, 0x56, 0x25, 0x72, 0xeb, 0x91, 0x9a, 0x5b, 0xeb, 0x34,
	0xe9, 0x0a, 0xc7, 0xc7, 0x68, 0x7b, 0x16, 0x51, 0xa6, 0x81, 0x6c, 0xb8, 0xb5, 0x8e, 0x43, 0xd3,
	0x0a, 0x9f, 0xa0, 0xbd, 0x59, 0x74, 0x39, 0x0b, 0x9e, 0x40, 0x0f, 0xf9, 0x6f, 0x20, 0x9b, 0x36,
	0x2d, 0x31, 0xdc, 0x33, 0x3d, 0xa6, 0xdb, 0x17, 0x11, 0x0f, 0x96, 0x64, 0xcb, 0xad, 0x75, 0x5a,
	0xbd, 0x56, 0x37, 0x56, 0xdd, 0x57, 0x4a, 0x4b, 0x3d, 0x66, 0x5e, 0x98, 0xcc, 0xab, 0x27, 0xf3,
	0xc2, 0x7c, 0x5e, 0x58, 0x9c, 0xb7, 0x9d, 0xcc, 0x0b, 0x2b, 0xf3, 0xc2, 0xe2, 0xbc, 0x9d, 0xf5,
	0xf3, 0x8a, 0x3d, 0xf8, 0x23, 0x72, 0x58, 0x18, 0xde, 0xdc, 0x0f, 0x40, 0xb3, 0x90, 0x69, 0x46,
	0x1a, 0x6e, 0xad, 0xd3, 0xa0, 0x65, 0x68, 0x2e, 0x16, 0xc2, 0x7c, 0xa8, 0x99, 0x9e, 0x29, 0x0a,
	0xcf, 0xd7, 0x12, 0x9e, 0x49, 0xd3, 0x6e, 0xb0, 0xc2, 0xf1, 0x57, 0x74, 0x2c, 0x61, 0x2a, 0xa4,
	0xf6, 0xb2, 0xe4, 0x92, 0x69, 0x0d, 0x72, 0x49, 0x90, 0x7d, 0xf5, 0x7f, 0x52, 0xfc, 0x05, 0xbd,
	0xad, 0x24, 0x03, 0x26, 0xc7, 0x3c, 0x26, 0xbb, 0xf6, 0xcf, 0xd6, 0x87, 0xf8, 0x08, 0xd5, 0x43,
	0x39, 0xe0, 0x31, 0xd9, 0xb3, 0xeb, 0x24, 0x45, 0x4a, 0xd9, 0x82, 0x38, 0x39, 0x65, 0x0b, 0xec,
	0xa2, 0xdd, 0xe0, 0x91, 0xc5, 0x31, 0x44, 0x03, 0xa6, 0x9e, 0x48, 0xcb, 0xad, 0x75, 0xf6, 0x68,
	0x11, 0xe1, 0xf7, 0xa8, 0x39, 0x95, 0x17, 0x51, 0x24, 0x5e, 0x20, 0x24, 0xfb, 0x76, 0xee, 0x2b,
	0x30, 0xe9, 0x63, 0x9e, 0xb6, 0x93, 0xf4, 0xb1, 0x98, 0x4a, 0x96, 0xa5, 0x07, 0x49, 0x9a, 0x03,
	0x93, 0xc6, 0x2f, 0x4f, 0x37, 0x20, 0x7e, 0x88, 0x80, 0xe0, 0x24, 0xcd, 0x81, 0x49, 0x35, 0x93,
	0x63, 0xd0, 0xfe, 0x15, 0x25, 0x87, 0x76, 0xe7, 0x57, 0x80, 0x3f, 0xa1, 0xd6, 0x84, 0xc7, 0x37,
	0xf7, 0x1e, 0x9f, 0x83, 0x54, 0x5c, 0x2f, 0xc9, 0x91, 0x6d, 0xa9, 0xd0, 0x93, 0x3f, 0x75, 0xe4,
	0x78, 0x50, 0x34, 0xbd, 0x83, 0xf6, 0x43, 0x58, 0x27, 0x7a, 0x15, 0x9b, 0x19, 0x6a, 0x36, 0x35,
	0x17, 0x56, 0xfd, 0x88, 0x29, 0x75, 0x69, 0x7d, 0x6f, 0xd0, 0x0a, 0x35, 0xbe, 0x04, 0xf6, 0xe9,
	0x27, 0x9f, 0x80, 0x98, 0xe9, 0x54, 0xfc, 0x32, 0x34, 0x6f, 0x9b, 0xf2, 0x78, 0x3c, 0x8c, 0x84,
	0xf6, 0x41, 0x72, 0x11, 0x5a, 0xf7, 0x1d, 0x5a, 0xa1, 0xf8, 0x03, 0x42, 0x19, 0xf1, 0x68, 0x6a,
	0x7c, 0x81, 0x18, 0xeb, 0xb3, 0xca, 0x3a, 0x97, 0x5a, 0x5f, 0x64, 0x2b, 0x9b, 0xf7, 0xad, 0xf7,
	0xd5, 0xcd, 0xfb, 0xf9, 0xe6, 0xfd, 0x6c, 0xf3, 0x46, 0x61, 0xf3, 0x0c, 0x9a, 0x8d, 0x26, 0x2c,
	0xf8, 0x65, 0x2e, 0x2a, 0x62, 0xeb, 0x78, 0x93, 0x16, 0x08, 0xfe, 0x8c, 0x0e, 0x24, 0x8c, 0x7d,
	0x26, 0xd9, 0x44, 0x51, 0x98, 0x73, 0xdb, 0x86, 0x6c, 0xdb, 0x6a, 0x80, 0xdf, 0xa1, 0x86, 0x5c,
	0x78, 0x10, 0xb1, 0xe5, 0x99, 0xd5, 0xd8, 0xa1, 0x79, 0x6d, 0x6c, 0x94, 0x0b, 0x8f, 0xde, 0x8d,
	0x46, 0x0a, 0xf4, 0x59, 0xea, 0x6f, 0x11, 0xa5, 0x1d, 0x4c, 0x33, 0xf3, 0x7b, 0xed, 0xa5, 0x2e,
	0x17, 0x11, 0x26, 0x68, 0x47, 0x2e, 0xcc, 0x15, 0x7a, 0xd6, 0x66, 0x87, 0x66, 0x25, 0xee, 0x22,
	0x3c, 0x62, 0x81, 0x16, 0x72, 0xe9, 0x4b, 0x50, 0x60, 0x4f, 0xa5, 0xc8, 0xbe, 0xbb, 0xd9, 0x71,
	0xe8, 0x9a, 0xc4, 0xbc, 0x69, 0xc2, 0x16, 0x57, 0xb7, 0xd4, 0xb7, 0x66, 0x3b, 0x34, 0x2b, 0xcd,
	0xff, 0x60, 0xc2, 0x16, 0xde, 0x4c, 0x2f, 0xfb, 0xcb, 0x20, 0x02, 0xab, 0xb6, 0x43, 0x4b, 0xcc,
	0xf4, 0x64, 0xd7, 0xfe, 0x26, 0x78, 0x9c, 0x0a, 0x5e, 0x62, 0xf6, 0x16, 0x23, 0x0a, 0x63, 0x73,
	0xb0, 0x43, 0x7b, 0xb0, 0xbc, 0x36, 0x57, 0xcd, 0x7a, 0xcf, 0x7b, 0x0f, 0x5c, 0x5f, 0xf7, 0x63,
	0x6d, 0x25, 0x6f, 0xd0, 0xd5, 0xe0, 0xd4, 0x45, 0xa8, 0xf0, 0x05, 0x6b, 0xa0, 0x2d, 0x8f, 0xde,
	0xf9, 0xed, 0x37, 0xe6, 0x69, 0x70, 0x41, 0xbf, 0xb7, 0x6b, 0x0f, 0xdb, 0xf6, 0xfb, 0x7f, 0xfe,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x09, 0x9c, 0x54, 0x76, 0x11, 0x06, 0x00, 0x00,
}

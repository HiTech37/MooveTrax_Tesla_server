// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: protos/vehicle_alert.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Audience the target audience for the alert.
type Audience int32

const (
	Audience_Unknown    Audience = 0
	Audience_Customer   Audience = 1
	Audience_Service    Audience = 2
	Audience_ServiceFix Audience = 3
)

// Enum value maps for Audience.
var (
	Audience_name = map[int32]string{
		0: "Unknown",
		1: "Customer",
		2: "Service",
		3: "ServiceFix",
	}
	Audience_value = map[string]int32{
		"Unknown":    0,
		"Customer":   1,
		"Service":    2,
		"ServiceFix": 3,
	}
)

func (x Audience) Enum() *Audience {
	p := new(Audience)
	*p = x
	return p
}

func (x Audience) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Audience) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_vehicle_alert_proto_enumTypes[0].Descriptor()
}

func (Audience) Type() protoreflect.EnumType {
	return &file_protos_vehicle_alert_proto_enumTypes[0]
}

func (x Audience) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Audience.Descriptor instead.
func (Audience) EnumDescriptor() ([]byte, []int) {
	return file_protos_vehicle_alert_proto_rawDescGZIP(), []int{0}
}

// VehicleAlerts is a collection of recent alerts.
type VehicleAlerts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts    []*VehicleAlert        `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Vin       string                 `protobuf:"bytes,3,opt,name=vin,proto3" json:"vin,omitempty"`
}

func (x *VehicleAlerts) Reset() {
	*x = VehicleAlerts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vehicle_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleAlerts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleAlerts) ProtoMessage() {}

func (x *VehicleAlerts) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vehicle_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleAlerts.ProtoReflect.Descriptor instead.
func (*VehicleAlerts) Descriptor() ([]byte, []int) {
	return file_protos_vehicle_alert_proto_rawDescGZIP(), []int{0}
}

func (x *VehicleAlerts) GetAlerts() []*VehicleAlert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

func (x *VehicleAlerts) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *VehicleAlerts) GetVin() string {
	if x != nil {
		return x.Vin
	}
	return ""
}

// VehicleAlert is a single alert and is active if ended_at is nil.
type VehicleAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Audiences []Audience             `protobuf:"varint,2,rep,packed,name=audiences,proto3,enum=telemetry.vehicle_alerts.Audience" json:"audiences,omitempty"`
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
}

func (x *VehicleAlert) Reset() {
	*x = VehicleAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vehicle_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleAlert) ProtoMessage() {}

func (x *VehicleAlert) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vehicle_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleAlert.ProtoReflect.Descriptor instead.
func (*VehicleAlert) Descriptor() ([]byte, []int) {
	return file_protos_vehicle_alert_proto_rawDescGZIP(), []int{1}
}

func (x *VehicleAlert) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VehicleAlert) GetAudiences() []Audience {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *VehicleAlert) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *VehicleAlert) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

var File_protos_vehicle_alert_proto protoreflect.FileDescriptor

var file_protos_vehicle_alert_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x76, 0x69, 0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x0c, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x2a,
	0x42, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x69,
	0x78, 0x10, 0x03, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x73, 0x6c, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2d, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_vehicle_alert_proto_rawDescOnce sync.Once
	file_protos_vehicle_alert_proto_rawDescData = file_protos_vehicle_alert_proto_rawDesc
)

func file_protos_vehicle_alert_proto_rawDescGZIP() []byte {
	file_protos_vehicle_alert_proto_rawDescOnce.Do(func() {
		file_protos_vehicle_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_vehicle_alert_proto_rawDescData)
	})
	return file_protos_vehicle_alert_proto_rawDescData
}

var file_protos_vehicle_alert_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_vehicle_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_vehicle_alert_proto_goTypes = []interface{}{
	(Audience)(0),                 // 0: telemetry.vehicle_alerts.Audience
	(*VehicleAlerts)(nil),         // 1: telemetry.vehicle_alerts.VehicleAlerts
	(*VehicleAlert)(nil),          // 2: telemetry.vehicle_alerts.VehicleAlert
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_protos_vehicle_alert_proto_depIdxs = []int32{
	2, // 0: telemetry.vehicle_alerts.VehicleAlerts.alerts:type_name -> telemetry.vehicle_alerts.VehicleAlert
	3, // 1: telemetry.vehicle_alerts.VehicleAlerts.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: telemetry.vehicle_alerts.VehicleAlert.audiences:type_name -> telemetry.vehicle_alerts.Audience
	3, // 3: telemetry.vehicle_alerts.VehicleAlert.started_at:type_name -> google.protobuf.Timestamp
	3, // 4: telemetry.vehicle_alerts.VehicleAlert.ended_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_vehicle_alert_proto_init() }
func file_protos_vehicle_alert_proto_init() {
	if File_protos_vehicle_alert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_vehicle_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleAlerts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_vehicle_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleAlert); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_vehicle_alert_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_vehicle_alert_proto_goTypes,
		DependencyIndexes: file_protos_vehicle_alert_proto_depIdxs,
		EnumInfos:         file_protos_vehicle_alert_proto_enumTypes,
		MessageInfos:      file_protos_vehicle_alert_proto_msgTypes,
	}.Build()
	File_protos_vehicle_alert_proto = out.File
	file_protos_vehicle_alert_proto_rawDesc = nil
	file_protos_vehicle_alert_proto_goTypes = nil
	file_protos_vehicle_alert_proto_depIdxs = nil
}
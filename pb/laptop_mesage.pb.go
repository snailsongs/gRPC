// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/laptop_mesage.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Laptop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are assignable to Weight:
	//
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight      isLaptop_Weight `protobuf_oneof:"weight"`
	Price       float64         `protobuf:"fixed64,12,opt,name=price,proto3" json:"price,omitempty"`
	ReleaseYear uint32          `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
}

func (x *Laptop) Reset() {
	*x = Laptop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptop_mesage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Laptop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Laptop) ProtoMessage() {}

func (x *Laptop) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptop_mesage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Laptop.ProtoReflect.Descriptor instead.
func (*Laptop) Descriptor() ([]byte, []int) {
	return file_proto_laptop_mesage_proto_rawDescGZIP(), []int{0}
}

func (x *Laptop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Laptop) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Laptop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Laptop) GetCpu() *CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *Laptop) GetRam() *Memory {
	if x != nil {
		return x.Ram
	}
	return nil
}

func (x *Laptop) GetGpus() []*GPU {
	if x != nil {
		return x.Gpus
	}
	return nil
}

func (x *Laptop) GetStorages() []*Storage {
	if x != nil {
		return x.Storages
	}
	return nil
}

func (x *Laptop) GetScreen() *Screen {
	if x != nil {
		return x.Screen
	}
	return nil
}

func (x *Laptop) GetKeyboard() *Keyboard {
	if x != nil {
		return x.Keyboard
	}
	return nil
}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (x *Laptop) GetWeightKg() float64 {
	if x, ok := x.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (x *Laptop) GetWeightLb() float64 {
	if x, ok := x.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (x *Laptop) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Laptop) GetReleaseYear() uint32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

var File_proto_laptop_mesage_proto protoreflect.FileDescriptor

var file_proto_laptop_mesage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x6d,
	0x65, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79,
	0x62, 0x6f, 0x72, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x03, 0x0a, 0x06, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x50,
	0x55, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x1f, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x52, 0x03, 0x72, 0x61, 0x6d, 0x12, 0x1e, 0x0a, 0x04, 0x67, 0x70, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x50,
	0x55, 0x52, 0x04, 0x67, 0x70, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x52, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x2b, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x6b, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x08, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x4b, 0x67, 0x12, 0x1d, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x6c, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x08, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x4c, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_laptop_mesage_proto_rawDescOnce sync.Once
	file_proto_laptop_mesage_proto_rawDescData = file_proto_laptop_mesage_proto_rawDesc
)

func file_proto_laptop_mesage_proto_rawDescGZIP() []byte {
	file_proto_laptop_mesage_proto_rawDescOnce.Do(func() {
		file_proto_laptop_mesage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_laptop_mesage_proto_rawDescData)
	})
	return file_proto_laptop_mesage_proto_rawDescData
}

var file_proto_laptop_mesage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_laptop_mesage_proto_goTypes = []interface{}{
	(*Laptop)(nil),   // 0: proto.Laptop
	(*CPU)(nil),      // 1: proto.CPU
	(*Memory)(nil),   // 2: proto.Memory
	(*GPU)(nil),      // 3: proto.GPU
	(*Storage)(nil),  // 4: proto.Storage
	(*Screen)(nil),   // 5: proto.Screen
	(*Keyboard)(nil), // 6: proto.Keyboard
}
var file_proto_laptop_mesage_proto_depIdxs = []int32{
	1, // 0: proto.Laptop.cpu:type_name -> proto.CPU
	2, // 1: proto.Laptop.ram:type_name -> proto.Memory
	3, // 2: proto.Laptop.gpus:type_name -> proto.GPU
	4, // 3: proto.Laptop.storages:type_name -> proto.Storage
	5, // 4: proto.Laptop.screen:type_name -> proto.Screen
	6, // 5: proto.Laptop.keyboard:type_name -> proto.Keyboard
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_laptop_mesage_proto_init() }
func file_proto_laptop_mesage_proto_init() {
	if File_proto_laptop_mesage_proto != nil {
		return
	}
	file_proto_processor_message_proto_init()
	file_proto_memory_message_proto_init()
	file_proto_storge_message_proto_init()
	file_proto_screen_message_proto_init()
	file_proto_keyborad_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_laptop_mesage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Laptop); i {
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
	file_proto_laptop_mesage_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_laptop_mesage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_laptop_mesage_proto_goTypes,
		DependencyIndexes: file_proto_laptop_mesage_proto_depIdxs,
		MessageInfos:      file_proto_laptop_mesage_proto_msgTypes,
	}.Build()
	File_proto_laptop_mesage_proto = out.File
	file_proto_laptop_mesage_proto_rawDesc = nil
	file_proto_laptop_mesage_proto_goTypes = nil
	file_proto_laptop_mesage_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.28.3
// source: ssl_simulation_robot_feedback.proto

package sim

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Feedback from a robot
type RobotFeedback struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of the robot
	Id *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Has the dribbler contact to the ball right now
	DribblerBallContact *bool `protobuf:"varint,2,opt,name=dribbler_ball_contact,json=dribblerBallContact" json:"dribbler_ball_contact,omitempty"`
	// Custom robot feedback for specific simulators (the protobuf files are managed by the simulators)
	Custom        *anypb.Any `protobuf:"bytes,3,opt,name=custom" json:"custom,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RobotFeedback) Reset() {
	*x = RobotFeedback{}
	mi := &file_ssl_simulation_robot_feedback_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RobotFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotFeedback) ProtoMessage() {}

func (x *RobotFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_robot_feedback_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotFeedback.ProtoReflect.Descriptor instead.
func (*RobotFeedback) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_robot_feedback_proto_rawDescGZIP(), []int{0}
}

func (x *RobotFeedback) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *RobotFeedback) GetDribblerBallContact() bool {
	if x != nil && x.DribblerBallContact != nil {
		return *x.DribblerBallContact
	}
	return false
}

func (x *RobotFeedback) GetCustom() *anypb.Any {
	if x != nil {
		return x.Custom
	}
	return nil
}

// Response to RobotControl from the simulator to the connected client
type RobotControlResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of errors, like using unsupported features
	Errors []*SimulatorError `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty"`
	// Feedback of the robots
	Feedback      []*RobotFeedback `protobuf:"bytes,2,rep,name=feedback" json:"feedback,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RobotControlResponse) Reset() {
	*x = RobotControlResponse{}
	mi := &file_ssl_simulation_robot_feedback_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RobotControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotControlResponse) ProtoMessage() {}

func (x *RobotControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_robot_feedback_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotControlResponse.ProtoReflect.Descriptor instead.
func (*RobotControlResponse) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_robot_feedback_proto_rawDescGZIP(), []int{1}
}

func (x *RobotControlResponse) GetErrors() []*SimulatorError {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *RobotControlResponse) GetFeedback() []*RobotFeedback {
	if x != nil {
		return x.Feedback
	}
	return nil
}

var File_ssl_simulation_robot_feedback_proto protoreflect.FileDescriptor

var file_ssl_simulation_robot_feedback_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x73, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x73, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a,
	0x0d, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32,
	0x0a, 0x15, 0x64, 0x72, 0x69, 0x62, 0x62, 0x6c, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x6c, 0x6c, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x64,
	0x72, 0x69, 0x62, 0x62, 0x6c, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x22, 0x6b, 0x0a, 0x14, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x12, 0x2a, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f,
	0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x73, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x69, 0x6d,
}

var (
	file_ssl_simulation_robot_feedback_proto_rawDescOnce sync.Once
	file_ssl_simulation_robot_feedback_proto_rawDescData = file_ssl_simulation_robot_feedback_proto_rawDesc
)

func file_ssl_simulation_robot_feedback_proto_rawDescGZIP() []byte {
	file_ssl_simulation_robot_feedback_proto_rawDescOnce.Do(func() {
		file_ssl_simulation_robot_feedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_simulation_robot_feedback_proto_rawDescData)
	})
	return file_ssl_simulation_robot_feedback_proto_rawDescData
}

var file_ssl_simulation_robot_feedback_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ssl_simulation_robot_feedback_proto_goTypes = []any{
	(*RobotFeedback)(nil),        // 0: RobotFeedback
	(*RobotControlResponse)(nil), // 1: RobotControlResponse
	(*anypb.Any)(nil),            // 2: google.protobuf.Any
	(*SimulatorError)(nil),       // 3: SimulatorError
}
var file_ssl_simulation_robot_feedback_proto_depIdxs = []int32{
	2, // 0: RobotFeedback.custom:type_name -> google.protobuf.Any
	3, // 1: RobotControlResponse.errors:type_name -> SimulatorError
	0, // 2: RobotControlResponse.feedback:type_name -> RobotFeedback
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ssl_simulation_robot_feedback_proto_init() }
func file_ssl_simulation_robot_feedback_proto_init() {
	if File_ssl_simulation_robot_feedback_proto != nil {
		return
	}
	file_ssl_simulation_error_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_simulation_robot_feedback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_simulation_robot_feedback_proto_goTypes,
		DependencyIndexes: file_ssl_simulation_robot_feedback_proto_depIdxs,
		MessageInfos:      file_ssl_simulation_robot_feedback_proto_msgTypes,
	}.Build()
	File_ssl_simulation_robot_feedback_proto = out.File
	file_ssl_simulation_robot_feedback_proto_rawDesc = nil
	file_ssl_simulation_robot_feedback_proto_goTypes = nil
	file_ssl_simulation_robot_feedback_proto_depIdxs = nil
}

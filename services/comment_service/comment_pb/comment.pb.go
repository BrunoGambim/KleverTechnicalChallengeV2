// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: grpc_proto_files/comment.proto

package comment_pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	post_pb "relembrando_2/services/post_service/post_pb"
	upvote_pb "relembrando_2/services/upvote_service/upvote_pb"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_files_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_files_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_files_comment_proto_rawDescGZIP(), []int{0}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	PostId  string `protobuf:"bytes,2,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *CommentRequest) Reset() {
	*x = CommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_files_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentRequest) ProtoMessage() {}

func (x *CommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_files_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentRequest.ProtoReflect.Descriptor instead.
func (*CommentRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_files_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommentRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type CommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Post    *post_pb.PostResponse       `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	Upvotes []*upvote_pb.UpvoteResponse `protobuf:"bytes,4,rep,name=upvotes,proto3" json:"upvotes,omitempty"`
}

func (x *CommentResponse) Reset() {
	*x = CommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_files_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentResponse) ProtoMessage() {}

func (x *CommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_files_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentResponse.ProtoReflect.Descriptor instead.
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_files_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommentResponse) GetPost() *post_pb.PostResponse {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *CommentResponse) GetUpvotes() []*upvote_pb.UpvoteResponse {
	if x != nil {
		return x.Upvotes
	}
	return nil
}

var File_grpc_proto_files_comment_proto protoreflect.FileDescriptor

var file_grpc_proto_files_comment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x70, 0x76, 0x6f, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x47,
	0x52, 0x50, 0x43, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x32, 0xb2,
	0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x52,
	0x50, 0x43, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x4c,
	0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x17, 0x2e, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x52, 0x50, 0x43, 0x2e,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x4c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x0d,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e,
	0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x33, 0x5a, 0x31, 0x72, 0x65, 0x6c, 0x65, 0x6d, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x5f, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_files_comment_proto_rawDescOnce sync.Once
	file_grpc_proto_files_comment_proto_rawDescData = file_grpc_proto_files_comment_proto_rawDesc
)

func file_grpc_proto_files_comment_proto_rawDescGZIP() []byte {
	file_grpc_proto_files_comment_proto_rawDescOnce.Do(func() {
		file_grpc_proto_files_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_files_comment_proto_rawDescData)
	})
	return file_grpc_proto_files_comment_proto_rawDescData
}

var file_grpc_proto_files_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_proto_files_comment_proto_goTypes = []interface{}{
	(*IdRequest)(nil),                // 0: LearningGRPC.IdRequest
	(*CommentRequest)(nil),           // 1: LearningGRPC.CommentRequest
	(*CommentResponse)(nil),          // 2: LearningGRPC.CommentResponse
	(*post_pb.PostResponse)(nil),     // 3: LearningGRPC.PostResponse
	(*upvote_pb.UpvoteResponse)(nil), // 4: LearningGRPC.UpvoteResponse
	(*empty.Empty)(nil),              // 5: google.protobuf.Empty
}
var file_grpc_proto_files_comment_proto_depIdxs = []int32{
	3, // 0: LearningGRPC.CommentResponse.post:type_name -> LearningGRPC.PostResponse
	4, // 1: LearningGRPC.CommentResponse.upvotes:type_name -> LearningGRPC.UpvoteResponse
	0, // 2: LearningGRPC.CommentService.GetCommentById:input_type -> LearningGRPC.IdRequest
	0, // 3: LearningGRPC.CommentService.DeleteCommentById:input_type -> LearningGRPC.IdRequest
	5, // 4: LearningGRPC.CommentService.GetAllComments:input_type -> google.protobuf.Empty
	1, // 5: LearningGRPC.CommentService.InsertComment:input_type -> LearningGRPC.CommentRequest
	2, // 6: LearningGRPC.CommentService.GetCommentById:output_type -> LearningGRPC.CommentResponse
	5, // 7: LearningGRPC.CommentService.DeleteCommentById:output_type -> google.protobuf.Empty
	2, // 8: LearningGRPC.CommentService.GetAllComments:output_type -> LearningGRPC.CommentResponse
	5, // 9: LearningGRPC.CommentService.InsertComment:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_proto_files_comment_proto_init() }
func file_grpc_proto_files_comment_proto_init() {
	if File_grpc_proto_files_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_files_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_grpc_proto_files_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentRequest); i {
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
		file_grpc_proto_files_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentResponse); i {
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
			RawDescriptor: file_grpc_proto_files_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_files_comment_proto_goTypes,
		DependencyIndexes: file_grpc_proto_files_comment_proto_depIdxs,
		MessageInfos:      file_grpc_proto_files_comment_proto_msgTypes,
	}.Build()
	File_grpc_proto_files_comment_proto = out.File
	file_grpc_proto_files_comment_proto_rawDesc = nil
	file_grpc_proto_files_comment_proto_goTypes = nil
	file_grpc_proto_files_comment_proto_depIdxs = nil
}

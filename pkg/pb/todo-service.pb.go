// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: todo-service.proto

package go_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_todo_service_proto protoreflect.FileDescriptor

var file_todo_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x12, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe3, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x18, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x1b,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x75, 0x6f, 0x6e, 0x67, 0x6e,
	0x64, 0x39, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_todo_service_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),   // 0: TodoService.CreateRequest
	(*ReadRequest)(nil),     // 1: TodoService.ReadRequest
	(*UpdateRequest)(nil),   // 2: TodoService.UpdateRequest
	(*DeleteRequest)(nil),   // 3: TodoService.DeleteRequest
	(*ReadAllRequest)(nil),  // 4: TodoService.ReadAllRequest
	(*CreateResponse)(nil),  // 5: TodoService.CreateResponse
	(*ReadResponse)(nil),    // 6: TodoService.ReadResponse
	(*UpdateResponse)(nil),  // 7: TodoService.UpdateResponse
	(*DeleteResponse)(nil),  // 8: TodoService.DeleteResponse
	(*ReadAllResponse)(nil), // 9: TodoService.ReadAllResponse
}
var file_todo_service_proto_depIdxs = []int32{
	0, // 0: TodoService.ToDoService.Create:input_type -> TodoService.CreateRequest
	1, // 1: TodoService.ToDoService.Read:input_type -> TodoService.ReadRequest
	2, // 2: TodoService.ToDoService.Update:input_type -> TodoService.UpdateRequest
	3, // 3: TodoService.ToDoService.Delete:input_type -> TodoService.DeleteRequest
	4, // 4: TodoService.ToDoService.ReadAll:input_type -> TodoService.ReadAllRequest
	5, // 5: TodoService.ToDoService.Create:output_type -> TodoService.CreateResponse
	6, // 6: TodoService.ToDoService.Read:output_type -> TodoService.ReadResponse
	7, // 7: TodoService.ToDoService.Update:output_type -> TodoService.UpdateResponse
	8, // 8: TodoService.ToDoService.Delete:output_type -> TodoService.DeleteResponse
	9, // 9: TodoService.ToDoService.ReadAll:output_type -> TodoService.ReadAllResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_todo_service_proto_init() }
func file_todo_service_proto_init() {
	if File_todo_service_proto != nil {
		return
	}
	file_todo_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_todo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_service_proto_goTypes,
		DependencyIndexes: file_todo_service_proto_depIdxs,
	}.Build()
	File_todo_service_proto = out.File
	file_todo_service_proto_rawDesc = nil
	file_todo_service_proto_goTypes = nil
	file_todo_service_proto_depIdxs = nil
}

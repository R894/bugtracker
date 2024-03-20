// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.26.0--rc3
// source: ports/bug.proto

package ports

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

type SaveBugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ProjectId   string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *SaveBugRequest) Reset() {
	*x = SaveBugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBugRequest) ProtoMessage() {}

func (x *SaveBugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBugRequest.ProtoReflect.Descriptor instead.
func (*SaveBugRequest) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{0}
}

func (x *SaveBugRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SaveBugRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SaveBugRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetBugByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BugId string `protobuf:"bytes,1,opt,name=bug_id,json=bugId,proto3" json:"bug_id,omitempty"`
}

func (x *GetBugByIDRequest) Reset() {
	*x = GetBugByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBugByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBugByIDRequest) ProtoMessage() {}

func (x *GetBugByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBugByIDRequest.ProtoReflect.Descriptor instead.
func (*GetBugByIDRequest) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{1}
}

func (x *GetBugByIDRequest) GetBugId() string {
	if x != nil {
		return x.BugId
	}
	return ""
}

type GetBugsByProjectIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetBugsByProjectIDRequest) Reset() {
	*x = GetBugsByProjectIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBugsByProjectIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBugsByProjectIDRequest) ProtoMessage() {}

func (x *GetBugsByProjectIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBugsByProjectIDRequest.ProtoReflect.Descriptor instead.
func (*GetBugsByProjectIDRequest) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{2}
}

func (x *GetBugsByProjectIDRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetBugsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBugsRequest) Reset() {
	*x = GetBugsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBugsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBugsRequest) ProtoMessage() {}

func (x *GetBugsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBugsRequest.ProtoReflect.Descriptor instead.
func (*GetBugsRequest) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{3}
}

type AssignBugToRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BugId    string `protobuf:"bytes,1,opt,name=bug_id,json=bugId,proto3" json:"bug_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AssignBugToRequest) Reset() {
	*x = AssignBugToRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignBugToRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignBugToRequest) ProtoMessage() {}

func (x *AssignBugToRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignBugToRequest.ProtoReflect.Descriptor instead.
func (*AssignBugToRequest) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{4}
}

func (x *AssignBugToRequest) GetBugId() string {
	if x != nil {
		return x.BugId
	}
	return ""
}

func (x *AssignBugToRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UpdateBugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BugId string `protobuf:"bytes,1,opt,name=bug_id,json=bugId,proto3" json:"bug_id,omitempty"`
}

func (x *UpdateBugRequest) Reset() {
	*x = UpdateBugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBugRequest) ProtoMessage() {}

func (x *UpdateBugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBugRequest.ProtoReflect.Descriptor instead.
func (*UpdateBugRequest) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBugRequest) GetBugId() string {
	if x != nil {
		return x.BugId
	}
	return ""
}

type DeleteBugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BugId string `protobuf:"bytes,1,opt,name=bug_id,json=bugId,proto3" json:"bug_id,omitempty"`
}

func (x *DeleteBugRequest) Reset() {
	*x = DeleteBugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBugRequest) ProtoMessage() {}

func (x *DeleteBugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBugRequest.ProtoReflect.Descriptor instead.
func (*DeleteBugRequest) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBugRequest) GetBugId() string {
	if x != nil {
		return x.BugId
	}
	return ""
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{7}
}

type BugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BugId       string                 `protobuf:"bytes,1,opt,name=bug_id,json=bugId,proto3" json:"bug_id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Priority    string                 `protobuf:"bytes,5,opt,name=priority,proto3" json:"priority,omitempty"`
	Assignee    string                 `protobuf:"bytes,6,opt,name=assignee,proto3" json:"assignee,omitempty"`
	ProjectId   string                 `protobuf:"bytes,7,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *BugResponse) Reset() {
	*x = BugResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BugResponse) ProtoMessage() {}

func (x *BugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BugResponse.ProtoReflect.Descriptor instead.
func (*BugResponse) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{8}
}

func (x *BugResponse) GetBugId() string {
	if x != nil {
		return x.BugId
	}
	return ""
}

func (x *BugResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BugResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BugResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BugResponse) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *BugResponse) GetAssignee() string {
	if x != nil {
		return x.Assignee
	}
	return ""
}

func (x *BugResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BugResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BugResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetBugsByProjectIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bugs []*BugResponse `protobuf:"bytes,1,rep,name=bugs,proto3" json:"bugs,omitempty"`
}

func (x *GetBugsByProjectIDResponse) Reset() {
	*x = GetBugsByProjectIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBugsByProjectIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBugsByProjectIDResponse) ProtoMessage() {}

func (x *GetBugsByProjectIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBugsByProjectIDResponse.ProtoReflect.Descriptor instead.
func (*GetBugsByProjectIDResponse) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{9}
}

func (x *GetBugsByProjectIDResponse) GetBugs() []*BugResponse {
	if x != nil {
		return x.Bugs
	}
	return nil
}

type GetBugsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bugs []*BugResponse `protobuf:"bytes,1,rep,name=bugs,proto3" json:"bugs,omitempty"`
}

func (x *GetBugsResponse) Reset() {
	*x = GetBugsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_bug_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBugsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBugsResponse) ProtoMessage() {}

func (x *GetBugsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ports_bug_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBugsResponse.ProtoReflect.Descriptor instead.
func (*GetBugsResponse) Descriptor() ([]byte, []int) {
	return file_ports_bug_proto_rawDescGZIP(), []int{10}
}

func (x *GetBugsResponse) GetBugs() []*BugResponse {
	if x != nil {
		return x.Bugs
	}
	return nil
}

var File_ports_bug_proto protoreflect.FileDescriptor

var file_ports_bug_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x62, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x62, 0x75, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x42,
	0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x2a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x67, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x12, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x42, 0x75, 0x67, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x62, 0x75, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x75, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x67, 0x49, 0x64, 0x22, 0x29,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x67, 0x49, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc1, 0x02, 0x0a, 0x0b, 0x42,
	0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x75,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x67, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x42,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x62, 0x75, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x75, 0x67,
	0x2e, 0x42, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x75,
	0x67, 0x73, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x75, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x42, 0x75, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x75, 0x67, 0x73, 0x32, 0xc9, 0x03, 0x0a, 0x14,
	0x42, 0x75, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x61, 0x76, 0x65, 0x42, 0x75, 0x67, 0x12,
	0x13, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x42, 0x75, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x67, 0x42, 0x79, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x42, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x42,
	0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x62, 0x75, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x75, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x62, 0x75, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x42,
	0x75, 0x67, 0x54, 0x6f, 0x12, 0x17, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x42, 0x75, 0x67, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x62, 0x75, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x67,
	0x12, 0x15, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x67, 0x12, 0x15, 0x2e, 0x62, 0x75, 0x67,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x62, 0x75, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x38, 0x39, 0x34, 0x2f, 0x62, 0x75, 0x67, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x75, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ports_bug_proto_rawDescOnce sync.Once
	file_ports_bug_proto_rawDescData = file_ports_bug_proto_rawDesc
)

func file_ports_bug_proto_rawDescGZIP() []byte {
	file_ports_bug_proto_rawDescOnce.Do(func() {
		file_ports_bug_proto_rawDescData = protoimpl.X.CompressGZIP(file_ports_bug_proto_rawDescData)
	})
	return file_ports_bug_proto_rawDescData
}

var file_ports_bug_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_ports_bug_proto_goTypes = []interface{}{
	(*SaveBugRequest)(nil),             // 0: bug.SaveBugRequest
	(*GetBugByIDRequest)(nil),          // 1: bug.GetBugByIDRequest
	(*GetBugsByProjectIDRequest)(nil),  // 2: bug.GetBugsByProjectIDRequest
	(*GetBugsRequest)(nil),             // 3: bug.GetBugsRequest
	(*AssignBugToRequest)(nil),         // 4: bug.AssignBugToRequest
	(*UpdateBugRequest)(nil),           // 5: bug.UpdateBugRequest
	(*DeleteBugRequest)(nil),           // 6: bug.DeleteBugRequest
	(*EmptyResponse)(nil),              // 7: bug.EmptyResponse
	(*BugResponse)(nil),                // 8: bug.BugResponse
	(*GetBugsByProjectIDResponse)(nil), // 9: bug.GetBugsByProjectIDResponse
	(*GetBugsResponse)(nil),            // 10: bug.GetBugsResponse
	(*timestamppb.Timestamp)(nil),      // 11: google.protobuf.Timestamp
}
var file_ports_bug_proto_depIdxs = []int32{
	11, // 0: bug.BugResponse.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: bug.BugResponse.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 2: bug.GetBugsByProjectIDResponse.bugs:type_name -> bug.BugResponse
	8,  // 3: bug.GetBugsResponse.bugs:type_name -> bug.BugResponse
	0,  // 4: bug.BugRepositoryService.SaveBug:input_type -> bug.SaveBugRequest
	1,  // 5: bug.BugRepositoryService.GetBugByID:input_type -> bug.GetBugByIDRequest
	2,  // 6: bug.BugRepositoryService.GetBugsByProjectID:input_type -> bug.GetBugsByProjectIDRequest
	3,  // 7: bug.BugRepositoryService.GetBugs:input_type -> bug.GetBugsRequest
	4,  // 8: bug.BugRepositoryService.AssignBugTo:input_type -> bug.AssignBugToRequest
	5,  // 9: bug.BugRepositoryService.UpdateBug:input_type -> bug.UpdateBugRequest
	6,  // 10: bug.BugRepositoryService.DeleteBug:input_type -> bug.DeleteBugRequest
	7,  // 11: bug.BugRepositoryService.SaveBug:output_type -> bug.EmptyResponse
	8,  // 12: bug.BugRepositoryService.GetBugByID:output_type -> bug.BugResponse
	9,  // 13: bug.BugRepositoryService.GetBugsByProjectID:output_type -> bug.GetBugsByProjectIDResponse
	10, // 14: bug.BugRepositoryService.GetBugs:output_type -> bug.GetBugsResponse
	7,  // 15: bug.BugRepositoryService.AssignBugTo:output_type -> bug.EmptyResponse
	7,  // 16: bug.BugRepositoryService.UpdateBug:output_type -> bug.EmptyResponse
	7,  // 17: bug.BugRepositoryService.DeleteBug:output_type -> bug.EmptyResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_ports_bug_proto_init() }
func file_ports_bug_proto_init() {
	if File_ports_bug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ports_bug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveBugRequest); i {
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
		file_ports_bug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBugByIDRequest); i {
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
		file_ports_bug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBugsByProjectIDRequest); i {
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
		file_ports_bug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBugsRequest); i {
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
		file_ports_bug_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignBugToRequest); i {
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
		file_ports_bug_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBugRequest); i {
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
		file_ports_bug_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBugRequest); i {
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
		file_ports_bug_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_ports_bug_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BugResponse); i {
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
		file_ports_bug_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBugsByProjectIDResponse); i {
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
		file_ports_bug_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBugsResponse); i {
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
			RawDescriptor: file_ports_bug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ports_bug_proto_goTypes,
		DependencyIndexes: file_ports_bug_proto_depIdxs,
		MessageInfos:      file_ports_bug_proto_msgTypes,
	}.Build()
	File_ports_bug_proto = out.File
	file_ports_bug_proto_rawDesc = nil
	file_ports_bug_proto_goTypes = nil
	file_ports_bug_proto_depIdxs = nil
}

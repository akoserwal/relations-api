// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: relations/v0/lookup.proto

package v0

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LookupResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource        *ObjectReference   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Relation        string             `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	SubjectType     *ObjectType        `protobuf:"bytes,3,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`
	SubjectRelation *string            `protobuf:"bytes,4,opt,name=subject_relation,json=subjectRelation,proto3,oneof" json:"subject_relation,omitempty"`
	Pagination      *RequestPagination `protobuf:"bytes,5,opt,name=pagination,proto3,oneof" json:"pagination,omitempty"`
}

func (x *LookupResourcesRequest) Reset() {
	*x = LookupResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relations_v0_lookup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupResourcesRequest) ProtoMessage() {}

func (x *LookupResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relations_v0_lookup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupResourcesRequest.ProtoReflect.Descriptor instead.
func (*LookupResourcesRequest) Descriptor() ([]byte, []int) {
	return file_relations_v0_lookup_proto_rawDescGZIP(), []int{0}
}

func (x *LookupResourcesRequest) GetResource() *ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *LookupResourcesRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *LookupResourcesRequest) GetSubjectType() *ObjectType {
	if x != nil {
		return x.SubjectType
	}
	return nil
}

func (x *LookupResourcesRequest) GetSubjectRelation() string {
	if x != nil && x.SubjectRelation != nil {
		return *x.SubjectRelation
	}
	return ""
}

func (x *LookupResourcesRequest) GetPagination() *RequestPagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type LookupResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource   *ObjectReference    `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Pagination *ResponsePagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *LookupResourcesResponse) Reset() {
	*x = LookupResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relations_v0_lookup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupResourcesResponse) ProtoMessage() {}

func (x *LookupResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relations_v0_lookup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupResourcesResponse.ProtoReflect.Descriptor instead.
func (*LookupResourcesResponse) Descriptor() ([]byte, []int) {
	return file_relations_v0_lookup_proto_rawDescGZIP(), []int{1}
}

func (x *LookupResourcesResponse) GetResource() *ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *LookupResourcesResponse) GetPagination() *ResponsePagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type LookupSubjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource        *ObjectReference   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Relation        string             `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	SubjectType     *ObjectType        `protobuf:"bytes,3,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`
	SubjectRelation *string            `protobuf:"bytes,4,opt,name=subject_relation,json=subjectRelation,proto3,oneof" json:"subject_relation,omitempty"`
	Pagination      *RequestPagination `protobuf:"bytes,5,opt,name=pagination,proto3,oneof" json:"pagination,omitempty"`
}

func (x *LookupSubjectsRequest) Reset() {
	*x = LookupSubjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relations_v0_lookup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupSubjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupSubjectsRequest) ProtoMessage() {}

func (x *LookupSubjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relations_v0_lookup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupSubjectsRequest.ProtoReflect.Descriptor instead.
func (*LookupSubjectsRequest) Descriptor() ([]byte, []int) {
	return file_relations_v0_lookup_proto_rawDescGZIP(), []int{2}
}

func (x *LookupSubjectsRequest) GetResource() *ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *LookupSubjectsRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *LookupSubjectsRequest) GetSubjectType() *ObjectType {
	if x != nil {
		return x.SubjectType
	}
	return nil
}

func (x *LookupSubjectsRequest) GetSubjectRelation() string {
	if x != nil && x.SubjectRelation != nil {
		return *x.SubjectRelation
	}
	return ""
}

func (x *LookupSubjectsRequest) GetPagination() *RequestPagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type LookupSubjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject    *SubjectReference   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Pagination *ResponsePagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *LookupSubjectsResponse) Reset() {
	*x = LookupSubjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relations_v0_lookup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupSubjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupSubjectsResponse) ProtoMessage() {}

func (x *LookupSubjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relations_v0_lookup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupSubjectsResponse.ProtoReflect.Descriptor instead.
func (*LookupSubjectsResponse) Descriptor() ([]byte, []int) {
	return file_relations_v0_lookup_proto_rawDescGZIP(), []int{3}
}

func (x *LookupSubjectsResponse) GetSubject() *SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *LookupSubjectsResponse) GetPagination() *ResponsePagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_relations_v0_lookup_proto protoreflect.FileDescriptor

var file_relations_v0_lookup_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x6c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf8, 0x02, 0x0a, 0x16, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c,
	0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x10,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa4, 0x01,
	0x0a, 0x17, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x30, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf7, 0x02, 0x0a, 0x15, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4c, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a,
	0x10, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa2,
	0x01, 0x0a, 0x16, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30,
	0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0xa1, 0x02, 0x0a, 0x13, 0x4b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0e,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x2a,
	0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30,
	0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x30, 0x01, 0x12,
	0x85, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x30, 0x01, 0x42, 0x61, 0x0a, 0x23, 0x6f, 0x72, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x30, 0x50, 0x01,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_relations_v0_lookup_proto_rawDescOnce sync.Once
	file_relations_v0_lookup_proto_rawDescData = file_relations_v0_lookup_proto_rawDesc
)

func file_relations_v0_lookup_proto_rawDescGZIP() []byte {
	file_relations_v0_lookup_proto_rawDescOnce.Do(func() {
		file_relations_v0_lookup_proto_rawDescData = protoimpl.X.CompressGZIP(file_relations_v0_lookup_proto_rawDescData)
	})
	return file_relations_v0_lookup_proto_rawDescData
}

var file_relations_v0_lookup_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_relations_v0_lookup_proto_goTypes = []any{
	(*LookupResourcesRequest)(nil),  // 0: kessel.relations.v0.LookupResourcesRequest
	(*LookupResourcesResponse)(nil), // 1: kessel.relations.v0.LookupResourcesResponse
	(*LookupSubjectsRequest)(nil),   // 2: kessel.relations.v0.LookupSubjectsRequest
	(*LookupSubjectsResponse)(nil),  // 3: kessel.relations.v0.LookupSubjectsResponse
	(*ObjectReference)(nil),         // 4: kessel.relations.v0.ObjectReference
	(*ObjectType)(nil),              // 5: kessel.relations.v0.ObjectType
	(*RequestPagination)(nil),       // 6: kessel.relations.v0.RequestPagination
	(*ResponsePagination)(nil),      // 7: kessel.relations.v0.ResponsePagination
	(*SubjectReference)(nil),        // 8: kessel.relations.v0.SubjectReference
}
var file_relations_v0_lookup_proto_depIdxs = []int32{
	4,  // 0: kessel.relations.v0.LookupResourcesRequest.resource:type_name -> kessel.relations.v0.ObjectReference
	5,  // 1: kessel.relations.v0.LookupResourcesRequest.subject_type:type_name -> kessel.relations.v0.ObjectType
	6,  // 2: kessel.relations.v0.LookupResourcesRequest.pagination:type_name -> kessel.relations.v0.RequestPagination
	4,  // 3: kessel.relations.v0.LookupResourcesResponse.resource:type_name -> kessel.relations.v0.ObjectReference
	7,  // 4: kessel.relations.v0.LookupResourcesResponse.pagination:type_name -> kessel.relations.v0.ResponsePagination
	4,  // 5: kessel.relations.v0.LookupSubjectsRequest.resource:type_name -> kessel.relations.v0.ObjectReference
	5,  // 6: kessel.relations.v0.LookupSubjectsRequest.subject_type:type_name -> kessel.relations.v0.ObjectType
	6,  // 7: kessel.relations.v0.LookupSubjectsRequest.pagination:type_name -> kessel.relations.v0.RequestPagination
	8,  // 8: kessel.relations.v0.LookupSubjectsResponse.subject:type_name -> kessel.relations.v0.SubjectReference
	7,  // 9: kessel.relations.v0.LookupSubjectsResponse.pagination:type_name -> kessel.relations.v0.ResponsePagination
	2,  // 10: kessel.relations.v0.KesselLookupService.LookupSubjects:input_type -> kessel.relations.v0.LookupSubjectsRequest
	0,  // 11: kessel.relations.v0.KesselLookupService.LookupResources:input_type -> kessel.relations.v0.LookupResourcesRequest
	3,  // 12: kessel.relations.v0.KesselLookupService.LookupSubjects:output_type -> kessel.relations.v0.LookupSubjectsResponse
	1,  // 13: kessel.relations.v0.KesselLookupService.LookupResources:output_type -> kessel.relations.v0.LookupResourcesResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_relations_v0_lookup_proto_init() }
func file_relations_v0_lookup_proto_init() {
	if File_relations_v0_lookup_proto != nil {
		return
	}
	file_relations_v0_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relations_v0_lookup_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LookupResourcesRequest); i {
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
		file_relations_v0_lookup_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LookupResourcesResponse); i {
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
		file_relations_v0_lookup_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LookupSubjectsRequest); i {
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
		file_relations_v0_lookup_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LookupSubjectsResponse); i {
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
	file_relations_v0_lookup_proto_msgTypes[0].OneofWrappers = []any{}
	file_relations_v0_lookup_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relations_v0_lookup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relations_v0_lookup_proto_goTypes,
		DependencyIndexes: file_relations_v0_lookup_proto_depIdxs,
		MessageInfos:      file_relations_v0_lookup_proto_msgTypes,
	}.Build()
	File_relations_v0_lookup_proto = out.File
	file_relations_v0_lookup_proto_rawDesc = nil
	file_relations_v0_lookup_proto_goTypes = nil
	file_relations_v0_lookup_proto_depIdxs = nil
}

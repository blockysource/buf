// Copyright 2020-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: buf/alpha/registry/v1alpha1/reference.proto

package registryv1alpha1

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

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reference:
	//
	//	*Reference_Branch
	//	*Reference_Tag
	//	*Reference_Commit
	//	*Reference_Main
	//	*Reference_Draft
	Reference isReference_Reference `protobuf_oneof:"reference"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{0}
}

func (m *Reference) GetReference() isReference_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (x *Reference) GetBranch() *RepositoryBranch {
	if x, ok := x.GetReference().(*Reference_Branch); ok {
		return x.Branch
	}
	return nil
}

func (x *Reference) GetTag() *RepositoryTag {
	if x, ok := x.GetReference().(*Reference_Tag); ok {
		return x.Tag
	}
	return nil
}

func (x *Reference) GetCommit() *RepositoryCommit {
	if x, ok := x.GetReference().(*Reference_Commit); ok {
		return x.Commit
	}
	return nil
}

func (x *Reference) GetMain() *RepositoryMainReference {
	if x, ok := x.GetReference().(*Reference_Main); ok {
		return x.Main
	}
	return nil
}

func (x *Reference) GetDraft() *RepositoryDraft {
	if x, ok := x.GetReference().(*Reference_Draft); ok {
		return x.Draft
	}
	return nil
}

type isReference_Reference interface {
	isReference_Reference()
}

type Reference_Branch struct {
	// The requested reference is a branch.
	Branch *RepositoryBranch `protobuf:"bytes,1,opt,name=branch,proto3,oneof"`
}

type Reference_Tag struct {
	// The requested reference is a tag.
	Tag *RepositoryTag `protobuf:"bytes,2,opt,name=tag,proto3,oneof"`
}

type Reference_Commit struct {
	// The requested reference is a commit.
	Commit *RepositoryCommit `protobuf:"bytes,3,opt,name=commit,proto3,oneof"`
}

type Reference_Main struct {
	// The requested reference is the default reference.
	Main *RepositoryMainReference `protobuf:"bytes,5,opt,name=main,proto3,oneof"`
}

type Reference_Draft struct {
	// The requested reference is a draft commit.
	Draft *RepositoryDraft `protobuf:"bytes,6,opt,name=draft,proto3,oneof"`
}

func (*Reference_Branch) isReference_Reference() {}

func (*Reference_Tag) isReference_Reference() {}

func (*Reference_Commit) isReference_Reference() {}

func (*Reference_Main) isReference_Reference() {}

func (*Reference_Draft) isReference_Reference() {}

type RepositoryMainReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is always 'main'.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The latest commit in this repository. If the repository has no commits,
	// this will be empty.
	Commit *RepositoryCommit `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *RepositoryMainReference) Reset() {
	*x = RepositoryMainReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryMainReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryMainReference) ProtoMessage() {}

func (x *RepositoryMainReference) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryMainReference.ProtoReflect.Descriptor instead.
func (*RepositoryMainReference) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{1}
}

func (x *RepositoryMainReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepositoryMainReference) GetCommit() *RepositoryCommit {
	if x != nil {
		return x.Commit
	}
	return nil
}

type RepositoryDraft struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the draft
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The commit this draft points to.
	Commit *RepositoryCommit `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *RepositoryDraft) Reset() {
	*x = RepositoryDraft{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryDraft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryDraft) ProtoMessage() {}

func (x *RepositoryDraft) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryDraft.ProtoReflect.Descriptor instead.
func (*RepositoryDraft) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{2}
}

func (x *RepositoryDraft) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepositoryDraft) GetCommit() *RepositoryCommit {
	if x != nil {
		return x.Commit
	}
	return nil
}

type GetReferenceByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the requested reference.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Owner of the repository the reference belongs to.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// Name of the repository the reference belongs to.
	RepositoryName string `protobuf:"bytes,3,opt,name=repository_name,json=repositoryName,proto3" json:"repository_name,omitempty"`
}

func (x *GetReferenceByNameRequest) Reset() {
	*x = GetReferenceByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReferenceByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReferenceByNameRequest) ProtoMessage() {}

func (x *GetReferenceByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReferenceByNameRequest.ProtoReflect.Descriptor instead.
func (*GetReferenceByNameRequest) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{3}
}

func (x *GetReferenceByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetReferenceByNameRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetReferenceByNameRequest) GetRepositoryName() string {
	if x != nil {
		return x.RepositoryName
	}
	return ""
}

type GetReferenceByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference *Reference `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
}

func (x *GetReferenceByNameResponse) Reset() {
	*x = GetReferenceByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReferenceByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReferenceByNameResponse) ProtoMessage() {}

func (x *GetReferenceByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReferenceByNameResponse.ProtoReflect.Descriptor instead.
func (*GetReferenceByNameResponse) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{4}
}

func (x *GetReferenceByNameResponse) GetReference() *Reference {
	if x != nil {
		return x.Reference
	}
	return nil
}

type ListGitCommitsForReferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String that represents the name of the reference.
	Reference string `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// Owner of the repository the reference belongs to.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// Name of the repository the reference belongs to.
	RepositoryName string `protobuf:"bytes,3,opt,name=repository_name,json=repositoryName,proto3" json:"repository_name,omitempty"`
	PageSize       uint32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The first page is returned if this is empty.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListGitCommitsForReferenceRequest) Reset() {
	*x = ListGitCommitsForReferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGitCommitsForReferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGitCommitsForReferenceRequest) ProtoMessage() {}

func (x *ListGitCommitsForReferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGitCommitsForReferenceRequest.ProtoReflect.Descriptor instead.
func (*ListGitCommitsForReferenceRequest) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{5}
}

func (x *ListGitCommitsForReferenceRequest) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *ListGitCommitsForReferenceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ListGitCommitsForReferenceRequest) GetRepositoryName() string {
	if x != nil {
		return x.RepositoryName
	}
	return ""
}

func (x *ListGitCommitsForReferenceRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListGitCommitsForReferenceRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListGitCommitsForReferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the BSR commit the reference resolved to.
	CommitId string `protobuf:"bytes,1,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	// List of git commits and metadata associated with the resolved reference.
	GitCommits []*GitCommitInformation `protobuf:"bytes,2,rep,name=git_commits,json=gitCommits,proto3" json:"git_commits,omitempty"`
	// There are no more pages if this is empty.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListGitCommitsForReferenceResponse) Reset() {
	*x = ListGitCommitsForReferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGitCommitsForReferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGitCommitsForReferenceResponse) ProtoMessage() {}

func (x *ListGitCommitsForReferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGitCommitsForReferenceResponse.ProtoReflect.Descriptor instead.
func (*ListGitCommitsForReferenceResponse) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{6}
}

func (x *ListGitCommitsForReferenceResponse) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *ListGitCommitsForReferenceResponse) GetGitCommits() []*GitCommitInformation {
	if x != nil {
		return x.GitCommits
	}
	return nil
}

func (x *ListGitCommitsForReferenceResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_buf_alpha_registry_v1alpha1_reference_proto protoreflect.FileDescriptor

var file_buf_alpha_registry_v1alpha1_reference_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62,
	0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2e, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x61, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x03, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x48, 0x00, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x3e, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x54, 0x61, 0x67, 0x48, 0x00, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x47, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x48, 0x00, 0x52, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x4a, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x44, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x72, 0x61, 0x66, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x52, 0x05, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x22, 0x74, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x4d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x45, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x6c, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x45, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x6e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x21, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x46, 0x6f, 0x72,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xbd, 0x01, 0x0a, 0x22, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x52, 0x0a,
	0x0b, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc4, 0x02, 0x0a, 0x10, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0xa2, 0x01, 0x0a, 0x1a,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x46, 0x6f,
	0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3e, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01,
	0x42, 0x9b, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x52, 0xaa, 0x02, 0x1b, 0x42, 0x75, 0x66, 0x2e, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68,
	0x61, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x27, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e,
	0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_registry_v1alpha1_reference_proto_rawDescOnce sync.Once
	file_buf_alpha_registry_v1alpha1_reference_proto_rawDescData = file_buf_alpha_registry_v1alpha1_reference_proto_rawDesc
)

func file_buf_alpha_registry_v1alpha1_reference_proto_rawDescGZIP() []byte {
	file_buf_alpha_registry_v1alpha1_reference_proto_rawDescOnce.Do(func() {
		file_buf_alpha_registry_v1alpha1_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_registry_v1alpha1_reference_proto_rawDescData)
	})
	return file_buf_alpha_registry_v1alpha1_reference_proto_rawDescData
}

var file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_buf_alpha_registry_v1alpha1_reference_proto_goTypes = []interface{}{
	(*Reference)(nil),                          // 0: buf.alpha.registry.v1alpha1.Reference
	(*RepositoryMainReference)(nil),            // 1: buf.alpha.registry.v1alpha1.RepositoryMainReference
	(*RepositoryDraft)(nil),                    // 2: buf.alpha.registry.v1alpha1.RepositoryDraft
	(*GetReferenceByNameRequest)(nil),          // 3: buf.alpha.registry.v1alpha1.GetReferenceByNameRequest
	(*GetReferenceByNameResponse)(nil),         // 4: buf.alpha.registry.v1alpha1.GetReferenceByNameResponse
	(*ListGitCommitsForReferenceRequest)(nil),  // 5: buf.alpha.registry.v1alpha1.ListGitCommitsForReferenceRequest
	(*ListGitCommitsForReferenceResponse)(nil), // 6: buf.alpha.registry.v1alpha1.ListGitCommitsForReferenceResponse
	(*RepositoryBranch)(nil),                   // 7: buf.alpha.registry.v1alpha1.RepositoryBranch
	(*RepositoryTag)(nil),                      // 8: buf.alpha.registry.v1alpha1.RepositoryTag
	(*RepositoryCommit)(nil),                   // 9: buf.alpha.registry.v1alpha1.RepositoryCommit
	(*GitCommitInformation)(nil),               // 10: buf.alpha.registry.v1alpha1.GitCommitInformation
}
var file_buf_alpha_registry_v1alpha1_reference_proto_depIdxs = []int32{
	7,  // 0: buf.alpha.registry.v1alpha1.Reference.branch:type_name -> buf.alpha.registry.v1alpha1.RepositoryBranch
	8,  // 1: buf.alpha.registry.v1alpha1.Reference.tag:type_name -> buf.alpha.registry.v1alpha1.RepositoryTag
	9,  // 2: buf.alpha.registry.v1alpha1.Reference.commit:type_name -> buf.alpha.registry.v1alpha1.RepositoryCommit
	1,  // 3: buf.alpha.registry.v1alpha1.Reference.main:type_name -> buf.alpha.registry.v1alpha1.RepositoryMainReference
	2,  // 4: buf.alpha.registry.v1alpha1.Reference.draft:type_name -> buf.alpha.registry.v1alpha1.RepositoryDraft
	9,  // 5: buf.alpha.registry.v1alpha1.RepositoryMainReference.commit:type_name -> buf.alpha.registry.v1alpha1.RepositoryCommit
	9,  // 6: buf.alpha.registry.v1alpha1.RepositoryDraft.commit:type_name -> buf.alpha.registry.v1alpha1.RepositoryCommit
	0,  // 7: buf.alpha.registry.v1alpha1.GetReferenceByNameResponse.reference:type_name -> buf.alpha.registry.v1alpha1.Reference
	10, // 8: buf.alpha.registry.v1alpha1.ListGitCommitsForReferenceResponse.git_commits:type_name -> buf.alpha.registry.v1alpha1.GitCommitInformation
	3,  // 9: buf.alpha.registry.v1alpha1.ReferenceService.GetReferenceByName:input_type -> buf.alpha.registry.v1alpha1.GetReferenceByNameRequest
	5,  // 10: buf.alpha.registry.v1alpha1.ReferenceService.ListGitCommitsForReference:input_type -> buf.alpha.registry.v1alpha1.ListGitCommitsForReferenceRequest
	4,  // 11: buf.alpha.registry.v1alpha1.ReferenceService.GetReferenceByName:output_type -> buf.alpha.registry.v1alpha1.GetReferenceByNameResponse
	6,  // 12: buf.alpha.registry.v1alpha1.ReferenceService.ListGitCommitsForReference:output_type -> buf.alpha.registry.v1alpha1.ListGitCommitsForReferenceResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_buf_alpha_registry_v1alpha1_reference_proto_init() }
func file_buf_alpha_registry_v1alpha1_reference_proto_init() {
	if File_buf_alpha_registry_v1alpha1_reference_proto != nil {
		return
	}
	file_buf_alpha_registry_v1alpha1_git_metadata_proto_init()
	file_buf_alpha_registry_v1alpha1_repository_branch_proto_init()
	file_buf_alpha_registry_v1alpha1_repository_commit_proto_init()
	file_buf_alpha_registry_v1alpha1_repository_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryMainReference); i {
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
		file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryDraft); i {
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
		file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReferenceByNameRequest); i {
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
		file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReferenceByNameResponse); i {
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
		file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGitCommitsForReferenceRequest); i {
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
		file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGitCommitsForReferenceResponse); i {
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
	file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Reference_Branch)(nil),
		(*Reference_Tag)(nil),
		(*Reference_Commit)(nil),
		(*Reference_Main)(nil),
		(*Reference_Draft)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_registry_v1alpha1_reference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_alpha_registry_v1alpha1_reference_proto_goTypes,
		DependencyIndexes: file_buf_alpha_registry_v1alpha1_reference_proto_depIdxs,
		MessageInfos:      file_buf_alpha_registry_v1alpha1_reference_proto_msgTypes,
	}.Build()
	File_buf_alpha_registry_v1alpha1_reference_proto = out.File
	file_buf_alpha_registry_v1alpha1_reference_proto_rawDesc = nil
	file_buf_alpha_registry_v1alpha1_reference_proto_goTypes = nil
	file_buf_alpha_registry_v1alpha1_reference_proto_depIdxs = nil
}

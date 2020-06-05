// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v1/services/label_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for [LabelService.GetLabel][google.ads.googleads.v1.services.LabelService.GetLabel].
type GetLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the label to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetLabelRequest) Reset() {
	*x = GetLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLabelRequest) ProtoMessage() {}

func (x *GetLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLabelRequest.ProtoReflect.Descriptor instead.
func (*GetLabelRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_label_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetLabelRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for [LabelService.MutateLabels][google.ads.googleads.v1.services.LabelService.MutateLabels].
type MutateLabelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the customer whose labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on labels.
	Operations []*LabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateLabelsRequest) Reset() {
	*x = MutateLabelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateLabelsRequest) ProtoMessage() {}

func (x *MutateLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateLabelsRequest.ProtoReflect.Descriptor instead.
func (*MutateLabelsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_label_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateLabelsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateLabelsRequest) GetOperations() []*LabelOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateLabelsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateLabelsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, remove, update) on a label.
type LabelOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*LabelOperation_Create
	//	*LabelOperation_Update
	//	*LabelOperation_Remove
	Operation isLabelOperation_Operation `protobuf_oneof:"operation"`
}

func (x *LabelOperation) Reset() {
	*x = LabelOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelOperation) ProtoMessage() {}

func (x *LabelOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelOperation.ProtoReflect.Descriptor instead.
func (*LabelOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_label_service_proto_rawDescGZIP(), []int{2}
}

func (x *LabelOperation) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *LabelOperation) GetOperation() isLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *LabelOperation) GetCreate() *resources.Label {
	if x, ok := x.GetOperation().(*LabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *LabelOperation) GetUpdate() *resources.Label {
	if x, ok := x.GetOperation().(*LabelOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *LabelOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*LabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isLabelOperation_Operation interface {
	isLabelOperation_Operation()
}

type LabelOperation_Create struct {
	// Create operation: No resource name is expected for the new label.
	Create *resources.Label `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type LabelOperation_Update struct {
	// Update operation: The label is expected to have a valid resource name.
	Update *resources.Label `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type LabelOperation_Remove struct {
	// Remove operation: A resource name for the label being removed, in
	// this format:
	//
	// `customers/{customer_id}/labels/{label_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*LabelOperation_Create) isLabelOperation_Operation() {}

func (*LabelOperation_Update) isLabelOperation_Operation() {}

func (*LabelOperation_Remove) isLabelOperation_Operation() {}

// Response message for a labels mutate.
type MutateLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateLabelsResponse) Reset() {
	*x = MutateLabelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateLabelsResponse) ProtoMessage() {}

func (x *MutateLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateLabelsResponse.ProtoReflect.Descriptor instead.
func (*MutateLabelsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_label_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateLabelsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateLabelsResponse) GetResults() []*MutateLabelResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for a label mutate.
type MutateLabelResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateLabelResult) Reset() {
	*x = MutateLabelResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateLabelResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateLabelResult) ProtoMessage() {}

func (x *MutateLabelResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_label_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateLabelResult.ProtoReflect.Descriptor instead.
func (*MutateLabelResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_label_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateLabelResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v1_services_label_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v1_services_label_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x26, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xe0, 0x01, 0x0a, 0x13, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x55, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xfc, 0x01, 0x0a, 0x0e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x42, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xad, 0x01, 0x0a, 0x14, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4d, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0xa8, 0x03, 0x0a, 0x0c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x40, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xce,
	0x01, 0x0a, 0x0c, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x2b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3a, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x1b, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0xf8, 0x01, 0x0a,
	0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x11, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v1_services_label_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v1_services_label_service_proto_rawDescData = file_google_ads_googleads_v1_services_label_service_proto_rawDesc
)

func file_google_ads_googleads_v1_services_label_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v1_services_label_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v1_services_label_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v1_services_label_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v1_services_label_service_proto_rawDescData
}

var file_google_ads_googleads_v1_services_label_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v1_services_label_service_proto_goTypes = []interface{}{
	(*GetLabelRequest)(nil),      // 0: google.ads.googleads.v1.services.GetLabelRequest
	(*MutateLabelsRequest)(nil),  // 1: google.ads.googleads.v1.services.MutateLabelsRequest
	(*LabelOperation)(nil),       // 2: google.ads.googleads.v1.services.LabelOperation
	(*MutateLabelsResponse)(nil), // 3: google.ads.googleads.v1.services.MutateLabelsResponse
	(*MutateLabelResult)(nil),    // 4: google.ads.googleads.v1.services.MutateLabelResult
	(*field_mask.FieldMask)(nil), // 5: google.protobuf.FieldMask
	(*resources.Label)(nil),      // 6: google.ads.googleads.v1.resources.Label
	(*status.Status)(nil),        // 7: google.rpc.Status
}
var file_google_ads_googleads_v1_services_label_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v1.services.MutateLabelsRequest.operations:type_name -> google.ads.googleads.v1.services.LabelOperation
	5, // 1: google.ads.googleads.v1.services.LabelOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 2: google.ads.googleads.v1.services.LabelOperation.create:type_name -> google.ads.googleads.v1.resources.Label
	6, // 3: google.ads.googleads.v1.services.LabelOperation.update:type_name -> google.ads.googleads.v1.resources.Label
	7, // 4: google.ads.googleads.v1.services.MutateLabelsResponse.partial_failure_error:type_name -> google.rpc.Status
	4, // 5: google.ads.googleads.v1.services.MutateLabelsResponse.results:type_name -> google.ads.googleads.v1.services.MutateLabelResult
	0, // 6: google.ads.googleads.v1.services.LabelService.GetLabel:input_type -> google.ads.googleads.v1.services.GetLabelRequest
	1, // 7: google.ads.googleads.v1.services.LabelService.MutateLabels:input_type -> google.ads.googleads.v1.services.MutateLabelsRequest
	6, // 8: google.ads.googleads.v1.services.LabelService.GetLabel:output_type -> google.ads.googleads.v1.resources.Label
	3, // 9: google.ads.googleads.v1.services.LabelService.MutateLabels:output_type -> google.ads.googleads.v1.services.MutateLabelsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v1_services_label_service_proto_init() }
func file_google_ads_googleads_v1_services_label_service_proto_init() {
	if File_google_ads_googleads_v1_services_label_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v1_services_label_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLabelRequest); i {
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
		file_google_ads_googleads_v1_services_label_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateLabelsRequest); i {
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
		file_google_ads_googleads_v1_services_label_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelOperation); i {
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
		file_google_ads_googleads_v1_services_label_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateLabelsResponse); i {
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
		file_google_ads_googleads_v1_services_label_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateLabelResult); i {
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
	file_google_ads_googleads_v1_services_label_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*LabelOperation_Create)(nil),
		(*LabelOperation_Update)(nil),
		(*LabelOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v1_services_label_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v1_services_label_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v1_services_label_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v1_services_label_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v1_services_label_service_proto = out.File
	file_google_ads_googleads_v1_services_label_service_proto_rawDesc = nil
	file_google_ads_googleads_v1_services_label_service_proto_goTypes = nil
	file_google_ads_googleads_v1_services_label_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LabelServiceClient is the client API for LabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LabelServiceClient interface {
	// Returns the requested label in full detail.
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*resources.Label, error)
	// Creates, updates, or removes labels. Operation statuses are returned.
	MutateLabels(ctx context.Context, in *MutateLabelsRequest, opts ...grpc.CallOption) (*MutateLabelsResponse, error)
}

type labelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelServiceClient(cc grpc.ClientConnInterface) LabelServiceClient {
	return &labelServiceClient{cc}
}

func (c *labelServiceClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*resources.Label, error) {
	out := new(resources.Label)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.LabelService/GetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) MutateLabels(ctx context.Context, in *MutateLabelsRequest, opts ...grpc.CallOption) (*MutateLabelsResponse, error) {
	out := new(MutateLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.LabelService/MutateLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelServiceServer is the server API for LabelService service.
type LabelServiceServer interface {
	// Returns the requested label in full detail.
	GetLabel(context.Context, *GetLabelRequest) (*resources.Label, error)
	// Creates, updates, or removes labels. Operation statuses are returned.
	MutateLabels(context.Context, *MutateLabelsRequest) (*MutateLabelsResponse, error)
}

// UnimplementedLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLabelServiceServer struct {
}

func (*UnimplementedLabelServiceServer) GetLabel(context.Context, *GetLabelRequest) (*resources.Label, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetLabel not implemented")
}
func (*UnimplementedLabelServiceServer) MutateLabels(context.Context, *MutateLabelsRequest) (*MutateLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateLabels not implemented")
}

func RegisterLabelServiceServer(s *grpc.Server, srv LabelServiceServer) {
	s.RegisterService(&_LabelService_serviceDesc, srv)
}

func _LabelService_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.LabelService/GetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_MutateLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).MutateLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.LabelService/MutateLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).MutateLabels(ctx, req.(*MutateLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.LabelService",
	HandlerType: (*LabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLabel",
			Handler:    _LabelService_GetLabel_Handler,
		},
		{
			MethodName: "MutateLabels",
			Handler:    _LabelService_MutateLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/label_service.proto",
}

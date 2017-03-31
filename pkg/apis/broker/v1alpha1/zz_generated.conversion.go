// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	broker "github.com/openshift/brokersdk/pkg/apis/broker"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/client-go/pkg/api"
	v1 "k8s.io/client-go/pkg/api/v1"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ServiceInstance_To_broker_ServiceInstance,
		Convert_broker_ServiceInstance_To_v1alpha1_ServiceInstance,
		Convert_v1alpha1_ServiceInstanceCondition_To_broker_ServiceInstanceCondition,
		Convert_broker_ServiceInstanceCondition_To_v1alpha1_ServiceInstanceCondition,
		Convert_v1alpha1_ServiceInstanceList_To_broker_ServiceInstanceList,
		Convert_broker_ServiceInstanceList_To_v1alpha1_ServiceInstanceList,
		Convert_v1alpha1_ServiceInstanceSpec_To_broker_ServiceInstanceSpec,
		Convert_broker_ServiceInstanceSpec_To_v1alpha1_ServiceInstanceSpec,
		Convert_v1alpha1_ServiceInstanceStatus_To_broker_ServiceInstanceStatus,
		Convert_broker_ServiceInstanceStatus_To_v1alpha1_ServiceInstanceStatus,
	)
}

func autoConvert_v1alpha1_ServiceInstance_To_broker_ServiceInstance(in *ServiceInstance, out *broker.ServiceInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ServiceInstanceSpec_To_broker_ServiceInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ServiceInstanceStatus_To_broker_ServiceInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_ServiceInstance_To_broker_ServiceInstance(in *ServiceInstance, out *broker.ServiceInstance, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceInstance_To_broker_ServiceInstance(in, out, s)
}

func autoConvert_broker_ServiceInstance_To_v1alpha1_ServiceInstance(in *broker.ServiceInstance, out *ServiceInstance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_broker_ServiceInstanceSpec_To_v1alpha1_ServiceInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_broker_ServiceInstanceStatus_To_v1alpha1_ServiceInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_broker_ServiceInstance_To_v1alpha1_ServiceInstance(in *broker.ServiceInstance, out *ServiceInstance, s conversion.Scope) error {
	return autoConvert_broker_ServiceInstance_To_v1alpha1_ServiceInstance(in, out, s)
}

func autoConvert_v1alpha1_ServiceInstanceCondition_To_broker_ServiceInstanceCondition(in *ServiceInstanceCondition, out *broker.ServiceInstanceCondition, s conversion.Scope) error {
	out.Type = broker.ServiceInstanceConditionType(in.Type)
	out.Status = api.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_v1alpha1_ServiceInstanceCondition_To_broker_ServiceInstanceCondition(in *ServiceInstanceCondition, out *broker.ServiceInstanceCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceInstanceCondition_To_broker_ServiceInstanceCondition(in, out, s)
}

func autoConvert_broker_ServiceInstanceCondition_To_v1alpha1_ServiceInstanceCondition(in *broker.ServiceInstanceCondition, out *ServiceInstanceCondition, s conversion.Scope) error {
	out.Type = ServiceInstanceConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_broker_ServiceInstanceCondition_To_v1alpha1_ServiceInstanceCondition(in *broker.ServiceInstanceCondition, out *ServiceInstanceCondition, s conversion.Scope) error {
	return autoConvert_broker_ServiceInstanceCondition_To_v1alpha1_ServiceInstanceCondition(in, out, s)
}

func autoConvert_v1alpha1_ServiceInstanceList_To_broker_ServiceInstanceList(in *ServiceInstanceList, out *broker.ServiceInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]broker.ServiceInstance)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1alpha1_ServiceInstanceList_To_broker_ServiceInstanceList(in *ServiceInstanceList, out *broker.ServiceInstanceList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceInstanceList_To_broker_ServiceInstanceList(in, out, s)
}

func autoConvert_broker_ServiceInstanceList_To_v1alpha1_ServiceInstanceList(in *broker.ServiceInstanceList, out *ServiceInstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]ServiceInstance, 0)
	} else {
		out.Items = *(*[]ServiceInstance)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_broker_ServiceInstanceList_To_v1alpha1_ServiceInstanceList(in *broker.ServiceInstanceList, out *ServiceInstanceList, s conversion.Scope) error {
	return autoConvert_broker_ServiceInstanceList_To_v1alpha1_ServiceInstanceList(in, out, s)
}

func autoConvert_v1alpha1_ServiceInstanceSpec_To_broker_ServiceInstanceSpec(in *ServiceInstanceSpec, out *broker.ServiceInstanceSpec, s conversion.Scope) error {
	out.Credential = in.Credential
	return nil
}

func Convert_v1alpha1_ServiceInstanceSpec_To_broker_ServiceInstanceSpec(in *ServiceInstanceSpec, out *broker.ServiceInstanceSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceInstanceSpec_To_broker_ServiceInstanceSpec(in, out, s)
}

func autoConvert_broker_ServiceInstanceSpec_To_v1alpha1_ServiceInstanceSpec(in *broker.ServiceInstanceSpec, out *ServiceInstanceSpec, s conversion.Scope) error {
	out.Credential = in.Credential
	return nil
}

func Convert_broker_ServiceInstanceSpec_To_v1alpha1_ServiceInstanceSpec(in *broker.ServiceInstanceSpec, out *ServiceInstanceSpec, s conversion.Scope) error {
	return autoConvert_broker_ServiceInstanceSpec_To_v1alpha1_ServiceInstanceSpec(in, out, s)
}

func autoConvert_v1alpha1_ServiceInstanceStatus_To_broker_ServiceInstanceStatus(in *ServiceInstanceStatus, out *broker.ServiceInstanceStatus, s conversion.Scope) error {
	out.Conditions = *(*[]broker.ServiceInstanceCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

func Convert_v1alpha1_ServiceInstanceStatus_To_broker_ServiceInstanceStatus(in *ServiceInstanceStatus, out *broker.ServiceInstanceStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ServiceInstanceStatus_To_broker_ServiceInstanceStatus(in, out, s)
}

func autoConvert_broker_ServiceInstanceStatus_To_v1alpha1_ServiceInstanceStatus(in *broker.ServiceInstanceStatus, out *ServiceInstanceStatus, s conversion.Scope) error {
	if in.Conditions == nil {
		out.Conditions = make([]ServiceInstanceCondition, 0)
	} else {
		out.Conditions = *(*[]ServiceInstanceCondition)(unsafe.Pointer(&in.Conditions))
	}
	return nil
}

func Convert_broker_ServiceInstanceStatus_To_v1alpha1_ServiceInstanceStatus(in *broker.ServiceInstanceStatus, out *ServiceInstanceStatus, s conversion.Scope) error {
	return autoConvert_broker_ServiceInstanceStatus_To_v1alpha1_ServiceInstanceStatus(in, out, s)
}

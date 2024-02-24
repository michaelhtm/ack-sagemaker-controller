// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package endpoint_config

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.EndpointConfig{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeEndpointConfigOutput
	resp, err = rm.sdkapi.DescribeEndpointConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeEndpointConfig", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ValidationException" && strings.HasPrefix(awsErr.Message(), "Could not find endpoint configuration") {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AsyncInferenceConfig != nil {
		f0 := &svcapitypes.AsyncInferenceConfig{}
		if resp.AsyncInferenceConfig.ClientConfig != nil {
			f0f0 := &svcapitypes.AsyncInferenceClientConfig{}
			if resp.AsyncInferenceConfig.ClientConfig.MaxConcurrentInvocationsPerInstance != nil {
				f0f0.MaxConcurrentInvocationsPerInstance = resp.AsyncInferenceConfig.ClientConfig.MaxConcurrentInvocationsPerInstance
			}
			f0.ClientConfig = f0f0
		}
		if resp.AsyncInferenceConfig.OutputConfig != nil {
			f0f1 := &svcapitypes.AsyncInferenceOutputConfig{}
			if resp.AsyncInferenceConfig.OutputConfig.KmsKeyId != nil {
				f0f1.KMSKeyID = resp.AsyncInferenceConfig.OutputConfig.KmsKeyId
			}
			if resp.AsyncInferenceConfig.OutputConfig.NotificationConfig != nil {
				f0f1f1 := &svcapitypes.AsyncInferenceNotificationConfig{}
				if resp.AsyncInferenceConfig.OutputConfig.NotificationConfig.ErrorTopic != nil {
					f0f1f1.ErrorTopic = resp.AsyncInferenceConfig.OutputConfig.NotificationConfig.ErrorTopic
				}
				if resp.AsyncInferenceConfig.OutputConfig.NotificationConfig.IncludeInferenceResponseIn != nil {
					f0f1f1f1 := []*string{}
					for _, f0f1f1f1iter := range resp.AsyncInferenceConfig.OutputConfig.NotificationConfig.IncludeInferenceResponseIn {
						var f0f1f1f1elem string
						f0f1f1f1elem = *f0f1f1f1iter
						f0f1f1f1 = append(f0f1f1f1, &f0f1f1f1elem)
					}
					f0f1f1.IncludeInferenceResponseIn = f0f1f1f1
				}
				if resp.AsyncInferenceConfig.OutputConfig.NotificationConfig.SuccessTopic != nil {
					f0f1f1.SuccessTopic = resp.AsyncInferenceConfig.OutputConfig.NotificationConfig.SuccessTopic
				}
				f0f1.NotificationConfig = f0f1f1
			}
			if resp.AsyncInferenceConfig.OutputConfig.S3FailurePath != nil {
				f0f1.S3FailurePath = resp.AsyncInferenceConfig.OutputConfig.S3FailurePath
			}
			if resp.AsyncInferenceConfig.OutputConfig.S3OutputPath != nil {
				f0f1.S3OutputPath = resp.AsyncInferenceConfig.OutputConfig.S3OutputPath
			}
			f0.OutputConfig = f0f1
		}
		ko.Spec.AsyncInferenceConfig = f0
	} else {
		ko.Spec.AsyncInferenceConfig = nil
	}
	if resp.DataCaptureConfig != nil {
		f2 := &svcapitypes.DataCaptureConfig{}
		if resp.DataCaptureConfig.CaptureContentTypeHeader != nil {
			f2f0 := &svcapitypes.CaptureContentTypeHeader{}
			if resp.DataCaptureConfig.CaptureContentTypeHeader.CsvContentTypes != nil {
				f2f0f0 := []*string{}
				for _, f2f0f0iter := range resp.DataCaptureConfig.CaptureContentTypeHeader.CsvContentTypes {
					var f2f0f0elem string
					f2f0f0elem = *f2f0f0iter
					f2f0f0 = append(f2f0f0, &f2f0f0elem)
				}
				f2f0.CsvContentTypes = f2f0f0
			}
			if resp.DataCaptureConfig.CaptureContentTypeHeader.JsonContentTypes != nil {
				f2f0f1 := []*string{}
				for _, f2f0f1iter := range resp.DataCaptureConfig.CaptureContentTypeHeader.JsonContentTypes {
					var f2f0f1elem string
					f2f0f1elem = *f2f0f1iter
					f2f0f1 = append(f2f0f1, &f2f0f1elem)
				}
				f2f0.JSONContentTypes = f2f0f1
			}
			f2.CaptureContentTypeHeader = f2f0
		}
		if resp.DataCaptureConfig.CaptureOptions != nil {
			f2f1 := []*svcapitypes.CaptureOption{}
			for _, f2f1iter := range resp.DataCaptureConfig.CaptureOptions {
				f2f1elem := &svcapitypes.CaptureOption{}
				if f2f1iter.CaptureMode != nil {
					f2f1elem.CaptureMode = f2f1iter.CaptureMode
				}
				f2f1 = append(f2f1, f2f1elem)
			}
			f2.CaptureOptions = f2f1
		}
		if resp.DataCaptureConfig.DestinationS3Uri != nil {
			f2.DestinationS3URI = resp.DataCaptureConfig.DestinationS3Uri
		}
		if resp.DataCaptureConfig.EnableCapture != nil {
			f2.EnableCapture = resp.DataCaptureConfig.EnableCapture
		}
		if resp.DataCaptureConfig.InitialSamplingPercentage != nil {
			f2.InitialSamplingPercentage = resp.DataCaptureConfig.InitialSamplingPercentage
		}
		if resp.DataCaptureConfig.KmsKeyId != nil {
			f2.KMSKeyID = resp.DataCaptureConfig.KmsKeyId
		}
		ko.Spec.DataCaptureConfig = f2
	} else {
		ko.Spec.DataCaptureConfig = nil
	}
	if resp.EnableNetworkIsolation != nil {
		ko.Spec.EnableNetworkIsolation = resp.EnableNetworkIsolation
	} else {
		ko.Spec.EnableNetworkIsolation = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.EndpointConfigArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.EndpointConfigArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.EndpointConfigName != nil {
		ko.Spec.EndpointConfigName = resp.EndpointConfigName
	} else {
		ko.Spec.EndpointConfigName = nil
	}
	if resp.ExecutionRoleArn != nil {
		ko.Spec.ExecutionRoleARN = resp.ExecutionRoleArn
	} else {
		ko.Spec.ExecutionRoleARN = nil
	}
	if resp.KmsKeyId != nil {
		ko.Spec.KMSKeyID = resp.KmsKeyId
	} else {
		ko.Spec.KMSKeyID = nil
	}
	if resp.ProductionVariants != nil {
		f8 := []*svcapitypes.ProductionVariant{}
		for _, f8iter := range resp.ProductionVariants {
			f8elem := &svcapitypes.ProductionVariant{}
			if f8iter.AcceleratorType != nil {
				f8elem.AcceleratorType = f8iter.AcceleratorType
			}
			if f8iter.ContainerStartupHealthCheckTimeoutInSeconds != nil {
				f8elem.ContainerStartupHealthCheckTimeoutInSeconds = f8iter.ContainerStartupHealthCheckTimeoutInSeconds
			}
			if f8iter.CoreDumpConfig != nil {
				f8elemf2 := &svcapitypes.ProductionVariantCoreDumpConfig{}
				if f8iter.CoreDumpConfig.DestinationS3Uri != nil {
					f8elemf2.DestinationS3URI = f8iter.CoreDumpConfig.DestinationS3Uri
				}
				if f8iter.CoreDumpConfig.KmsKeyId != nil {
					f8elemf2.KMSKeyID = f8iter.CoreDumpConfig.KmsKeyId
				}
				f8elem.CoreDumpConfig = f8elemf2
			}
			if f8iter.EnableSSMAccess != nil {
				f8elem.EnableSSMAccess = f8iter.EnableSSMAccess
			}
			if f8iter.InitialInstanceCount != nil {
				f8elem.InitialInstanceCount = f8iter.InitialInstanceCount
			}
			if f8iter.InitialVariantWeight != nil {
				f8elem.InitialVariantWeight = f8iter.InitialVariantWeight
			}
			if f8iter.InstanceType != nil {
				f8elem.InstanceType = f8iter.InstanceType
			}
			if f8iter.ManagedInstanceScaling != nil {
				f8elemf7 := &svcapitypes.ProductionVariantManagedInstanceScaling{}
				if f8iter.ManagedInstanceScaling.MaxInstanceCount != nil {
					f8elemf7.MaxInstanceCount = f8iter.ManagedInstanceScaling.MaxInstanceCount
				}
				if f8iter.ManagedInstanceScaling.MinInstanceCount != nil {
					f8elemf7.MinInstanceCount = f8iter.ManagedInstanceScaling.MinInstanceCount
				}
				if f8iter.ManagedInstanceScaling.Status != nil {
					f8elemf7.Status = f8iter.ManagedInstanceScaling.Status
				}
				f8elem.ManagedInstanceScaling = f8elemf7
			}
			if f8iter.ModelDataDownloadTimeoutInSeconds != nil {
				f8elem.ModelDataDownloadTimeoutInSeconds = f8iter.ModelDataDownloadTimeoutInSeconds
			}
			if f8iter.ModelName != nil {
				f8elem.ModelName = f8iter.ModelName
			}
			if f8iter.RoutingConfig != nil {
				f8elemf10 := &svcapitypes.ProductionVariantRoutingConfig{}
				if f8iter.RoutingConfig.RoutingStrategy != nil {
					f8elemf10.RoutingStrategy = f8iter.RoutingConfig.RoutingStrategy
				}
				f8elem.RoutingConfig = f8elemf10
			}
			if f8iter.ServerlessConfig != nil {
				f8elemf11 := &svcapitypes.ProductionVariantServerlessConfig{}
				if f8iter.ServerlessConfig.MaxConcurrency != nil {
					f8elemf11.MaxConcurrency = f8iter.ServerlessConfig.MaxConcurrency
				}
				if f8iter.ServerlessConfig.MemorySizeInMB != nil {
					f8elemf11.MemorySizeInMB = f8iter.ServerlessConfig.MemorySizeInMB
				}
				if f8iter.ServerlessConfig.ProvisionedConcurrency != nil {
					f8elemf11.ProvisionedConcurrency = f8iter.ServerlessConfig.ProvisionedConcurrency
				}
				f8elem.ServerlessConfig = f8elemf11
			}
			if f8iter.VariantName != nil {
				f8elem.VariantName = f8iter.VariantName
			}
			if f8iter.VolumeSizeInGB != nil {
				f8elem.VolumeSizeInGB = f8iter.VolumeSizeInGB
			}
			f8 = append(f8, f8elem)
		}
		ko.Spec.ProductionVariants = f8
	} else {
		ko.Spec.ProductionVariants = nil
	}
	if resp.VpcConfig != nil {
		f10 := &svcapitypes.VPCConfig{}
		if resp.VpcConfig.SecurityGroupIds != nil {
			f10f0 := []*string{}
			for _, f10f0iter := range resp.VpcConfig.SecurityGroupIds {
				var f10f0elem string
				f10f0elem = *f10f0iter
				f10f0 = append(f10f0, &f10f0elem)
			}
			f10.SecurityGroupIDs = f10f0
		}
		if resp.VpcConfig.Subnets != nil {
			f10f1 := []*string{}
			for _, f10f1iter := range resp.VpcConfig.Subnets {
				var f10f1elem string
				f10f1elem = *f10f1iter
				f10f1 = append(f10f1, &f10f1elem)
			}
			f10.Subnets = f10f1
		}
		ko.Spec.VPCConfig = f10
	} else {
		ko.Spec.VPCConfig = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.EndpointConfigName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeEndpointConfigInput, error) {
	res := &svcsdk.DescribeEndpointConfigInput{}

	if r.ko.Spec.EndpointConfigName != nil {
		res.SetEndpointConfigName(*r.ko.Spec.EndpointConfigName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateEndpointConfigOutput
	_ = resp
	resp, err = rm.sdkapi.CreateEndpointConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateEndpointConfig", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.EndpointConfigArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.EndpointConfigArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateEndpointConfigInput, error) {
	res := &svcsdk.CreateEndpointConfigInput{}

	if r.ko.Spec.AsyncInferenceConfig != nil {
		f0 := &svcsdk.AsyncInferenceConfig{}
		if r.ko.Spec.AsyncInferenceConfig.ClientConfig != nil {
			f0f0 := &svcsdk.AsyncInferenceClientConfig{}
			if r.ko.Spec.AsyncInferenceConfig.ClientConfig.MaxConcurrentInvocationsPerInstance != nil {
				f0f0.SetMaxConcurrentInvocationsPerInstance(*r.ko.Spec.AsyncInferenceConfig.ClientConfig.MaxConcurrentInvocationsPerInstance)
			}
			f0.SetClientConfig(f0f0)
		}
		if r.ko.Spec.AsyncInferenceConfig.OutputConfig != nil {
			f0f1 := &svcsdk.AsyncInferenceOutputConfig{}
			if r.ko.Spec.AsyncInferenceConfig.OutputConfig.KMSKeyID != nil {
				f0f1.SetKmsKeyId(*r.ko.Spec.AsyncInferenceConfig.OutputConfig.KMSKeyID)
			}
			if r.ko.Spec.AsyncInferenceConfig.OutputConfig.NotificationConfig != nil {
				f0f1f1 := &svcsdk.AsyncInferenceNotificationConfig{}
				if r.ko.Spec.AsyncInferenceConfig.OutputConfig.NotificationConfig.ErrorTopic != nil {
					f0f1f1.SetErrorTopic(*r.ko.Spec.AsyncInferenceConfig.OutputConfig.NotificationConfig.ErrorTopic)
				}
				if r.ko.Spec.AsyncInferenceConfig.OutputConfig.NotificationConfig.IncludeInferenceResponseIn != nil {
					f0f1f1f1 := []*string{}
					for _, f0f1f1f1iter := range r.ko.Spec.AsyncInferenceConfig.OutputConfig.NotificationConfig.IncludeInferenceResponseIn {
						var f0f1f1f1elem string
						f0f1f1f1elem = *f0f1f1f1iter
						f0f1f1f1 = append(f0f1f1f1, &f0f1f1f1elem)
					}
					f0f1f1.SetIncludeInferenceResponseIn(f0f1f1f1)
				}
				if r.ko.Spec.AsyncInferenceConfig.OutputConfig.NotificationConfig.SuccessTopic != nil {
					f0f1f1.SetSuccessTopic(*r.ko.Spec.AsyncInferenceConfig.OutputConfig.NotificationConfig.SuccessTopic)
				}
				f0f1.SetNotificationConfig(f0f1f1)
			}
			if r.ko.Spec.AsyncInferenceConfig.OutputConfig.S3FailurePath != nil {
				f0f1.SetS3FailurePath(*r.ko.Spec.AsyncInferenceConfig.OutputConfig.S3FailurePath)
			}
			if r.ko.Spec.AsyncInferenceConfig.OutputConfig.S3OutputPath != nil {
				f0f1.SetS3OutputPath(*r.ko.Spec.AsyncInferenceConfig.OutputConfig.S3OutputPath)
			}
			f0.SetOutputConfig(f0f1)
		}
		res.SetAsyncInferenceConfig(f0)
	}
	if r.ko.Spec.DataCaptureConfig != nil {
		f1 := &svcsdk.DataCaptureConfig{}
		if r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader != nil {
			f1f0 := &svcsdk.CaptureContentTypeHeader{}
			if r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader.CsvContentTypes != nil {
				f1f0f0 := []*string{}
				for _, f1f0f0iter := range r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader.CsvContentTypes {
					var f1f0f0elem string
					f1f0f0elem = *f1f0f0iter
					f1f0f0 = append(f1f0f0, &f1f0f0elem)
				}
				f1f0.SetCsvContentTypes(f1f0f0)
			}
			if r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader.JSONContentTypes != nil {
				f1f0f1 := []*string{}
				for _, f1f0f1iter := range r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader.JSONContentTypes {
					var f1f0f1elem string
					f1f0f1elem = *f1f0f1iter
					f1f0f1 = append(f1f0f1, &f1f0f1elem)
				}
				f1f0.SetJsonContentTypes(f1f0f1)
			}
			f1.SetCaptureContentTypeHeader(f1f0)
		}
		if r.ko.Spec.DataCaptureConfig.CaptureOptions != nil {
			f1f1 := []*svcsdk.CaptureOption{}
			for _, f1f1iter := range r.ko.Spec.DataCaptureConfig.CaptureOptions {
				f1f1elem := &svcsdk.CaptureOption{}
				if f1f1iter.CaptureMode != nil {
					f1f1elem.SetCaptureMode(*f1f1iter.CaptureMode)
				}
				f1f1 = append(f1f1, f1f1elem)
			}
			f1.SetCaptureOptions(f1f1)
		}
		if r.ko.Spec.DataCaptureConfig.DestinationS3URI != nil {
			f1.SetDestinationS3Uri(*r.ko.Spec.DataCaptureConfig.DestinationS3URI)
		}
		if r.ko.Spec.DataCaptureConfig.EnableCapture != nil {
			f1.SetEnableCapture(*r.ko.Spec.DataCaptureConfig.EnableCapture)
		}
		if r.ko.Spec.DataCaptureConfig.InitialSamplingPercentage != nil {
			f1.SetInitialSamplingPercentage(*r.ko.Spec.DataCaptureConfig.InitialSamplingPercentage)
		}
		if r.ko.Spec.DataCaptureConfig.KMSKeyID != nil {
			f1.SetKmsKeyId(*r.ko.Spec.DataCaptureConfig.KMSKeyID)
		}
		res.SetDataCaptureConfig(f1)
	}
	if r.ko.Spec.EnableNetworkIsolation != nil {
		res.SetEnableNetworkIsolation(*r.ko.Spec.EnableNetworkIsolation)
	}
	if r.ko.Spec.EndpointConfigName != nil {
		res.SetEndpointConfigName(*r.ko.Spec.EndpointConfigName)
	}
	if r.ko.Spec.ExecutionRoleARN != nil {
		res.SetExecutionRoleArn(*r.ko.Spec.ExecutionRoleARN)
	}
	if r.ko.Spec.KMSKeyID != nil {
		res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	}
	if r.ko.Spec.ProductionVariants != nil {
		f6 := []*svcsdk.ProductionVariant{}
		for _, f6iter := range r.ko.Spec.ProductionVariants {
			f6elem := &svcsdk.ProductionVariant{}
			if f6iter.AcceleratorType != nil {
				f6elem.SetAcceleratorType(*f6iter.AcceleratorType)
			}
			if f6iter.ContainerStartupHealthCheckTimeoutInSeconds != nil {
				f6elem.SetContainerStartupHealthCheckTimeoutInSeconds(*f6iter.ContainerStartupHealthCheckTimeoutInSeconds)
			}
			if f6iter.CoreDumpConfig != nil {
				f6elemf2 := &svcsdk.ProductionVariantCoreDumpConfig{}
				if f6iter.CoreDumpConfig.DestinationS3URI != nil {
					f6elemf2.SetDestinationS3Uri(*f6iter.CoreDumpConfig.DestinationS3URI)
				}
				if f6iter.CoreDumpConfig.KMSKeyID != nil {
					f6elemf2.SetKmsKeyId(*f6iter.CoreDumpConfig.KMSKeyID)
				}
				f6elem.SetCoreDumpConfig(f6elemf2)
			}
			if f6iter.EnableSSMAccess != nil {
				f6elem.SetEnableSSMAccess(*f6iter.EnableSSMAccess)
			}
			if f6iter.InitialInstanceCount != nil {
				f6elem.SetInitialInstanceCount(*f6iter.InitialInstanceCount)
			}
			if f6iter.InitialVariantWeight != nil {
				f6elem.SetInitialVariantWeight(*f6iter.InitialVariantWeight)
			}
			if f6iter.InstanceType != nil {
				f6elem.SetInstanceType(*f6iter.InstanceType)
			}
			if f6iter.ManagedInstanceScaling != nil {
				f6elemf7 := &svcsdk.ProductionVariantManagedInstanceScaling{}
				if f6iter.ManagedInstanceScaling.MaxInstanceCount != nil {
					f6elemf7.SetMaxInstanceCount(*f6iter.ManagedInstanceScaling.MaxInstanceCount)
				}
				if f6iter.ManagedInstanceScaling.MinInstanceCount != nil {
					f6elemf7.SetMinInstanceCount(*f6iter.ManagedInstanceScaling.MinInstanceCount)
				}
				if f6iter.ManagedInstanceScaling.Status != nil {
					f6elemf7.SetStatus(*f6iter.ManagedInstanceScaling.Status)
				}
				f6elem.SetManagedInstanceScaling(f6elemf7)
			}
			if f6iter.ModelDataDownloadTimeoutInSeconds != nil {
				f6elem.SetModelDataDownloadTimeoutInSeconds(*f6iter.ModelDataDownloadTimeoutInSeconds)
			}
			if f6iter.ModelName != nil {
				f6elem.SetModelName(*f6iter.ModelName)
			}
			if f6iter.RoutingConfig != nil {
				f6elemf10 := &svcsdk.ProductionVariantRoutingConfig{}
				if f6iter.RoutingConfig.RoutingStrategy != nil {
					f6elemf10.SetRoutingStrategy(*f6iter.RoutingConfig.RoutingStrategy)
				}
				f6elem.SetRoutingConfig(f6elemf10)
			}
			if f6iter.ServerlessConfig != nil {
				f6elemf11 := &svcsdk.ProductionVariantServerlessConfig{}
				if f6iter.ServerlessConfig.MaxConcurrency != nil {
					f6elemf11.SetMaxConcurrency(*f6iter.ServerlessConfig.MaxConcurrency)
				}
				if f6iter.ServerlessConfig.MemorySizeInMB != nil {
					f6elemf11.SetMemorySizeInMB(*f6iter.ServerlessConfig.MemorySizeInMB)
				}
				if f6iter.ServerlessConfig.ProvisionedConcurrency != nil {
					f6elemf11.SetProvisionedConcurrency(*f6iter.ServerlessConfig.ProvisionedConcurrency)
				}
				f6elem.SetServerlessConfig(f6elemf11)
			}
			if f6iter.VariantName != nil {
				f6elem.SetVariantName(*f6iter.VariantName)
			}
			if f6iter.VolumeSizeInGB != nil {
				f6elem.SetVolumeSizeInGB(*f6iter.VolumeSizeInGB)
			}
			f6 = append(f6, f6elem)
		}
		res.SetProductionVariants(f6)
	}
	if r.ko.Spec.Tags != nil {
		f7 := []*svcsdk.Tag{}
		for _, f7iter := range r.ko.Spec.Tags {
			f7elem := &svcsdk.Tag{}
			if f7iter.Key != nil {
				f7elem.SetKey(*f7iter.Key)
			}
			if f7iter.Value != nil {
				f7elem.SetValue(*f7iter.Value)
			}
			f7 = append(f7, f7elem)
		}
		res.SetTags(f7)
	}
	if r.ko.Spec.VPCConfig != nil {
		f8 := &svcsdk.VpcConfig{}
		if r.ko.Spec.VPCConfig.SecurityGroupIDs != nil {
			f8f0 := []*string{}
			for _, f8f0iter := range r.ko.Spec.VPCConfig.SecurityGroupIDs {
				var f8f0elem string
				f8f0elem = *f8f0iter
				f8f0 = append(f8f0, &f8f0elem)
			}
			f8.SetSecurityGroupIds(f8f0)
		}
		if r.ko.Spec.VPCConfig.Subnets != nil {
			f8f1 := []*string{}
			for _, f8f1iter := range r.ko.Spec.VPCConfig.Subnets {
				var f8f1elem string
				f8f1elem = *f8f1iter
				f8f1 = append(f8f1, &f8f1elem)
			}
			f8.SetSubnets(f8f1)
		}
		res.SetVpcConfig(f8)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return nil, ackerr.NewTerminalError(ackerr.NotImplemented)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteEndpointConfigOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteEndpointConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteEndpointConfig", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteEndpointConfigInput, error) {
	res := &svcsdk.DeleteEndpointConfigInput{}

	if r.ko.Spec.EndpointConfigName != nil {
		res.SetEndpointConfigName(*r.ko.Spec.EndpointConfigName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.EndpointConfig,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidParameterCombination",
		"InvalidParameterValue",
		"MissingParameter":
		return true
	default:
		return false
	}
}

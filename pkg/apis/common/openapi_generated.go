// +build !ignore_autogenerated

/*
Copyright 2020 BlackRock, Inc.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package common

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/pkg/apis/common.Event":         schema_argo_events_pkg_apis_common_Event(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.EventContext":  schema_argo_events_pkg_apis_common_EventContext(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.EventProtocol": schema_argo_events_pkg_apis_common_EventProtocol(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.Http":          schema_argo_events_pkg_apis_common_Http(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.Nats":          schema_argo_events_pkg_apis_common_Nats(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.S3Artifact":    schema_argo_events_pkg_apis_common_S3Artifact(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.S3Bucket":      schema_argo_events_pkg_apis_common_S3Bucket(ref),
		"github.com/argoproj/argo-events/pkg/apis/common.S3Filter":      schema_argo_events_pkg_apis_common_S3Filter(ref),
	}
}

func schema_argo_events_pkg_apis_common_Event(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Event is a data and its context. Adheres to the CloudEvents v0.3 specification",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"context": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.EventContext"),
						},
					},
					"data": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "bytes",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "byte",
						},
					},
				},
				Required: []string{"context", "data"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/common.EventContext"},
	}
}

func schema_argo_events_pkg_apis_common_EventContext(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventContext contains metadata that provides circumstantial information about the occurrence.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "The type of occurrence which has happened. Often this attribute is used for routing, observability, policy enforcement, etc. should be prefixed with a reverse-DNS name. The prefixed domain dictates the organization which defines the semantics of this event type. ex: com.github.pull.create",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"specVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "The version of the CloudEvents specification which the event uses. Enables the interpretation of the context.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"source": {
						SchemaProps: spec.SchemaProps{
							Description: "This describes the event producer.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "ID of the event. The semantics are explicitly undefined to ease the implementation of producers.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"time": {
						SchemaProps: spec.SchemaProps{
							Description: "Time when the event happened. Must adhere to format specified in RFC 3339.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.MicroTime"),
						},
					},
					"dataContentType": {
						SchemaProps: spec.SchemaProps{
							Description: "Content type of the data attribute value. Enables the data attribute to carry any type of content, whereby format and encoding might differ from that of the chosen event format. For example, the data attribute may carry an XML or JSON payload and the consumer is informed by this attribute being set to \"application/xml\" or \"application/json\" respectively.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"subject": {
						SchemaProps: spec.SchemaProps{
							Description: "Subject of the event",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "specVersion", "source", "id", "time", "dataContentType", "subject"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.MicroTime"},
	}
}

func schema_argo_events_pkg_apis_common_EventProtocol(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Dispatch protocol contains configuration necessary to dispatch an event to sensor over different communication protocols",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"http": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.Http"),
						},
					},
					"nats": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.Nats"),
						},
					},
				},
				Required: []string{"type", "http", "nats"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/common.Http", "github.com/argoproj/argo-events/pkg/apis/common.Nats"},
	}
}

func schema_argo_events_pkg_apis_common_Http(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Http contains the information required to setup a http server and listen to incoming events",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"port": {
						SchemaProps: spec.SchemaProps{
							Description: "Port on which server will run",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels to be set for the service generated",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Description: "Annotations to be set for the service generated",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"port"},
			},
		},
	}
}

func schema_argo_events_pkg_apis_common_Nats(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Nats contains the information required to connect to nats server and get subscriptions",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL is nats server/service URL",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startWithLastReceived": {
						SchemaProps: spec.SchemaProps{
							Description: "Subscribe starting with most recently published value. Refer https://github.com/nats-io/go-nats-streaming",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"deliverAllAvailable": {
						SchemaProps: spec.SchemaProps{
							Description: "Receive all stored values in order.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"startAtSequence": {
						SchemaProps: spec.SchemaProps{
							Description: "Receive messages starting at a specific sequence number",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startAtTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Subscribe starting at a specific time",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startAtTimeDelta": {
						SchemaProps: spec.SchemaProps{
							Description: "Subscribe starting a specific amount of time in the past (e.g. 30 seconds ago)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"durable": {
						SchemaProps: spec.SchemaProps{
							Description: "Durable subscriptions allow clients to assign a durable name to a subscription when it is created",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"clusterId": {
						SchemaProps: spec.SchemaProps{
							Description: "The NATS Streaming cluster ID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"clientId": {
						SchemaProps: spec.SchemaProps{
							Description: "The NATS Streaming cluster ID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of the connection. either standard or streaming",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"url", "type"},
			},
		},
	}
}

func schema_argo_events_pkg_apis_common_S3Artifact(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Artifact contains information about an S3 connection and bucket",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"bucket": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.S3Bucket"),
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"insecure": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"accessKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"secretKey": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
					"events": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "string",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"filter": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/common.S3Filter"),
						},
					},
				},
				Required: []string{"endpoint", "bucket", "accessKey", "secretKey"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/common.S3Bucket", "github.com/argoproj/argo-events/pkg/apis/common.S3Filter", "k8s.io/api/core/v1.SecretKeySelector"},
	}
}

func schema_argo_events_pkg_apis_common_S3Bucket(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Bucket contains information to describe an S3 Bucket",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"key": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_argo_events_pkg_apis_common_S3Filter(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Filter represents filters to apply to bucket nofifications for specifying constraints on objects",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"prefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"suffix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"prefix", "suffix"},
			},
		},
	}
}
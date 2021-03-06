/*
Copyright 2018 the Heptio Ark contributors.

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

package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServerStatusRequest is a request to access current status information about
// the Ark server.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type ServerStatusRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ServerStatusRequestSpec   `json:"spec"`
	Status ServerStatusRequestStatus `json:"status,omitempty"`
}

// ServerStatusRequestSpec is the specification for a ServerStatusRequest.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type ServerStatusRequestSpec struct {
}

// ServerStatusRequestPhase represents the lifecycle phase of a ServerStatusRequest.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type ServerStatusRequestPhase string

const (
	// ServerStatusRequestPhaseNew means the ServerStatusRequest has not been processed yet.
	ServerStatusRequestPhaseNew ServerStatusRequestPhase = "New"
	// ServerStatusRequestPhaseProcessed means the ServerStatusRequest has been processed.
	ServerStatusRequestPhaseProcessed ServerStatusRequestPhase = "Processed"
)

// ServerStatusRequestStatus is the current status of a ServerStatusRequest.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type ServerStatusRequestStatus struct {
	// Phase is the current lifecycle phase of the ServerStatusRequest.
	Phase ServerStatusRequestPhase `json:"phase"`

	// ProcessedTimestamp is when the ServerStatusRequest was processed
	// by the ServerStatusRequestController.
	ProcessedTimestamp metav1.Time `json:"processedTimestamp"`

	// ServerVersion is the Ark server version.
	ServerVersion string `json:"serverVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServerStatusRequestList is a list of ServerStatusRequests.
//
// Deprecated: Consumers should switch to the same type in
// github.com/heptio/velero/pkg/apis/velero/v1 instead. This
// type will be removed in the v1.0 release.
type ServerStatusRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ServerStatusRequest `json:"items"`
}

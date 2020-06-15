/*******************************************************************************
 * Copyright (c) 2020 IBM Corporation and others.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 *     IBM Corporation - initial API and implementation
 *******************************************************************************/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KeycloakSpec defines the desired state of Keycloak
type KeycloakSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// StorageSize : Size of the Keycloak PVC
	// +kubebuilder:validation:Pattern=[0-9]*Gi$
	StorageSize string `json:"storageSize"`
}

// KeycloakStatus defines the observed state of Keycloak
type KeycloakStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	Phase        string `json:"phase"`
	AccessURL    string `json:"url"`
	DefaultRealm string `json:"defaultRealm"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Keycloak is the Schema for the keycloaks API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=keycloaks,scope=Namespaced
// +kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".metadata.namespace",priority=0,description="Deployment namespace"
// +kubebuilder:printcolumn:name="AuthID",type="string",JSONPath=".metadata.annotations.authID",priority=0,description="Deployment AuthID"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",priority=0,description="Age of the resource"
// +kubebuilder:printcolumn:name="Access",type="string",JSONPath=".status.url",priority=0,description="Exposed route"
type Keycloak struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeycloakSpec   `json:"spec,omitempty"`
	Status            KeycloakStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeycloakList contains a list of Keycloak
type KeycloakList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Keycloak `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Keycloak{}, &KeycloakList{})
}

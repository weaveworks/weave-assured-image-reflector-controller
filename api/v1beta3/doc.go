/*
Copyright 2023 The Flux authors

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

// Package v1beta3 contains API types for the image API group, version
// v1beta3. These types are concerned with reflecting metadata from
// OCI image repositories into a cluster, so they can be consulted for
// e.g., automation.
//
// +kubebuilder:object:generate=true
// +groupName=image.toolkit.fluxcd.io
package v1beta3

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
  This file is part of Astarte.

  Copyright 2020-23 SECO Mind Srl

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Astarte) DeepCopyInto(out *Astarte) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Astarte.
func (in *Astarte) DeepCopy() *Astarte {
	if in == nil {
		return nil
	}
	out := new(Astarte)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Astarte) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteAPISpec) DeepCopyInto(out *AstarteAPISpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteAPISpec.
func (in *AstarteAPISpec) DeepCopy() *AstarteAPISpec {
	if in == nil {
		return nil
	}
	out := new(AstarteAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteAppengineAPISpec) DeepCopyInto(out *AstarteAppengineAPISpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AstarteGenericAPISpec.DeepCopyInto(&out.AstarteGenericAPISpec)
	if in.MaxResultsLimit != nil {
		in, out := &in.MaxResultsLimit, &out.MaxResultsLimit
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteAppengineAPISpec.
func (in *AstarteAppengineAPISpec) DeepCopy() *AstarteAppengineAPISpec {
	if in == nil {
		return nil
	}
	out := new(AstarteAppengineAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCARootConfigSigningCAConstraintSpec) DeepCopyInto(out *AstarteCFSSLCARootConfigSigningCAConstraintSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCARootConfigSigningCAConstraintSpec.
func (in *AstarteCFSSLCARootConfigSigningCAConstraintSpec) DeepCopy() *AstarteCFSSLCARootConfigSigningCAConstraintSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCARootConfigSigningCAConstraintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCARootConfigSigningDefaultSpec) DeepCopyInto(out *AstarteCFSSLCARootConfigSigningDefaultSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CAConstraint != nil {
		in, out := &in.CAConstraint, &out.CAConstraint
		*out = new(AstarteCFSSLCARootConfigSigningCAConstraintSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCARootConfigSigningDefaultSpec.
func (in *AstarteCFSSLCARootConfigSigningDefaultSpec) DeepCopy() *AstarteCFSSLCARootConfigSigningDefaultSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCARootConfigSigningDefaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCARootConfigSigningSpec) DeepCopyInto(out *AstarteCFSSLCARootConfigSigningSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(AstarteCFSSLCARootConfigSigningDefaultSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCARootConfigSigningSpec.
func (in *AstarteCFSSLCARootConfigSigningSpec) DeepCopy() *AstarteCFSSLCARootConfigSigningSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCARootConfigSigningSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCARootConfigSpec) DeepCopyInto(out *AstarteCFSSLCARootConfigSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Signing != nil {
		in, out := &in.Signing, &out.Signing
		*out = new(AstarteCFSSLCARootConfigSigningSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCARootConfigSpec.
func (in *AstarteCFSSLCARootConfigSpec) DeepCopy() *AstarteCFSSLCARootConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCARootConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCSRRootCACASpec) DeepCopyInto(out *AstarteCFSSLCSRRootCACASpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCSRRootCACASpec.
func (in *AstarteCFSSLCSRRootCACASpec) DeepCopy() *AstarteCFSSLCSRRootCACASpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCSRRootCACASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCSRRootCAKeySpec) DeepCopyInto(out *AstarteCFSSLCSRRootCAKeySpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCSRRootCAKeySpec.
func (in *AstarteCFSSLCSRRootCAKeySpec) DeepCopy() *AstarteCFSSLCSRRootCAKeySpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCSRRootCAKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCSRRootCANamesSpec) DeepCopyInto(out *AstarteCFSSLCSRRootCANamesSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCSRRootCANamesSpec.
func (in *AstarteCFSSLCSRRootCANamesSpec) DeepCopy() *AstarteCFSSLCSRRootCANamesSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCSRRootCANamesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLCSRRootCASpec) DeepCopyInto(out *AstarteCFSSLCSRRootCASpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(AstarteCFSSLCSRRootCAKeySpec)
		**out = **in
	}
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]AstarteCFSSLCSRRootCANamesSpec, len(*in))
		copy(*out, *in)
	}
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(AstarteCFSSLCSRRootCACASpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLCSRRootCASpec.
func (in *AstarteCFSSLCSRRootCASpec) DeepCopy() *AstarteCFSSLCSRRootCASpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLCSRRootCASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLDBConfigSpec) DeepCopyInto(out *AstarteCFSSLDBConfigSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLDBConfigSpec.
func (in *AstarteCFSSLDBConfigSpec) DeepCopy() *AstarteCFSSLDBConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLDBConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCFSSLSpec) DeepCopyInto(out *AstarteCFSSLSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	out.CASecret = in.CASecret
	if in.DBConfig != nil {
		in, out := &in.DBConfig, &out.DBConfig
		*out = new(AstarteCFSSLDBConfigSpec)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(AstartePersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CSRRootCa != nil {
		in, out := &in.CSRRootCa, &out.CSRRootCa
		*out = new(AstarteCFSSLCSRRootCASpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CARootConfig != nil {
		in, out := &in.CARootConfig, &out.CARootConfig
		*out = new(AstarteCFSSLCARootConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCFSSLSpec.
func (in *AstarteCFSSLSpec) DeepCopy() *AstarteCFSSLSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCFSSLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCassandraConnectionSpec) DeepCopyInto(out *AstarteCassandraConnectionSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.PoolSize != nil {
		in, out := &in.PoolSize, &out.PoolSize
		*out = new(int)
		**out = **in
	}
	if in.Autodiscovery != nil {
		in, out := &in.Autodiscovery, &out.Autodiscovery
		*out = new(bool)
		**out = **in
	}
	in.SSLConfiguration.DeepCopyInto(&out.SSLConfiguration)
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(LoginCredentialsSecret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCassandraConnectionSpec.
func (in *AstarteCassandraConnectionSpec) DeepCopy() *AstarteCassandraConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCassandraConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCassandraSSLConfigurationSpec) DeepCopyInto(out *AstarteCassandraSSLConfigurationSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.CustomCASecret = in.CustomCASecret
	if in.SNI != nil {
		in, out := &in.SNI, &out.SNI
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCassandraSSLConfigurationSpec.
func (in *AstarteCassandraSSLConfigurationSpec) DeepCopy() *AstarteCassandraSSLConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCassandraSSLConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteCassandraSpec) DeepCopyInto(out *AstarteCassandraSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(AstartePersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(AstarteCassandraConnectionSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteCassandraSpec.
func (in *AstarteCassandraSpec) DeepCopy() *AstarteCassandraSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteCassandraSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteComponentsSpec) DeepCopyInto(out *AstarteComponentsSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.Flow.DeepCopyInto(&out.Flow)
	in.Housekeeping.DeepCopyInto(&out.Housekeeping)
	in.RealmManagement.DeepCopyInto(&out.RealmManagement)
	in.Pairing.DeepCopyInto(&out.Pairing)
	in.DataUpdaterPlant.DeepCopyInto(&out.DataUpdaterPlant)
	in.AppengineAPI.DeepCopyInto(&out.AppengineAPI)
	in.TriggerEngine.DeepCopyInto(&out.TriggerEngine)
	in.Dashboard.DeepCopyInto(&out.Dashboard)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteComponentsSpec.
func (in *AstarteComponentsSpec) DeepCopy() *AstarteComponentsSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteComponentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDashboardConfigAuthSpec) DeepCopyInto(out *AstarteDashboardConfigAuthSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDashboardConfigAuthSpec.
func (in *AstarteDashboardConfigAuthSpec) DeepCopy() *AstarteDashboardConfigAuthSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDashboardConfigAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDashboardConfigSpec) DeepCopyInto(out *AstarteDashboardConfigSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = make([]AstarteDashboardConfigAuthSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDashboardConfigSpec.
func (in *AstarteDashboardConfigSpec) DeepCopy() *AstarteDashboardConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDashboardConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDashboardSpec) DeepCopyInto(out *AstarteDashboardSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDashboardSpec.
func (in *AstarteDashboardSpec) DeepCopy() *AstarteDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteDataUpdaterPlantSpec) DeepCopyInto(out *AstarteDataUpdaterPlantSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.DataQueueCount != nil {
		in, out := &in.DataQueueCount, &out.DataQueueCount
		*out = new(int)
		**out = **in
	}
	if in.PrefetchCount != nil {
		in, out := &in.PrefetchCount, &out.PrefetchCount
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteDataUpdaterPlantSpec.
func (in *AstarteDataUpdaterPlantSpec) DeepCopy() *AstarteDataUpdaterPlantSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteDataUpdaterPlantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteFeatures) DeepCopyInto(out *AstarteFeatures) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.AstartePodPriorities != nil {
		in, out := &in.AstartePodPriorities, &out.AstartePodPriorities
		*out = new(AstartePodPrioritiesSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteFeatures.
func (in *AstarteFeatures) DeepCopy() *AstarteFeatures {
	if in == nil {
		return nil
	}
	out := new(AstarteFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteGenericAPISpec) DeepCopyInto(out *AstarteGenericAPISpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.DisableAuthentication != nil {
		in, out := &in.DisableAuthentication, &out.DisableAuthentication
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteGenericAPISpec.
func (in *AstarteGenericAPISpec) DeepCopy() *AstarteGenericAPISpec {
	if in == nil {
		return nil
	}
	out := new(AstarteGenericAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteGenericClusteredResource) DeepCopyInto(out *AstarteGenericClusteredResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Deploy != nil {
		in, out := &in.Deploy, &out.Deploy
		*out = new(bool)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.AntiAffinity != nil {
		in, out := &in.AntiAffinity, &out.AntiAffinity
		*out = new(bool)
		**out = **in
	}
	if in.CustomAffinity != nil {
		in, out := &in.CustomAffinity, &out.CustomAffinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentStrategy != nil {
		in, out := &in.DeploymentStrategy, &out.DeploymentStrategy
		*out = new(appsv1.DeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalEnv != nil {
		in, out := &in.AdditionalEnv, &out.AdditionalEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(AstarteGenericClusteredResourceAutoscalerSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteGenericClusteredResource.
func (in *AstarteGenericClusteredResource) DeepCopy() *AstarteGenericClusteredResource {
	if in == nil {
		return nil
	}
	out := new(AstarteGenericClusteredResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteGenericClusteredResourceAutoscalerSpec) DeepCopyInto(out *AstarteGenericClusteredResourceAutoscalerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteGenericClusteredResourceAutoscalerSpec.
func (in *AstarteGenericClusteredResourceAutoscalerSpec) DeepCopy() *AstarteGenericClusteredResourceAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteGenericClusteredResourceAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteGenericComponentSpec) DeepCopyInto(out *AstarteGenericComponentSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.API.DeepCopyInto(&out.API)
	in.Backend.DeepCopyInto(&out.Backend)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteGenericComponentSpec.
func (in *AstarteGenericComponentSpec) DeepCopy() *AstarteGenericComponentSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteGenericComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteList) DeepCopyInto(out *AstarteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Astarte, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteList.
func (in *AstarteList) DeepCopy() *AstarteList {
	if in == nil {
		return nil
	}
	out := new(AstarteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AstarteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstartePersistentStorageSpec) DeepCopyInto(out *AstartePersistentStorageSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.VolumeDefinition != nil {
		in, out := &in.VolumeDefinition, &out.VolumeDefinition
		*out = new(v1.Volume)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstartePersistentStorageSpec.
func (in *AstartePersistentStorageSpec) DeepCopy() *AstartePersistentStorageSpec {
	if in == nil {
		return nil
	}
	out := new(AstartePersistentStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstartePodPrioritiesSpec) DeepCopyInto(out *AstartePodPrioritiesSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.AstarteHighPriority != nil {
		in, out := &in.AstarteHighPriority, &out.AstarteHighPriority
		*out = new(int)
		**out = **in
	}
	if in.AstarteMidPriority != nil {
		in, out := &in.AstarteMidPriority, &out.AstarteMidPriority
		*out = new(int)
		**out = **in
	}
	if in.AstarteLowPriority != nil {
		in, out := &in.AstarteLowPriority, &out.AstarteLowPriority
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstartePodPrioritiesSpec.
func (in *AstartePodPrioritiesSpec) DeepCopy() *AstartePodPrioritiesSpec {
	if in == nil {
		return nil
	}
	out := new(AstartePodPrioritiesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteRabbitMQConnectionSpec) DeepCopyInto(out *AstarteRabbitMQConnectionSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int16)
		**out = **in
	}
	in.SSLConfiguration.DeepCopyInto(&out.SSLConfiguration)
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(LoginCredentialsSecret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteRabbitMQConnectionSpec.
func (in *AstarteRabbitMQConnectionSpec) DeepCopy() *AstarteRabbitMQConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteRabbitMQConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteRabbitMQSSLConfigurationSpec) DeepCopyInto(out *AstarteRabbitMQSSLConfigurationSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.CustomCASecret = in.CustomCASecret
	if in.SNI != nil {
		in, out := &in.SNI, &out.SNI
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteRabbitMQSSLConfigurationSpec.
func (in *AstarteRabbitMQSSLConfigurationSpec) DeepCopy() *AstarteRabbitMQSSLConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteRabbitMQSSLConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteRabbitMQSpec) DeepCopyInto(out *AstarteRabbitMQSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(AstarteRabbitMQConnectionSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(AstartePersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalPlugins != nil {
		in, out := &in.AdditionalPlugins, &out.AdditionalPlugins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteRabbitMQSpec.
func (in *AstarteRabbitMQSpec) DeepCopy() *AstarteRabbitMQSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteRabbitMQSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteSpec) DeepCopyInto(out *AstarteSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.DeploymentStrategy != nil {
		in, out := &in.DeploymentStrategy, &out.DeploymentStrategy
		*out = new(appsv1.DeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
	in.Features.DeepCopyInto(&out.Features)
	if in.RBAC != nil {
		in, out := &in.RBAC, &out.RBAC
		*out = new(bool)
		**out = **in
	}
	in.API.DeepCopyInto(&out.API)
	in.RabbitMQ.DeepCopyInto(&out.RabbitMQ)
	in.Cassandra.DeepCopyInto(&out.Cassandra)
	in.VerneMQ.DeepCopyInto(&out.VerneMQ)
	in.CFSSL.DeepCopyInto(&out.CFSSL)
	in.Components.DeepCopyInto(&out.Components)
	out.AstarteSystemKeyspace = in.AstarteSystemKeyspace
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteSpec.
func (in *AstarteSpec) DeepCopy() *AstarteSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteStatus) DeepCopyInto(out *AstarteStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteStatus.
func (in *AstarteStatus) DeepCopy() *AstarteStatus {
	if in == nil {
		return nil
	}
	out := new(AstarteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteSystemKeyspaceSpec) DeepCopyInto(out *AstarteSystemKeyspaceSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteSystemKeyspaceSpec.
func (in *AstarteSystemKeyspaceSpec) DeepCopy() *AstarteSystemKeyspaceSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteSystemKeyspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteTriggerEngineSpec) DeepCopyInto(out *AstarteTriggerEngineSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteTriggerEngineSpec.
func (in *AstarteTriggerEngineSpec) DeepCopy() *AstarteTriggerEngineSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteTriggerEngineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstarteVerneMQSpec) DeepCopyInto(out *AstarteVerneMQSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AstarteGenericClusteredResource.DeepCopyInto(&out.AstarteGenericClusteredResource)
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int16)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(AstartePersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxOfflineMessages != nil {
		in, out := &in.MaxOfflineMessages, &out.MaxOfflineMessages
		*out = new(int)
		**out = **in
	}
	if in.SSLListener != nil {
		in, out := &in.SSLListener, &out.SSLListener
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstarteVerneMQSpec.
func (in *AstarteVerneMQSpec) DeepCopy() *AstarteVerneMQSpec {
	if in == nil {
		return nil
	}
	out := new(AstarteVerneMQSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockWorker) DeepCopyInto(out *BlockWorker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.DataProvider.DeepCopyInto(&out.DataProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockWorker.
func (in *BlockWorker) DeepCopy() *BlockWorker {
	if in == nil {
		return nil
	}
	out := new(BlockWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerBlockSpec) DeepCopyInto(out *ContainerBlockSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = make([]BlockWorker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerBlockSpec.
func (in *ContainerBlockSpec) DeepCopy() *ContainerBlockSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerBlockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProvider) DeepCopyInto(out *DataProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.RabbitMQ != nil {
		in, out := &in.RabbitMQ, &out.RabbitMQ
		*out = new(RabbitMQDataProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProvider.
func (in *DataProvider) DeepCopy() *DataProvider {
	if in == nil {
		return nil
	}
	out := new(DataProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flow) DeepCopyInto(out *Flow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flow.
func (in *Flow) DeepCopy() *Flow {
	if in == nil {
		return nil
	}
	out := new(Flow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Flow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowList) DeepCopyInto(out *FlowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Flow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowList.
func (in *FlowList) DeepCopy() *FlowList {
	if in == nil {
		return nil
	}
	out := new(FlowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSpec) DeepCopyInto(out *FlowSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Astarte = in.Astarte
	if in.NativeBlocksResources != nil {
		in, out := &in.NativeBlocksResources, &out.NativeBlocksResources
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	out.FlowPool = in.FlowPool
	if in.ContainerBlocks != nil {
		in, out := &in.ContainerBlocks, &out.ContainerBlocks
		*out = make([]ContainerBlockSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSpec.
func (in *FlowSpec) DeepCopy() *FlowSpec {
	if in == nil {
		return nil
	}
	out := new(FlowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowStatus) DeepCopyInto(out *FlowStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.UnrecoverableFailures != nil {
		in, out := &in.UnrecoverableFailures, &out.UnrecoverableFailures
		*out = make([]v1.ContainerState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowStatus.
func (in *FlowStatus) DeepCopy() *FlowStatus {
	if in == nil {
		return nil
	}
	out := new(FlowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginCredentialsSecret) DeepCopyInto(out *LoginCredentialsSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginCredentialsSecret.
func (in *LoginCredentialsSecret) DeepCopy() *LoginCredentialsSecret {
	if in == nil {
		return nil
	}
	out := new(LoginCredentialsSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitMQConfig) DeepCopyInto(out *RabbitMQConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitMQConfig.
func (in *RabbitMQConfig) DeepCopy() *RabbitMQConfig {
	if in == nil {
		return nil
	}
	out := new(RabbitMQConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitMQDataProvider) DeepCopyInto(out *RabbitMQDataProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Queues != nil {
		in, out := &in.Queues, &out.Queues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exchange != nil {
		in, out := &in.Exchange, &out.Exchange
		*out = new(RabbitMQExchange)
		**out = **in
	}
	if in.RabbitMQConfig != nil {
		in, out := &in.RabbitMQConfig, &out.RabbitMQConfig
		*out = new(RabbitMQConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitMQDataProvider.
func (in *RabbitMQDataProvider) DeepCopy() *RabbitMQDataProvider {
	if in == nil {
		return nil
	}
	out := new(RabbitMQDataProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitMQExchange) DeepCopyInto(out *RabbitMQExchange) {
	*out = *in
	out.TypeMeta = in.TypeMeta
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitMQExchange.
func (in *RabbitMQExchange) DeepCopy() *RabbitMQExchange {
	if in == nil {
		return nil
	}
	out := new(RabbitMQExchange)
	in.DeepCopyInto(out)
	return out
}

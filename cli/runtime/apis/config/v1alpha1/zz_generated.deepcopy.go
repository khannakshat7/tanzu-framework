//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CLIOptions) DeepCopyInto(out *CLIOptions) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]PluginRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DiscoverySources != nil {
		in, out := &in.DiscoverySources, &out.DiscoverySources
		*out = make([]PluginDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CLIOptions.
func (in *CLIOptions) DeepCopy() *CLIOptions {
	if in == nil {
		return nil
	}
	out := new(CLIOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConfig) DeepCopyInto(out *ClientConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.KnownServers != nil {
		in, out := &in.KnownServers, &out.KnownServers
		*out = make([]*Server, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Server)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KnownContexts != nil {
		in, out := &in.KnownContexts, &out.KnownContexts
		*out = make([]*Context, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Context)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CurrentContext != nil {
		in, out := &in.CurrentContext, &out.CurrentContext
		*out = make(map[ContextType]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ClientOptions != nil {
		in, out := &in.ClientOptions, &out.ClientOptions
		*out = new(ClientOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConfig.
func (in *ClientConfig) DeepCopy() *ClientConfig {
	if in == nil {
		return nil
	}
	out := new(ClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConfigList) DeepCopyInto(out *ClientConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClientConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConfigList.
func (in *ClientConfigList) DeepCopy() *ClientConfigList {
	if in == nil {
		return nil
	}
	out := new(ClientConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClientConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientOptions) DeepCopyInto(out *ClientOptions) {
	*out = *in
	if in.CLI != nil {
		in, out := &in.CLI, &out.CLI
		*out = new(CLIOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make(map[string]FeatureMap, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(FeatureMap, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientOptions.
func (in *ClientOptions) DeepCopy() *ClientOptions {
	if in == nil {
		return nil
	}
	out := new(ClientOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterServer) DeepCopyInto(out *ClusterServer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterServer.
func (in *ClusterServer) DeepCopy() *ClusterServer {
	if in == nil {
		return nil
	}
	out := new(ClusterServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMetadata) DeepCopyInto(out *ConfigMetadata) {
	*out = *in
	if in.PatchStrategy != nil {
		in, out := &in.PatchStrategy, &out.PatchStrategy
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMetadata.
func (in *ConfigMetadata) DeepCopy() *ConfigMetadata {
	if in == nil {
		return nil
	}
	out := new(ConfigMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	if in.GlobalOpts != nil {
		in, out := &in.GlobalOpts, &out.GlobalOpts
		*out = new(GlobalServer)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterOpts != nil {
		in, out := &in.ClusterOpts, &out.ClusterOpts
		*out = new(ClusterServer)
		**out = **in
	}
	if in.DiscoverySources != nil {
		in, out := &in.DiscoverySources, &out.DiscoverySources
		*out = make([]PluginDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EnvMap) DeepCopyInto(out *EnvMap) {
	{
		in := &in
		*out = make(EnvMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvMap.
func (in EnvMap) DeepCopy() EnvMap {
	if in == nil {
		return nil
	}
	out := new(EnvMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FeatureMap) DeepCopyInto(out *FeatureMap) {
	{
		in := &in
		*out = make(FeatureMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureMap.
func (in FeatureMap) DeepCopy() FeatureMap {
	if in == nil {
		return nil
	}
	out := new(FeatureMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPDiscovery) DeepCopyInto(out *GCPDiscovery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPDiscovery.
func (in *GCPDiscovery) DeepCopy() *GCPDiscovery {
	if in == nil {
		return nil
	}
	out := new(GCPDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPPluginRepository) DeepCopyInto(out *GCPPluginRepository) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPPluginRepository.
func (in *GCPPluginRepository) DeepCopy() *GCPPluginRepository {
	if in == nil {
		return nil
	}
	out := new(GCPPluginRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericRESTDiscovery) DeepCopyInto(out *GenericRESTDiscovery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericRESTDiscovery.
func (in *GenericRESTDiscovery) DeepCopy() *GenericRESTDiscovery {
	if in == nil {
		return nil
	}
	out := new(GenericRESTDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalServer) DeepCopyInto(out *GlobalServer) {
	*out = *in
	in.Auth.DeepCopyInto(&out.Auth)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalServer.
func (in *GlobalServer) DeepCopy() *GlobalServer {
	if in == nil {
		return nil
	}
	out := new(GlobalServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalServerAuth) DeepCopyInto(out *GlobalServerAuth) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Expiration.DeepCopyInto(&out.Expiration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalServerAuth.
func (in *GlobalServerAuth) DeepCopy() *GlobalServerAuth {
	if in == nil {
		return nil
	}
	out := new(GlobalServerAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesDiscovery) DeepCopyInto(out *KubernetesDiscovery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesDiscovery.
func (in *KubernetesDiscovery) DeepCopy() *KubernetesDiscovery {
	if in == nil {
		return nil
	}
	out := new(KubernetesDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDiscovery) DeepCopyInto(out *LocalDiscovery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDiscovery.
func (in *LocalDiscovery) DeepCopy() *LocalDiscovery {
	if in == nil {
		return nil
	}
	out := new(LocalDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementClusterServer) DeepCopyInto(out *ManagementClusterServer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementClusterServer.
func (in *ManagementClusterServer) DeepCopy() *ManagementClusterServer {
	if in == nil {
		return nil
	}
	out := new(ManagementClusterServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.ConfigMetadata != nil {
		in, out := &in.ConfigMetadata, &out.ConfigMetadata
		*out = new(ConfigMetadata)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIDiscovery) DeepCopyInto(out *OCIDiscovery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIDiscovery.
func (in *OCIDiscovery) DeepCopy() *OCIDiscovery {
	if in == nil {
		return nil
	}
	out := new(OCIDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginDiscovery) DeepCopyInto(out *PluginDiscovery) {
	*out = *in
	if in.GCP != nil {
		in, out := &in.GCP, &out.GCP
		*out = new(GCPDiscovery)
		**out = **in
	}
	if in.OCI != nil {
		in, out := &in.OCI, &out.OCI
		*out = new(OCIDiscovery)
		**out = **in
	}
	if in.REST != nil {
		in, out := &in.REST, &out.REST
		*out = new(GenericRESTDiscovery)
		**out = **in
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(KubernetesDiscovery)
		**out = **in
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalDiscovery)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginDiscovery.
func (in *PluginDiscovery) DeepCopy() *PluginDiscovery {
	if in == nil {
		return nil
	}
	out := new(PluginDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginRepository) DeepCopyInto(out *PluginRepository) {
	*out = *in
	if in.GCPPluginRepository != nil {
		in, out := &in.GCPPluginRepository, &out.GCPPluginRepository
		*out = new(GCPPluginRepository)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginRepository.
func (in *PluginRepository) DeepCopy() *PluginRepository {
	if in == nil {
		return nil
	}
	out := new(PluginRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	if in.GlobalOpts != nil {
		in, out := &in.GlobalOpts, &out.GlobalOpts
		*out = new(GlobalServer)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagementClusterOpts != nil {
		in, out := &in.ManagementClusterOpts, &out.ManagementClusterOpts
		*out = new(ManagementClusterServer)
		**out = **in
	}
	if in.DiscoverySources != nil {
		in, out := &in.DiscoverySources, &out.DiscoverySources
		*out = make([]PluginDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

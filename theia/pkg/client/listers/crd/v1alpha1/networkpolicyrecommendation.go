// Copyright 2022 Antrea Authors
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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "antrea.io/theia/pkg/apis/crd/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkPolicyRecommendationLister helps list NetworkPolicyRecommendations.
// All objects returned here must be treated as read-only.
type NetworkPolicyRecommendationLister interface {
	// List lists all NetworkPolicyRecommendations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkPolicyRecommendation, err error)
	// NetworkPolicyRecommendations returns an object that can list and get NetworkPolicyRecommendations.
	NetworkPolicyRecommendations(namespace string) NetworkPolicyRecommendationNamespaceLister
	NetworkPolicyRecommendationListerExpansion
}

// networkPolicyRecommendationLister implements the NetworkPolicyRecommendationLister interface.
type networkPolicyRecommendationLister struct {
	indexer cache.Indexer
}

// NewNetworkPolicyRecommendationLister returns a new NetworkPolicyRecommendationLister.
func NewNetworkPolicyRecommendationLister(indexer cache.Indexer) NetworkPolicyRecommendationLister {
	return &networkPolicyRecommendationLister{indexer: indexer}
}

// List lists all NetworkPolicyRecommendations in the indexer.
func (s *networkPolicyRecommendationLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkPolicyRecommendation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkPolicyRecommendation))
	})
	return ret, err
}

// NetworkPolicyRecommendations returns an object that can list and get NetworkPolicyRecommendations.
func (s *networkPolicyRecommendationLister) NetworkPolicyRecommendations(namespace string) NetworkPolicyRecommendationNamespaceLister {
	return networkPolicyRecommendationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkPolicyRecommendationNamespaceLister helps list and get NetworkPolicyRecommendations.
// All objects returned here must be treated as read-only.
type NetworkPolicyRecommendationNamespaceLister interface {
	// List lists all NetworkPolicyRecommendations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkPolicyRecommendation, err error)
	// Get retrieves the NetworkPolicyRecommendation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NetworkPolicyRecommendation, error)
	NetworkPolicyRecommendationNamespaceListerExpansion
}

// networkPolicyRecommendationNamespaceLister implements the NetworkPolicyRecommendationNamespaceLister
// interface.
type networkPolicyRecommendationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkPolicyRecommendations in the indexer for a given namespace.
func (s networkPolicyRecommendationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkPolicyRecommendation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkPolicyRecommendation))
	})
	return ret, err
}

// Get retrieves the NetworkPolicyRecommendation from the indexer for a given namespace and name.
func (s networkPolicyRecommendationNamespaceLister) Get(name string) (*v1alpha1.NetworkPolicyRecommendation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkpolicyrecommendation"), name)
	}
	return obj.(*v1alpha1.NetworkPolicyRecommendation), nil
}

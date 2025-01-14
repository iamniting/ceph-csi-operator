/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	policyv1 "k8s.io/api/policy/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// PodDisruptionBudgetLister helps list PodDisruptionBudgets.
// All objects returned here must be treated as read-only.
type PodDisruptionBudgetLister interface {
	// List lists all PodDisruptionBudgets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*policyv1.PodDisruptionBudget, err error)
	// PodDisruptionBudgets returns an object that can list and get PodDisruptionBudgets.
	PodDisruptionBudgets(namespace string) PodDisruptionBudgetNamespaceLister
	PodDisruptionBudgetListerExpansion
}

// podDisruptionBudgetLister implements the PodDisruptionBudgetLister interface.
type podDisruptionBudgetLister struct {
	listers.ResourceIndexer[*policyv1.PodDisruptionBudget]
}

// NewPodDisruptionBudgetLister returns a new PodDisruptionBudgetLister.
func NewPodDisruptionBudgetLister(indexer cache.Indexer) PodDisruptionBudgetLister {
	return &podDisruptionBudgetLister{listers.New[*policyv1.PodDisruptionBudget](indexer, policyv1.Resource("poddisruptionbudget"))}
}

// PodDisruptionBudgets returns an object that can list and get PodDisruptionBudgets.
func (s *podDisruptionBudgetLister) PodDisruptionBudgets(namespace string) PodDisruptionBudgetNamespaceLister {
	return podDisruptionBudgetNamespaceLister{listers.NewNamespaced[*policyv1.PodDisruptionBudget](s.ResourceIndexer, namespace)}
}

// PodDisruptionBudgetNamespaceLister helps list and get PodDisruptionBudgets.
// All objects returned here must be treated as read-only.
type PodDisruptionBudgetNamespaceLister interface {
	// List lists all PodDisruptionBudgets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*policyv1.PodDisruptionBudget, err error)
	// Get retrieves the PodDisruptionBudget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*policyv1.PodDisruptionBudget, error)
	PodDisruptionBudgetNamespaceListerExpansion
}

// podDisruptionBudgetNamespaceLister implements the PodDisruptionBudgetNamespaceLister
// interface.
type podDisruptionBudgetNamespaceLister struct {
	listers.ResourceIndexer[*policyv1.PodDisruptionBudget]
}

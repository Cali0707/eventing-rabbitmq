/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/rabbitmq/messaging-topology-operator/api/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UserLister helps list Users.
// All objects returned here must be treated as read-only.
type UserLister interface {
	// List lists all Users in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.User, err error)
	// Users returns an object that can list and get Users.
	Users(namespace string) UserNamespaceLister
	UserListerExpansion
}

// userLister implements the UserLister interface.
type userLister struct {
	indexer cache.Indexer
}

// NewUserLister returns a new UserLister.
func NewUserLister(indexer cache.Indexer) UserLister {
	return &userLister{indexer: indexer}
}

// List lists all Users in the indexer.
func (s *userLister) List(selector labels.Selector) (ret []*v1alpha2.User, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.User))
	})
	return ret, err
}

// Users returns an object that can list and get Users.
func (s *userLister) Users(namespace string) UserNamespaceLister {
	return userNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserNamespaceLister helps list and get Users.
// All objects returned here must be treated as read-only.
type UserNamespaceLister interface {
	// List lists all Users in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.User, err error)
	// Get retrieves the User from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.User, error)
	UserNamespaceListerExpansion
}

// userNamespaceLister implements the UserNamespaceLister
// interface.
type userNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Users in the indexer for a given namespace.
func (s userNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.User, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.User))
	})
	return ret, err
}

// Get retrieves the User from the indexer for a given namespace and name.
func (s userNamespaceLister) Get(name string) (*v1alpha2.User, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("user"), name)
	}
	return obj.(*v1alpha2.User), nil
}
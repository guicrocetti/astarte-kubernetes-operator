/*
This file is part of Astarte.

Copyright 2020-25 SECO Mind Srl.

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

package reconcile

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	apiv2alpha1 "github.com/astarte-platform/astarte-kubernetes-operator/api/api/v2alpha1"
	"github.com/astarte-platform/astarte-kubernetes-operator/internal/misc"
)

func GetAstarteClusteredServicePolicyRules() []rbacv1.PolicyRule {
	// This is needed for Astarte > 1.2.0, as DUP/AppEngine/VerneMQ are clustered using Erlang.
	return []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{"pods", "endpoints"},
			Verbs:     []string{"list", "get"},
		},
	}
}

// EnsureErlangClusteringRBAC makes sure that the Service Account used by an Erlang-clustered
// Astarte component can access the resources it needs.
func EnsureErlangClusteringRBAC(cr *apiv2alpha1.Astarte, component apiv2alpha1.AstarteComponent, c client.Client, scheme *runtime.Scheme) error {
	roleName := cr.Name + "-" + component.DashedString()
	roleBindingName := cr.Name + "-" + component.DashedString()
	serviceAccountName := cr.Name + "-" + component.DashedString()

	// Let's reconcile the Role first
	role := &rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: roleName, Namespace: cr.Namespace}}
	result, err := controllerutil.CreateOrUpdate(context.TODO(), c, role, func() error {
		if err := controllerutil.SetControllerReference(cr, role, scheme); err != nil {
			return err
		}

		role.Rules = GetAstarteClusteredServicePolicyRules()

		return nil
	})
	if err != nil {
		return err
	}
	misc.LogCreateOrUpdateOperationResult(log, result, cr, role)

	// And then the Role Binding
	roleBinding := &rbacv1.RoleBinding{ObjectMeta: metav1.ObjectMeta{Name: roleBindingName, Namespace: cr.Namespace}}
	result, err = controllerutil.CreateOrUpdate(context.TODO(), c, roleBinding, func() error {
		if err := controllerutil.SetControllerReference(cr, roleBinding, scheme); err != nil {
			return err
		}

		roleBinding.RoleRef = rbacv1.RoleRef{APIGroup: "rbac.authorization.k8s.io", Kind: "Role", Name: roleName}
		roleBinding.Subjects = []rbacv1.Subject{
			{Kind: "ServiceAccount", Name: serviceAccountName, Namespace: cr.Namespace},
		}

		return nil
	})
	if err != nil {
		return err
	}
	misc.LogCreateOrUpdateOperationResult(log, result, cr, roleBinding)

	return nil
}

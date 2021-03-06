// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package component

import (
	"context"
)

// Deployer is used to control the life-cycle of a component.
type Deployer interface {
	// Deploy a component.
	Deploy(ctx context.Context) error
	// Destroy already deployed component.
	Destroy(ctx context.Context) error
}

// Waiter waits for life-cycle operations of a component to finish.
type Waiter interface {
	// Wait for deployment to finish and component to report ready.
	Wait(ctx context.Context) error
	// WaitCleanup for destruction to finish and component to be fully removed.
	WaitCleanup(ctx context.Context) error
}

// DeployWaiter controls and waits for life-cycle operations of a component.
type DeployWaiter interface {
	Deployer
	Waiter
}

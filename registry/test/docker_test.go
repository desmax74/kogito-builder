/*
Copyright 2022.
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
package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistryWithDocker(t *testing.T) {

	d := Docker{}
	defer d.RemoveRegistryContainerAndImage()
	assert.Truef(t, d.StartRegistry(), "Registry not started")
	assert.Truef(t, d.IsRegistryImagePresent(), "Registry image not present")
	assert.Truef(t, d.IsRegistryRunning(), "Registry container not running")
	assert.Truef(t, d.StopRegistry(), "Registry not stopped")
}
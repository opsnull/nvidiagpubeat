/***********************************************************************
This module was automatically generated using the framework found below:
https://www.elastic.co/guide/en/beats/devguide/current/new-beat.html

Modifications to auto-generated code - Copyright 2018 eBay Inc.
Architect/Developer: Deepak Vasthimal

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
************************************************************************/

package config

import (
	"testing"
	"time"
)

func Test_Config(t *testing.T) {

	if DefaultConfig.Env != "local" {
		t.Errorf("Expected %s, Actual %s", "local", DefaultConfig.Env)
	}

	if DefaultConfig.Period != 1*time.Second {
		t.Errorf("Expected %d, Actual %d", 1*time.Second, DefaultConfig.Period)
	}
}

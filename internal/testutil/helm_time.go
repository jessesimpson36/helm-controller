/*
Copyright 2022 The Flux authors

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

package testutil

import (
	"time"

	helmtime "github.com/jessesimpson36/helm/v4/pkg/time"
)

// MustParseHelmTime parses a string into a Helm time.Time, panicking if it
// fails.
func MustParseHelmTime(t string) helmtime.Time {
	res, err := helmtime.Parse(time.RFC3339, t)
	if err != nil {
		panic(err)
	}
	return res
}

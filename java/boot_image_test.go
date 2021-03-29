// Copyright (C) 2021 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package java

import (
	"testing"

	"android/soong/android"
	"android/soong/dexpreopt"
)

// Contains some simple tests for boot_image logic, additional tests can be found in
// apex/boot_image_test.go as the ART boot image requires modules from the ART apex.

var prepareForTestWithBootImage = android.GroupFixturePreparers(
	PrepareForTestWithJavaDefaultModules,
	dexpreopt.PrepareForTestByEnablingDexpreopt,
)

func TestUnknownBootImage(t *testing.T) {
	prepareForTestWithBootImage.
		ExtendWithErrorHandler(android.FixtureExpectsAtLeastOneErrorMatchingPattern(
			`\Qimage_name: Unknown image name "unknown", expected one of art, boot\E`)).
		RunTestWithBp(t, `
			boot_image {
				name: "unknown-boot-image",
				image_name: "unknown",
			}
		`)
}

func TestUnknownBootclasspathFragmentImageName(t *testing.T) {
	prepareForTestWithBootImage.
		ExtendWithErrorHandler(android.FixtureExpectsAtLeastOneErrorMatchingPattern(
			`\Qimage_name: Unknown image name "unknown", expected one of art, boot\E`)).
		RunTestWithBp(t, `
			bootclasspath_fragment {
				name: "unknown-boot-image",
				image_name: "unknown",
			}
		`)
}

func TestUnknownPrebuiltBootImage(t *testing.T) {
	prepareForTestWithBootImage.
		ExtendWithErrorHandler(android.FixtureExpectsAtLeastOneErrorMatchingPattern(
			`\Qimage_name: Unknown image name "unknown", expected one of art, boot\E`)).
		RunTestWithBp(t, `
			prebuilt_boot_image {
				name: "unknown-boot-image",
				image_name: "unknown",
			}
		`)
}

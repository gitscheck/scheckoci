// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"
)

// issue-routing-tag: terraform/default
func TestUnitSafe_splitSizeToOffsetsAndLimits(t *testing.T) {

	offsets, _, _ := splitSizeToOffsetsAndLimits(defaultFilePartSize*8 + 1)
	if len(offsets) != 9 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), defaultFilePartSize*8+1)
		return
	}

	offsets, _, _ = splitSizeToOffsetsAndLimits(defaultFilePartSize * 7)
	if len(offsets) != 7 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), defaultFilePartSize*7)
		return
	}

	offsets, _, _ = splitSizeToOffsetsAndLimits(defaultFilePartSize + 1)
	if len(offsets) != 2 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), defaultFilePartSize+1)
		return
	}

	offsets, _, _ = splitSizeToOffsetsAndLimits(defaultFilePartSize)
	if len(offsets) != 1 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), defaultFilePartSize)
		return
	}

	offsets, _, _ = splitSizeToOffsetsAndLimits(defaultFilePartSize / 2)
	if len(offsets) != 1 {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), defaultFilePartSize/2)
		return
	}

	offsets, _, _ = splitSizeToOffsetsAndLimits(defaultFilePartSize * maxCount)
	if len(offsets) != int(maxCount) {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), defaultFilePartSize*maxCount)
		return
	}

	offsets, _, _ = splitSizeToOffsetsAndLimits(defaultFilePartSize*maxCount + 1)
	if len(offsets) != int(maxCount) {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), defaultFilePartSize*maxCount)
		return
	}

	offsets, _, _ = splitSizeToOffsetsAndLimits(defaultFilePartSize * maxCount * 2)
	if len(offsets) != int(maxCount) {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), defaultFilePartSize*maxCount*2)
		return
	}

	offsets, _, _ = splitSizeToOffsetsAndLimits(maxPartSize * maxCount)
	if len(offsets) != int(maxCount) {
		t.Errorf("The reported %v number of parts is wrong for the size %v", len(offsets), maxCount*maxPartSize)
		return
	}

	_, _, err := splitSizeToOffsetsAndLimits(maxPartSize*maxCount + 1)
	if err == nil {
		t.Errorf("The error should be returned for the too large file size")
		return
	}

	return
}

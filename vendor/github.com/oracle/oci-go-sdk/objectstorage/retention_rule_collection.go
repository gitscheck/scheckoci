// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RetentionRuleCollection Retention rule collection.
type RetentionRuleCollection struct {

	// An array of retention rule summaries.
	Items []RetentionRuleSummary `mandatory:"true" json:"items"`
}

func (m RetentionRuleCollection) String() string {
	return common.PointerString(m)
}

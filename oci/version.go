// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"log"
)

const Version = "4.41.0"
const ReleaseDate = "2021-08-25"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}

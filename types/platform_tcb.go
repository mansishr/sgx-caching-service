/*
 * Copyright (C) 2020 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package types

import (
	"time"
)

// PlatformTcb struct is the database schema for platform_tcbs table
type PlatformTcb struct {
	QeId        string    `json:"-" gorm:"primary_key"`
	PceId       string    `json:"-"`
	CpuSvn      string    `json:"-"`
	PceSvn      string    `json:"-"`
	Tcbm        string    `json:"-"`
	CreatedTime time.Time `json:"-"`
	UpdatedTime time.Time `json:"-"`
}

type PlatformTcbs []PlatformTcb

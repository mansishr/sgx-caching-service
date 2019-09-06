/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package repository

import "intel/isecl/sgx-caching-service/types"

type FmspcTcbInfoRepository interface {
	Create(types.FmspcTcbInfo) (*types.FmspcTcbInfo, error)
	Retrieve(types.FmspcTcbInfo) (*types.FmspcTcbInfo, error)
	RetrieveAll(user types.FmspcTcbInfo) (types.FmspcTcbInfos, error)
	Update(types.FmspcTcbInfo) error
	Delete(types.FmspcTcbInfo) error
}

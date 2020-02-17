/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package resource

import (

        "intel/isecl/sgx-caching-service/config"
        "intel/isecl/sgx-caching-service/repository"
        "intel/isecl/sgx-caching-service/types"
	"github.com/pkg/errors"
)

func GetLazyCachePlatformInfo( db repository.SCSDatabase, encryptedPPIDType string, cpuSvnType string, 
			PceSvnType string, pceIdType string, qeIdType string) (*types.Platform, error ) {
	log.Trace("resource/lazy_cache_ops.go:GetLazyCachePlatformInfo() Entering")
	defer log.Trace("resource/lazy_cache_ops.go:GetLazyCachePlatformInfo() Leaving")

	var data SgxData
	data.PlatformInfo.EncryptedPPID = encryptedPPIDType
	data.PlatformInfo.CpuSvn = cpuSvnType
	data.PlatformInfo.PceSvn = PceSvnType
	data.PlatformInfo.PceId = pceIdType
	data.PlatformInfo.QeId = qeIdType

        err := FetchPCKCertInfo(&data)
	if err != nil {
		return nil, errors.New("FetchPCKCertInfo:" + err.Error())
	}

	err = CachePlatformInfo(db, &data)
	if err != nil {
		return nil, errors.New("CachePlatformInfo:" + err.Error())
	}
	log.Warn("CachePlatformInfo done")

	err = CachePckCertChainInfo(db, &data)
	if err != nil {
		return nil, errors.New("CachePckCertChainInfo:" + err.Error())
	}
	log.Warn("CachePckCertChainInfo done")

	err = CachePckCertInfo(db, &data)
	if err != nil {
		return nil, errors.New("CachePckCertInfo:" + err.Error())
	}
	log.Warn("CachePckCertInfo done")
	
	log.Debug("GetLazyCachePlatformInfo fetch and cache operation completed successfully")
	return data.Platform, nil
}

func GetLazyCacheFmspcTcbInfo(db repository.SCSDatabase, fmspcType string) ( *types.FmspcTcbInfo, error ) {
	log.Trace("resource/lazy_cache_ops.go:GetLazyCacheFmspcTcbInfo() Entering")
	defer log.Trace("resource/lazy_cache_ops.go:GetLazyCacheFmspcTcbInfo() Leaving")

	var data SgxData
	data.FmspcTcbInfo.Fmspc = fmspcType

	err := FetchFmspcTcbInfo(&data)
	if err != nil {
		return nil, errors.New("FetchFmspcTcbInfo:" + err.Error())
	}

	err = CacheFmspcTcbInfo(db, &data)
	if err != nil {
		return nil, errors.New("CacheFmspcTcbInfo:" + err.Error())
	}

	log.Debug("GetLazyCacheFmspcTcbInfo fetch and cache operation completed successfully")
	return data.FmspcTcb, nil
}

func GetLazyCachePckCrl(db repository.SCSDatabase, CaType string) ( *types.PckCrl, error ) {
	log.Trace("resource/lazy_cache_ops.go:GetLazyCachePckCrl() Entering")
	defer log.Trace("resource/lazy_cache_ops.go:GetLazyCachePckCrl() Leaving")

	var data SgxData
	data.PlatformInfo.Ca = CaType

	err := FetchPCKCRLInfo(&data)
	if err != nil {
		return nil, errors.New("FetchPCKCRLInfo:" + err.Error())
	}

	err = CachePckCRLInfo(db, &data)
	if err != nil {
		return nil, errors.New("CachePckCRLInfo:" + err.Error())
	}
 	
	log.Debug("GetLazyCachePckCrl fetch and cache operation completed successfully")
	return data.PckCrl, nil
}

func GetLazyCacheQEIdentityInfo(db repository.SCSDatabase) ( types.QEIdentities, error ) {
	log.Trace("resource/lazy_cache_ops.go:GetLazyCacheQEIdentityInfo() Entering")
	defer log.Trace("resource/lazy_cache_ops.go:GetLazyCacheQEIdentityInfo() Leaving")

	var data SgxData

	err := FetchQEIdentityInfo(&data)
	if err != nil {
		return nil, errors.New("FetchQEIdentityInfo:" + err.Error())
	}

	err = CacheQEIdentityInfo(db, &data)
	if err != nil {
		return nil, errors.New("CacheQEIdentityInfo:" + err.Error())
	}

	existingQeInfo, err := db.QEIdentityRepository().RetrieveAll()
	if existingQeInfo == nil && err == nil {
		return nil, errors.New("GetLazyCacheQEIdentityInfo: Retrive data error" +  err.Error() )
	}

	log.Debug("GetLazyCacheQEIdentityInfo fetch and cache operation completed successfully")
	return existingQeInfo, nil
}

func GetCacheModel() ( int, error ) {
	log.Trace("resource/lazy_cache_ops.go:GetCacheModel() Entering")
	defer log.Trace("resource/lazy_cache_ops.go:GetCacheModel() Leaving")

	conf := config.Global()
	if conf == nil {
		return 0, errors.New("GetLazyCacheModel Configuration pointer is null")
        }
	
	log.Debug("Caching Model is: ",conf.CachingModel)	
	return conf.CachingModel, nil
}

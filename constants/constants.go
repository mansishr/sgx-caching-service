/*
 * Copyright (C) 2020 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package constants

import (
	"time"
)

const (
	HomeDir                        = "/opt/scs/"
	ConfigDir                      = "/etc/scs/"
	ExecLinkPath                   = "/usr/bin/scs"
	RunDirPath                     = "/run/scs"
	LogDir                         = "/var/log/scs/"
	LogFile                        = LogDir + "scs.log"
	SecLogFile                     = LogDir + "scs-security.log"
	HTTPLogFile                    = LogDir + "http.log"
	ConfigFile                     = "config.yml"
	DefaultTLSCertFile             = ConfigDir + "tls-cert.pem"
	DefaultTLSKeyFile              = ConfigDir + "tls.key"
	TrustedJWTSigningCertsDir      = ConfigDir + "certs/trustedjwt/"
	TrustedCAsStoreDir             = ConfigDir + "certs/trustedca/"
	ServiceRemoveCmd               = "systemctl disable scs"
	DefaultSSLCertFilePath         = ConfigDir + "scsdbcert.pem"
	ServiceName                    = "SCS"
	HostDataUpdaterGroupName       = "HostDataUpdater"
	HostDataReaderGroupName        = "HostDataReader"
	CacheManagerGroupName          = "CacheManager"
	SCSUserName                    = "scs"
	DefaultHttpsPort               = 9000
	DefaultKeyAlgorithm            = "rsa"
	DefaultKeyAlgorithmLength      = 3072
	DefaultScsTlsSan               = "127.0.0.1,localhost"
	DefaultScsTlsCn                = "SCS TLS Certificate"
	DefaultIntelProvServerURL      = "https://sbx.api.trustedservices.intel.com/sgx/certification/v3/"
	EncPPID_Key                    = "encrypted_ppid"
	CpuSvn_Key                     = "cpu_svn"
	PceSvn_Key                     = "pce_svn"
	PceId_Key                      = "pce_id"
	QeId_Key                       = "qe_id"
	Ca_Key                         = "ca"
	Type_Key                       = "type"
	Ca_Processor                   = "processor"
	Fmspc_Key                      = "fmspc"
	DefaultScsRefreshHours         = 24
	DefaultJwtValidateCacheKeyMins = 60
	SCSLogLevel                    = "SCS_LOGLEVEL"
	DefaultReadTimeout             = 30 * time.Second
	DefaultReadHeaderTimeout       = 10 * time.Second
	DefaultWriteTimeout            = 10 * time.Second
	DefaultIdleTimeout             = 10 * time.Second
	DefaultMaxHeaderBytes          = 1 << 20
	DefaultLogEntryMaxLength       = 300
	Type_Refresh_Cert              = "certs"
	Type_Refresh_Tcb               = "tcbs"
	MaxTcbLevels                   = 16
	DefaultRetrycount              = 3
	DefaultWaitTime                = 1
)

// State represents whether or not a daemon is running or not
type State bool
type CacheType int

const (
	CacheInsert = iota + 1
	CacheRefresh
)

const (
	// Stopped is the default nil value, indicating not running
	Stopped State = false
	// Running means the daemon is active
	Running State = true
)

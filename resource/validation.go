/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
 package resource

 import (
	 "regexp"
	 "intel/isecl/scs/constants"
 )

var regExMap = map[string]*regexp.Regexp{
				constants.EncPPID_Key: regexp.MustCompile(`^[0-9a-fA-F]{768}$`),
				constants.CpuSvn_Key: regexp.MustCompile(`^[0-9a-fA-F]{32}$`),
				constants.PceSvn_Key: regexp.MustCompile(`^[0-9a-fA-F]{4}$`),
				constants.PceId_Key: regexp.MustCompile(`^[0-9a-fA-F]{4}$`),
				constants.Ca_Key: regexp.MustCompile(`^(processor)$`),
				constants.Type_Key: regexp.MustCompile(`^(certs)$`),
				constants.Fmspc_Key: regexp.MustCompile(`^[0-9a-fA-F]{12}$`),
				constants.QeId_Key: regexp.MustCompile(`^[0-9a-fA-F]{32}$`)}

func ValidateInputString(key string, inString string) bool {
	regEx := regExMap[key]
	if len(key)<=0 || !regEx.MatchString(inString) {
		log.WithField(key, inString).Error("Input Validation")
		return false
	}
	return true
}

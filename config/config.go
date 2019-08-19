package config

import "time"

const (
	mainNetUrl             = "https://mainnet.infura.io"
	version                = "v3"
	projectId              = "189d58fe65f14de6bbe6cb78c0ce8aea"
	FullMainNetPath        = mainNetUrl + "/" + version + "/" + projectId
	DefaultRequestsTimeout = 2 * time.Second
	// cache config
	CacheSize                = 100 * 1024 * 1024
	CacheExpireTime   = time.Minute
	CacheUpdateLastBlockTime = time.Minute
	// service setting
	DefaultAddr = ":8123"

	// configure usage limits by source

)

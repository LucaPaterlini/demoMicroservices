package config

import "time"

const (
	mainNetURL = "https://mainnet.infura.io"
	version    = "v3"
	projectID  = "189d58fe65f14de6bbe6cb78c0ce8aea"
	// FullMainNetPath contain the main url of the path to call to access the 3rd party api.
	FullMainNetPath = mainNetURL + "/" + version + "/" + projectID
	// DefaultRequestsTimeout contains the timeout time for the request of the 3rd party api.
	DefaultRequestsTimeout = 2 * time.Second
	//CacheSize size of the cache to use to store the api call to the 3rd party api
	CacheSize = 100 * 1024 * 1024
	//CacheExpireTime the ttl of the api calls cached
	CacheExpireTime = time.Minute
	//CacheUpdateLastBlockTime the ticker to update the value of the last block of the eth chain"
	CacheUpdateLastBlockTime = time.Minute
	// DefaultAddr contains the default address to bind to run the api server.
	DefaultAddr = ":8123"
)

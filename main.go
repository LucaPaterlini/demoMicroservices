package main

import (
	"flag"
	"github.com/LucaPaterlini/infura/API"
	"github.com/LucaPaterlini/infura/config"
	"github.com/LucaPaterlini/infura/middlewares/limit"
	"github.com/LucaPaterlini/infura/middlewares/logger"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	cache "github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
	"log"
	"net/http"
	"os"
	"time"
)

var limiterActive = flag.Bool("limiter", false, "activate limiter filter")

func main() {
	// declaring the routes
	router := mux.NewRouter().PathPrefix("/v1/").Subrouter()
	router.HandleFunc("/block/{blockId:[0-9]+}", API.GetBlockHandler).Methods(http.MethodGet)
	router.HandleFunc("/tx/{blockId:[0-9]+}/{txId:[0-9]+}", API.GetTransactionHandler).Methods(http.MethodGet)

	// allowing cors
	router.Use(mux.CORSMethodMiddleware(router))
	// prepare the limiter middleware
	accessLimit := limit.Visitors{
		CleanupRefreshTime: 10 * time.Second,
		CleanupExpiry:      time.Minute,
		// allow 10 new request for each user each second
		R: 10,
		B: 15,
	}

	// allocate the memory for caching (

	memcached, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LRU),
		memory.AdapterWithCapacity(config.CacheSize),
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(memcached),
		cache.ClientWithTTL(config.CacheExpireTime),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// add handler compression
	handler := handlers.CompressHandler(router)

	// add Panic Logger middleware
	handler = logger.LogRequestPanic(handler)

	// add http response caching
	handler = cacheClient.Middleware(handler)

	// add Request Logger middleware
	handler = logger.LogRequest(handler)

	srv := &http.Server{
		Addr:         config.DefaultAddr,
		WriteTimeout: time.Minute * 10,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		// log the requests compress the handler (its safe as its now no user input data or tls),
		// limit the access for each user
		Handler: accessLimit.Limit(handler, *limiterActive),
	}

	log.Fatal(srv.ListenAndServe())
}

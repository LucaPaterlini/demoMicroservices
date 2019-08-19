// Package API provides the functions handlers.
package API

import (
	"fmt"
	"github.com/LucaPaterlini/infura/config"
	"github.com/LucaPaterlini/infura/dataCollection"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var lastBlock uint64

// UpdateRoutine updates the value of the last block atomically avery freq interval, with a set timeout.
func UpdateRoutine(freq, timeout time.Duration) {
	ticker := time.NewTicker(freq)
	// new request every config.DefaultRequestsTimeout time
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		confirm := 1
		for ; true; <-ticker.C {
			// retrieve the last block
			lastBlockTmp, err := dataCollection.GetLastBlockNumber(timeout)
			if confirm > 0 {
				confirm--
				wg.Done()
			}
			if err != nil {
				log.Println(err)
				continue
			}
			atomic.StoreUint64(&lastBlock, lastBlockTmp)
		}
	}()
	wg.Wait()
}

// init handle the creation of the cache and the update of the lastBlock
func init() {
	debug.SetGCPercent(10)
	UpdateRoutine(config.CacheUpdateLastBlockTime, config.DefaultRequestsTimeout)
}

// GetBlockHandler is the handler that manage the caching and execution of the GetBlock function that will contact
// the third party api in case its not able to satisfy a legit request.
func GetBlockHandler(w http.ResponseWriter, r *http.Request) {
	// update in case last update is newer
	vars := mux.Vars(r)
	// convert block index string to uint64
	blockID, _ := strconv.ParseUint(vars["blockId"], 10, 64)

	tmp := atomic.LoadUint64(&lastBlock)
	if blockID > tmp {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("requested id %d latest %d\n", blockID, lastBlock)
		log.Println(err.Error())
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	// retuning anything in the body regardless of any error code
	// it may contain
	_, _, body, _ := dataCollection.GetBlock(blockID, config.DefaultRequestsTimeout)
	writeResponse(body,&w)
}

// GetTransactionHandler is the handler that manage the caching and execution of the  GetTransaction function that will contact
//// the third party api in case its not able to satisfy a legit request.
func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// retrieve the parameters
	param := make(map[string]uint64)
	for _, key := range []string{"blockId", "txId"} {
		param[key], _ = strconv.ParseUint(vars["blockId"], 10, 64)
	}

	tmp := atomic.LoadUint64(&lastBlock)
	if param["blockId"] > tmp {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("requested id %d latest %d", param["blockId"], lastBlock)
		log.Println( err.Error())
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	// retuning anything in the body regardless of any error code
	// it may contain
	_, _, body, _ := dataCollection.GetTransaction(param["blockId"], param["txId"], config.DefaultRequestsTimeout)
	writeResponse(body,&w)
}

// writeResponse writes the response.
func writeResponse(body []byte,w *http.ResponseWriter){
	(*w).Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := (*w).Write(body)
	if err != nil {
		log.Println(err.Error())
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
}

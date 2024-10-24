package templates

import (
	"my5G-RANTester/config"
	"my5G-RANTester/internal/control_test_engine/gnb"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	log_time "my5G-RANTester/internal/analytics/log_time"
	"strconv"
)

func TestMultiUesDivision(numUes int, initialDelay int, delayStart int, numDivisor int, intervalDivision int, showAnalytics bool) {

	wg := sync.WaitGroup{}

	cfg, err := config.GetConfig()
	if err != nil {
		//return nil
		log.Fatal("Error in get configuration")
	} else if numUes > 100 {
		log.Fatal("Exceeded the limit of UEs")
	} else if numDivisor == 0 {
		log.Fatal("Divisor cannot be zero")
	}

	gnbid, err := strconv.Atoi(cfg.GNodeB.PlmnList.GnbId) // Parse gNB ID
	log_time.SetGnodebId(gnbid)                           // Set gNB ID
	log_time.ChangeAnalyticsState(showAnalytics)          // Enable/Disable analytics

	go gnb.InitGnb(cfg, &wg)

	wg.Add(1)

	time.Sleep(time.Duration(delayStart) * time.Second)
	multiplier := (100 * getHostNumber(getHostname()))
	msin := cfg.Ue.Msin

	delayPerUe := initialDelay

	for i := 1; i <= numUes; i++ {
		if i%intervalDivision == 1 && i > 1 {
			delayPerUe = initialDelay / numDivisor
		}

		go registerSingleUe(cfg, &wg, msin, i, multiplier)
		time.Sleep(time.Duration(delayPerUe) * time.Millisecond)
	}

	wg.Wait()
}

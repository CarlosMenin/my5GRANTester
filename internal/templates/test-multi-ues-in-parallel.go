package templates

import (
	"my5G-RANTester/config"
	"my5G-RANTester/internal/control_test_engine/gnb"
	"my5G-RANTester/internal/control_test_engine/ue"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	log_time "my5G-RANTester/internal/analytics/log_time"
	"strconv"
)

func TestMultiUesInParallel(numUes int, delayUes int, delayStart int, showAnalytics bool) {

	wg := sync.WaitGroup{}

	cfg, err := config.GetConfig()
	if err != nil {
		//return nil
		log.Fatal("Error in get configuration")
	} else if numUes > 100 {
		log.Fatal("Exceeded the limit of UEs")
	}

	gnbid, err := strconv.Atoi(cfg.GNodeB.PlmnList.GnbId) // Parse gNB ID
	log_time.SetGnodebId(gnbid)                           // Set gNB ID
	log_time.ChangeAnalyticsState(showAnalytics)          // Enable/Disable analytics

	go gnb.InitGnb(cfg, &wg)

	wg.Add(1)

	time.Sleep(time.Duration(delayStart) * time.Second)
	multiplier := (100 * getHostNumber(getHostname()))
	msin := cfg.Ue.Msin

	for i := 1; i <= numUes; i++ {
		go registerSingleUe(cfg, &wg, msin, i, multiplier)
		time.Sleep(time.Duration(delayUes) * time.Millisecond)
	}

	wg.Wait()
}

func registerSingleUe(cfg config.Config, wg *sync.WaitGroup, msin string, i int, multip int) {
	imsi := imsiGenerator(i, msin, multip)
	log.Info("[TESTER] TESTING REGISTRATION USING IMSI ", imsi, " UE")
	cfg.Ue.Msin = imsi
	log_time.LogUeTime(imsi, "StartRegistration")
	go ue.RegistrationUe(cfg, int64(i), wg)
	//wg.Add(1)
}

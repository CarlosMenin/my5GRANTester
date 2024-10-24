package templates

import (
	"fmt"
	"my5G-RANTester/config"
	"my5G-RANTester/internal/control_test_engine/gnb"
	"my5G-RANTester/internal/control_test_engine/ue"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestMultiUesInQueue(numUes int) {

	wg := sync.WaitGroup{}

	cfg, err := config.GetConfig()
	if err != nil {
		//return nil
		log.Fatal("Error in get configuration")
	} else if numUes > 100 {
		log.Fatal("Exceeded the limit of UEs")
	}

	go gnb.InitGnb(cfg, &wg)

	wg.Add(1)

	time.Sleep(1 * time.Second)
	multiplier := (100 * getHostNumber(getHostname()))
	msin := cfg.Ue.Msin
	for i := 1; i <= numUes; i++ {

		imsi := imsiGenerator(i, msin, multiplier)
		log.Info("[TESTER] TESTING REGISTRATION USING IMSI ", imsi, " UE")
		cfg.Ue.Msin = imsi
		go ue.RegistrationUe(cfg, int64(i), &wg)
		wg.Add(1)

		time.Sleep(10 * time.Second)
	}

	wg.Wait()
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Erro ao obter o nome do host:", err)
	}
	return hostname
}

func getHostNumber(hostname string) int {
	parts := strings.Split(hostname, "-")
	if len(parts) < 2 {
		log.Fatal("Nome do host inválido")
	}

	numberStr := strings.TrimSpace(parts[len(parts)-1])
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Fatal("Número do host inválido:", numberStr)
	}

	return number
}

func imsiGenerator(i int, msin string, multip int) string {

	msin_int, err := strconv.Atoi(msin)
	if err != nil {
		log.Fatal("Error in get configuration")
	}

	base := (msin_int + (i - 1) + multip)

	imsi := fmt.Sprintf("%010d", base)
	return imsi
}

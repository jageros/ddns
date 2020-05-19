package main

import (
	"ddns_pro/config"
	"ddns_pro/ddns"
	extip "ddns_pro/ext_ip"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	wait := &sync.WaitGroup{}
	wait.Add(1)
	go func() {
		checkDns()
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Kill, os.Interrupt, syscall.SIGINT)
		t := time.Tick(time.Second * time.Duration(config.CFG.CheckTime))
		for {
			select {
			case <-sig:
				log.Printf("Goroutine has exit !")
				wait.Done()
				return

			case <-t:
				checkDns()
			}
		}
	}()
	wait.Wait()
}

func checkDns() {
	ip, update := extip.GetExternalIP()
	if update {
		for _, subDomain := range config.CFG.SubDomains {
			ddns.SetDns(subDomain, ip)
			time.Sleep(time.Second)
		}
	}
}

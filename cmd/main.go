package main

import (
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
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Kill, os.Interrupt, syscall.SIGINT)
		t := time.Tick(time.Second * 60)
		for {
			select {
			case <-sig:
				log.Printf("Goroutine has exit !")
				wait.Done()
				return

			case <-t:
				ip, update := extip.GetExternalIP()
				if update {
					ddns.SetDns(ip)

				}
			}
		}
	}()
	wait.Wait()
}

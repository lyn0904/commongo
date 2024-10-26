package common

import (
	"os"
	"os/signal"
	"syscall"
)

func Blocking() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}

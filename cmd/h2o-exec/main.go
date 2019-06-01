package main

import (
	"time"
	"os/exec"
)

func main() {
	t := time.NewTicker(17*time.Minute)
	select {
	case <- t.C:
		startH2O()
	}
}

func startH2O() {
	cmd := exec.Command("bash", "-c", "sudo /home/pi/h2o/h2o -c /usr/local/etc/h2o/h2o.conf &")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
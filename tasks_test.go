package tasks

import (
	"testing"
	"time"

	"github.com/quartermeat/tasks/controller"
)

func TestRun(t *testing.T) {
	powerOn := make(chan bool)
	go controller.Run(powerOn)
	time.Sleep(1 * time.Second)
	powerOn <- false

}

package analyze

import (
	"forex/go-soft4fx/internal/simulator"
	"forex/go-soft4fx/internal/simulator/analyze/drawdown"
)

type Result struct {
	simulator *simulator.Simulator
	weekday   *Weekday
	dd        *drawdown.Result
}

func (r Result) Simulator() *simulator.Simulator {
	return r.simulator
}

func (r Result) Weekday() *Weekday {
	return r.weekday
}

func (r Result) Drawdown() *drawdown.Result {
	return r.dd
}

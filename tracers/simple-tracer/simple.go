package live

import (
	"encoding/json"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/log"
)

func init() {
	tracers.DefaultDirectory.Register("simpleTracer", newSimpleTracer, false)
}

type simpleTracer struct {
	result    map[string]any
	interrupt atomic.Bool // Atomic flag to signal execution interruption
	reason    error       // Textual reason for the interruption
}

func newSimpleTracer(_ *tracers.Context, _ json.RawMessage) (*tracers.Tracer, error) {
	t := &simpleTracer{result: make(map[string]any)}
	log.Info("Simple tracer initialized")
	return &tracers.Tracer{
		Hooks: &tracing.Hooks{
			OnTxStart: t.OnTxStart,
		},
		GetResult: t.GetResult,
		Stop:      t.Stop,
	}, nil
}

func (t *simpleTracer) OnTxStart(env *tracing.VMContext, tx *types.Transaction, from common.Address) {
	t.result["from"] = from
}

func (t *simpleTracer) GetResult() (json.RawMessage, error) {
	res, err := json.Marshal(t.result)
	if err != nil {
		return nil, err
	}
	return res, t.reason
}

// Stop terminates execution of the tracer at the first opportune moment.
func (t *simpleTracer) Stop(err error) {
	t.reason = err
	t.interrupt.Store(true)
}

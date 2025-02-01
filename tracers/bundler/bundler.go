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
	tracers.LiveDirectory.Register("bundler", newBundlerTracer)
}

type bundlerTracer struct {
	result    map[string]any
	interrupt atomic.Bool // Atomic flag to signal execution interruption
	reason    error       // Textual reason for the interruption
}

func newBundlerTracer(_ json.RawMessage) (*tracing.Hooks, error) {
	t := &bundlerTracer{result: make(map[string]any)}
	log.Info("Bundler tracer initialized")
	return &tracing.Hooks{
		OnTxStart: t.OnTxStart,
	}, nil
}

func (t *bundlerTracer) OnTxStart(env *tracing.VMContext, tx *types.Transaction, from common.Address) {
	t.result["from"] = from
}

func (t *bundlerTracer) GetResult() (json.RawMessage, error) {
	res, err := json.Marshal(t.result)
	if err != nil {
		return nil, err
	}
	return res, t.reason
}

// Stop terminates execution of the tracer at the first opportune moment.
func (t *bundlerTracer) Stop(err error) {
	t.reason = err
	t.interrupt.Store(true)
}

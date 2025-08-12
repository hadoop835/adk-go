package checkpoint

import (
	"errors"
	"iter"

	"google.golang.org/adk/session"
)

type Manager struct {
	InvocationID string
	Branch       string
}

type CheckpointStatus string

const (
	CheckpointStatusPending   CheckpointStatus = "pending"
	CheckpointStatusCompleted CheckpointStatus = "completed"
)

type Checkpoint struct {
	manager *Manager
	id      string
	name    string
	status  CheckpointStatus
}

type AppendFunc func(events ...*session.Event)

// checkpoint creates a new checkpoint and stores them.
// If an existing checkpoint is found, it is updated.
func (m *Manager) checkpoint() error {
	panic("not implemented")
}

func (m *Manager) load(name string) (*Checkpoint, iter.Seq2[*session.Event, error]) {
	panic("not implemented")
}

type RunStatus struct {
	err error
}

func (r *RunStatus) Error() error {
	return r.err
}

func (m *Manager) Run(op string, fn func(append AppendFunc) error) (iter.Seq2[*session.Event, error], RunStatus) {
	status := RunStatus{}
	if op == "" {
		status.err = errors.New("operation cannot be empty")
		return nil, status
	}

	resumedOp, resumedIt := m.load(op)
	if resumedOp.status == CheckpointStatusCompleted {
		return resumedIt, status
	}
	panic("not implemented")

	// var yield iter.Seq2[*session.Event, error]
	// // appender := newAppender(yield)

	// if err := fn(appender.appendFunc); err != nil {
	// 	status.err = err
	// 	return nil, status
	// }
	// if err := m.checkpoint(); err != nil {
	// 	status.err = err
	// }
	// TODO: Return the yielder
	return nil, status
}

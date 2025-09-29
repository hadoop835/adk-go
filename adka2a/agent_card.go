package adka2a

import (
	"github.com/a2aproject/a2a-go/a2a"
	"github.com/a2aproject/a2a-go/a2asrv"
	"google.golang.org/adk/agent"
)

var _ a2asrv.AgentCardProducer = (*CardProducer)(nil)

type CardProducer struct {
	Agent agent.Agent
}

func (cp *CardProducer) Card() *a2a.AgentCard {
	// TODO(yarolegovich): implement
	return &a2a.AgentCard{}
}

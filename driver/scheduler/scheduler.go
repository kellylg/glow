package scheduler

import (
	"sync"

	"github.com/chrislusf/glow/driver/scheduler/market"
	"github.com/chrislusf/glow/resource"
)

type Scheduler struct {
	Leader                    string
	EventChan                 chan interface{}
	resourceFetcher           *ResourceFetcher
	Market                    *market.Market
	datasetShard2Location     map[string]resource.Location
	datasetShard2LocationLock sync.Mutex
}

type SchedulerOption struct {
	DataCenter string
	Rack       string
}

func NewScheduler(leader string, option *SchedulerOption) *Scheduler {
	s := &Scheduler{
		Leader:                leader,
		EventChan:             make(chan interface{}),
		resourceFetcher:       NewResourceFetcher(),
		Market:                market.NewMarket(),
		datasetShard2Location: make(map[string]resource.Location),
	}
	s.InitMarket()
	return s
}

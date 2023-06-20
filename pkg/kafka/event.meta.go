package kafka



type SimEventStatus int

const (
	SimEventStatusRequested SimEventStatus = iota
	SimEventStatusPending
	SimEventStatusAllocated
	SimEventStatusFailed
	SimEventStatusFinished
)

var SimEventTopic = map[SimEventStatus]string{
	SimEventStatusRequested: "sim.requested",
	SimEventStatusPending:   "sim.pending",
	SimEventStatusAllocated: "sim.allocated",
	SimEventStatusFailed:    "sim.failed",
	SimEventStatusFinished:  "sim.finished",
}

var SimEventRollbackTopic = map[SimEventStatus]string{
	SimEventStatusRequested: "sim.requested_rollback",
	SimEventStatusPending:   "sim.pending_rollback",
	SimEventStatusAllocated: "sim.allocated_rollback",
	SimEventStatusFailed:    "sim.failed_rollback",
	SimEventStatusFinished:  "sim.finished_rollback",
}



type ValEventStatus int

const (
	ValEventStatusRequested ValEventStatus = iota
	ValEventStatusPending
	ValEventStatusAllocated
	ValEventStatusValidated
	ValEventStatusInvalidated
	ValEventStatusFailed	
	ValEventStatusFinished
)

var ValEventTopic = map[ValEventStatus]string{
	ValEventStatusRequested:   "val.requested",
	ValEventStatusPending:     "val.pending",
	ValEventStatusAllocated:   "val.allocated",
	ValEventStatusFailed:      "val.failed",
	ValEventStatusFinished:    "val.finished",
}

var ValEventRollbackTopic = map[ValEventStatus]string{
	ValEventStatusRequested:   "val.requested_rollback",
	ValEventStatusPending:     "val.pending_rollback",
	ValEventStatusAllocated:   "val.allocated_rollback",
	ValEventStatusFailed:      "val.failed_rollback",
	ValEventStatusFinished:    "val.finished_rollback",
}



type RealEventStatus int

const (
	RealEventStatusRequested RealEventStatus = iota
	RealEventStatusPending
	RealEventStatusAllocated
)

var RealEventTopic = map[RealEventStatus]string{
	RealEventStatusRequested: "real.requested",
	RealEventStatusPending:   "real.pending",
	RealEventStatusAllocated: "real.allocated",
}

var RealEventRollbackTopic = map[RealEventStatus]string{
	RealEventStatusRequested: "real.requested_rollback",
	RealEventStatusPending:   "real.pending_rollback",
	RealEventStatusAllocated: "real.allocated_rollback",
}
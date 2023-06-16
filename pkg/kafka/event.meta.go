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
	SimEventStatusRequested: "sim_requested",
	SimEventStatusPending:   "sim_pending",
	SimEventStatusAllocated: "sim_allocated",
	SimEventStatusFailed:    "sim_failed",
	SimEventStatusFinished:  "sim_finished",
}

var SimEventRollbackTopic = map[SimEventStatus]string{
	SimEventStatusRequested: "sim_requested_rollback",
	SimEventStatusPending:   "sim_pending_rollback",
	SimEventStatusAllocated: "sim_allocated_rollback",
	SimEventStatusFailed:    "sim_failed_rollback",
	SimEventStatusFinished:  "sim_finished_rollback",
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
	ValEventStatusRequested:   "val_requested",
	ValEventStatusPending:     "val_pending",
	ValEventStatusAllocated:   "val_allocated",
	ValEventStatusValidated:   "val_validated",
	ValEventStatusInvalidated: "val_invalidated",
	ValEventStatusFailed:      "val_failed",
	ValEventStatusFinished:    "val_finished",
}

var ValEventRollbackTopic = map[ValEventStatus]string{
	ValEventStatusRequested:   "val_requested_rollback",
	ValEventStatusPending:     "val_pending_rollback",
	ValEventStatusAllocated:   "val_allocated_rollback",
	ValEventStatusValidated:   "val_validated_rollback",
	ValEventStatusInvalidated: "val_invalidated_rollback",
	ValEventStatusFailed:      "val_failed_rollback",
	ValEventStatusFinished:    "val_finished_rollback",
}



type RealEventStatus int

const (
	RealEventStatusRequested RealEventStatus = iota
	RealEventStatusPending
	RealEventStatusAllocated
)

var RealEventTopic = map[RealEventStatus]string{
	RealEventStatusRequested: "real_requested",
	RealEventStatusPending:   "real_pending",
	RealEventStatusAllocated: "real_allocated",
}

var RealEventRollbackTopic = map[RealEventStatus]string{
	RealEventStatusRequested: "real_requested_rollback",
	RealEventStatusPending:   "real_pending_rollback",
	RealEventStatusAllocated: "real_allocated_rollback",
}
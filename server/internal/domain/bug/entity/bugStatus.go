package entity

type BugStatus string

const (
	BugStatusOpen    BugStatus = "Open"
	BugStatusClosed  BugStatus = "Closed"
	BugStatusPending BugStatus = "Pending"
)

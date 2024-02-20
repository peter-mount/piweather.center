package util

import "errors"

// VisitorStop is an error which causes the current step in a Visitor to stop processing.
// It's used to enable a Visitor to handle all processing of a node within itself rather
// than the Visitor proceeding to any child nodes of that node.
var VisitorStop = errors.New("visitor stop")

func IsVisitorStop(err error) bool {
	return err != nil && errors.Is(err, VisitorStop)
}

// VisitorExit is an error which will terminate the Visitor.
// This is the same as any error occurring within a Visitor except that the final error
// returned from specific handlers will become nil.
var VisitorExit = errors.New("visitor exit")

func IsVisitorExit(err error) bool {
	return err != nil && errors.Is(err, VisitorExit)
}

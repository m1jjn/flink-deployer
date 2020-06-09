package operations

import (
	"github.com/spf13/afero"
	"github.com/therako/flink-deployer/cmd/cli/flink"
)

// Operator is an interface which contains all the functionality
// that the deployer exposes
type Operator interface {
	Deploy(d Deploy) error
	Update(u UpdateJob) error
	RetrieveJobs() ([]flink.Job, error)
	Terminate(t TerminateJob) error
}

// RealOperator is the Operator used in the production code
type RealOperator struct {
	Filesystem   afero.Fs
	FlinkRestAPI flink.FlinkRestAPI
}

package service

import "fmt"

var (
	ErrNoSuchSource   = fmt.Errorf("no such source")
	ErrSubjectInvalid = fmt.Errorf("Subject Invalid")
)

package model

//go:generate stringer -type IDEStatus -linecomment
type IDEStatus int

const (
	_                             IDEStatus = iota + 99
	IDEStatusSucceeded                      // succeeded
	IDEStatusUnknownError                   // unknown error
	_                                       // deprecated: 102
	_                                       // deprecated: 103
	IDEStatusAPIFail                        // API request fail
	IDEStatusTokenInvalid                   // token invalid
	IDEStatusAPITimeout                     // API timeout
	IDEStatusScanDirInvalid                 // scan dir invalid
	IDEStatusLogFileCreationError           // log file creation error
	IDEStatusServerFail                     // server fail
	IDEStatusGeneralAPIError                // general API error
	IDEStatusTaskNotExists                  // task not exists
)

func (i IDEStatus) Error() string {
	return i.String()
}

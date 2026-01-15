package helpers

import "fmt"

type FactPulseError struct { Message string }
func (e *FactPulseError) Error() string { return e.Message }

type FactPulseAuthError struct { FactPulseError }
func NewFactPulseAuthError(msg string) *FactPulseAuthError { return &FactPulseAuthError{FactPulseError{Message: msg}} }

type FactPulsePollingTimeout struct { FactPulseError; TaskID string; Timeout int64 }
func NewFactPulsePollingTimeout(taskID string, timeout int64) *FactPulsePollingTimeout {
    return &FactPulsePollingTimeout{FactPulseError: FactPulseError{Message: fmt.Sprintf("Timeout (%dms) for task %s", timeout, taskID)}, TaskID: taskID, Timeout: timeout}
}

type ValidationErrorDetail struct { Level, Item, Reason, Source, Code string }
func (e ValidationErrorDetail) String() string {
    item := e.Item; if item == "" { item = "unknown" }
    reason := e.Reason; if reason == "" { reason = "Unknown error" }
    return fmt.Sprintf("[%s] %s", item, reason)
}

type FactPulseValidationError struct { FactPulseError; Errors []ValidationErrorDetail }
func NewFactPulseValidationError(msg string, errors []ValidationErrorDetail) *FactPulseValidationError {
    return &FactPulseValidationError{FactPulseError: FactPulseError{Message: msg}, Errors: errors}
}

// Package helpers fournit un client simplifié pour l'API FactPulse avec
// authentification JWT et polling intégrés.
package helpers

import (
	"fmt"
	"strings"
)

// FactPulseError est la classe de base pour toutes les erreurs FactPulse.
type FactPulseError struct {
	Message string
}

func (e *FactPulseError) Error() string {
	return e.Message
}

// FactPulseAuthError est levée lors d'erreurs d'authentification.
type FactPulseAuthError struct {
	FactPulseError
}

// NewFactPulseAuthError crée une nouvelle erreur d'authentification.
func NewFactPulseAuthError(message string) *FactPulseAuthError {
	return &FactPulseAuthError{FactPulseError{Message: message}}
}

// FactPulsePollingTimeout est levée lors d'un timeout de polling.
type FactPulsePollingTimeout struct {
	FactPulseError
	TaskID  string
	Timeout int64
}

// NewFactPulsePollingTimeout crée une nouvelle erreur de timeout.
func NewFactPulsePollingTimeout(taskID string, timeout int64) *FactPulsePollingTimeout {
	return &FactPulsePollingTimeout{
		FactPulseError: FactPulseError{
			Message: fmt.Sprintf("Timeout (%dms) atteint pour la tâche %s", timeout, taskID),
		},
		TaskID:  taskID,
		Timeout: timeout,
	}
}

// ValidationErrorDetail représente un détail d'erreur de validation.
type ValidationErrorDetail struct {
	Level  string `json:"level"`
	Item   string `json:"item"`
	Reason string `json:"reason"`
	Source string `json:"source,omitempty"`
	Code   string `json:"code,omitempty"`
}

func (e ValidationErrorDetail) String() string {
	item := e.Item
	if item == "" {
		item = "unknown"
	}
	reason := e.Reason
	if reason == "" {
		reason = "Unknown error"
	}
	return fmt.Sprintf("[%s] %s", item, reason)
}

// FactPulseValidationError est levée lors d'erreurs de validation.
type FactPulseValidationError struct {
	FactPulseError
	Errors []ValidationErrorDetail
}

// NewFactPulseValidationError crée une nouvelle erreur de validation.
func NewFactPulseValidationError(message string, errors []ValidationErrorDetail) *FactPulseValidationError {
	return &FactPulseValidationError{
		FactPulseError: FactPulseError{Message: message},
		Errors:         errors,
	}
}

func (e *FactPulseValidationError) Error() string {
	if len(e.Errors) > 0 {
		var details []string
		for _, err := range e.Errors {
			details = append(details, "  - "+err.String())
		}
		return e.Message + "\n\nDétails:\n" + strings.Join(details, "\n")
	}
	return e.Message
}

package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestValidateConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		email        string
		apiKey       string
		organization string
		wantErr      bool
		errContains  string
	}{
		{
			name:         "all values present",
			email:        "test@example.com",
			apiKey:       "test-api-key",
			organization: "test-org",
			wantErr:      false,
		},
		{
			name:         "missing email",
			email:        "",
			apiKey:       "test-api-key",
			organization: "test-org",
			wantErr:      true,
			errContains:  "email is not configured",
		},
		{
			name:         "missing api_key",
			email:        "test@example.com",
			apiKey:       "",
			organization: "test-org",
			wantErr:      true,
			errContains:  "api_key is not configured",
		},
		{
			name:         "missing organization",
			email:        "test@example.com",
			apiKey:       "test-api-key",
			organization: "",
			wantErr:      true,
			errContains:  "organization is not configured",
		},
		{
			name:         "all values missing",
			email:        "",
			apiKey:       "",
			organization: "",
			wantErr:      true,
			errContains:  "email is not configured",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset viper for each test
			viper.Reset()

			// Set test values
			viper.Set("email", tt.email)
			viper.Set("api_key", tt.apiKey)
			viper.Set("organization", tt.organization)

			// Call validateConfiguration with nil command and args (not used)
			err := validateConfiguration(&cobra.Command{}, []string{})

			if tt.wantErr {
				if err == nil {
					t.Errorf("validateConfiguration() expected error, got nil")
					return
				}
				if tt.errContains != "" && err.Error() != tt.errContains {
					t.Errorf("validateConfiguration() error = %q, want %q", err.Error(), tt.errContains)
				}
			} else {
				if err != nil {
					t.Errorf("validateConfiguration() unexpected error: %v", err)
				}
			}
		})
	}
}

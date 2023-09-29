package config

import (
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()

	if config.Consumer.Return.Errors != true {
		t.Errorf("Unexpected value for Return.Errors. got: %v, want: %v", config.Consumer.Return.Errors, true)
	}

	if config.Consumer.Fetch.Min != 1073741824 {
		t.Errorf("Unexpected value for Fetch.Min. got: %v, want: %v", config.Consumer.Fetch.Min, 1073741824)
	}

	if config.Consumer.Fetch.Default != 2147483647 {
		t.Errorf("Unexpected value for Fetch.Default. got: %v, want: %v", config.Consumer.Fetch.Default, 2147483647)
	}

	if config.Consumer.MaxWaitTime != 100*time.Millisecond {
		t.Errorf("Unexpected value for MaxWaitTime. got: %v, want: %v", config.Consumer.MaxWaitTime, 100*time.Millisecond)
	}
}

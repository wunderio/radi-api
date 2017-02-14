package operation

import (
	api_usage "github.com/wunderkraut/radi-api/usage"
)

const (
	// This property should be considered externally executable
	USAGE_EXTERNAL_EXEC = "external"
)

/**
 * Simple Usage constructors for Operations
 */

// Make a default usage map
func Usage_Default() api_usage.Usage {
	return api_usage.MakeUsageFromMap(map[string]bool{
		USAGE_EXTERNAL_EXEC: true,
	})
}

// Make a usage map for marking a propert as internally usable only.
func Usage_Internal() api_usage.Usage {
	return api_usage.MakeUsageFromMap(map[string]bool{
		USAGE_EXTERNAL_EXEC: true,
	})
}

// Make a usage map for giving external only external read only permissions
func Usage_External() api_usage.Usage {
	return api_usage.MakeUsageFromMap(map[string]bool{
		USAGE_EXTERNAL_EXEC: true,
	})
}

/**
 * Abstracted tests
 */

// Should this operations be made avaialble to external consumers
func IsUsage_External(usage api_usage.Usage) bool {
	return usage.Has(USAGE_EXTERNAL_EXEC)
}

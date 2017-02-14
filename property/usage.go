package property

import (
	api_usage "github.com/wunderkraut/radi-api/usage"
)

// A number of constants used to define an initial set of of possible usage request keys
const (
	// This property should be considered externally required to write before it"s operation is executed
	USAGE_EXTERNAL_REQUIRED = "external.required"
	// This property should be considered externally readable before it"s operation is executed
	USAGE_EXTERNAL_READ_BEFOREEXEC = "external.read.beforeexec"
	// This property should be considered externally writeable before it"s operation is executed
	USAGE_EXTERNAL_WRITE_BEFOREEXEC = "external.write.beforeexec"
	// This property should be considered externally readable after it"s operation is executed
	USAGE_EXTERNAL_READ_AFTEREXEC = "external.read.afterexec"
	// This property should be considered externally writeable after it"s operation is executed
	USAGE_EXTERNAL_WRITE_AFTEREEXEC = "external.write.afterexec"
	// This property should be considered internally readable before it"s operation is executed
	USAGE_INTERNAL_READ_BEFOREEXEC = "internal.read.beforeexec"
	// This property should be considered externally readable before it"s operation is executed
	USAGE_INTERNAL_WRITE_BEFOREEXEC = "internal.write.beforeexec"
	// This property should be considered internally readable before it"s operation is executed
	USAGE_INTERNAL_READ_AFTEREXEC = "internal.read.afterexec"
	// This property should be considered externally writeable before it"s operation is executed
	USAGE_INTERNAL_WRITE_AFTEREXEC = "internal.write.afterexec"
)

/**
 * Usage constructors
 *
 * The following functions can be used to give Usage interface objects
 * for various internal and external property usages
 */

// Defalt permissive Usage
func Usage_Default() api_usage.Usage {
	return api_usage.MakeUsageFromMap(map[string]bool{
		USAGE_EXTERNAL_REQUIRED:         false,
		USAGE_EXTERNAL_READ_BEFOREEXEC:  true,
		USAGE_EXTERNAL_WRITE_BEFOREEXEC: true,
		USAGE_EXTERNAL_READ_AFTEREXEC:   true,
		USAGE_EXTERNAL_WRITE_AFTEREEXEC: true,
		USAGE_INTERNAL_READ_BEFOREEXEC:  true,
		USAGE_INTERNAL_WRITE_BEFOREEXEC: true,
		USAGE_INTERNAL_READ_AFTEREXEC:   true,
		USAGE_INTERNAL_WRITE_AFTEREXEC:  true,
	})
}

// Internal only Usage
func Usage_Internal() api_usage.Usage {
	return api_usage.MakeUsageFromMap(map[string]bool{
		USAGE_EXTERNAL_REQUIRED:         false,
		USAGE_EXTERNAL_READ_BEFOREEXEC:  false,
		USAGE_EXTERNAL_WRITE_BEFOREEXEC: false,
		USAGE_EXTERNAL_READ_AFTEREXEC:   false,
		USAGE_EXTERNAL_WRITE_AFTEREEXEC: false,
		USAGE_INTERNAL_READ_BEFOREEXEC:  true,
		USAGE_INTERNAL_WRITE_BEFOREEXEC: true,
		USAGE_INTERNAL_READ_AFTEREXEC:   true,
		USAGE_INTERNAL_WRITE_AFTEREXEC:  true,
	})
}

// External write requried usage
func Usage_Required() api_usage.Usage {
	return api_usage.MakeUsageFromMap(map[string]bool{
		USAGE_EXTERNAL_REQUIRED:         true,
		USAGE_EXTERNAL_READ_BEFOREEXEC:  true,
		USAGE_EXTERNAL_WRITE_BEFOREEXEC: true,
		USAGE_EXTERNAL_READ_AFTEREXEC:   true,
		USAGE_EXTERNAL_WRITE_AFTEREEXEC: true,
		USAGE_INTERNAL_READ_BEFOREEXEC:  true,
		USAGE_INTERNAL_WRITE_BEFOREEXEC: true,
		USAGE_INTERNAL_READ_AFTEREXEC:   true,
		USAGE_INTERNAL_WRITE_AFTEREXEC:  true,
	})
}

// External write optional Usage
func Usage_Optional() api_usage.Usage {
	return api_usage.MakeUsageFromMap(map[string]bool{
		USAGE_EXTERNAL_REQUIRED:         false,
		USAGE_EXTERNAL_READ_BEFOREEXEC:  true,
		USAGE_EXTERNAL_WRITE_BEFOREEXEC: true,
		USAGE_EXTERNAL_READ_AFTEREXEC:   true,
		USAGE_EXTERNAL_WRITE_AFTEREEXEC: true,
		USAGE_INTERNAL_READ_BEFOREEXEC:  true,
		USAGE_INTERNAL_WRITE_BEFOREEXEC: true,
		USAGE_INTERNAL_READ_AFTEREXEC:   true,
		USAGE_INTERNAL_WRITE_AFTEREXEC:  true,
	})
}

// External read only Usage
func Usage_ReadOnly() api_usage.Usage {
	return api_usage.MakeUsageFromMap(map[string]bool{
		USAGE_EXTERNAL_REQUIRED:         false,
		USAGE_EXTERNAL_READ_BEFOREEXEC:  true,
		USAGE_EXTERNAL_WRITE_BEFOREEXEC: false,
		USAGE_EXTERNAL_READ_AFTEREXEC:   true,
		USAGE_EXTERNAL_WRITE_AFTEREEXEC: false,
		USAGE_INTERNAL_READ_BEFOREEXEC:  true,
		USAGE_INTERNAL_WRITE_BEFOREEXEC: false,
		USAGE_INTERNAL_READ_AFTEREXEC:   true,
		USAGE_INTERNAL_WRITE_AFTEREXEC:  false,
	})
}

/**
 * Usage tests
 */

// Is this an external required field
func IsUsage_Required(usage api_usage.Usage) bool {
	return usage.Has(USAGE_EXTERNAL_REQUIRED)
}

// Is this an external required field
func IsUsage_ExternalRequired(usage api_usage.Usage) bool {
	return usage.Has(USAGE_EXTERNAL_WRITE_BEFOREEXEC) && usage.Has(USAGE_EXTERNAL_REQUIRED)
}

// Is this an external required field
func IsUsage_ExternalOptional(usage api_usage.Usage) bool {
	return usage.Has(USAGE_EXTERNAL_WRITE_BEFOREEXEC)
}

// Is this an external required field
func IsUsage_ExternalVisibleAfter(usage api_usage.Usage) bool {
	return usage.Has(USAGE_INTERNAL_READ_AFTEREXEC)
}

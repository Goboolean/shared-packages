package resolver

import (
	"fmt"
)


// ConfigMap is a struct that holds the configuration values as dynamic
// its value type can be string, int, float64.
type ConfigMap map[string]interface{}

// One way to set key on ConfigMap is to use Constructor as
// &ConfigMap{"KEY": value, ...}
// The other way is to use SetKey method
func (c *ConfigMap) SetKey(key string, value interface{}) error {
	(*c)[key] = value
	return nil
}


// GetStringKey returns the value of the key as string
func (c *ConfigMap) GetStringKey(key string) (string, error) {

	// It throws an error if the key is not found
	// It throws an error if the value is not string
	errMsg := fmt.Sprintf("value %s is required as string value", key)

	value, exists := (*c)[key]
	if !exists {
		return "", fmt.Errorf("failed to get key: %s", errMsg)
	}

	stringValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("failed to parse as string: %s", errMsg)
	}

	return stringValue, nil
}


// GetIntKey returns the value of the key as int
func (c *ConfigMap) GetIntKey(key string) (int, error) {

	// It throws an error if the key is not found
	// It throws an error if the value is not int
	errMsg := fmt.Sprintf("value %s is required as int value", key)

	value, exists := (*c)[key]
	if !exists {
		return 0, fmt.Errorf("failed to get key: %s", errMsg)
	}

	intValue, ok := value.(int)
	if !ok {
		return 0, fmt.Errorf("failed to parse as int: %s", errMsg)
	}

	return intValue, nil
}


// GetFloatKey returns the value of the key as float64
func (c *ConfigMap) GetFloatKey(key string) (float64, error) {

	// It throws an error if the key is not found
	// It throws an error if the value is not float
	errMsg := fmt.Sprintf("value %s is required as float value", key)

	value, exists := (*c)[key]
	if !exists {
		return 0, fmt.Errorf("failed to get key: %s", errMsg)
	}

	floatValue, ok := value.(float64)
	if !ok {
		return 0, fmt.Errorf("failed to parse as float: %s", errMsg)
	}

	return floatValue, nil
}
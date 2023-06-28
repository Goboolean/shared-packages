package resolver

import (
	"fmt"
)



type ConfigMap map[string]interface{}

func (c *ConfigMap) SetKey(key string, value interface{}) error {
	(*c)[key] = value
	return nil
}


func (c *ConfigMap) GetStringKey(key string) (string, error) {

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


func (c *ConfigMap) GetIntKey(key string) (int, error) {

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


func (c *ConfigMap) GetFloatKey(key string) (float64, error) {

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



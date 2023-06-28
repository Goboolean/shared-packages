package resolver

import "fmt"



// Config is a struct that holds the configuration values as a const
// Use ConfigMap rather than Config. It is now a lagecy
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Address  string
	Key      string
	Path     string
}

func (c *Config) ShouldHostExist() error {
	if c.Host == "" {
		return fmt.Errorf("config value HOST required")
	}
	return nil
}

func (c *Config) ShouldPortExist() error {
	if c.Port == "" {
		return fmt.Errorf("config value PORT required")
	}
	return nil
}

func (c *Config) ShouldUserExist() error {
	if c.User == "" {
		return fmt.Errorf("config vlalue USER required")
	}
	return nil
}

func (c *Config) ShouldPWExist() error {
	if c.Password == "" {
		return fmt.Errorf("config vlalue PASSWORD required")
	}
	return nil
}

func (c *Config) ShouldDBExist() error {
	if c.Database == "" {
		return fmt.Errorf("config vlalue DATABASE required")
	}
	return nil
}

func (c *Config) ShouldAddrExist() error {
	if c.Address == "" {
		return fmt.Errorf("config vlalue ADDRESS required")
	}
	return nil
}

func (c *Config) ShouldKeyExist() error {
	if c.Key == "" {
		return fmt.Errorf("config vlalue KEY required")
	}
	return nil
}

func (c *Config) ShouldPathExist() error {
	if c.Path == "" {
		return fmt.Errorf("config vlalue PATH required")
	}
	return nil
}


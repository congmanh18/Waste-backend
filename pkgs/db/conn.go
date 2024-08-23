package db

import (
	"errors"
	"fmt"
)

type Connection struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func (c Connection) HasError() error {
	var errs []error
	if c.Host == "" {
		errs = append(errs, fmt.Errorf("host required"))
	}
	if c.User == "" {
		errs = append(errs, fmt.Errorf("user required"))
	}
	if c.Password == "" {
		errs = append(errs, fmt.Errorf("password required"))
	}
	if c.DBName == "" {
		errs = append(errs, fmt.Errorf("dBName required"))
	}
	if c.Port == "" {
		errs = append(errs, fmt.Errorf("port required"))
	}

	return errors.Join(errs...)
}

func (c Connection) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.Host, c.User, c.Password, c.DBName, c.Port)
}

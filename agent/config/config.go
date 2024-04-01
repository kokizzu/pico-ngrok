// Copyright 2024 Andrew Dunstall. All rights reserved.
//
// Use of this source code is governed by a MIT style license that can be
// found in the LICENSE file.

package config

import "fmt"

type LogConfig struct {
	Level string `json:"level"`
	// Subsystems enables debug logging on logs the given subsystems (which
	// overrides level).
	Subsystems []string `json:"subsystems"`
}

func (c *LogConfig) Validate() error {
	if c.Level == "" {
		return fmt.Errorf("missing level")
	}
	return nil
}

type Config struct {
	Log LogConfig `json:"log"`
}

func (c *Config) Validate() error {
	if err := c.Log.Validate(); err != nil {
		return fmt.Errorf("log: %w", err)
	}
	return nil
}

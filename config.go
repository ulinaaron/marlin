// config.go

package main

import "github.com/spf13/viper"

// Config stores editor settings 
type Config struct {
  // Add config fields  
}

// LoadConfig parses config file
func LoadConfig() (*Config, error) {

  viper.SetConfigName("config")

  if err := viper.ReadInConfig(); err != nil {
    return nil, err
  } 

  // Parse config into struct
  var c Config 
  if err := viper.Unmarshal(&c); err != nil {
    return nil, err
  }

  return &c, nil
}

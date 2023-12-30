package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/glebziz/fs_db/pkg/model"
)

const (
	minDirCount     = 100
	defaultDirCount = 1_000_000
)

type Storage struct {
	DbPath      string   `yaml:"dbPath"`
	MaxDirCount uint64   `yaml:"maxDirCount"`
	RootDirs    []string `yaml:"rootDirs"`
}

func (s *Storage) Valid() error {
	if s.DbPath == "" {
		return model.EmptyDbPathErr
	}

	if s.MaxDirCount == 0 {
		s.MaxDirCount = defaultDirCount
	} else if s.MaxDirCount < minDirCount {
		s.MaxDirCount = minDirCount
	}

	if len(s.RootDirs) == 0 {
		return model.EmptyRootDirs
	}

	return nil
}

type Config struct {
	Port    int     `yaml:"port"`
	Storage Storage `yaml:"storage"`
}

func ParseConfig(confFile string) (*Config, error) {
	f, err := os.Open(confFile)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer f.Close()

	conf := Config{}

	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	return &conf, nil
}

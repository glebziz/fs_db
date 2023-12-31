package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/glebziz/fs_db/pkg/model"
)

const (
	envPort     = "PORT"
	envDbPath   = "DB_PATH"
	envDirCount = "DIR_COUNT"
	envRootDirs = "ROOT_DIRS"
)

const (
	defaultPort     = 8888
	defaultDbPath   = "test.db"
	defaultDirCount = 1_000_000
	defaultRootDir  = "./testStorage"

	minDirCount = 100
)

var (
	defaultConfig = Config{
		Port: defaultPort,
		Storage: Storage{
			DbPath:      defaultDbPath,
			MaxDirCount: defaultDirCount,
			RootDirs:    []string{defaultRootDir},
		},
	}
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

	if s.MaxDirCount < minDirCount {
		s.MaxDirCount = minDirCount
	}

	if len(s.RootDirs) == 0 {
		return model.EmptyRootDirs
	}

	return nil
}

func (s *Storage) ParseEnv() error {
	if env, ok := os.LookupEnv(envDbPath); ok && env != "" {
		s.DbPath = env
	}
	if env, ok := os.LookupEnv(envDirCount); ok && env != "" {
		dirCount, err := strconv.ParseUint(env, 10, 64)
		if err != nil {
			return fmt.Errorf("parse dir count: %w", err)
		}
		s.MaxDirCount = dirCount
	}
	if env, ok := os.LookupEnv(envRootDirs); ok && env != "" {
		s.RootDirs = strings.Split(env, ";")
	}

	return nil
}

type Config struct {
	Port    int     `yaml:"port"`
	Storage Storage `yaml:"storage"`
}

func (c *Config) ParseEnv() error {
	if env, ok := os.LookupEnv(envPort); ok && env != "" {
		port, err := strconv.Atoi(env)
		if err != nil {
			return fmt.Errorf("parse port: %w", err)
		}
		c.Port = port
	}

	err := c.Storage.ParseEnv()
	if err != nil {
		return fmt.Errorf("storage parse env: %w", err)
	}

	return nil
}

func ParseConfig(confFile string) (*Config, error) {
	conf := defaultConfig

	if confFile != "" {
		f, err := os.Open(confFile)
		if err != nil {
			return nil, fmt.Errorf("open: %w", err)
		}
		defer f.Close()

		err = yaml.NewDecoder(f).Decode(&conf)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("decode: %w", err)
		}
	}

	err := conf.ParseEnv()
	if err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return &conf, nil
}

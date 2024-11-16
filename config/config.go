package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/glebziz/fs_db"
)

const (
	envPort         = "PORT"
	envDbPath       = "DB_PATH"
	envDirCount     = "DIR_COUNT"
	envRootDirs     = "ROOT_DIRS"
	envGCPeriod     = "GC_PERIOD"
	envNumWorkers   = "NUM_WORKERS"
	envSendDuration = "SEND_DURATION"
)

const (
	defaultPort         = 8888
	defaultDbPath       = "test_db"
	defaultDirCount     = 1_000_000
	defaultRootDir      = "./testStorage"
	defaultGCPeriod     = time.Minute
	defaultSendDuration = time.Millisecond

	minDirCount = 100
)

var (
	defaultConfig = Config{
		Port: defaultPort,
		Storage: Storage{
			DbPath:      defaultDbPath,
			MaxDirCount: defaultDirCount,
			RootDirs:    []string{defaultRootDir},
			GCPeriod:    defaultGCPeriod,
		},
		WPool: WPool{
			NumWorkers:   runtime.GOMAXPROCS(0),
			SendDuration: defaultSendDuration,
		},
	}
)

// Storage provides configuration options for fs db storage.
//
//	Default:
//	  dbPath: test_db
//	  maxDirCount: 1000000
//	  rootDirs:
//	    - ./testStorage
//	  gcPeriod: 1m
type Storage struct {
	// DbPath path to badger database directory.
	//	Default: test_db
	//	Env: DB_PATH
	DbPath string `yaml:"dbPath"`

	// MaxDirCount max number of files in one subdirectory.
	//	Default: 1_000_000
	//	Env: DIR_COUNT
	MaxDirCount uint64 `yaml:"maxDirCount"`

	// RootDirs slice with root directories.
	//	Default: ["./testStorage"]
	//	Env: ROOT_DIRS
	RootDirs []string `yaml:"rootDirs"`

	// GCPeriod period for run the mechanism for deleting old versions of files.
	//	Default: 1m
	//	Env: GC_PERIOD
	GCPeriod time.Duration `yaml:"gcPeriod"`
}

// Valid validates the storage options.
func (s *Storage) Valid() error {
	if s.DbPath == "" {
		return fs_db.EmptyDbPathErr
	}

	if s.MaxDirCount < minDirCount {
		s.MaxDirCount = minDirCount
	}

	if len(s.RootDirs) == 0 {
		return fs_db.EmptyRootDirs
	}

	return nil
}

// ParseEnv fills the storage options with environment variables.
func (s *Storage) ParseEnv() (err error) {
	if env, ok := os.LookupEnv(envDbPath); ok && env != "" {
		s.DbPath = env
	}
	if env, ok := os.LookupEnv(envDirCount); ok && env != "" {
		s.MaxDirCount, err = strconv.ParseUint(env, 10, 64)
		if err != nil {
			return fmt.Errorf("parse dir count: %w", err)
		}
	}
	if env, ok := os.LookupEnv(envRootDirs); ok && env != "" {
		s.RootDirs = strings.Split(env, ";")
	}
	if env, ok := os.LookupEnv(envGCPeriod); ok && env != "" {
		s.GCPeriod, err = time.ParseDuration(env)
		if err != nil {
			return fmt.Errorf("parse GC period: %w", err)
		}
	}

	return nil
}

// WPool provides configuration options for a worker pool.
//
//	Default:
//	  numWorkers: runtime.GOMAXPROCS(0)
//	  sendDuration: 1ms
type WPool struct {
	// NumWorkers the number of workers in a worker pool that can work at the same time.
	//	Default: runtime.GOMAXPROCS(0)
	//	Env: NUM_WORKERS
	NumWorkers int `yaml:"numWorkers"`

	// SendDuration the maximum duration of wait for an event to be sent to the worker pool.
	//	Default: 1ms
	//	Env: SEND_DURATION
	SendDuration time.Duration `yaml:"sendDuration"`
}

// ParseEnv fills the storage options with environment variables.
func (wp *WPool) ParseEnv() (err error) {
	if env, ok := os.LookupEnv(envNumWorkers); ok && env != "" {
		wp.NumWorkers, err = strconv.Atoi(env)
		if err != nil {
			return fmt.Errorf("parse num workers: %w", err)
		}
	}
	if env, ok := os.LookupEnv(envSendDuration); ok && env != "" {
		wp.SendDuration, err = time.ParseDuration(env)
		if err != nil {
			return fmt.Errorf("parse send duration: %w", err)
		}
	}

	return nil
}

// Config provides fs db configuration options.
//
//	Default:
//	  port: 8888
//	  storage:
//	    dbPath: test_db
//	    maxDirCount: 1000000
//	    rootDirs:
//	      - ./testStorage
//	    gcPeriod: 1m
//	  wPool:
//	    numWorkers: runtime.GOMAXPROCS(0)
//	    sendDuration: 1ms
type Config struct {
	// Port fs db server port.
	//  Default: 8888
	//  Env: PORT
	Port int `yaml:"port"`

	// Storage fs db storage options.
	Storage Storage `yaml:"storage"`

	// WPool worker pool options.
	WPool WPool `yaml:"wPool"`
}

// ParseEnv fills the config options with environment variables.
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

	err = c.WPool.ParseEnv()
	if err != nil {
		return fmt.Errorf("wPool parse env: %w", err)
	}

	return nil
}

// ParseConfig returns filled Config options from a configuration file.
func ParseConfig(confFile string) (Config, error) {
	conf := defaultConfig

	if confFile != "" {
		f, err := os.Open(confFile)
		if err != nil {
			return Config{}, fmt.Errorf("open: %w", err)
		}
		defer f.Close()

		err = yaml.NewDecoder(f).Decode(&conf)
		if err != nil && !errors.Is(err, io.EOF) {
			return Config{}, fmt.Errorf("decode: %w", err)
		}
	}

	err := conf.ParseEnv()
	if err != nil {
		return Config{}, fmt.Errorf("parse env: %w", err)
	}

	return conf, nil
}

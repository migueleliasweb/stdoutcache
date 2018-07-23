package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"os/exec"
	"time"
)

//StdoutCacher Represents the cacher
type StdoutCacher struct {
	//TTL in seconds the cache is still valid
	//Zero (0) disables cache (usefull for debugging)
	ttl int

	//The underlying command to be ran
	command string

	//Arguments to be passed to the "command"
	args []string

	//Slice of string compatible with os.Environ()
	environment []string
}

func (cacher *StdoutCacher) generateCacheFilename() string {
	hash := sha256.New()

	hash.Write([]byte(cacher.command))

	for _, arg := range cacher.args {
		hash.Write([]byte(arg))
	}

	for _, envvar := range cacher.environment {
		hash.Write([]byte(envvar))
	}

	md := hash.Sum(nil)

	return cacher.command + "_" + hex.EncodeToString(md) + ".cache"
}

func (cacher *StdoutCacher) getCacheReader() (io.ReadCloser, error) {
	fileCacheName := cacher.generateCacheFilename()
	absPath := os.TempDir() + "/" + fileCacheName

	return os.Open(absPath)
}

func (cacher *StdoutCacher) isCacheValid() bool {
	fileCacheName := cacher.generateCacheFilename()
	fullPath := os.TempDir() + "/" + fileCacheName

	cacheStat, err := os.Stat(fullPath)

	if os.IsNotExist(err) {
		return false
	}

	cacheMaxCreateTime := time.Now().Add(time.Second * time.Duration(-1*cacher.ttl))

	//ModTime is used here as the cache file is never updated.
	//Instead, the file is always recreated
	if cacheStat.ModTime().After(cacheMaxCreateTime) {
		return false
	}

	return true
}

func (cacher *StdoutCacher) readCache() (string, error) {
	return "", nil
}

//RunCommand Executes the underlying command
func (cacher *StdoutCacher) RunCommand() (string, error) {

	if cacher.isCacheValid() {
		return cacher.readCache()
	}

	cmd := exec.Command(cacher.command)
	cmd.Args = cacher.args
	cmd.Env = append(cacher.environment)

	stdout, stderr := cmd.Output()

	return string(stdout), stderr
}

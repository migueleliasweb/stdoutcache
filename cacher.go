package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os/exec"
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

//RunCommand Executes the underlying command
func (cacher *StdoutCacher) RunCommand() (string, error) {
	cmd := exec.Command(cacher.command)
	cmd.Args = cacher.args
	cmd.Env = append(cacher.environment)

	stdout, stderr := cmd.Output()

	return string(stdout), stderr
}

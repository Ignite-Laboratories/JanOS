package common

import (
	"log"
	"os"
	"path/filepath"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetJanOSDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		log.Printf("[Config] Error getting working directory: %s\n", err)
	}
	return filepath.Dir(path)
}

func GetSystemDirectory() string {
	return filepath.Join(GetJanOSDirectory(), "systems", GetSystemName())
}

func GetSocketDirectory() string {
	return filepath.Join(GetJanOSDirectory(), "socket")
}

func GetSystemName() string {
	return os.Args[1]
}

func GetComponentName() string {
	return os.Args[2]
}

func GetRemoteSocketAddress(names ...string) string {
	componentName := names[0]
	var hasInstanceName bool
	var instanceName string
	if len(names) > 1 {
		instanceName = names[1]
		hasInstanceName = true
	}

	socketAddress := filepath.Join(GetSocketDirectory(), GetSystemName()+"-"+componentName)
	if hasInstanceName {
		socketAddress += "-" + instanceName
	}
	return socketAddress
}

func GetSocketAddress() string {
	var hasInstanceName bool
	var instanceName string

	if len(os.Args) > 3 {
		instanceName = os.Args[3]
		hasInstanceName = true
	}

	socketAddress := filepath.Join(GetSocketDirectory(), GetSystemName()+"-"+GetComponentName())
	if hasInstanceName {
		socketAddress += "-" + instanceName
	}

	// If the socket file already exists we clobber it and restart
	if FileExists(socketAddress) {
		err := os.Remove(socketAddress)
		if err != nil {
			log.Println("[nexus] Error removing existing socket")
		}
	}
	return socketAddress
}

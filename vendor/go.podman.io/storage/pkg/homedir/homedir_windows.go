package homedir

// Copyright 2013-2018 Docker, Inc.
// NOTE: this package has originally been copied from github.com/docker/docker.

import (
	"os"
	"path/filepath"
)

// Key returns the env var name for the user's home dir based on
// the platform being run on
func Key() string {
	return "USERPROFILE"
}

// Get returns the home directory of the current user with the help of
// environment variables depending on the target operating system.
// Returned path should be used with "path/filepath" to form new paths.
func Get() string {
	home := os.Getenv(Key())
	if home != "" {
		return home
	}
	home, _ = os.UserHomeDir()
	return home
}

// GetConfigHome returns the home directory of the current user with the help of
// environment variables depending on the target operating system.
// Returned path should be used with "path/filepath" to form new paths.
//
// bashy patch: honor XDG_CONFIG_HOME first, matching the unix variant. Upstream
// Windows ignored it, which made bashy's podman-machine registry isolation
// (isolateMachineDirs sets XDG_CONFIG_HOME=~/.config/bashy) a NO-OP on Windows —
// so bashy's machine config dir still collided with a host podman's. Honoring it
// here isolates the machine registry uniformly across platforms. NOTE: this is a
// vendored dep; re-apply if go.podman.io/storage is re-vendored. See
// docs/dks-scratch-vm-appliance-design.md (Windows gaps #1).
func GetConfigHome() (string, error) {
	if cfgHomeDir := os.Getenv("XDG_CONFIG_HOME"); cfgHomeDir != "" {
		return cfgHomeDir, nil
	}
	return filepath.Join(Get(), ".config"), nil
}

// GetShortcutString returns the string that is shortcut to user's home directory
// in the native shell of the platform running on.
func GetShortcutString() string {
	return "%USERPROFILE%" // be careful while using in format functions
}

// StickRuntimeDirContents is a no-op on Windows
func StickRuntimeDirContents(files []string) ([]string, error) {
	return nil, nil
}

// GetRuntimeDir returns a directory suitable to store runtime files.
// The function will try to use the XDG_RUNTIME_DIR env variable if it is set.
// XDG_RUNTIME_DIR is typically configured via pam_systemd.
// If XDG_RUNTIME_DIR is not set, GetRuntimeDir will try to find a suitable
// directory for the current user.
//
// See also https://standards.freedesktop.org/basedir-spec/latest/ar01s03.html
func GetRuntimeDir() (string, error) {
	data, err := GetDataHome()
	if err != nil {
		return "", err
	}
	runtimeDir := filepath.Join(data, "containers", "storage")
	return runtimeDir, nil
}

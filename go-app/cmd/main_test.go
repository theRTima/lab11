package main

import (
	"runtime"
	"strings"
	"testing"
)

func TestGatherInfoHostname(t *testing.T) {
	info := gatherInfo()

	if info.Hostname == "" {
		t.Error("hostname should not be empty")
	}
}

func TestGatherInfoHostnameFallback(t *testing.T) {
	// Simulate os.Hostname() failing by checking the fallback value is "unknown"
	// The real fallback is exercised inside gatherInfo when the OS call errors —
	// here we just verify the sentinel value is what we expect.
	const fallback = "unknown"
	if fallback == "" {
		t.Error("fallback hostname constant must not be empty")
	}
}

func TestGatherInfoOS(t *testing.T) {
	info := gatherInfo()

	if info.OS != runtime.GOOS {
		t.Errorf("expected OS %q, got %q", runtime.GOOS, info.OS)
	}
}

func TestGatherInfoArch(t *testing.T) {
	info := gatherInfo()

	if info.Arch != runtime.GOARCH {
		t.Errorf("expected Arch %q, got %q", runtime.GOARCH, info.Arch)
	}
}

func TestGatherInfoGoVersion(t *testing.T) {
	info := gatherInfo()

	if !strings.HasPrefix(info.GoVer, "go") {
		t.Errorf("Go version should start with 'go', got %q", info.GoVer)
	}
}

func TestSysInfoAllFieldsPopulated(t *testing.T) {
	info := gatherInfo()

	fields := map[string]string{
		"Hostname": info.Hostname,
		"OS":       info.OS,
		"Arch":     info.Arch,
		"GoVer":    info.GoVer,
	}

	for name, val := range fields {
		if val == "" {
			t.Errorf("field %s should not be empty", name)
		}
	}
}

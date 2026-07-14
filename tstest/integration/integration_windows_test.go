// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build windows

package integration

import (
	"testing"

	"tailscale.com/tstest"
)

// Issue 2137: Windows tailscaled works from the CLI alone, without the GUI to
// start it. Runs as a Windows service; see NewTestEnv's Windows gating.
func TestOneNodeUpWindowsStyle(t *testing.T) {
	tstest.Parallel(t)
	env := NewTestEnv(t, UnskipOnWindows())
	n1 := NewTestNode(t, env)
	n1.upFlagGOOS = "windows"

	d1 := n1.StartDaemonAsIPNGOOS("windows")
	n1.AwaitResponding()
	n1.MustUp("--unattended")

	t.Logf("Got IP: %v", n1.AwaitIP4())
	n1.AwaitRunning()

	d1.MustCleanShutdown(t)
}

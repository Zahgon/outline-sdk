// Copyright 2023 The Outline Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// todo: find a more portable way of configuring DNS (e.g. resolved)
const (
	resolvConfFile           = "/etc/resolv.conf"
	resolvConfHeadFile       = "/etc/resolv.conf.head"
	resolvConfBackupFile     = "/etc/resolv.outlinecli.backup"
	resolvConfHeadBackupFile = "/etc/resolv.head.outlinecli.backup"
)

type systemDNSBackup struct {
	original, backup string
}

var systemDNSBackups = make([]systemDNSBackup, 0, 2)

func setSystemDNSServer(serverHost string) error { _ = "STUB: not implemented"; return nil }

func backupAndWriteFile(original, backup string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// original file exist - move it into backup

// if not exist - it's ok, just write to it

// save data to original

func restoreSystemDNSServer() { _ = "STUB: not implemented"; return }

// backup exist - replace original with it

// backup not exist - just remove original, because it's created by ourselves

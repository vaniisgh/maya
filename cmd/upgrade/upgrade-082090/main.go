/*
Copyright 2018 The OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	upgrade "github.com/openebs/maya/cmd/upgrade/app/v1alpha1"
	log "k8s.io/klog"
)

func main() {
	err := flag.Set("logtostderr", "true")
	if err != nil {
		fmt.Printf("failed to upgrade: %s", err)
		os.Exit(1)
	}
	configPath := flag.String("config-path", "/etc/config/upgrade", "path to upgrade config file.")
	defer log.Flush()
	flag.Parse()

	u, err := upgrade.NewUpgradeForConfigPath(*configPath)
	if err != nil {
		log.Errorf("failed to upgrade: %+v", err)
		os.Exit(1)
	}

	err = u.Run()
	if err != nil {
		log.Errorf("failed to upgrade: %+v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

/*
Copyright 2021 The Caoyingjunz Authors.

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

package localstorage

import (
	"fmt"
	"k8s.io/klog/v2"
	"sync"
)

const (
	DefaultDriverName = "localstorage.csi.pixiu.io"
)

type localStorage struct {
	config Config

	mutex sync.Mutex
}

type Config struct {
	DriverName string
	Endpoint   string
	Version    string
}

func NewLocalStorage(cfg Config) (*localStorage, error) {
	if len(cfg.DriverName) == 0 {
		return nil, fmt.Errorf("no driver name provided")
	}

	if len(cfg.Endpoint) == 0 {
		return nil, fmt.Errorf("no driver endpoint provided")
	}

	klog.V(2).Infof("Driver: %v version: %v", cfg.DriverName, cfg.Version)

	ls := &localStorage{
		config: cfg,
	}
	return ls, nil
}

func (ls *localStorage) Run() error {

	return nil
}

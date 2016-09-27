//
// Copyright © 2016 Ikey Doherty
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ylib

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

// PackageYML is a struct representation of the ypkg package.yml file
type PackageYML struct {
	Name        string              `yaml:"name"`
	Version     string              `yaml:"version"`
	Release     uint                `yaml:"release"`
	Source      []map[string]string `yaml:"source"`
	License     []string            `yaml:"license"`
	Summary     string              `yaml:"summary"`
	Component   string              `yaml:"component"`
	Description string              `yaml:"description,omitempty"`
	BuildDeps   []string            `yaml:"builddeps,omitempty"`
	RunDeps     []string            `yaml:"rundeps,omitempty"`
	Emul32      string              `yaml:"emul32,omitempty"`
	Optimize    string              `yaml:"optimize,omitempty"`
	Patterns    map[string][]string `yaml:"patterns,omitempty"`
	Setup       string              `yaml:"setup,omitempty"`
	Build       string              `yaml:"build,omitempty"`
	Install     string              `yaml:"install"`
}

func (pkg *PackageYML) Read(in io.Reader) (err error) {
	raw, err := ioutil.ReadAll(in)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(raw, pkg)
	return
}

func (pkg *PackageYML) Write(out io.Writer) (n int, err error) {
	raw, err := yaml.Marshal(pkg)
	if err != nil {
		return
	}
	n, err = out.Write(raw)
	return
}

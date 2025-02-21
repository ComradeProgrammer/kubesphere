/*

 Copyright 2021 The KubeSphere Authors.

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

package authorization

import (
	"fmt"

	"github.com/spf13/pflag"
	"k8s.io/klog"

	"kubesphere.io/kubesphere/pkg/utils/sliceutil"
)

type Options struct {
	Mode string `json:"mode"  yaml:"mode"`
}

func NewOptions() *Options {
	return &Options{Mode: RBAC}
}

var (
	AlwaysDeny  = "AlwaysDeny"
	AlwaysAllow = "AlwaysAllow"
	RBAC        = "RBAC"
)

func (o *Options) AddFlags(fs *pflag.FlagSet, s *Options) {
	fs.StringVar(&o.Mode, "authorization", s.Mode, "Authorization setting, allowed values: AlwaysDeny, AlwaysAllow, RBAC.")
}

func (o *Options) Validate() []error {
	errs := make([]error, 0)
	if !sliceutil.HasString([]string{AlwaysAllow, AlwaysDeny, RBAC}, o.Mode) {
		err := fmt.Errorf("authorization mode %s not support", o.Mode)
		klog.Error(err)
		errs = append(errs, err)
	}
	return errs
}

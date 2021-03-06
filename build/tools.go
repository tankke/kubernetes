// +build tools

/*
Copyright 2019 The Kubernetes Authors.

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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/bazelbuild/bazel-gazelle/cmd/gazelle"
	_ "github.com/bazelbuild/buildtools/buildozer"
	_ "github.com/cespare/prettybench"
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/cloudflare/cfssl/cmd/cfssl"
	_ "github.com/cloudflare/cfssl/cmd/cfssljson"
	_ "github.com/jstemmer/go-junit-report"
	_ "github.com/jteeuwen/go-bindata/go-bindata"
	_ "github.com/onsi/ginkgo/ginkgo"
	_ "golang.org/x/lint/golint"
	_ "k8s.io/code-generator/cmd/go-to-protobuf"
	_ "k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo"
	_ "k8s.io/gengo/examples/deepcopy-gen/generators"
	_ "k8s.io/gengo/examples/defaulter-gen/generators"
	_ "k8s.io/gengo/examples/import-boss/generators"
	_ "k8s.io/gengo/examples/set-gen/generators"
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
	_ "k8s.io/repo-infra/kazel"
)

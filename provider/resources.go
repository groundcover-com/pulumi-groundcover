// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package groundcover

import (
	"path"

	_ "embed"

	"github.com/groundcover-com/terraform-provider-groundcover/shim"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumiverse/pulumi-groundcover/provider/pkg/version"
)

const (
	mainPkg = "groundcover"
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-groundcover/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:           pfbridge.ShimProvider(shim.NewProvider(version.Version)),
		Name:        "groundcover",
		Version:     version.Version,
		DisplayName: "Groundcover",
		Publisher:   "Pulumiverse",
		LogoURL:     "https://raw.githubusercontent.com/pulumiverse/pulumi-groundcover/main/docs/groundcover.png",
		PluginDownloadURL: "github://api.github.com/pulumiverse/pulumi-groundcover",
		Description:       "A Pulumi package for creating and managing groundcover resources.",
		Keywords:          []string{"groundcover", "observability", "monitoring", "category/monitoring"},
		License:           "Apache-2.0",
		Homepage:          "https://groundcover.com",
		Repository:        "https://github.com/pulumiverse/pulumi-groundcover",
		GitHubOrg:         "groundcover-com",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config:            map[string]*tfbridge.SchemaInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/groundcover",
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_groundcover",
			RespectSchemaVersion: true,
			PyProject:            struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				"github.com/pulumiverse/pulumi-groundcover/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RespectSchemaVersion: true,
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("groundcover_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}

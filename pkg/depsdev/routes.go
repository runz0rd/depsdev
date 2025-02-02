/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package depsdev

const (
	// API base path.
	BasePath = `https://api.deps.dev/v3alpha`

	// API routes.
	GetPackagePath      = `/systems/%s/packages/%s`
	GetVersionPath      = `/systems/%s/packages/%s/versions/%s`
	GetDependenciesPath = `/systems/%s/packages/%s/versions/%s:dependencies`
	GetProjectPath      = `/projects/%s`
	GetAdvisoryPath     = `/advisories/%s`
	GetQueryPath        = `/query`
	GetRequirementsPath = `/systems/%s/packages/%s/versions/%s:requirements`
)

package depsdev

type Requirements struct {
	Nuget *NugetRequirement `json:"nuget,omitempty"`
	Npm   *NpmRequirement   `json:"npm,omitempty"`
	Maven *MavenRequirement `json:"maven,omitempty"`
}

type NugetRequirement struct {
	DependependencyGroups []DependependencyGroup `json:"dependencyGroups,omitempty"`
}

type DependependencyGroup struct {
	TargetFramework string       `json:"targetFramework,omitempty"`
	Dependencies    []Dependency `json:"dependencies,omitempty"`
}

type Dependency struct {
	Name        string   `json:"name,omitempty"`
	Requirement string   `json:"requirement,omitempty"`
	Version     string   `json:"version,omitempty"`
	System      string   `json:"system,omitempty"`
	Classifier  string   `json:"classifier,omitempty"`
	Type        string   `json:"type,omitempty"`
	Scope       string   `json:"scope,omitempty"`
	Optional    string   `json:"optional,omitempty"`
	Exclusions  []string `json:"exclusions,omitempty"`
}

type NpmRequirement struct {
	Dependencies NpmDependency `json:"dependencies,omitempty"`
	Bundled      []Bundled     `json:"bundled,omitempty"`
}

type Bundled struct {
	Path         string          `json:"path,omitempty"`
	Name         string          `json:"name,omitempty"`
	Version      string          `json:"version,omitempty"`
	Dependencies []NpmDependency `json:"dependencies,omitempty"`
}

type NpmDependency struct {
	Dependencies         []Dependency `json:"dependencies"`
	DevDependencies      []Dependency `json:"devDependencies"`
	OptionalDependencies []Dependency `json:"optionalDependencies"`
	PeerDependencies     []Dependency `json:"peerDependencies"`
	BundleDependencies   []string     `json:"bundleDependencies"`
}

type MavenRequirement struct {
	ID                   string             `json:"id,omitempty"`
	Parent               Dependency         `json:"parent,omitempty"`
	Activation           Activation         `json:"activation,omitempty"`
	Dependencies         []Dependency       `json:"dependencies,omitempty"`
	DependencyManagement []Dependency       `json:"dependency_management,omitempty"`
	Properties           []Property         `json:"properties,omitempty"`
	Repositories         []MavenRepository  `json:"repositories,omitempty"`
	Profiles             []MavenRequirement `json:"profiles,omitempty"`
}

type Property struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type MavenRepository struct {
	ID               string `json:"id,omitempty"`
	URL              string `json:"url,omitempty"`
	Layout           string `json:"layout,omitempty"`
	ReleasesEnabled  string `json:"releases_enabled,omitempty"`
	SnapshotsEnabled string `json:"snapshots_enabled,omitempty"`
}

type Activation struct {
	JDK      JDK             `json:"jdk,omitempty"`
	OS       OS              `json:"os,omitempty"`
	Property PropertyWrapper `json:"property,omitempty"`
	File     File            `json:"file,omitempty"`
}

type JDK struct {
	JDK string `json:"jdk,omitempty"`
}

type OS struct {
	Name    string `json:"name,omitempty"`
	Family  string `json:"family,omitempty"`
	Arch    string `json:"arch,omitempty"`
	Version string `json:"version,omitempty"`
}

type PropertyWrapper struct {
	Property Property `json:"property,omitempty"`
}

type File struct {
	Exists  string `json:"exists,omitempty"`
	Missing string `json:"missing,omitempty"`
}

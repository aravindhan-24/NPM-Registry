package constants

type InstallPackage struct {
	Name      string                `json:"name"`
	Dist_tags map[string]string     `json:"dist-tags"`
	Versions  map[string]NpmVersion `json:"versions"`
}

type NpmVersion struct {
	Name         string             `json:"name"`
	Version      string             `json:"version"`
	Main         string             `json:"main"`
	Scripts      map[string]string  `json:"scripts"`
	Author       map[string]string  `json:"author"`
	License      string             `json:"license"`
	Dependencies *map[string]string `json:"dependencies"`
	Dist         Distribution       `json:"dist"`
	ID           string             `json:"_id"`
	Readme       string             `json:"readme"`
	NodeVersion  string             `json:"_nodeVersion"`
	NpmVersion   string             `json:"_npmVersion"`
}
type Distribution struct {
	Sha256sum string `json:"shasum"`
	Tarball   string `json:"tarball"`
	Integrity string `json:"integrity"`
}

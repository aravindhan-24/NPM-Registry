package constants

type Package struct {
	ID          string                    `json:"_id"`
	Name        string                    `json:"name"`
	DistTags    map[string]string         `json:"dist-tags"`
	Versions    map[string]PackageVersion `json:"versions"`
	Access      interface{}               `json:"access"` // null in your JSON
	Attachments map[string]Attachment     `json:"_attachments"`
}

type PackageVersion struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Main         string            `json:"main"`
	Scripts      map[string]string `json:"scripts"`
	Keywords     []string          `json:"keywords"`
	Author       map[string]string `json:"author"`
	License      string            `json:"license"`
	Dependencies map[string]string `json:"dependencies"`
	ID           string            `json:"_id"`
	Readme       string            `json:"readme"`
	NodeVersion  string            `json:"_nodeVersion"`
	NpmVersion   string            `json:"_npmVersion"`
	Dist         Dist              `json:"dist"`
}

type Dist struct {
	Integrity string `json:"integrity"`
	Shasum    string `json:"shasum"`
	Tarball   string `json:"tarball"`
}

type Attachment struct {
	ContentType string `json:"content_type"`
	Data        string `json:"data"`
	Length      int    `json:"length"`
}

package risk

// RiskAssessment represents the top-level structure of a risk assessment definition file.
type RiskAssessment struct {
	Name        string      `yaml:"name" json:"name"`
	Description string      `yaml:"description" json:"description"`
	Risks       []RiskItem  `yaml:"risks" json:"risks"`
	Controls    []ControlItem `yaml:"controls,omitempty" json:"controls,omitempty"` // Controls applied to this specific assessment
}

// RiskItem defines an individual risk.
type RiskItem struct {
	ID          string   `yaml:"id" json:"id"`
	Name        string   `yaml:"name" json:"name"`
	Description string   `yaml:"description" json:"description"`
	Impact      int      `yaml:"impact" json:"impact"`           // e.g., 1-5, higher is worse
	Likelihood  int      `yaml:"likelihood" json:"likelihood"`   // e.g., 1-5, higher is more likely
	AffectedAssets []string `yaml:"affected_assets" json:"affected_assets"`
	Threats     []string `yaml:"threats" json:"threats"`
	Vulnerabilities []string `yaml:"vulnerabilities" json:"vulnerabilities"`
	// Additional fields like owner, status, etc., can go here
}

// ControlItem represents a security control applied in a specific assessment.
// This might be a subset of the ControlCatalogItem or have assessment-specific fields.
type ControlItem struct {
	ID          string `yaml:"id" json:"id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Effectiveness int    `yaml:"effectiveness,omitempty" json:"effectiveness,omitempty"` // Per-assessment effectiveness (e.g., 1-5)
	// Status (e.g., implemented, partially implemented, planned)
}

// ControlCatalog represents a collection of default or predefined controls.
type ControlCatalog struct {
	Controls []ControlCatalogItem `yaml:"controls" json:"controls"`
}

// ControlCatalogItem defines a standard security control from a catalog.
type ControlCatalogItem struct {
	ID          string `yaml:"id" json:"id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Category    string `yaml:"category" json:"category"` // e.g., "Access Control", "Configuration Management"
	DefaultEffectiveness int `yaml:"default_effectiveness" json:"default_effectiveness"` // Default effectiveness if no specific assessment value
	Source      string `yaml:"source,omitempty" json:"source,omitempty"` // e.g., "NIST SP 800-53", "ISO 27002"
}
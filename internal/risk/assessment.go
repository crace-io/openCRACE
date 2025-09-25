package risk

import (
	"fmt"
	"io/ioutil" // Using ioutil for simplicity, but os.ReadFile is preferred in Go 1.16+

	"gopkg.in/yaml.v3" // For YAML unmarshalling
)

// LoadRiskAssessment reads and parses a risk assessment definition file (YAML or JSON).
func LoadRiskAssessment(filePath string) (*RiskAssessment, error) {
	data, err := ioutil.ReadFile(filePath) // os.ReadFile is preferred for Go 1.16+
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var assessment RiskAssessment
	// For now, assuming YAML. You'll add logic to detect file type later if needed.
	err = yaml.Unmarshal(data, &assessment)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML from %s: %w", filePath, err)
	}

	// TODO: Add schema validation here (using the static/schemas)

	return &assessment, nil
}

// CalculateInitialRisk calculates a simple initial risk score for the assessment.
// This is a basic multiplication. Real-world scoring can be more sophisticated.
func (ra *RiskAssessment) CalculateInitialRisk() int {
	totalRiskScore := 0
	for _, r := range ra.Risks {
		// A common simple risk calculation: Impact * Likelihood
		totalRiskScore += r.Impact * r.Likelihood
	}
	return totalRiskScore
}

// LoadControlCatalog loads a catalog of predefined controls from a YAML file.
func LoadControlCatalog(filePath string) (*ControlCatalog, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read control catalog file %s: %w", filePath, err)
	}

	var catalog ControlCatalog
	err = yaml.Unmarshal(data, &catalog)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal control catalog from %s: %w", filePath, err)
	}
	return &catalog, nil
}

// CalculateResidualRisk calculates the residual risk after applying controls.
// This is a simplified example. A real implementation would involve:
// 1. Mapping ControlItems in assessment to ControlCatalogItems for their default effectiveness.
// 2. Adjusting effectiveness based on assessment-specific overrides in ra.Controls.
// 3. Applying control effectiveness to reduce impact/likelihood of affected risks.
func (ra *RiskAssessment) CalculateResidualRisk(controlCatalog *ControlCatalog) int {
	// For demonstration, let's assume a very basic reduction
	initialRisk := ra.CalculateInitialRisk()
	if initialRisk == 0 {
		return 0
	}

	totalEffectiveness := 0
	for _, appliedControl := range ra.Controls {
		// Find the control in the catalog to get its default properties
		for _, catalogControl := range controlCatalog.Controls {
			if appliedControl.ID == catalogControl.ID {
				// Use the assessment's specific effectiveness if provided, otherwise the default
				effectiveness := appliedControl.Effectiveness
				if effectiveness == 0 {
					effectiveness = catalogControl.DefaultEffectiveness
				}
				totalEffectiveness += effectiveness
				break
			}
		}
	}

	// Very simple, conceptual residual risk calculation:
	// Total effectiveness could be mapped to a percentage reduction or a fixed deduction.
	// This needs careful definition in your risk model.
	// For now, let's say each point of effectiveness reduces risk by a small amount.
	reductionFactor := float64(totalEffectiveness) * 0.05 // Example: 5% reduction per effectiveness point
	if reductionFactor > 1.0 { // Cap reduction at 100%
		reductionFactor = 1.0
	}

	residualRisk := int(float64(initialRisk) * (1.0 - reductionFactor))
	if residualRisk < 0 { // Ensure risk doesn't go negative
		residualRisk = 0
	}
	return residualRisk
}
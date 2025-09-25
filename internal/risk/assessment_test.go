package risk

import (
	"io/ioutil"
	"os" // Import the strings package for checking error messages
	"testing"
)

// TestLoadRiskAssessment_ValidYAML tests loading a valid YAML risk definition.
func TestLoadRiskAssessment_ValidYAML(t *testing.T) {
	// Create a temporary valid YAML file
	validYAML := `
name: "Test Assessment"
description: "A simple test assessment."
risks:
  - id: "TEST-001"
    name: "Test Risk 1"
    description: "First test risk"
    impact: 3
    likelihood: 2
    affected_assets: ["AssetA"]
    threats: ["ThreatX"]
    vulnerabilities: ["VulnY"]
  - id: "TEST-002"
    name: "Test Risk 2"
    description: "Second test risk"
    impact: 5
    likelihood: 4
    affected_assets: ["AssetB"]
    threats: ["ThreatZ"]
    vulnerabilities: []
`
	tmpfile, err := ioutil.TempFile("", "valid_risk_*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	if _, err := tmpfile.WriteString(validYAML); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Load the assessment
	assessment, err := LoadRiskAssessment(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadRiskAssessment failed: %v", err)
	}

	// Assertions
	if assessment.Name != "Test Assessment" {
		t.Errorf("Expected Name 'Test Assessment', got '%s'", assessment.Name)
	}
	if len(assessment.Risks) != 2 {
		t.Errorf("Expected 2 risks, got %d", len(assessment.Risks))
	}
	if assessment.Risks[0].ID != "TEST-001" {
		t.Errorf("Expected first risk ID 'TEST-001', got '%s'", assessment.Risks[0].ID)
	}
	if assessment.Risks[1].Impact != 5 {
		t.Errorf("Expected second risk impact 5, got %d", assessment.Risks[1].Impact)
	}
	if len(assessment.Risks[0].AffectedAssets) != 1 || assessment.Risks[0].AffectedAssets[0] != "AssetA" {
		t.Errorf("Expected 'AssetA' in affected assets, got %v", assessment.Risks[0].AffectedAssets)
	}
}

// TestLoad
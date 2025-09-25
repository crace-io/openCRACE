package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// Import your internal packages here as you create them
	"github.com/crace-io/openCRACE/internal/risk"
	"github.com/crace-io/openCRACE/pkg/config"
)

var (
	cfgFile string // Path to an optional configuration file
	verbose bool   // Global verbose flag
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openCRACE",
	Short: "openCRACE: EU CRA Cybersecurity Risk Assessment CLI",
	Long: `openCRACE is an open-source command-line tool designed to facilitate
EU Cyber Resilience Act (CRA) compliant cybersecurity risk assessments.

It allows for defining risks, applying controls, and calculating residual risk scores,
supporting robust reporting and compliance efforts.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// This runs before any command.
		// Use it for global setup, e.g., loading configuration.
		if cfgFile != "" {
			fmt.Printf("Loading config from: %s\n", cfgFile)
		} else {
			// Try to load default config or set default values
			fmt.Println("No config file specified, using defaults or env vars.")
		}
		// Example of loading some basic config (implement in pkg/config)
		appConfig, err := config.LoadAppConfig(cfgFile) // Assuming this function exists
		if err != nil {
			return fmt.Errorf("failed to load application configuration: %w", err)
		}
		_ = appConfig // Use appConfig later, for now just load it.

		if verbose {
			fmt.Println("Verbose mode enabled.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is specified, print help.
		cmd.Help()
	},
}

// assessCmd represents the 'assess' command
var assessCmd = &cobra.Command{
	Use:   "assess <risk-definition-file>",
	Short: "Performs a cybersecurity risk assessment based on a definition file",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	RunE: func(cmd *cobra.Command, args []string) error {
		riskFile := args[0]
		fmt.Printf("Processing risk definition file: %s\n", riskFile)

		// 1. Load Risk Assessment from file
		assessment, err := risk.LoadRiskAssessment(riskFile)
		if err != nil {
			return fmt.Errorf("failed to load risk assessment from '%s': %w", riskFile, err)
		}

		// 2. Calculate Initial Risk
		initialScore := assessment.CalculateInitialRisk()
		fmt.Printf("Initial Risk Score: %d\n", initialScore)

		// TODO:
		// 3. Load default controls from static/controls/default_controls.yaml
		// 4. Allow user to select/override controls
		// 5. Calculate Residual Risk based on selected controls
		// 6. Generate reports (PDF, Excel, JSON, YAML)

		fmt.Println("Assessment process started. (Further steps to be implemented)")
		return nil
	},
}

// init function runs automatically before main()
func init() {
	cobra.OnInitialize(initConfig) // Function to run after flags are parsed

	// Define global flags (flags available to all commands)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.opencrace.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	// Add subcommands
	rootCmd.AddCommand(assessCmd)
	// You will add more subcommands here later, e.g.:
	// rootCmd.AddCommand(reportCmd)
	// rootCmd.AddCommand(controlCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// TODO: Implement actual config loading logic in pkg/config
	// For now, this is a placeholder.
	// You might use viper for robust config management (env vars, default paths, etc.)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %v\n", err)
		os.Exit(1)
	}
}
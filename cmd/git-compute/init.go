package git_compute

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration file",
	Long: `Interactively create a .git-compute.yaml configuration file 
in your home directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		var token, owner, repo string

		fmt.Print("Enter GitHub Personal Access Token (repo scope): ")
		fmt.Scanln(&token)

		fmt.Print("Enter Repository Owner (e.g., 10xdev4u-alt): ")
		fmt.Scanln(&owner)

		fmt.Print("Enter Repository Name (e.g., cicd-as-a-service): ")
		fmt.Scanln(&repo)

		if token == "" || owner == "" || repo == "" {
			fmt.Println("❌ All fields are required.")
			return
		}

		viper.Set("github_token", token)
		viper.Set("owner", owner)
		viper.Set("repo", repo)
		viper.Set("workflow_id", "worker.yml")

		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("❌ Could not find home directory: %v\n", err)
			return
		}

		configPath := filepath.Join(home, ".git-compute.yaml")
		err = viper.WriteConfigAs(configPath)
		if err != nil {
			fmt.Printf("❌ Failed to write config: %v\n", err)
			return
		}

		fmt.Printf("✅ Configuration saved to %s\n", configPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

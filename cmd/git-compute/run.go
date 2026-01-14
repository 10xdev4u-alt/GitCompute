package git_compute

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/10xdev4u-alt/git-compute/pkg/config"
	"github.com/10xdev4u-alt/git-compute/pkg/runner"
	"github.com/10xdev4u-alt/git-compute/pkg/utils"
)

var (
	cmdFlag   string
	osFlag    string
	watchFlag bool
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a command on a remote runner",
	Long: `Trigger a GitHub Action to run the specified command on a remote VM.
Example: git-compute run --cmd "echo Hello" --os ubuntu-latest`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			fmt.Println("Please ensure .git-compute.yaml exists with github_token, owner, and repo.")
			os.Exit(1)
		}

		client := runner.NewClient(cfg.GitHubToken, cfg.Owner, cfg.Repo)
		ctx := context.Background()

		fmt.Printf("🚀 Triggering command '%s' on %s...\n", cmdFlag, osFlag)
		
		inputs := map[string]interface{}{
			"command": cmdFlag,
			"os":      osFlag,
		}

		err = client.TriggerWorkflow(ctx, cfg.WorkflowID, inputs)
		if err != nil {
			fmt.Printf("❌ Error triggering workflow: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ Workflow dispatched successfully.")

		if watchFlag {
			fmt.Println("👀 Waiting for run to start...")
			run, err := client.WaitForRun(ctx, cfg.WorkflowID)
			if err != nil {
				fmt.Printf("⚠️ Could not track run: %v\n", err)
				return
			}
			
			fmt.Printf("🏃 Run started: %s (ID: %d)\n", run.GetHTMLURL(), run.GetID())
			fmt.Println("⏳ Polling status...")

			for {
				r, err := client.GetRunStatus(ctx, run.GetID())
				if err != nil {
					fmt.Printf("Error polling status: %v\n", err)
					break
				}

				status := r.GetStatus()
				conclusion := r.GetConclusion()

				fmt.Printf("\rStatus: %s | Conclusion: %s   ", status, conclusion)

				if status == "completed" {
					fmt.Println("\n🏁 Run completed!")
					if conclusion == "success" {
						fmt.Println("🎉 Success! Downloading results...")
						
						zipName := fmt.Sprintf("result-%d.zip", run.GetID())
						err := client.DownloadArtifact(ctx, run.GetID(), "execution-result", zipName)
						if err != nil {
							fmt.Printf("❌ Failed to download artifact: %v\n", err)
							break
						}
						
						fmt.Printf("📦 Artifact downloaded to %s. Unzipping...\n", zipName)
						
						// Unzip to current directory
						err = utils.Unzip(zipName, ".")
						if err != nil {
							fmt.Printf("⚠️ Failed to unzip: %v\n", err)
						} else {
							fmt.Println("📂 Results extracted to output.log")
							
							// Read and display content
							content, _ := os.ReadFile("output.log")
							fmt.Println("\n--- Remote Output ---")
							fmt.Println(string(content))
							fmt.Println("---------------------")
						}
						
						// Clean up zip
						os.Remove(zipName)
						
					} else {
						fmt.Println("❌ Failed.")
					}
					break
				}

				time.Sleep(3 * time.Second)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&cmdFlag, "cmd", "c", "echo Hello World", "Command to run remotely")
	runCmd.Flags().StringVarP(&osFlag, "os", "o", "ubuntu-latest", "Operating System (ubuntu-latest, windows-latest, macos-latest)")
	runCmd.Flags().BoolVarP(&watchFlag, "watch", "w", true, "Watch execution status")
	
	// Bind flags to viper if needed, but here we just use flags directly for the run command
}

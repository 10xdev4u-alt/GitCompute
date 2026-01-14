package runner

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

type Client struct {
	ghClient *github.Client
	owner    string
	repo     string
}

func NewClient(token, owner, repo string) *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	ghClient := github.NewClient(tc)

	return &Client{
		ghClient: ghClient,
		owner:    owner,
		repo:     repo,
	}
}

func (c *Client) TriggerWorkflow(ctx context.Context, workflowID string, inputs map[string]interface{}) error {
	// inputs for CreateWorkflowDispatchEvent must be map[string]interface{} but the value should be string/int/bool
	// The library expects map[string]interface{}
	
	ref := "main" // Trigger on main branch by default

	opts := github.CreateWorkflowDispatchEventRequest{
		Ref:    ref,
		Inputs: inputs,
	}

	_, _, err := c.ghClient.Actions.CreateWorkflowDispatchEvent(ctx, c.owner, c.repo, workflowID, opts)
	if err != nil {
		return fmt.Errorf("failed to dispatch workflow: %w", err)
	}
	return nil
}

// WaitForRun locates the run triggered by our dispatch.
// Since dispatch doesn't return the Run ID, we have to query for recent runs and match.
// This is a known limitation of the GitHub API.
func (c *Client) WaitForRun(ctx context.Context, workflowID string) (*github.WorkflowRun, error) {
	// Wait a moment for the run to appear
	time.Sleep(3 * time.Second)

	opts := &github.ListWorkflowRunsOptions{
		Event: "workflow_dispatch",
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 1, // Get the most recent one
		},
	}

	// Simple retry loop to find a very recent run (created in the last minute)
	for i := 0; i < 5; i++ {
		runs, _, err := c.ghClient.Actions.ListWorkflowRunsByFileName(ctx, c.owner, c.repo, workflowID, opts)
		if err != nil {
			return nil, err
		}

		if len(runs.WorkflowRuns) > 0 {
			run := runs.WorkflowRuns[0]
			// Check if it was created very recently (e.g. within last 10 seconds) to ensure it's ours
			// This is heuristic-based because of API limitations
			if time.Since(run.GetCreatedAt().Time) < 1*time.Minute {
				return run, nil
			}
		}
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("timeout waiting for workflow run to appear")
}

func (c *Client) GetRunStatus(ctx context.Context, runID int64) (*github.WorkflowRun, error) {
	run, _, err := c.ghClient.Actions.GetWorkflowRunByID(ctx, c.owner, c.repo, runID)
	return run, err
}

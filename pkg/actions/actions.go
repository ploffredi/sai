package actions

import (
	"errors"
	"fmt"
	"sai/pkg/providers"
)

// ExecuteAction executes a specific action for a given software and provider.
// If software is '*', the action is executed on all managed software.
func ExecuteAction(software, action, provider string) error {
	if software == "*" {
		// Handle wildcard: execute action on all managed software
		allSoftware := providers.GetAllManagedSoftware()
		if len(allSoftware) == 0 {
			return errors.New("no managed software found")
		}

		for _, sw := range allSoftware {
			err := ExecuteAction(sw, action, provider)
			if err != nil {
				fmt.Printf("Error executing action '%s' on software '%s': %v\n", action, sw, err)
			}
		}
		return nil
	}

	if provider == "" {
		switch action {
		case "install", "test", "build", "log", "check", "observe", "trace", "config", "info", "debug", "troubleshoot", "monitor", "upgrade", "uninstall", "status", "start", "stop", "restart", "enable", "disable", "list", "search", "update", "ask", "help":
			return errors.New("provider is required for this action")
		default:
			fmt.Printf("Executing action '%s' on '%s' without a specific provider\n", action, software)
			return nil
		}
	}

	p := providers.GetProvider(provider)
	if p == nil {
		return fmt.Errorf("provider '%s' not found", provider)
	}

	// Call specific methods based on the action.
	switch action {
	// ...existing code for other actions...
	default:
		return fmt.Errorf("action '%s' is not supported", action)
	}
}

// ExecuteActionWithDefault executes an action using a default provider if none is specified.
func ExecuteActionWithDefault(software, action, provider string) error {
	if provider == "" {
		provider = "default"
	}
	return ExecuteAction(software, action, provider)
}

// ...existing code...
func IsActionSupported(provider, action string) bool {
	// Check if the provider supports the action
	return true // placeholder
}

// Used in main.go
func PerformAction(software, provider, action string) error {
	// Run the action using the specified provider, if supported
	// Otherwise, return an error
	if IsActionSupported(provider, action) {
		// Check if the action is supported by the provider
		if !IsActionSupported(provider, action) {
			return fmt.Errorf("action '%s' is not supported by provider '%s'", action, provider)
		}
		// Run the a function called $actionProvider,
		// if not available generate an error
		// and return the error	

   

/*
		switch action {
		case "install":
			return action.Install(software, provider)

    case "test":
			return providers.Test(software, provider)
		case "build":
			return providers.Build(software, provider)
		case "log":
			return providers.Log(software, provider)
		case "check":
			return providers.Check(software, provider)
		case "observe":
			return providers.Observe(software, provider)
		case "trace":
			return providers.Trace(software, provider)
		case "config":
			return providers.Config(software, provider)
		case "info":
			return providers.Info(software, provider)
		case "debug":
			return providers.Debug(software, provider)
		case "troubleshoot":
			return providers.Troubleshoot(software, provider)
		case "monitor":
			return providers.Monitor(software, provider)
		case "upgrade":
			return providers.Upgrade(software, provider)
		case "uninstall":
			return providers.Uninstall(software, provider)
		case "status":
			return providers.Status(software, provider)
		case "start":
			return providers.Start(software, provider)
		case "stop":
			return providers.Stop(software, provider)
		case "restart":
			return providers.Restart(software, provider)
		case "enable":
			return providers.Enable(software, provider)
		case "disable":
			return providers.Disable(software, provider)
		case "list":
			return providers.List(software, provider)
		case "search":
			return providers.Search(software, provider)
		case "update":
			return providers.Update(software, provider)
		case "ask":
			return providers.Ask(software, provider)
		case "help":
			return providers.Help(software, provider)
		default:
			return fmt.Errorf("action '%s' is not supported", action)
		}
	}
*/
	}
  return fmt.Errorf("Request support with: sai %s support %s %s ", software, action, provider)
}

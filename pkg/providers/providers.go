package providers

import (
	"fmt"
	"sync"
)

// ActionFunc defines the signature for an action function.
type ActionFunc func(software string) error

// Provider represents a collection of actions.
type Provider struct {
	actions map[string]ActionFunc
}

// RegisterAction registers an action for the provider.
func (p *Provider) RegisterAction(actionName string, action ActionFunc) {
	if p.actions == nil {
		p.actions = make(map[string]ActionFunc)
	}
	p.actions[actionName] = action
}

// ExecuteAction executes the specified action for the provider.
func (p *Provider) ExecuteAction(actionName, software string) error {
	if action, exists := p.actions[actionName]; exists {
		return action(software)
	}
	return fmt.Errorf("action '%s' is not supported for this provider", actionName)
}

// ProviderRegistry manages the registration of providers.
type ProviderRegistry struct {
	mu        sync.RWMutex
	providers map[string]*Provider
}

// NewProviderRegistry creates a new ProviderRegistry.
func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{
		providers: make(map[string]*Provider),
	}
}

// RegisterProvider registers a provider with the registry.
func (r *ProviderRegistry) RegisterProvider(name string, provider *Provider) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.providers[name] = provider
}

// GetProvider retrieves a provider by name. Returns a default provider if not found.
func (r *ProviderRegistry) GetProvider(name string) *Provider {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if provider, exists := r.providers[name]; exists {
		return provider
	}
	return &Provider{} // Return a default provider with no actions.
}

// Global registry instance
var registry = NewProviderRegistry()

// RegisterProvider is a helper function to register a provider globally.
func RegisterProvider(name string, provider *Provider) {
	registry.RegisterProvider(name, provider)
}

// GetProvider is a helper function to retrieve a provider globally.
func GetProvider(name string) *Provider {
	return registry.GetProvider(name)
}

// ExecuteAction is a helper function to execute an action for a provider globally.
func ExecuteAction(providerName, actionName, software string) error {
	provider := GetProvider(providerName)
	err := provider.ExecuteAction(actionName, software)
	if err != nil {
		return fmt.Errorf("error executing action '%s' for provider '%s': %w", actionName, providerName, err)
	}
	return nil
}

// GetAllManagedSoftware returns a list of all managed software.
func GetAllManagedSoftware() []string {
	// Example: Return a static list of managed software.
	return []string{"nginx", "docker", "mysql", "redis", "kubernetes"}
}

// Example usage of dynamic registration
func init() {
	// Example: Registering a "brew" provider with some actions.
	brewProvider := &Provider{}
	brewProvider.RegisterAction("install", func(software string) error {
		fmt.Printf("Installing %s using brew...\n", software)
		return nil
	})
	brewProvider.RegisterAction("observe", func(software string) error {
		fmt.Printf("Observing %s using brew...\n", software)
		return nil
	})
	RegisterProvider("brew", brewProvider)

	// Example: Registering a "rpm" provider with some actions.
	rpmProvider := &Provider{}
	rpmProvider.RegisterAction("install", func(software string) error {
		fmt.Printf("Installing %s using rpm...\n", software)
		return nil
	})
	RegisterProvider("rpm", rpmProvider)
}

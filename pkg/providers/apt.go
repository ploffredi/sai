package providers

import (
	"fmt"
	"os/exec"
)

type AptProvider struct{}

func (p *AptProvider) Install(software string) error {
	fmt.Printf("Installing %s using apt...\n", software)
	cmd := exec.Command("apt", "install", "-y", software)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install %s: %v, output: %s", software, err, string(output))
	}
	fmt.Printf("Successfully installed %s\n", software)
	return nil
}

func (p *AptProvider) Test(software string) error {
	fmt.Printf("Testing %s using apt...\n", software)
	return nil
}

func (p *AptProvider) Build(software string) error {
	fmt.Printf("Building %s using apt...\n", software)
	return nil
}

func (p *AptProvider) Log(software string) error {
	fmt.Printf("Retrieving logs for %s using apt...\n", software)
	return nil
}

func (p *AptProvider) Check(software string) error {
	fmt.Printf("Checking %s using apt...\n", software)
	return nil
}

func (p *AptProvider) Observe(software string) error {
	fmt.Printf("Observing %s using apt...\n", software)
	return nil
}

func (p *AptProvider) Trace(software string) error {
	fmt.Printf("Tracing %s using apt...\n", software)
	return nil
}

func (p *AptProvider) Config(software string) error {
	fmt.Printf("Configuring %s using apt...\n", software)
	return nil
}

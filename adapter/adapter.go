package adapter

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os/exec"
	"strings"
	"sync"
)

const (
	registryPath = `SYSTEM\CurrentControlSet\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}`
)

type AdapterInfo struct {
	Name        string `json:"Name"`
	Description string `json:"InterfaceDescription"`
	MacAddress  string `json:"MacAddress"`
}

// GenerateMac generates a random MAC address
func GenerateMac() (string, error) {
	mac := make([]byte, 6)
	_, err := rand.Read(mac)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}
	mac[0] = (mac[0] | 2) & 0xFe // set locally administered and unicast bit
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5]), nil
}

// GetOriginalMAC retrieves the current MAC address of the specified adapter
func GetOriginalMAC(adapterName string) (string, error) {
	cmd := exec.Command("powershell", "-Command",
		fmt.Sprintf(`(Get-NetAdapter -Name "%s").MacAddress`, adapterName))
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get original MAC: %v", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// ListAdapters returns a list of all network adapters
func ListAdapters() ([]AdapterInfo, error) {
	cmd := exec.Command("powershell", "-Command",
		`Get-NetAdapter | Select-Object Name, InterfaceDescription, MacAddress | ConvertTo-Json`)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list adapters: %v", err)
	}

	var adapters []AdapterInfo
	if err := json.Unmarshal(out, &adapters); err != nil {
		return nil, fmt.Errorf("failed to parse adapter info: %v", err)
	}

	var wg sync.WaitGroup

	for i := range adapters {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			if adapters[i].MacAddress == "" {
				adapters[i].MacAddress = "-"
			}
		}(i)
	}

	wg.Wait()

	return adapters, nil
}

// findAdapterRegistry finds the registry key for the specified adapter
func findAdapterRegistry(adapterName string) (registry.Key, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, registryPath, registry.ALL_ACCESS)
	if err != nil {
		return 0, fmt.Errorf("failed to open registry key: %v", err)
	}
	defer k.Close()

	subkeys, err := k.ReadSubKeyNames(0)
	if err != nil {
		return 0, fmt.Errorf("failed to read subkeys: %v", err)
	}

	for _, subkey := range subkeys {
		fullPath := registryPath + "\\" + subkey
		sk, err := registry.OpenKey(registry.LOCAL_MACHINE, fullPath, registry.ALL_ACCESS)
		if err != nil {
			continue
		}

		driverDesc, _, err := sk.GetStringValue("DriverDesc")
		if err != nil || !strings.Contains(driverDesc, adapterName) {
			sk.Close()
			continue
		}
		return sk, nil
	}

	return 0, fmt.Errorf("adapter not found in registry")
}

// ChangeMACAddress changes the MAC address of the specified adapter
func ChangeMACAddress(adapterName, newMAC string) error {
	sk, err := findAdapterRegistry(adapterName)
	if err != nil {
		return err
	}
	defer sk.Close()

	// Backup original MAC if not already saved
	_, _, err = sk.GetStringValue("OriginalMAC")
	if err != nil {
		currentMAC, err := GetOriginalMAC(adapterName)
		if err != nil {
			return fmt.Errorf("failed to backup original MAC: %v", err)
		}
		if err := sk.SetStringValue("OriginalMAC", currentMAC); err != nil {
			return fmt.Errorf("failed to save original MAC: %v", err)
		}
	}

	// Set the new MAC address
	if err := sk.SetStringValue("NetworkAddress", strings.ReplaceAll(newMAC, ":", "")); err != nil {
		return fmt.Errorf("failed to set new MAC: %v", err)
	}

	return RestartAdapter(adapterName)
}

// RestoreOriginalMAC restores the original MAC address of the specified adapter
func RestoreOriginalMAC(adapterName string) error {
	sk, err := findAdapterRegistry(adapterName)
	if err != nil {
		return err
	}
	defer sk.Close()

	_, _, err = sk.GetStringValue("OriginalMAC")
	if err != nil {
		return fmt.Errorf("original MAC not found: %v", err)
	}

	if err := sk.DeleteValue("NetworkAddress"); err != nil && !errors.Is(err, registry.ErrNotExist) {
		return fmt.Errorf("failed to remove custom MAC: %v", err)
	}

	return RestartAdapter(adapterName)
}

// RestartAdapter disables and re-enables the network adapter
func RestartAdapter(adapterName string) error {
	// Disable adapter
	cmd := exec.Command("powershell", "-Command",
		fmt.Sprintf(`Disable-NetAdapter -Name "%s" -Confirm:$false`, adapterName))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to disable adapter: %v", err)
	}

	// Enable adapter
	cmd = exec.Command("powershell", "-Command",
		fmt.Sprintf(`Enable-NetAdapter -Name "%s" -Confirm:$false`, adapterName))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to enable adapter: %v", err)
	}

	return nil
}

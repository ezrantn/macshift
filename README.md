# Macshift

## Overview

Macshift is a simple command-line tool designed to help users easily manage and manipulate MAC addresses of network adapters on **Windows**. It allows you to list available network adapters, change the MAC address of a specified adapter to a randomly generated one, and restore the original MAC address of the adapter. This tool is especially useful when you need to troubleshoot or prevent issues caused by fixed MAC addresses (e.g., network lockouts or privacy concerns).

## Problem Statement

In many situations, a network device (e.g., laptop or PC) might get locked out of a network due to its MAC address being stuck on a switch or router. This can occur when a user connects to a neighbor's Wi-Fi and then the adapter's MAC address is associated with a specific port, preventing the device from reconnecting to other networks or causing issues with network access. Additionally, some users may want to mask their original MAC addresses for privacy or security reasons, especially when using public Wi-Fi networks.

### Why Do We Need Macshift?

- Network Access Issues: Sometimes, when a device's MAC address gets stuck on a network switch, it may prevent the device from reconnecting to a network. Macshift can change the MAC address to resolve such issues.

- Privacy Concerns: Changing your MAC address can help ensure anonymity on public or shared networks, preventing tracking of your device's unique identifier.

- Convenience: This tool automates the process of generating random MAC addresses, making it easy for users to change or restore their MAC addresses without having to manually modify network settings or dive into advanced configurations.

## Installation

Run the following command:

```shell
go install github.com/ezrantn/macshift@latest
```

## Commands and Usage

`macshift list`

- Description: Lists all available network adapters along with their names, descriptions, and MAC addresses.

**Usage**

```shell
macshift list
```

**Example Output**

```shell
Available network adapters:

Name: Local Area Connection
Description: TAP-Windows Adapter V9 for OpenVPN Connect
MAC: XX-XX-XX-XX-XX-XX

Name: Wi-Fi
Description: Killer(R) Wi-Fi 6 AX1650i 160MHz Wireless Network Adapter
MAC: XX-XX-XX-XX-XX-XX
```

`macshift change -i "Adapter Name"`

- Description: Generates a random MAC address and changes the MAC address of the specified adapter. The `-i` flag is used to specify the adapter name.

**Usage**

```shell
macshift change -i "Wi-Fi"
```

**Example Output**

```shell
Generated random MAC address: 26:a5:a3:5d:b3:39
MAC address changed successfully to 26:a5:a3:5d:b3:39 on interface Wi-Fi
```

This command will generate a new, random MAC address for the specified network adapter (e.g., "Wi-Fi") and apply the change.

`macshift restore -i "Adapter Name"`

- Description: Restores the original MAC address of the specified adapter. The `-i` flag is used to specify the adapter name.

**Usage**

```shell
macshift restore -i "Wi-Fi"
```

**Example Output**

```shell
Original MAC address restored successfully on interface Wi-Fi
```

This command will revert the MAC address of the "Wi-Fi" adapter to its original state that was previously recorded before any changes were made.

## Important Notes

1. Running as Administrator

    - Press `Windows + X`
    - Click "Windows PowerShell (Admin)"
    - Click "Yes" when prompted

2. Network Changes

    - Your network will disconnect briefly when changing MAC addresses
    - This is normal - just wait a moment
    - You may need to reconnect to Wi-Fi

## Credits

This tool was inspired by a video by [Dr. Jonas Birch](https://www.youtube.com/@dr-Jonas-Birch) where he demonstrates how to change and restore a network adapter's MAC address using the C programming language. The concept and the core approach were implemented in Go for simplicity and portability.

You can find Dr. Jonas Birch's video [here](https://www.youtube.com/watch?v=n4t13B7xVJM&t=734s).

## License

This tool is open-source and available under the [MIT](https://github.com/ezrantn/macshift/LICENSE) License.

## Contributions

Contributions are welcome! Please feel free to submit a pull request.
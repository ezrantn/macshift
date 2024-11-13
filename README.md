# Macshift

## Overview

A command-line tool for managing MAC addresses of network adapters on Windows.

## What is Macshift

Macshift helps you change the MAC addresses of your network adapters. It's especially useful when:
- Your device is having trouble connecting to networks (e.g., MAC address stuck on a switch/router)
- You want to enhance your privacy on public networks
- You need to bypass MAC address filtering
- You're testing network security configurations

## Prerequisites

- Windows operating system
- Go 1.23 or higher
- Administrator privileges
- Network adapter that supports MAC address changes

## Installation

Run the following command:

```shell
go install github.com/ezrantn/macshift@latest
```

**Important**: Always run `macshift` as Administrator. To do this:
1. Press `Windows + X`
2. Select "Terminal (Admin)"
3. Click "Yes" when prompted

## Commands 

### List Network Adapters

Shows all available network adapters with their names and MAC addresses.

```shell
macshift list
```

**Example Output**

```shell
Available network adapters:

Name: Wi-Fi
Description: Killer(R) Wi-Fi 6 AX1650i Wireless Network Adapter
MAC: XX-XX-XX-XX-XX-XX

Name: Local Area Connection
Description: TAP-Windows Adapter V9
MAC: XX-XX-XX-XX-XX-XX
```

### Change MAC Address

Generates and applies a random MAC address for the specified adapter.

```shell
macshift change -i "Wi-Fi"
```

**Example Output**

```shell
Generated random MAC address: 26:a5:a3:5d:b3:39
MAC address changed successfully to 26:a5:a3:5d:b3:39 on interface Wi-Fi
```

### Restore Original MAC Address

Returns the adapter to its original MAC address.

```shell
macshift restore -i "Wi-Fi"
```

**Example Output**

```shell
Original MAC address restored successfully on interface Wi-Fi
```

## Important Notes

1. Network Changes

   - Your network will disconnect briefly when changing MAC addresses
   - This is normal - just wait a moment
   - You may need to reconnect to Wi-Fi


2. Some network adapters don't support MAC address changes

## Credits

This tool was inspired by a video by [Dr. Jonas Birch](https://www.youtube.com/@dr-Jonas-Birch) where he demonstrates how to change and restore a network adapter's MAC address using the C programming language. The concept and the core approach were implemented in Go for simplicity and portability.

You can find Dr. Jonas Birch's video [here](https://www.youtube.com/watch?v=n4t13B7xVJM&t=734s).

## License

This tool is open-source and available under the [MIT](https://github.com/ezrantn/macshift/LICENSE) License.

## Contributions

Contributions are welcome! Please feel free to submit a pull request.
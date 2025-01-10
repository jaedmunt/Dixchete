package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Command represents a single command entry in our cheat sheet.
type Command struct {
	Name          string
	Description   string
	Usage         string
	ExampleOutput string
}

// categories holds our commands grouped by category.
var categories = map[string][]Command{
	"Storage Commands": {
		{
			Name:        "wmic diskdrive get model,name,size,mediaType",
			Description: "Displays the model, name, size, and media type of all disk drives.",
			Usage:       "Use this command to get detailed information about the disk drives connected to the system.",
			ExampleOutput: `MediaType              Model                       Name                Size
Fixed hard disk media  ST1000DM003-1SB10C          \\.\PHYSICALDRIVE0  1000202273280
Fixed hard disk media  addlink M.2 PCIE G3x4 NVMe  \\.\PHYSICALDRIVE1  256052966400`,
		},
		{
			Name:        "wmic diskdrive list brief",
			Description: "Lists brief information about all disk drives.",
			Usage:       "Use this command to get a quick overview of the disk drives connected to the system.",
			ExampleOutput: `Caption                     DeviceID            Model                       Partitions  Size
ST1000DM003-1SB10C          \\.\PHYSICALDRIVE0  ST1000DM003-1SB10C          3           1000202273280
addlink M.2 PCIE G3x4 NVMe  \\.\PHYSICALDRIVE1  addlink M.2 PCIE G3x4 NVMe  3           256052966400`,
		},
	},
	"System Information Commands": {
		{
			Name:        "systeminfo",
			Description: "Displays detailed configuration information about the computer and its operating system.",
			Usage:       "Use this command to get a complete overview of the system's hardware and software configuration.",
		},
		{
			Name:        "wmic baseboard get product,Manufacturer",
			Description: "Displays the manufacturer and product name of the baseboard (motherboard).",
			Usage:       "Use this command to get information about your motherboard.",
			ExampleOutput: `Manufacturer                        Product
Micro-Star International Co., Ltd.  MPG Z490 GAMING PLUS (MS-7C75)`,
		},
	},
	"PCI & Device Commands": {
		{
			Name:        `wmic path win32_pnpentity where "PNPDeviceID like '%PCI%'" get Name,DeviceID`,
			Description: "Lists the names and device IDs of all Plug and Play entities with PCI in their device ID.",
			Usage:       "Use this command to get information about PCI devices.",
			ExampleOutput: `DeviceID                                                        Name
PCI\VEN_8086&DEV_06E0&SUBSYS_7C751462&REV_00\3&11583659&0&B0    Intel(R) Management Engine Interface #1
PCI\VEN_10DE&DEV_1C82&SUBSYS_33511462&REV_A1\4&F8D4272&0&0008   NVIDIA GeForce GTX 1050 Ti
...`,
		},
	},
	"Network Commands": {
		{
			Name:        "ipconfig",
			Description: "Displays all current TCP/IP network configuration values.",
			Usage:       "Use this command to view network adapter settings (IPv4, IPv6, DNS, etc.).",
		},
		{
			Name:        "ping <hostname or IP>",
			Description: "Sends ICMP ECHO_REQUEST packets to network hosts.",
			Usage:       "Use this command to check the connectivity to a specific host or IP.",
		},
		{
			Name:        "tracert <hostname or IP>",
			Description: "Prints the route (path) that packets take to a network host.",
			Usage:       "Use this command to diagnose where network problems might be occurring along the route.",
		},
	},
}

// printCategories prints the list of available categories.
func printCategories() {
	fmt.Println("Available Command Categories:")
	idx := 1
	for cat := range categories {
		fmt.Printf("  %d. %s\n", idx, cat)
		idx++
	}
	fmt.Println()
}

// printCommandsInCategory prints all commands for the specified category name.
func printCommandsInCategory(categoryName string) {
	cmds, exists := categories[categoryName]
	if !exists {
		fmt.Printf("Category '%s' does not exist.\n", categoryName)
		return
	}
	fmt.Printf("\n=== %s ===\n\n", categoryName)
	for _, c := range cmds {
		fmt.Printf("Command:      %s\n", c.Name)
		fmt.Printf("Description:  %s\n", c.Description)
		fmt.Printf("Usage:        %s\n", c.Usage)
		if c.ExampleOutput != "" {
			fmt.Println("Example Output:")
			for _, line := range strings.Split(c.ExampleOutput, "\n") {
				fmt.Printf("  %s\n", line)
			}
		}
		fmt.Println("------------------------------------------------------\n")
	}
}

// findCategoryByPartialName attempts to match the user input to a category.
func findCategoryByPartialName(input string) (string, bool) {
	input = strings.ToLower(input)
	for cat := range categories {
		if strings.Contains(strings.ToLower(cat), input) {
			return cat, true
		}
	}
	return "", false
}

func main() {
	args := os.Args
	// If no arguments are provided, show category list and prompt user.
	if len(args) == 1 {
		printCategories()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the number or exact name of the category (or 'q' to quit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.ToLower(text) == "q" {
			return
		}

		// Check if user typed a number or a category name
		num, err := strconv.Atoi(text)
		if err == nil {
			// It's a number, find the corresponding category
			if num < 1 || num > len(categories) {
				fmt.Println("Invalid category number.")
				return
			}
			// Convert categories map to slice for indexing
			i := 1
			for cat := range categories {
				if i == num {
					printCommandsInCategory(cat)
					return
				}
				i++
			}
		} else {
			// It's possibly a category name
			// Exact match first
			if _, exists := categories[text]; exists {
				printCommandsInCategory(text)
				return
			}
			// Fallback to partial match
			cat, ok := findCategoryByPartialName(text)
			if ok {
				printCommandsInCategory(cat)
				return
			}
			fmt.Printf("Category '%s' not found.\n", text)
		}
	} else {
		// If arguments are provided, try to interpret them as a category name
		userInput := strings.Join(args[1:], " ")
		// Exact match
		if _, exists := categories[userInput]; exists {
			printCommandsInCategory(userInput)
			return
		}
		// Partial match
		cat, ok := findCategoryByPartialName(userInput)
		if ok {
			printCommandsInCategory(cat)
		} else {
			fmt.Printf("No category found matching '%s'.\n", userInput)
		}
	}
}

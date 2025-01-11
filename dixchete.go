package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/qeesung/image2ascii/convert"

	"bytes"
	_ "embed"
	"image/gif"
)

// This is the path to the dope ASCII art of my face
// I tried quite hard to make it refresh but the clear effect is quite displeasing and reduces the quality of the actual tool

//go:embed images/dixchete.gif
var dixcheteGif []byte

func displayASCIIArt() {
	img, err := gif.Decode(bytes.NewReader(dixcheteGif))
	if err != nil {
		fmt.Println("Failed to decode GIF:", err)
		return
	}
	converter := convert.NewImageConverter()
	fmt.Println(converter.Image2ASCIIString(img, &convert.Options{
		Ratio: 1.0, // Adjust the ratio as needed
	}))
}

//left the variable name as jaedonAvatarJpg but I made it equal to a gif
// My face wobbles in it so each time you run the program it looks a little different

//go:embed images/jaedon_avatar_wobble.gif
var jaedonAvatarJpg []byte

func displayJaedonAvatar() {
	img, err := gif.Decode(bytes.NewReader(jaedonAvatarJpg))
	if err != nil {
		fmt.Println("Failed to decode JPEG:", err)
		return
	}
	converter := convert.NewImageConverter()
	fmt.Println(converter.Image2ASCIIString(img, &convert.Options{
		Ratio:   1.0,  // No clue what this really does. It has been redundant in my testing
		Colored: true, // Enable colored ASCII art - it was very boring before
	}))
	fmt.Println("\n\033[34m--------------------------------------\033[0m")
	fmt.Println("Local Current Time:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("\033[--------------------------------------]\033")
	fmt.Println("\033[34mAuthor: Jaedon Munton\033[0m")
	fmt.Println("\n\033[34mWebsite: https://www.jaedonmunton.pro/\033[0m")
	fmt.Println("\033[34mGithub: https://www.github.com/jaedmunt\033[0m")
	fmt.Println("\033[34mLinkedIn: https://www.linkedin.com/in/jaedonmunton/\033[0m")
}

func init() {
	fmt.Println("\033[34m--------------------------------------\033[0m")
	displayASCIIArt()
	fmt.Println("\033[34m--------------------------------------\033[0m")
	displayJaedonAvatar()
	fmt.Println("\033[34m--------------------------------------\033[0m")

}

// Command represents a single command entry in our cheat sheet.
type Command struct {
	Name          string
	Description   string
	Usage         string
	ExampleOutput string
}

// categories holds our commands grouped by category. Left it as a var so I can add more commands later
// This is quite easy to change so if there is any interest at all, I will expand it
var categories = map[string][]Command{
	"Storage Commands": {
		{
			Name:        "\033[wmic diskdrive get model,name,size,mediaType]\033",
			Description: "Displays the model, name, size, and media type of all disk drives.",
			Usage:       "Use this command to get detailed information about the disk drives connected to the system.",
			ExampleOutput: `MediaType              Model                       Name                Size
Fixed hard disk media  ST1000DM003-1SB10C          \\.\PHYSICALDRIVE0  1000202273280
Fixed hard disk media  addlink M.2 PCIE G3x4 NVMe  \\.\PHYSICALDRIVE1  256052966400`,
		},
		{
			Name:        "\033[wmic diskdrive list brief]\033",
			Description: "Lists brief information about all disk drives.",
			Usage:       "Use this command to get a quick overview of the disk drives connected to the system.",
			ExampleOutput: `Caption                     DeviceID            Model                       Partitions  Size
ST1000DM003-1SB10C          \\.\PHYSICALDRIVE0  ST1000DM003-1SB10C          3           1000202273280
addlink M.2 PCIE G3x4 NVMe  \\.\PHYSICALDRIVE1  addlink M.2 PCIE G3x4 NVMe  3           256052966400`,
		},
	},
	"System Information Commands": {
		{
			Name:        "\033[systeminfo]\033",
			Description: "Displays detailed configuration information about the computer and its operating system.",
			Usage:       "Use this command to get a complete overview of the system's hardware and software configuration.",
		},
		{
			Name:        "\033[wmic baseboard get product,Manufacturer]\033",
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
			Name:        "\033[ipconfig]\033",
			Description: "Displays all current TCP/IP network configuration values.",
			Usage:       "Use this command to view network adapter settings (IPv4, IPv6, DNS, etc.).",
		},
		{
			Name:        "\033[ping <hostname or IP>]\033",
			Description: "Sends ICMP ECHO_REQUEST packets to network hosts.",
			Usage:       "Use this command to check the connectivity to a specific host or IP.",
		},
		{
			Name:        "\033[tracert <hostname or IP>]\033",
			Description: "Prints the route (path) that packets take to a network host.",
			Usage:       "Use this command to diagnose where network problems might be occurring along the route.",
		},
	},
}

// printCategories prints the list of available categories. I thought I would clarify that 0_0

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
		fmt.Println("------------------------------------------------------")
	}
}

// findCategoryByPartialName attempts to match the user input to a category.
// This is the coolest function. I had to idiot proof it because I can't touch type lol
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
	// If no arguments are provided, show category list and prompt user to quit
	// A bit like a bad barrister except no category lists
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
		//This pairs up with the partial naming so it is idiot proof squared because I also have fat fingers
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

<p align="center">
    <strong><span style="font-size: 24px; color: #1E90FF; text-shadow: 0 0 5px #1E90FF;">Dixchete</span></strong>
</p>

*Note: The nice title and fonts don't render on GitHub browser.*

### Includes: 
    - Copy paste commands on this README
    - A command line tool to quickly accessing your computer's specs

###  Command line tool Features 

1. Interactive CLI interface with ASCII art display
2. Color-coded command categories and outputs
3. Fuzzy matching for category names
4. Execute commands directly or copy to clipboard
5. A navigatable menu to choose multiple options without restarting command prompt
6. Real-time command output display

> #### WHY?: I was having trouble quickly getting my computer specs when shopping for components, so I decided to create a cheat sheet of commands for easy reference, accessible by the command line.

## Dependencies:

1. [Go](https://go.dev/dl/) - Download and install Go for your OS
2. Go must be added to PATH:
    ```cmd
    setx GOPATH "%USERPROFILE%\go"
    setx PATH "%PATH%;%USERPROFILE%\go\bin"
    ```

3. Add dixchete to PATH:
    ```cmd
    setx PATH "%PATH%;C:\Program Files\dixchete"
    ```

4. Install the ASCII converter:
    ```go
    go get -u github.com/qeesung/image2ascii/convert
    ```

## Installation

Build it yourself since the .exe isn't signed:

```go
go mod init dixchete
go build -o dixchete.exe dixchete.go
mkdir "C:\Program Files\dixchete"
move dixchete.exe "C:\Program Files\dixchete"
```

To use, simply run:
```bash
dixchete
```

## Navigation Guide

#### Main Menu
- Select a category by:
    1. Entering the category number (1-4)
    2. Typing partial category name
    3. Press 'q' to quit

#### Category Menu
- Within a category:
    1. Enter command number to select it
    2. Press Enter to go back

#### Command Menu
- For selected commands:
    1. Press Enter to go back
    2. Enter 'p' to print command
    3. Enter 'e' to execute command
    4. Enter 'q' to quit

## Available Categories & Commands
## Storage Space Commands

- **`df -h`**
    - **Description:** Displays the amount of disk space used and available on all mounted filesystems in a human-readable format.
    - **When to use:** Use this command to quickly check how much disk space is being used and how much is available.
    mmand to quickly check how much disk space is being used and how much is available.

    ## Windows Specific Commands (Examples)

    - **`wmic diskdrive get model,name,size,mediaType`**
        - **Description:** Displays the model, name, size, and media type of all disk drives.
        - **When to use:** Use this command to get detailed information about the disk drives connected to the system.
        - **Example Output:**
            ```plaintext
            C:\Users\Jaedon>wmic diskdrive get model,name,size,mediaType
            MediaType              Model                       Name                Size
            Fixed hard disk media  ST1000DM003-1SB10C          \\.\PHYSICALDRIVE0  1000202273280
            Fixed hard disk media  addlink M.2 PCIE G3x4 NVMe  \\.\PHYSICALDRIVE1  256052966400
            ```

    - **`wmic baseboard get product,Manufacturer`**
        - **Description:** Displays the manufacturer and product name of the baseboard.
        - **When to use:** Use this command to get information about the motherboard.
        - **Example Output:**
            ```plaintext
            C:\Users\Jaedon>wmic baseboard get product,Manufacturer
            Manufacturer                        Product
            Micro-Star International Co., Ltd.  MPG Z490 GAMING PLUS (MS-7C75)
            ```

    - **`wmic path win32_pnpentity where "PNPDeviceID like '%PCI%'" get Name, DeviceID`**
        - **Description:** Lists the names and device IDs of all Plug and Play entities with PCI in their device ID.
        - **When to use:** Use this command to get information about PCI devices.
        - **Example Output:**
            ```plaintext
            C:\Users\Jaedon>wmic path win32_pnpentity where "PNPDeviceID like '%PCI%'" get Name, DeviceID
            DeviceID                                                        Name
            PCI\VEN_8086&DEV_06E0&SUBSYS_7C751462&REV_00\3&11583659&0&B0    Intel(R) Management Engine Interface #1
            PCI\VEN_8086&DEV_9BC8&SUBSYS_7C751462&REV_03\3&11583659&0&10    Intel(R) UHD Graphics 630
            PCI\VEN_8086&DEV_06F9&SUBSYS_7C751462&REV_00\3&11583659&0&90    Intel(R) Thermal Subsystem - 06F9
            PCI\VEN_8086&DEV_9B53&SUBSYS_7C751462&REV_03\3&11583659&0&00    Intel(R) Host Bridge/DRAM Registers - 9B53
            PCI\VEN_8086&DEV_1911&SUBSYS_7C751462&REV_00\3&11583659&0&40    Intel(R) Xeon(R) E3 - 1200/1500 v5/6th Gen Intel(R) Core(TM) Gaussian Mixture Model - 1911
            PCI\VEN_8086&DEV_06A3&SUBSYS_7C751462&REV_00\3&11583659&0&FC    Intel(R) SMBus - 06A3
            PCI\VEN_8086&DEV_06BC&SUBSYS_7C751462&REV_F0\3&11583659&0&E4    Intel(R) PCI Express Root Port #5 - 06BC
            PCI\VEN_8086&DEV_06B0&SUBSYS_7C751462&REV_F0\3&11583659&0&E8    Intel(R) PCI Express Root Port #9 - 06B0
            PCI\VEN_8086&DEV_06B8&SUBSYS_7C751462&REV_F0\3&11583659&0&E0    Intel(R) PCI Express Root Port #1 - 06B8
            ROOT\VPCIVSP\0000                                               Microsoft Hyper-V PCI Server
            PCI\VEN_14E4&DEV_43A0&SUBSYS_061914E4&REV_03\4&220F100&0&00E6   TP-LINK 802.11ac Network Adapter #2
            PCI\VEN_10DE&DEV_0FB9&SUBSYS_33511462&REV_A1\4&F8D4272&0&0108   High Definition Audio Controller
            PCI\VEN_8086&DEV_06BE&SUBSYS_7C751462&REV_F0\3&11583659&0&E6    Intel(R) PCI Express Root Port #7 - 06BE
            PCI\VEN_8086&DEV_06ED&SUBSYS_7C751462&REV_00\3&11583659&0&A0    Intel(R) USB 3.1 eXtensible Host Controller - 1.10 (Microsoft)
            PCI\VEN_1987&DEV_5012&SUBSYS_50121987&REV_01\4&24D42239&0&00E8  Standard NVM Express Controller
            PCI\VEN_10DE&DEV_1C82&SUBSYS_33511462&REV_A1\4&F8D4272&0&0008   NVIDIA GeForce GTX 1050 Ti
            PCI\VEN_10EC&DEV_8125&SUBSYS_7C751462&REV_04\4&28BD9FE6&0&00E4  Realtek PCIe 2.5GbE Family Controller
            PCI\VEN_8086&DEV_1901&SUBSYS_7C751462&REV_03\3&11583659&0&08    Intel(R) Xeon(R) E3 - 1200/1500 v5/6th Gen Intel(R) Core(TM) PCIe Controller (x16) - 1901
            PCI\VEN_8086&DEV_06EF&SUBSYS_72708086&REV_00\3&11583659&0&A2    Intel(R) Shared SRAM - 06EF
            SCSI\DISK&VEN_NVME&PROD_ADDLINK_M.2_PCIE\5&269496E0&0&000000    addlink M.2 PCIE G3x4 NVMe
            PCI\VEN_8086&DEV_0685&SUBSYS_7C751462&REV_00\3&11583659&0&F8    Intel(R) LPC Controller (Z490) - 0685
            PCI\VEN_8086&DEV_06D2&SUBSYS_7C751462&REV_00\3&11583659&0&B8    Standard SATA AHCI Controller
            PCI\VEN_8086&DEV_06C8&SUBSYS_EC751462&REV_00\3&11583659&0&FB    High Definition Audio Controller
            PCI\VEN_8086&DEV_06A4&SUBSYS_7C751462&REV_00\3&11583659&0&FD    Intel(R) SPI (flash) Controller - 06A4
            ```

- **`du -sh /path/to/directory`**
    - **Description:** Summarizes the disk usage of the specified directory in a human-readable format.
    - **When to use:** Use this command to find out how much space a specific directory is consuming.

## Drive Setup Commands

- **`fdisk -l`**
    - **Description:** Lists all available partitions and their details.
    - **When to use:** Use this command to view information about the disk partitions on your system.

- **`lsblk`**
    - **Description:** Lists information about all available or the specified block devices.
    - **When to use:** Use this command to get a detailed overview of all block devices and their mount points.

## System Information Commands

- **`systeminfo`**
    - **Description:** Displays detailed configuration information about the computer and its operating system.
    - **When to use:** Use this command to get a complete overview of the system's hardware and software configuration (Windows only).

- **`uname -a`**
    - **Description:** Prints all system information including kernel name, node name, kernel release, kernel version, machine, processor, hardware platform, and operating system.
    - **When to use:** Use this command to get comprehensive system information (Linux/Unix only).

- **`lshw`**
    - **Description:** Lists information about all hardware components of the system.
    - **When to use:** Use this command to get detailed information about the system's hardware configuration.

- **`lscpu`**
    - **Description:** Displays information about the CPU architecture.
    - **When to use:** Use this command to get details about the CPU, including number of CPUs, threads, cores, sockets, and architecture.

## Network Information Commands

- **`ifconfig`**
    - **Description:** Configures network interfaces and displays information about them.
    - **When to use:** Use this command to view and configure network interfaces (deprecated, use `ip` command instead).

- **`ip a`**
    - **Description:** Shows and manipulates routing, devices, policy routing, and tunnels.
    - **When to use:** Use this command to view detailed information about network interfaces and configure them.

- **`ping hostname`**
    - **Description:** Sends ICMP ECHO_REQUEST packets to network hosts.
    - **When to use:** Use this command to check the connectivity to a specific host.

- **`traceroute hostname`**
    - **Description:** Prints the route packets take to the network host.
    - **When to use:** Use this command to diagnose the route and measure transit delays of packets across the network.

## Additional Commands

- **`top`**
    - **Description:** Displays real-time system information including tasks, CPU usage, memory usage, and more.
    - **When to use:** Use this command to monitor system performance and resource usage.

- **`htop`**
    - **Description:** An interactive process viewer for Unix systems.
    - **When to use:** Use this command for an improved, user-friendly interface for monitoring system processes and resource usage.

- **`free -h`**
    - **Description:** Displays the total amount of free and used memory in the system in a human-readable format.
    - **When to use:** Use this command to check the system's memory usage.

## Windows Specific Commands

- **`wmic diskdrive list brief`**
    - **Description:** Lists brief information about all disk drives.
    - **When to use:** Use this command to get a quick overview of the disk drives connected to the system.
    - **Example Output:**
        ```plaintext
        Caption                     DeviceID            Model                       Partitions  Size
        ST1000DM003-1SB10C          \\.\PHYSICALDRIVE0  ST1000DM003-1SB10C          3           1000202273280
        addlink M.2 PCIE G3x4 NVMe  \\.\PHYSICALDRIVE1  addlink M.2 PCIE G3x4 NVMe  3           256052966400
        ```

- **`wmic diskdrive get model,name,size,mediaType`**
    - **Description:** Displays the model, name, size, and media type of all disk drives.
    - **When to use:** Use this command to get detailed information about the disk drives connected to the system.
    - **Example Output:**
        ```plaintext
        MediaType              Model                       Name                Size
        Fixed hard disk media  ST1000DM003-1SB10C          \\.\PHYSICALDRIVE0  1000202273280
        Fixed hard disk media  addlink M.2 PCIE G3x4 NVMe  \\.\PHYSICALDRIVE1  256052966400
        ```

- **`wmic baseboard get product,Manufacturer`**
    - **Description:** Displays the manufacturer and product name of the baseboard.
    - **When to use:** Use this command to get information about the motherboard.
    - **Example Output:**
        ```plaintext
        Manufacturer                        Product
        Micro-Star International Co., Ltd.  MPG Z490 GAMING PLUS (MS-7C75)
        ```

- **`wmic path win32_pnpentity where "PNPDeviceID like '%PCI%'" get Name, DeviceID`**
    - **Description:** Lists the names and device IDs of all Plug and Play entities with PCI in their device ID.
    - **When to use:** Use this command to get information about PCI devices.
    - **Example Output:**
        ```plaintext
        DeviceID                                                        Name
        ```

Feel free to add more commands!

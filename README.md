# computer_hardware_commands

Having trouble quickly understanding space, drive setup, etc., I decided to create a cheat sheet of commands for easy reference.

## Storage Space Commands

- **`df -h`**
    - **Description:** Displays the amount of disk space used and available on all mounted filesystems in a human-readable format.
    - **When to use:** Use this command to quickly check how much disk space is being used and how much is available.

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

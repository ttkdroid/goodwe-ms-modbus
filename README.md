## Modbus Reader for Goodwe MS Series Inverter

This Go program is designed to communicate with Goodwe MS Series inverters over Modbus TCP. It reads and processes various registers from the inverter, based on the Modbus register configuration specified in a JSON file. The program supports different types of data from the inverter, such as 16-bit and 32-bit unsigned and signed integers, as well as strings.

## Overview
**Purpose**:  This program allows users to retrieve data from Goodwe MS Series inverters by polling Modbus registers. It connects to the inverter via Modbus TCP and queries the registers specified in a JSON configuration file.
**Inverter Communication:** The program reads data from the inverter using the Modbus protocol, parses the responses, and then prints the results.
**Customization:** Users can modify the JSON configuration file to poll different registers or change the register parameters (address, type, scaling factor, etc.) as per their specific requirements.

## Features
 - Supports multiple data types:
	 - Unsigned 16-bit integer (U16) 
	 - Signed 16-bit integer (S16) 
	 - Unsigned 32-bit integer (U32) 
	 - Signed 32-bit integer (S32) 
	 - String (STR)
 - Scales the register values using a scaling factor defined in the JSON
   configuration. 
  - Outputs formatted data, including units (e.g., volts,    amps,
   watts), based on the register definitions.
 - Simple, configurable interface via a JSON file for easy adjustment of
   registers to be polled.

## Requirements
 - Go 1.18+ (recommended)
 - A Goodwe MS Series inverter with Modbus TCP enabled.
 - The github.com/goburrow/modbus Go package for Modbus communication.

## Installation
To run this program, follow these steps:

 1. Install Go (if you haven't already):
	Download and install Go from https://golang.org/dl/.
 2. Clone the repository:
	git clone https://github.com/yourusername/goodwe-modbus-reader.git
	cd goodwe-modbus-reader

 3. Install the required Go packages:
    go get github.com/goburrow/modbus

 4. Configure the JSON file:
	Create a JSON configuration file (e.g., src/inverter.json) to define the registers you want to poll from the 	inverter. Below is an example of a JSON configuration:
  

      `[
        {
        "address": "0x1000",
        "name": "Voltage",
        "property": "Line Voltage",
        "type": "U16",
        "length": 2,
        "sf_gain": 1,
        "units": "V",
        "range": "0-500",
        "note": "AC Line Voltage"
        },
        {
        "address": "0x1002",
        "name": "Current",
        "property": "Line Current",
        "type": "S16",
        "length": 2,
        "sf_gain": 1,
        "units": "A",
        "range": "-100-100",
        "note": "AC Line Current"
        }
        ]`


`address`: The Modbus register address (in hexadecimal).

`name`: A descriptive name for the register.

`property`: A description of what the register represents (e.g., "Voltage", "Current").

`type`: The data type of the register (U16, S16, U32, S32, STR).

`length`: The number of bytes to read for the register.

`sf_gain`: The scaling factor used to convert raw data into meaningful values.

`units`: The units of the register value (e.g., "V", "A", "W").

`range`: The valid range of the register values (if known).

`note`: Additional information about the register.

 5. Modify the code if necessary:

If you want to poll different registers, consult the Modbus documentation for your inverter model.

Modify the JSON file to include the addresses, data types, and other parameters corresponding to the registers you wish to poll.

 6. Run the program

`go run main.go` 

The program will connect to the inverter at the IP address `10.0.0.111` (as specified in the code) and read the registers defined in the JSON file. It will then print the values along with units and scaling information.

## Example Output

When the program successfully retrieves data, it will display something like this:
 
`
    MODBUS READER (Goodwe MS Series) by TTKLabs
    ------------------------------------------
    Voltage: 230.50 V
    Current: 12.30 A
`
## Modifying the JSON Configuration for Different Registers

If you need to poll different registers from the inverter, you must:  
 - Refer to the inverter's Modbus documentation to find the correct
   register addresses, data types, and scaling factors for the data you
   want to retrieve.
   
 - Update the JSON configuration accordingly with the correct addresses 
   and other parameters. You may add more registers as necessary.

For example, to add a register for "Power" with address 0x1004, you could add the following entry to the JSON file: 

    `{
      "address": "0x1004",
      "name": "Power",
      "property": "Instantaneous Power",
      "type": "U32",
      "length": 4,
      "sf_gain": 1000,
      "units": "W",
      "range": "0-10000",
      "note": "Instantaneous Power in Watts"
    }`

## License
This project is licensed under the BSD 2-Clause License. See the LICENSE file for more information.


## Contact
 If you have any questions, issues, or improvements, feel free to open an issue on the GitHub repository.
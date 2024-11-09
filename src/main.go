/*
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
 * OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
	"encoding/json" // for JSON parsing
	"fmt"           // for printing output
	"io"            // for reading data from the file
	"log"           // for logging errors
	"os"            // for file handling
	"strconv"       // for converting string to integer

	"github.com/goburrow/modbus" // for interacting with Modbus devices
)

// ModbusData struct holds the metadata for each Modbus register address.
type ModbusData struct {
	Address  string `json:"address"`  // Modbus register address as a string (hex format)
	Name     string `json:"name"`     // Name of the register (e.g., "Power", "Voltage")
	Property string `json:"property"` // The property of the register
	Type     string `json:"type"`     // Data type of the register (e.g., "U16", "S32", "STR")
	Length   int    `json:"length"`   // The number of bytes to read for this register
	SfGain   int    `json:"sf_gain"`  // Scaling factor for the data
	Units    string `json:"units"`    // Units of measurement (e.g., "V", "A", "W")
	Range    string `json:"range"`    // Valid range of values
	Note     string `json:"note"`     // Additional notes related to the register
}

func main() {
	// Display the header with formatting
	fmt.Println(_MAGENTA + _BOLD + "MODBUS READER (Goodwe MS Series) by TTKLabs" + _NC)

	// Open the JSON file that contains the Modbus register configurations
	jsonFile, err := os.Open("src/inverter.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %s", err) // If the file can't be opened, log an error and exit
	}
	defer jsonFile.Close() // Ensure the file is closed when the function ends

	// Read the entire content of the JSON file into a byte slice
	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshal the JSON data into the modbusData slice, each element represents a Modbus register
	var modbusData []ModbusData
	json.Unmarshal(byteValue, &modbusData)

	// Set up the Modbus TCP client to connect to the inverter (IP address and port)
	client := modbus.TCPClient("10.0.0.111:502")

	// Iterate over each entry in the modbusData slice to read the corresponding register
	for _, data := range modbusData {
		// Convert the hex string address to a uint16 integer for Modbus communication
		address, err := strconv.ParseUint(data.Address, 0, 16)
		if err != nil {
			log.Fatalf("Failed to parse address: %s", err) // If address parsing fails, log an error and exit
		}

		// Read the Modbus holding registers starting from the given address and read length
		results, err := client.ReadHoldingRegisters(uint16(address), uint16(data.Length))
		if err != nil {
			// If there's an error reading the register, print the error and skip to the next register
			fmt.Printf("Error reading %s: %s\n", data.Name, err)
			continue
		}

		// Process the results based on the type of data specified in the JSON (U16, S16, U32, S32, STR)
		switch data.Type {
		case "U16": // Unsigned 16-bit integer
			value := uint16(results[0])<<8 | uint16(results[1]) // Combine the two bytes into a 16-bit value
			// Print the result, scaled by the scale factor and with appropriate units
			fmt.Printf("%s%s:%s %.2f %s\n", _CYAN, data.Name, _NC, float64(value)/float64(data.SfGain), data.Units)
		case "S16": // Signed 16-bit integer
			value := int16(results[0])<<8 | int16(results[1]) // Combine the two bytes into a 16-bit signed value
			// Print the result, scaled by the scale factor and with appropriate units
			fmt.Printf("%s%s:%s %.2f %s\n", _CYAN, data.Name, _NC, float64(value)/float64(data.SfGain), data.Units)
		case "U32": // Unsigned 32-bit integer
			// Combine the four bytes into a 32-bit unsigned value
			value := uint32(results[0])<<24 | uint32(results[1])<<16 | uint32(results[2])<<8 | uint32(results[3])
			// Print the result, scaled by the scale factor and with appropriate units
			fmt.Printf("%s%s:%s %.2f %s\n", _CYAN, data.Name, _NC, float64(value)/float64(data.SfGain), data.Units)
		case "S32": // Signed 32-bit integer
			// Combine the four bytes into a 32-bit signed value
			value := int32(results[0])<<24 | int32(results[1])<<16 | int32(results[2])<<8 | int32(results[3])
			// Print the result, scaled by the scale factor and with appropriate units
			fmt.Printf("%s%s:%s %.2f %s\n", _CYAN, data.Name, _NC, float64(value)/float64(data.SfGain), data.Units)
		case "STR": // String (interprets the register as a string)
			// Convert the byte slice into a string
			strValue := string(results[:])
			// Print the string value
			fmt.Printf("%s%s:%s %s\n", _CYAN, data.Name, _NC, strValue)
		default:
			// Handle unsupported types, if any type is not in the predefined set (U16, S16, U32, S32, STR)
			fmt.Printf("Unsupported type %s for %s\n", data.Type, data.Name)
		}
	}
}

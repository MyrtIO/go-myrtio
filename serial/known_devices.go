package serial

var knownDevices = []string{
	"10C4",      // CP210X USB UART
	"0403:6015", // FT231XS USB UART
	"067B:2303", // Prolific Technology, Inc. PL2303 Serial Port
	"1A86:7523", // QinHeng Electronics HL-340 USB-Serial adapter
	"1A86:55D3", // QinHeng Electronics CH343 USB-Serial adapter
	"1A86:55D4", // QinHeng Electronics CH9102 USB-Serial adapter
	"2341",      // Arduino boards
	"2A03",      // Arduino boards
	"03EB:6124", // Arduino SAM-BA
	"16D0:0753", // Digistump boards
	"1EAF",      // Maple with DFU
	"0C9F:1781", // USBtiny
	"16C0",      // Teensy boards
	"0451:F432", // TI MSP430 Launchpad
	"28E9:0189", // GD32V DFU Bootloader
	"1A86:7522", // FireBeetle-ESP32
	"2886",      // Wio Terminal
	"2E8A",      // Raspberry Pi Pico
	"0D28:0204", // AIR32F103
	"0403",      // Many providers is using it
	"0451:C32A", // TI ICDI
	"0483",      // STLink probes
	"0640:0028", // Hilscher NXHX Boards
	"0640",      // Hitex probes
	"09FB:6001", // Altera USB Blaster
	"138E:9000", // Raisonance RLink
	"1457:5118", // Debug Board for Neo1973
	"15BA",      // Olimex probes
	"1781:0C63", // USBprog with OpenOCD firmware
	"1CBE:00FD", // TI/Luminary Stellaris In-Circuit Debug Interface (ICDI) Board
	"9E88:9E8F", // Marvell Sheevaplug
	"C251:2710", // Keil Software, Inc. ULink
	"03EB:2107", // Atmel AVR Dragon
	"303A:1001", // Espressif USB JTAG/serial debug unit
}

func isKnownDevice(id string) bool {
	for _, device := range knownDevices {
		if device == id {
			return true
		}
	}
	return false
}

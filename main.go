package main

func main() {
	var devices = []Device{}

	devices = append(devices, Device{Name: "Temp sensor 1", WeightInGrams: 10, Manufacturer: "Hej A"})
	devices = append(devices, Device{Name: "Temp sensor 2", WeightInGrams: 12, Manufacturer: "Hej B"})
	devices = append(devices, Device{Name: "Temp sensor 3", WeightInGrams: 45, Manufacturer: "Hej A"})

	for {
		showMenu()
		var selection int

		selection = inputIntData(1, 5, "Val")

		switch selection {
		case 1:
			newDevice := createDevice()
			devices = append(devices, *newDevice)
		case 2:
			listAll(devices)
		case 3:
			updateDevice(devices)
		case 4:
			searchDevice(devices)
			//1. fråga vad letar du efter  - > q
		case 5:
			//break would just break the switch!
			return

		}
	}
	// Kom ihåg lite rast också!
	//SES kl 14:00 !
	// Då kör vi  funktioner
	// och refaktoriserar dessutom denna kod
}

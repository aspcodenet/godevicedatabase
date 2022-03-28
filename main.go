package main

var devices = []string{}

func main() {

	devices = append(devices, "Temperature sensor A")
	devices = append(devices, "Temperature sensor B")

	for {
		showMenu()
		var selection int

		selection = inputIntData(1, 5, "Val")

		switch selection {
		case 1:
			newDevice := createDevice()
			devices = append(devices, newDevice)
		case 2:
			listAll()
		case 3:
			updateDevice()
		case 4:
			searchDevice()
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

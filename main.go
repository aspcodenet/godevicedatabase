package main

func main() {

	for {
		showMenu()
		var selection int

		selection = inputIntData(1, 5, "Val")

		switch selection {
		case 1:
			createDevice()
		case 2:
			listAll()
		case 3:
			updateDevice()
		case 4:
			searchDevice()
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

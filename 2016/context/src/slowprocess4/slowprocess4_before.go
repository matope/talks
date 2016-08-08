package main
func slowProcess(doneCh <-chan struct{}) error { // HL
		// ...
		case <-doneCh: // HL
			log.Println("slowProcess done.", i)
			return nil
		// ... 

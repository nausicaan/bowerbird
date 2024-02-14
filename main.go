package main

// Launch the program and execute the appropriate code
func main() {
	switch flag {
	case "-v", "--version":
		version()
	case "-h", "--help":
		about()
	case "-w", "--wpackagist", "-r", "--release":
		mrtest(flag)
	case "-p", "--premium":
		flag = "-p"
		premtest()
	case "--zero":
		alert("No flag detected -")
		about()
	default:
		alert("Bad flag detected -")
		about()
	}
}

// Determine which function to call based on the passed variable
func mrtest(flag string) {
	changedir()
	if inputs >= 4 {
		prepare()
		switch flag {
		case "-w", "--wpackagist":
			flag = "-w"
			wpackagist()
		case "-r", "--release":
			flag = "-r"
			released()
		}
	} else {
		alert(zero)
		about()
	}
}

// Call the Premium function if the required arguments are supplied
func premtest() {
	if inputs < 4 {
		alert(zero)
		about()
	} else if inputs > 4 {
		alert("Too many arguments supplied -")
		about()
	} else {
		premium()
	}
}

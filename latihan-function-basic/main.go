package main

import "fmt"

func main() {
	var car = map[string]string{}
	car["name"] = "BWM"
	car["color"] = "Black"
	message := kembalian(car["name"], car["color"])
	cetak(message)
}

func kembalian(x string, y string) string {
	return "Mobil " + x + " berwarna " + y
}

func cetak(pesan interface{}) {
	fmt.Println(pesan)
}

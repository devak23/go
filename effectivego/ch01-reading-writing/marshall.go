package main

// Task: Using In-Memory Readers and Writers to Support []byte
// You’re working on the back end for a ride-sharing site. The back end is composed of several processes that communicate
// with each other. The existing implementation has code to serialize data, but it works with io.Reader and io.Writer.
// You want to explore shared memory as a medium to exchange data, and in shared memory you need to work with []byte.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// Ride defines a struct with
type Ride struct {
	ID       int
	Time     time.Time
	Duration time.Duration
	Distance float64
	Price    float64
}

// Decoder - defines a new struct which contains a field of type json.Decoder. Why do we need a wrapper struct like this?
// Because we want to create a custom decoder that can read from an io.Reader and unmarshal JSON data into our Ride struct.
// This allows us to encapsulate the decoding logic and provide a clean interface for decoding JSON data.
// The Decoder struct will have a method DecodeRide that takes a pointer to a Ride struct and decodes the JSON data into it.
// This approach allows us to easily extend or modify the decoding behavior in the future if needed.
type Decoder struct {
	decoder *json.Decoder
}

// NewDecoder takes in an io.Reader (any data source you can read from, like files, buffers, etc.) and returns a new
// Decoder instance. This is in conjunction with what we mentioned above
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{json.NewDecoder(r)}
}

// UnmarshalRide is the custom function that we wanted to build to encapsulate the decoding logic. It returns a Ride
// from serialized data. Here we need to convert a []byte to an io.Reader. This is where bytes.NewReader comes in handy.
func UnmarshalRide(data []byte, ride *Ride) error {
	r := bytes.NewReader(data)
	return NewDecoder(r).DecodeRide(ride)
}

// DecodeRide is a method on the Decoder struct that decodes JSON data into the provided Ride struct. This is again a
// custom method that we defined to encapsulate the decoding logic.
func (d *Decoder) DecodeRide(v interface{}) error {
	return d.decoder.Decode(v)
}

func main() {
	data := []byte(`{"ID":1,"Time":"2023-10-01T12:00:00Z","Duration":3600,"Distance":10.5,"Price":15.75}`)

	var ride Ride
	if err := UnmarshalRide(data, &ride); err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v\n", ride)
	}

	//var ride2 Ride
	//json.Unmarshal(data, &ride2)
	//fmt.Printf("%+v\n", ride2)

	// We could have just used json.Unmarshal directly, right? Why all this encapsulation?
	// Yes, we could have used json.Unmarshal directly. It is the simplest and most idiomatic way to parse a JSON slice into
	// a struct and generally used when the entire data is already in memory as a byte slice. However, json.NewDecoder is
	// useful when you want to read JSON data from an io.Reader, such as a file or network connection, where the data may not
	// be fully available in memory at once. This is more of a "streaming" use case of decoding JSON, allowing you to handle
	// large datasets or data that is being read incrementally.

	// So then why not just write: err := json.NewDecoder(bytes.NewReader(data)).Decode(&ride) ?
	// We could have done that as well. The reason for encapsulating it in a custom Decoder struct is to provide a cleaner and
	// more maintainable interface for decoding JSON data. Wrapping with your own Decoder struct adds a tiny (almost negligible)
	// overhead—one extra struct and method call.Directly using the standard library is a bit simpler and more idiomatic for
	// one-off decoding tasks. The custom wrapper (Decoder) is only worth it if you want to add more custom decode logic,
	// helper methods, or enforce an interface.

	// if we were to read a file consisting of all such rides, we could use the Decoder struct to read each ride:
	file, err := os.OpenFile("rides.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()
	if err != nil {
		panic(err)
	}

}

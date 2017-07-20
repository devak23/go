package creational

import (
	"fmt"
	"sync"
)

type MichaelJackson struct {
	count int
}

var mjInstance *MichaelJackson
var once sync.Once

func (m *MichaelJackson) thriller() {
	fmt.Print("... cause it's the Thriller!")
}

func GetMJ() *MichaelJackson {
	once.Do(func() {
		mjInstance = &MichaelJackson{1}
	})

	return mjInstance
}


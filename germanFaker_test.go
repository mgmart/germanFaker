package germanFaker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestAllFaker(t *testing.T) {

	CustomGenerator()
	var sample Sample
	_ = faker.FakeData(&sample)
	sampleP, _ := json.MarshalIndent(sample, "", "\t")
	fmt.Println(string(sampleP))

}

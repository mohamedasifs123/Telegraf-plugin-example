package randomnumber

import (
    "math/rand"
    "time"

    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/inputs"
)

type RandomNumber struct {
    Min int `toml:"min"`
    Max int `toml:"max"`
}

var sampleConfig = `
  ## Minimum value
  min = 0
  ## Maximum value
  max = 100
`

func (r *RandomNumber) SampleConfig() string {
    return sampleConfig
}

func (r *RandomNumber) Description() string {
    return "Generates a random number"
}

func (r *RandomNumber) Gather(acc telegraf.Accumulator) error {
    rand.Seed(time.Now().UnixNano())
    value := r.Min + rand.Intn(r.Max-r.Min+1)
    fields := map[string]interface{}{
        "value": value,
    }
    acc.AddFields("randomnumber", fields, nil)
    return nil
}

func init() {
    inputs.Add("randomnumber", func() telegraf.Input {
        return &RandomNumber{}
    })
}

package simulation

import (
	"encoding/json"
	"io/ioutil"
)

func (sim *Simulation) SaveToFile(filename string) error {
	data, err := json.Marshal(sim); if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644); if err != nil {
		return err
	}

	return nil
}

func LoadFromFile(filename string) (*Simulation, error) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	} else {
		var simulation Simulation

		err := json.Unmarshal(data, &simulation)

		if err != nil {
			return nil, err
		} else {
			return &simulation, nil
		}
	}
}
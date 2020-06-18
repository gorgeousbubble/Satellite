package parse

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"satellite/app/erika/logs"
)

func unmarshalAutoMonitor(path string, out *TAutoMonitor, style string) (err error) {
	// open the file path
	file, err := os.Open(path)
	if err != nil {
		logs.Error("Error open file:", err)
		return err
	}
	defer file.Close()
	// read the file data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logs.Error("Error read data from file:", err)
		return err
	}
	// unmarshal data to struct
	switch style {
	case "json":
		err = json.Unmarshal(data, out)
		if err != nil {
			logs.Error("Error unmarshal json data to struct:", err)
			return err
		}
	case "yaml":
		err = yaml.Unmarshal(data, out)
		if err != nil {
			logs.Error("Error unmarshal yaml data to struct:", err)
			return err
		}
	default:
		err = json.Unmarshal(data, out)
		if err != nil {
			logs.Error("Error unmarshal json data to struct:", err)
			return err
		}
	}
	// show conversion result
	logs.Info("Unmarshal Auto Monitor successfully.")
	return err
}

func marshalAutoMonitor(path string, in *TAutoMonitor, style string) (err error) {
	var data []byte
	// marshal struct to data
	switch style {
	case "json":
		data, err = json.MarshalIndent(in, "", "\t")
		if err != nil {
			logs.Error("Error marshal struct to json data")
			return err
		}
	case "yaml":
		data, err = yaml.Marshal(in)
		if err != nil {
			logs.Error("Error marshal struct to yaml data")
			return err
		}
	default:
		data, err = json.MarshalIndent(in, "", "\t")
		if err != nil {
			logs.Error("Error marshal struct to json data")
			return err
		}
	}
	// write data to file path
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		logs.Error("Error write data to file:", err)
		return err
	}
	// show conversion result
	logs.Info("Marshal Auto Monitor successfully.")
	return err
}

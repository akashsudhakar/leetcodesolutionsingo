package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Manager struct {
	FullName       string `json:"full_name,omitempty"`
	Position       string `json:"position,omitempty"`
	Age            int32  `json:"age,omitempty"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	j, e := json.Marshal(manager)
	//json.Unmarshal(j, &manager)
	if e != nil {
		fmt.Println("Error while marshalling")
		return nil, e
	}
	return strings.NewReader(string(j)), nil
}

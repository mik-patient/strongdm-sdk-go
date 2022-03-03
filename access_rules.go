// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// This file was generated by protogen. DO NOT EDIT.
package sdm

import (
	"encoding/json"
	"strings"
)

type AccessRule struct {
	IDs  []string `json:"ids,omitempty"`
	Type string   `json:"type,omitempty"`
	Tags Tags     `json:"tags,omitempty"`
}

type AccessRules []AccessRule

func convertAccessRulesToPorcelain(rules string) (AccessRules, error) {
	if rules == "" {
		return nil, nil
	}
	result := AccessRules{}
	decoder := json.NewDecoder(strings.NewReader(rules))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func convertAccessRulesToPlumbing(rules AccessRules) string {
	if rules == nil {
		rules = AccessRules{}
	}
	result, _ := json.Marshal(rules)
	return string(result)
}

// ParseAccessRulesJSON parses the given access rules JSON string.
func ParseAccessRulesJSON(data string) (AccessRules, error) {
	result := AccessRules{}
	decoder := json.NewDecoder(strings.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

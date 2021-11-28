package util

import (
	"encoding/json"
	"os"
	"strings"
)

// SetEnvironmentFromSopsURL Retrieve a sops file from specified SOPS_FILE_URL environment variable.
// Then decrypt the content into key value pairs and set environment variables for each.
func SetEnvironmentFromSopsURL() error {
	var sopsVersion, sopsVersionError = RunCommand("sops --version")

	if sopsVersionError != nil {
		return sopsVersionError
	}

	Log(sopsVersion)

	var fileURL = os.Getenv("SOPS_FILE_URL")
	Log("Set to retrieve sops file at: " + fileURL)
	var encryptedContent, sopsURLError = GetURL(fileURL)

	if sopsURLError != nil {
		return sopsURLError
	}

	encryptedContent = strings.ReplaceAll(encryptedContent, "\n", "")

	var decryptedJSON, commandError = RunCommand(encryptedContent + " | " + "sops --input-type json --output type json /dev/stdin")
	if commandError != nil {
		return commandError
	}

	var secrets map[string]string
	var jsonError = json.Unmarshal([]byte(decryptedJSON), &secrets)

	if jsonError != nil {
		return jsonError
	}

	for key, value := range secrets {
		os.Setenv(key, value)
		Log("Set environment variable: " + key)
	}

	return nil
}

package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
)

// tmpDir is a temporary directory UNIQUE to this hyperlane node. make sure to use a different one per node.
// chainName e.g. simd1 or simd2
// rpcUrl is the node RPC endpoint for e.g. simd1
// hyperlaneDomain is the chain's hyperlane domain, as configured in the chain app state or genesis
func preconfigureHyperlane(node *hyperlane.HyperlaneChainConfig, tmpDir string, chainName string, chainRpcUrl string, hyperlaneDomain uint32) error {
	path := filepath.Dir(tmpDir)
	hyperlaneConfigPath := filepath.Join(path, chainName+".json")

	// Write the hyperlane CONFIG_FILES to disk where the bind mount will expect it.
	// See also https://docs.hyperlane.xyz/docs/operators/agent-configuration#config-files-with-docker.
	valJson := generateHyperlaneValidatorConfig(chainName, chainRpcUrl, hyperlaneDomain)
	err := os.WriteFile(hyperlaneConfigPath, []byte(valJson), 777)
	if err != nil {
		return err
	}

	// Search and replace for the Docker env vars, cmd-flags, and bind-mounts, see hyperlane.yaml
	node.SetReplacements(map[string]string{
		"${val_config_mount}": hyperlaneConfigPath,
		"${chainName}":        chainName,
		"${CHAINNAME}":        strings.ToUpper(chainName),
	})

	return nil
}

func generateHyperlaneValidatorConfig(chainName string, connectionUrl string, domain uint32) string {
	rawJson := `{
		"chains": {
		  "%s": {
			"connection": { "url": "%s" },
			"name": "%s",
			"domain": %d,
			"addresses": {
			  "mailbox": "0x35231d4c2D8B8ADcB5617A638A0c4548684c7C70",
			  "interchainGasPaymaster": "0x6cA0B6D22da47f091B7613223cD4BB03a2d77918",
			  "validatorAnnounce": "0x9bBdef63594D5FFc2f370Fe52115DdFFe97Bc524"
			},
			"protocol": "ethereum",
			"finalityBlocks": 1,
			"index": {
			  "from": 0
			}
		  }
		}
	  }`
	return fmt.Sprintf(rawJson, chainName, connectionUrl, chainName, domain)
}
package actions

import (
	"encoding/json"
	"log"

	"github.com/urfave/cli"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/alt-research/avs/core/config"
	"github.com/alt-research/avs/operator"
	"github.com/alt-research/avs/types"
)

func RegisterOperatorWithEigenlayer(ctx *cli.Context) error {

	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	err = operator.RegisterOperatorWithEigenlayer()
	if err != nil {
		return err
	}

	return nil
}
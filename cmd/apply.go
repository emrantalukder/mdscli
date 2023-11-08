package cmd

import (
	"log"
	"os"

	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var yamlFile string

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply a YAML file with role bindings",
	Long:  `Reads a YAML file and processes the role bindings against Confluent Platform's Metadata Service API`,
	Run: func(cmd *cobra.Command, args []string) {
		applyRoleBindings(yamlFile)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVarP(&yamlFile, "file", "f", "", "YAML file with role bindings")
	applyCmd.MarkFlagRequired("file")
}

type RoleBinding struct {
	Principal               string  `yaml:"principal"`
	Role                    string  `yaml:"role"`
	Resource                *string `yaml:"resource,omitempty"`
	KafkaClusterId          *string `yaml:"kafka_cluster_id"`
	SchemaRegistryClusterId *string `yaml:"schema_registry_cluster_id,omitempty"`
	KsqlClusterId           *string `yaml:"ksql_cluster_id,omitempty"`
	ConnectClusterId        *string `yaml:"connect_cluster_id,omitempty"`
}

type Scope struct {
	KafkaClusterId          *string `yaml:"kafka_cluster_id,omitempty"`
	SchemaRegistryClusterId *string `yaml:"schema_registry_cluster_id,omitempty"`
	KsqlClusterId           *string `yaml:"ksql_cluster_id,omitempty"`
	ConnectClusterId        *string `yaml:"connect_cluster_id,omitempty"`
}

type Config struct {
	RoleBindings []RoleBinding `yaml:"role_bindings"`
}

func applyRoleBindings(yamlFilePath string) {
	yamlFile, err := os.ReadFile(yamlFilePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}

	for _, roleBinding := range config.RoleBindings {
		b, err := yaml.Marshal(roleBinding)
		if err != nil {
			fmt.Println(err)
		}
		log.Println("\n" + string(b))
	}
}

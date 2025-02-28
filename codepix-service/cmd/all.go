/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/robertheory/codepix/app/grpc"
	"github.com/robertheory/codepix/app/kafka"
	"github.com/robertheory/codepix/infra/db"
	"github.com/spf13/cobra"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

var grpcPortNumber int

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run all services",
	Long:  `Command to run all services of the application.`,
	Run: func(cmd *cobra.Command, args []string) {

		database := db.ConnectDB(os.Getenv("env"))

		go grpc.StartGrpcServer(database, grpcPortNumber)

		deliveryChan := make(chan ckafka.Event)
		producer := kafka.NewKafkaProducer()
		go kafka.DeliveryReport(deliveryChan)
		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	allCmd.Flags().IntVarP(&grpcPortNumber, "port", "p", 50051, "gRPC server port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

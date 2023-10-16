// from https://github.com/GoogleCloudPlatform/golang-samples/blob/main/bigtable/helloworld/main.go
// my instance - https://console.cloud.google.com/bigtable/instances?project=modern-unison-305820
// ./bigtable -project="selam-project-id" -instance="test-instance"
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/bigtable"
)

const (
	tableName        = "Hello-Bigtable"
	columnFamilyName = "cf1"
	columnName       = "greetings"
)

var greetings = []string{"Hello world!", "Hello cloud BigTable!", "Hello golang!"}

func sliceContains(slice []string, target string) bool {
	for _, s := range slice {
		if target == s {
			return true
		}
	}
	return false
}

func main() {
	project := flag.String("project", "", "The Google Cloud Platform project ID. Required")
	instance := flag.String("instance", "", "The Google Cloud Platform instance ID. Required")
	flag.Parse()

	for _, f := range []string{"project", "instance"} {
		if flag.Lookup(f).Value.String() == "" {
			log.Fatal("Required flag: ", f)
		}
	}
	fmt.Println("project: ", *project)
	fmt.Println("instance: ", *instance)

	ctx := context.Background()

	adminClient, err := bigtable.NewAdminClient(ctx, *project, *instance)
	if err != nil {
		log.Fatal("couldn't create admin client")
	}
	defer adminClient.Close()

	tables, err := adminClient.Tables(ctx)
	if err != nil {
		log.Fatal("couldn't get tables. Error: ", err)
	}

	if !sliceContains(tables, tableName) {
		log.Println("Creating table: ", tableName)
		if err := adminClient.CreateTable(ctx, tableName); err != nil {
			log.Fatalf("couldn't create table %s: %v", tableName, err)
		}
	}

	tableInfo, err := adminClient.TableInfo(ctx, tableName)
	if err != nil {
		log.Fatal("couldn't get table info")
	}

	fmt.Println("table info: ", tableInfo.Families)

	if !sliceContains(tableInfo.Families, columnFamilyName) {
		log.Println("Creating colum family name: ", columnFamilyName)
		if err := adminClient.CreateColumnFamily(ctx, tableName, columnFamilyName); err != nil {
			log.Fatal("couldn't create column family")
		}
	}

	client, err := bigtable.NewClient(ctx, *project, *instance)
	if err != nil {
		log.Fatal("couldn't create client")
	}
	defer client.Close()

	table := client.Open(tableName)
	muts := make([]*bigtable.Mutation, len(greetings))
	rowKeys := make([]string, len(greetings))

	log.Println("writing greeting logs to Bigtable")
	for i, greeting := range greetings {
		muts[i] = bigtable.NewMutation()
		muts[i].Set(columnFamilyName, columnName, bigtable.Now(), []byte(greeting))
		rowKeys[i] = fmt.Sprintf("%s%d", columnName, i)
	}

	rowErr, err := table.ApplyBulk(ctx, rowKeys, muts)
	if err != nil {
		log.Fatal("Applying bulk failed entirely: ", err)
	}

	if rowErr != nil {
		for i, re := range rowErr {
			log.Printf("Error %v writing %d row", re, i)
		}
		log.Fatal("couldn't write some rows")
	}

	log.Println("Reading a single greeting by row key")

	//the filter - include only rows that have the specific columns in them
	row, err := table.ReadRow(ctx, rowKeys[0], bigtable.RowFilter(bigtable.ColumnFilter(columnName)))
	if err != nil {
		log.Fatalf("couldn't read row with key %s: %v", rowKeys[0], err)
	}

	log.Printf("\t%s = %s\n", rowKeys[0], string(row[columnFamilyName][0].Value))

	log.Println("Reading all greeting")

	//we want all prefixed with "greeting"
	table.ReadRows(ctx, bigtable.PrefixRange(columnName), func(row bigtable.Row) bool {
		item := row[columnFamilyName][0]
		log.Printf("\t%s = %s\n", item.Row, item.Value)
		return true
	}, bigtable.RowFilter(bigtable.ColumnFilter(columnName)))

	/* log.Println("Deleting the table")
	if err := adminClient.DeleteTable(ctx, tableName); err != nil {
		log.Fatalf("couldn't delete table %s: %v", tableName, err)
	} */

}

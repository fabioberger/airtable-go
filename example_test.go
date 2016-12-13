package airtable_test

import (
	"fmt"
	"os"

	airtable "github.com/fabioberger/airtable-go"
)

func ExampleNew() {
	airtableAPIKey := os.Getenv("AIRTABLE_API_KEY")
	baseID := "apphllLCpWnySSF7q"
	shouldRetryIfRateLimited := true

	client := airtable.New(airtableAPIKey, baseID, shouldRetryIfRateLimited)

	fmt.Println(client)
}

func ExampleClient_DestroyRecord() {
	client := airtable.New("AIRTABLE_API_KEY", "BASE_ID", true)

	if err := client.DestroyRecord("TABLE_NAME", "RECORD_ID"); err != nil {
		panic(err)
	}
}

func ExampleClient_ListRecords() {
	client := airtable.New("AIRTABLE_API_KEY", "BASE_ID", true)

	type task struct {
		AirtableID string
		Fields     struct {
			Name  string
			Notes string
		}
	}
	tasks := []task{}
	err := client.ListRecords("TABLE_NAME", &tasks)
	if err != nil {
		panic(err)
	}

	fmt.Println(tasks)
}

func ExampleClient_RetrieveRecord() {
	client := airtable.New("AIRTABLE_API_KEY", "BASE_ID", true)

	type task struct {
		AirtableID string
		Fields     struct {
			Name  string
			Notes string
		}
	}
	retrievedTask := task{}
	err := client.RetrieveRecord("TABLE_NAME", "RECORD_ID", &retrievedTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(retrievedTask)
}

func ExampleClient_UpdateRecord() {
	client := airtable.New("AIRTABLE_API_KEY", "BASE_ID", true)

	type task struct {
		AirtableID string
		Fields     struct {
			Name  string
			Notes string
		}
	}
	aTask := task{
		AirtableID: "RECORD_ID",
		Fields: struct {
			Name  string
			Notes string
		}{
			Name:  "Clean kitchen",
			Notes: "Make sure to clean all the counter tops",
		},
	}
	UpdatedFields := map[string]interface{}{
		"Name": "Clean entire kitchen",
	}
	err := client.UpdateRecord("TABLE_NAME", "RECORD_ID", UpdatedFields, &aTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(aTask)
}

func ExampleListParameters() {
	client := airtable.New("AIRTABLE_API_KEY", "BASE_ID", true)

	listParams := airtable.ListParameters{
		Fields:          []string{"Name", "Notes", "Priority"},
		FilterByFormula: "{Priority} < 2",
		MaxRecords:      50,
		Sort: []*airtable.SortParameter{
			airtable.NewSortParameter("Priority", "desc"),
		},
		View: "Main View",
	}

	type task struct {
		AirtableID string
		Fields     struct {
			Name     string
			Notes    string
			Priority int
		}
	}
	tasks := []task{}
	if err := client.ListRecords("TABLE_NAME", &tasks, listParams); err != nil {
		panic(err)
	}

	fmt.Println(tasks)
}

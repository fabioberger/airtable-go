package integrationTests

import (
	"testing"

	. "github.com/fabioberger/airtable-go"
	"github.com/fabioberger/airtable-go/test_configs"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ClientSuite struct{}

var _ = Suite(&ClientSuite{})

var shouldRetryIfRateLimited = true
var client = New(testConfigs.AirtableTestAPIKey, testConfigs.AirtableTestBaseID, shouldRetryIfRateLimited)

var tasksTableName = "Tasks"

type task struct {
	AirtableID string `json:"id,omitempty"`
	Fields     struct {
		Name      string
		Notes     string
		Completed bool
		TimeEst   float64 `json:"Time Estimate (days)"`
	} `json:"fields"`
}

var teamMatesTableName = "Teammates"

type teamMate struct {
	AirtableID string `json:"id,omitempty"`
	Fields     struct {
		Name string
	} `json:"fields"`
}

var LogTableName = "Log"

type Log struct {
	AirtableID string `json:"id,omitempty"`
	Fields     struct {
		AutoNumber int      `json:"Auto Number"`
		Projects   []string `json:"Projects"`
	}
}

func (s *ClientSuite) TearDownTest(c *C) {
	client.RestoreAPIResponseStub()
}

func (s *ClientSuite) TestListRecords(c *C) {
	tasks := []task{}
	err := client.ListRecords(tasksTableName, &tasks)
	c.Assert(err, Equals, nil)
	c.Assert(len(tasks), Equals, 3)
}

func (s *ClientSuite) TestRetrieveRecord(c *C) {
	tasks := []task{}
	client.ListRecords(tasksTableName, &tasks)
	t := tasks[0]

	aTask := task{}
	client.RetrieveRecord(tasksTableName, t.AirtableID, &aTask)
	c.Assert(aTask.Fields.Name, Equals, t.Fields.Name)
}

func (s *ClientSuite) TestCreateAndDestroyRecord(c *C) {
	tm := teamMate{}
	tm.Fields.Name = "Bob"
	err := client.CreateRecord(teamMatesTableName, &tm)
	c.Assert(err, Equals, nil)

	err = client.DestroyRecord(teamMatesTableName, tm.AirtableID)
	c.Assert(err, Equals, nil)
}

func (s *ClientSuite) TestUpdateRecord(c *C) {
	tasks := []task{}
	client.ListRecords(tasksTableName, &tasks)
	t := tasks[0]
	oldName := t.Fields.Name

	updatedFields := map[string]interface{}{
		"Name": "John Coltrain",
	}
	err := client.UpdateRecord(teamMatesTableName, t.AirtableID, updatedFields, &t)
	c.Assert(err, Equals, nil)
	c.Assert(t.Fields.Name, Equals, "John Coltrain")
	c.Assert(t.Fields.Name, Not(Equals), oldName)

	// revert to old name
	updatedFields = map[string]interface{}{
		"Name": oldName,
	}
	err = client.UpdateRecord(teamMatesTableName, t.AirtableID, updatedFields, &t)
	c.Assert(err, Equals, nil)
}

func (s *ClientSuite) TestListRecordsRequiringMultipleRequests(c *C) {
	logs := []Log{}
	if err := client.ListRecords(LogTableName, &logs); err != nil {
		c.Error(err)
	}
	c.Assert(len(logs), Equals, 117)
}

func (s *ClientSuite) TestListRecordsInSpecificView(c *C) {
	tasks := []task{}
	listParameters := ListParameters{
		View: "Tea Packaging Tasks",
	}
	err := client.ListRecords(tasksTableName, &tasks, listParameters)
	c.Assert(err, Equals, nil)
	c.Assert(len(tasks), Equals, 1)
}

func (s *ClientSuite) TestListRecordsUsingFilterByFormula(c *C) {
	tasks := []task{}
	listParameters := ListParameters{
		FilterByFormula: "{Time Estimate (days)} > 2",
	}
	err := client.ListRecords(tasksTableName, &tasks, listParameters)
	c.Assert(err, Equals, nil)
	c.Assert(len(tasks), Equals, 2)
}

func (s *ClientSuite) TestListRecordsWithASortedOrder(c *C) {
	tasks := []task{}
	listParameters := ListParameters{
		Sort: []*SortParameter{
			NewSortParameter("Time Estimate (days)", "asc"),
		},
	}
	err := client.ListRecords(tasksTableName, &tasks, listParameters)
	c.Assert(err, Equals, nil)
	c.Assert(tasks[0].Fields.TimeEst, Equals, 1.5)
	c.Assert(tasks[1].Fields.TimeEst, Equals, 2.5)
	c.Assert(tasks[2].Fields.TimeEst, Equals, 15.0)
}

func (s *ClientSuite) TestListRecordsWithSpecifiedMaxRecords(c *C) {
	tasks := []task{}
	listParameters := ListParameters{
		MaxRecords: 1,
	}
	err := client.ListRecords(tasksTableName, &tasks, listParameters)
	c.Assert(err, Equals, nil)
	c.Assert(len(tasks), Equals, 1)
}

func (s *ClientSuite) TestListRecordsWithSpecifiedFields(c *C) {
	tasks := []task{}
	listParameters := ListParameters{
		Fields: []string{"Time Estimate (days)"},
	}
	err := client.ListRecords(tasksTableName, &tasks, listParameters)
	c.Assert(err, Equals, nil)
	for _, task := range tasks {
		c.Assert(task.Fields.Name, Equals, "")
		c.Assert(task.Fields.Notes, Equals, "")
		c.Assert(task.Fields.Completed, Equals, false)
	}
}

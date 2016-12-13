package airtable

import (
	"testing"
	"time"

	"github.com/fabioberger/airtable-go/test_configs"
	. "gopkg.in/check.v1"
)

var shouldRetryIfRateLimited = true
var client = New(testConfigs.AirtableTestAPIKey, testConfigs.AirtableTestBaseID, shouldRetryIfRateLimited)

var tasksTableName = "Tasks"
var teamMatesTableName = "Teammates"
var fakeRecordID = "recSG8Ytl8KWpNEVA"

type task struct {
	AirtableID string `json:"id,omitempty"`
	Fields     struct {
		Name      string `json:"name"`
		Notes     string `json:"notes"`
		Completed bool   `json:"Completed"`
	} `json:"fields"`
}

type teamMate struct {
	AirtableID string `json:"id,omitempty"`
	Fields     struct {
		Name string
	} `json:"fields"`
}

func Test(t *testing.T) { TestingT(t) }

type ClientSuite struct{}

var _ = Suite(&ClientSuite{})

func (s *ClientSuite) TearDownTest(c *C) {
	client.RestoreAPIResponseStub()
}

func (s *ClientSuite) TestListRecords(c *C) {
	tasks := []task{}
	client.StubAPIResponseWithFileContentsOrPanic(200, "./test_responses/list_tasks.json")
	err := client.ListRecords(tasksTableName, &tasks)
	c.Assert(err, Equals, nil)
	c.Assert(len(tasks), Equals, 3)
}

func (s *ClientSuite) TestAirtableError(c *C) {
	tasks := []task{}
	client.StubAPIResponseWithFileContentsOrPanic(404, "./test_responses/404_error.json")
	err := client.ListRecords(tasksTableName, &tasks)
	c.Assert(err.Error(), Equals, "NOT_FOUND: Could not find table x in application appmUJMUx1SyZYQYX [HTTP code 404]")
}

func (s *ClientSuite) TestRetrieveRecord(c *C) {
	aTask := task{}
	client.StubAPIResponseWithFileContentsOrPanic(200, "./test_responses/retrieve_task.json")
	client.RetrieveRecord(tasksTableName, fakeRecordID, &aTask)
	c.Assert("Research other tea packaging", Equals, aTask.Fields.Name)
}

func (s *ClientSuite) TestCreateRecord(c *C) {
	tm := teamMate{}
	tm.Fields.Name = "Bob"
	client.StubAPIResponseWithFileContentsOrPanic(200, "./test_responses/create_teammate.json")
	err := client.CreateRecord(teamMatesTableName, &tm)
	c.Assert(err, Equals, nil)
}

func (s *ClientSuite) TestDestroyRecord(c *C) {
	client.StubAPIResponseWithFileContentsOrPanic(200, "./test_responses/delete_teammate.json")
	err := client.DestroyRecord(teamMatesTableName, fakeRecordID)
	c.Assert(err, Equals, nil)
}

func (s *ClientSuite) TestUpdateRecord(c *C) {
	updatedFields := map[string]interface{}{
		"Name": "John Coltrain",
	}
	client.StubAPIResponseWithFileContentsOrPanic(200, "./test_responses/update_teammate.json")
	t := teamMate{}
	err := client.UpdateRecord(teamMatesTableName, fakeRecordID, updatedFields, &t)
	c.Assert(err, Equals, nil)
	c.Assert(t.Fields.Name, Equals, "John Coltrain")
}

func (s *ClientSuite) TestRetryLogicIfRateLimited(c *C) {
	channel := make(chan bool)
	go func() {
		updatedFields := map[string]interface{}{
			"Name": "Bill Bob",
		}
		client.StubAPIResponseWithFileContentsOrPanic(rateLimitStatusCode, "./test_responses/update_teammate.json")
		err := client.UpdateRecord(teamMatesTableName, fakeRecordID, updatedFields, nil)
		c.Assert(err, Equals, nil)
		channel <- true
	}()

	for {
		select {
		case _ = <-channel:
			c.Error("Request terminated before rateLimit sleep completed.")
		case _ = <-time.After(2 * time.Second):
			return // Request correctly still waiting to retry after rateLimited
		}
	}
}

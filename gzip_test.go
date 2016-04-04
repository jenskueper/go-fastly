package fastly

import "testing"

func TestClient_Gzips(t *testing.T) {
	t.Parallel()

	var err error
	var tv *Version
	record(t, "gzips/version", func(c *Client) {
		tv = testVersion(t, c)
	})

	// Create
	var gzip *Gzip
	record(t, "gzips/create", func(c *Client) {
		gzip, err = c.CreateGzip(&CreateGzipInput{
			Service:      testServiceID,
			Version:      tv.Number,
			Name:         "test-gzip",
			ContentTypes: "text/html text/css",
			Extensions:   "html css",
		})
	})
	if err != nil {
		t.Fatal(err)
	}

	// Ensure deleted
	defer func() {
		record(t, "gzips/cleanup", func(c *Client) {
			c.DeleteGzip(&DeleteGzipInput{
				Service: testServiceID,
				Version: tv.Number,
				Name:    "test-gzip",
			})

			c.DeleteGzip(&DeleteGzipInput{
				Service: testServiceID,
				Version: tv.Number,
				Name:    "new-test-gzip",
			})
		})
	}()

	if gzip.Name != "test-gzip" {
		t.Errorf("bad name: %q", gzip.Name)
	}
	if gzip.ContentTypes != "text/html text/css" {
		t.Errorf("bad content_types: %q", gzip.ContentTypes)
	}
	if gzip.Extensions != "html css" {
		t.Errorf("bad extensions: %q", gzip.Extensions)
	}

	// List
	var gzips []*Gzip
	record(t, "gzips/list", func(c *Client) {
		gzips, err = c.ListGzips(&ListGzipsInput{
			Service: testServiceID,
			Version: tv.Number,
		})
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(gzips) < 1 {
		t.Errorf("bad gzips: %v", gzips)
	}

	// Get
	var ngzip *Gzip
	record(t, "gzips/get", func(c *Client) {
		ngzip, err = c.GetGzip(&GetGzipInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "test-gzip",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
	if gzip.Name != ngzip.Name {
		t.Errorf("bad name: %q", gzip.Name)
	}
	if gzip.ContentTypes != "text/html text/css" {
		t.Errorf("bad content_types: %q", gzip.ContentTypes)
	}
	if gzip.Extensions != "html css" {
		t.Errorf("bad extensions: %q", gzip.Extensions)
	}

	// Update
	var ugzip *Gzip
	record(t, "gzips/update", func(c *Client) {
		ugzip, err = c.UpdateGzip(&UpdateGzipInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "test-gzip",
			NewName: "new-test-gzip",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
	if ugzip.Name != "new-test-gzip" {
		t.Errorf("bad name: %q", ugzip.Name)
	}

	// Delete
	record(t, "gzips/delete", func(c *Client) {
		err = c.DeleteGzip(&DeleteGzipInput{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "new-test-gzip",
		})
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_ListGzips_validation(t *testing.T) {
	var err error
	_, err = testClient.ListGzips(&ListGzipsInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.ListGzips(&ListGzipsInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_CreateGzip_validation(t *testing.T) {
	var err error
	_, err = testClient.CreateGzip(&CreateGzipInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.CreateGzip(&CreateGzipInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_GetGzip_validation(t *testing.T) {
	var err error
	_, err = testClient.GetGzip(&GetGzipInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetGzip(&GetGzipInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetGzip(&GetGzipInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_UpdateGzip_validation(t *testing.T) {
	var err error
	_, err = testClient.UpdateGzip(&UpdateGzipInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateGzip(&UpdateGzipInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateGzip(&UpdateGzipInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_DeleteGzip_validation(t *testing.T) {
	var err error
	err = testClient.DeleteGzip(&DeleteGzipInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteGzip(&DeleteGzipInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteGzip(&DeleteGzipInput{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

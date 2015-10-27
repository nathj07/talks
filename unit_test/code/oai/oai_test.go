package oai

import "testing"

var testResponse = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="http://www.worldsciencepublisher.org/journals/lib/pkp/xml/oai2.xsl" ?>
<OAI-PMH xmlns="http://www.openarchives.org/OAI/2.0/"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/
		http://www.openarchives.org/OAI/2.0/OAI-PMH.xsd">
	<responseDate>2015-10-16T19:24:52Z</responseDate>
	<request verb="ListMetadataFormats">http://www.worldsciencepublisher.org/journals/index.php/ACSA/oai</request>
	<ListMetadataFormats>
		<metadataFormat>
			<metadataPrefix>nlm</metadataPrefix>
			<schema>http://dtd.nlm.nih.gov/publishing/2.3/xsd/journalpublishing.xsd</schema>
			<metadataNamespace>http://td.nlm.nih.gov/publishing/2.3</metadataNamespace>
		</metadataFormat>
		<metadataFormat>
			<metadataPrefix>oai_dc</metadataPrefix>
			<schema>http://www.openarchives.org/OAI/2.0/oai_dc.xsd</schema>
			<metadataNamespace>http://www.openarchives.org/OAI/2.0/oai_dc/</metadataNamespace>
		</metadataFormat>
		<metadataFormat>
			<metadataPrefix>rfc1807</metadataPrefix>
			<schema>http://www.openarchives.org/OAI/1.1/rfc1807.xsd</schema>
			<metadataNamespace>http://info.internet.isi.edu:80/in-notes/rfc/files/rfc1807.txt</metadataNamespace>
		</metadataFormat>
		<metadataFormat>
			<metadataPrefix>marcxml</metadataPrefix>
			<schema>http://www.loc.gov/standards/marcxml/schema/MARC21slim.xsd</schema>
			<metadataNamespace>http://www.loc.gov/MARC21/slim</metadataNamespace>
		</metadataFormat>
		<metadataFormat>
			<metadataPrefix>oai_marc</metadataPrefix>
			<schema>http://www.openarchives.org/OAI/1.1/oai_marc.xsd</schema>
			<metadataNamespace>http://www.openarchives.org/OAI/1.1/oai_marc</metadataNamespace>
		</metadataFormat>
	</ListMetadataFormats>
</OAI-PMH>`)

// TESTMOCK OMIT
type mockFetcher struct {
}

func (mf mockFetcher) Fetch(url string) ([]byte, error) {
	return testResponse, nil // testResponse local []byte holding known good response
}

var oai = &OAI{
	OAIFetcher: mockFetcher{}, // inject the mock
}

// TESTMOCK OMIT

func TestFetchMetadataFormats(t *testing.T) {
	got, err := oai.FetchMetadataFormats("http://test.com")
	if err != nil {
		t.Errorf("Unepxted error from FetchMetadataFormats: %v", err)
	}

	expected := []MDFormat{
		{
			Prefix:    "nlm",
			Schema:    "http://dtd.nlm.nih.gov/publishing/2.3/xsd/journalpublishing.xsd",
			Namespace: "http://td.nlm.nih.gov/publishing/2.3",
		},
		{
			Prefix:    "oai_dc",
			Schema:    "http://www.openarchives.org/OAI/2.0/oai_dc.xsd",
			Namespace: "http://www.openarchives.org/OAI/2.0/oai_dc/",
		},
		{
			Prefix:    "rfc1807",
			Schema:    "http://www.openarchives.org/OAI/1.1/rfc1807.xsd",
			Namespace: "http://info.internet.isi.edu:80/in-notes/rfc/files/rfc1807.txt",
		},
		{
			Prefix:    "marcxml",
			Schema:    "http://www.loc.gov/standards/marcxml/schema/MARC21slim.xsd",
			Namespace: "http://www.loc.gov/MARC21/slim",
		},
		{
			Prefix:    "oai_marc",
			Schema:    "http://www.openarchives.org/OAI/1.1/oai_marc.xsd",
			Namespace: "http://www.openarchives.org/OAI/1.1/oai_marc",
		},
	}

	for i, format := range got {
		if format.Prefix != expected[i].Prefix {
			t.Errorf("Test %d. Unexpected Prefix: %s, expected: %s", i, format.Prefix, expected[i].Prefix)
		}
		if format.Schema != expected[i].Schema {
			t.Errorf("Test %d. Unexpected Schema: %s, expected: %s", i, format.Schema, expected[i].Schema)
		}
		if format.Namespace != expected[i].Namespace {
			t.Errorf("Test %d. Unexpected Namespace: %s, expected: %s", i, format.Namespace, expected[i].Namespace)
		}
		// could replace all the above if with a reflect.DeepEquals call:
		/*
			if !reflect.DeepEqual(format, expected[i]) {
				t.Errorf("Mismatched format record:\nexp:%+v\ngot:%+v", expected[i], format)
			}
		*/
	}

}

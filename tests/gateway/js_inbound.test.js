import common from "../common";
import request from "supertest";


const testNamespace = "js-inbound";


const endpointJSFile = `
direktiv_api: endpoint/v1
allow_anonymous: true
plugins:
  target:
    type: target-flow
    configuration:
        flow: /target.yaml
        content_type: application/json
  inbound:
    - type: js-inbound
      configuration:
        script: |
            input["Headers"].Del("Header1")
            input["Headers"].Add("Header3", "value3")
            input["Queries"].Del("Query2")
            b = JSON.parse(input["Body"])
            b["addquery"] = input["Queries"].Get("Query1")
            b["addquerydel"] = input["Queries"].Get("Query2")

            b["addheader"] = input["Headers"].Get("Header3")
            b["addheaderdel"] = input["Headers"].Get("Header1")
            input["Body"] = JSON.stringify(b) 
methods: 
  - POST
path: /target`


const wf = `
direktiv_api: workflow/v1
states:
- id: helloworld
  type: noop
  transform:
    result: jq(.)
`

describe("Test js inbound plugin", () => {
    beforeAll(common.helpers.deleteAllNamespaces);
  
    common.helpers.itShouldCreateNamespace(it, expect, testNamespace);
    // common.helpers.itShouldCreateNamespace(it, expect, testNamespace);
  
    common.helpers.itShouldCreateFile(
      it,
      expect,
      testNamespace,
      "/endpoint1.yaml",
      endpointJSFile
    );
  
    common.helpers.itShouldCreateFile(
      it,
      expect,
      testNamespace,
      "/target.yaml",
      wf
    );

    it(`should have expected body after js`, async () => {
      const req = await request(common.config.getDirektivHost()).post(
        `/ns/` + testNamespace + `/target?Query1=value1&Query2=value2`
      ).set('Header1', 'Value1').send({"hello":"world"});

      expect(req.statusCode).toEqual(200);
      expect(req.body.result.addheader).toEqual("value3")
      expect(req.body.result.addheaderdel).toEqual("")
      expect(req.body.result.addquery).toEqual("value1")
      expect(req.body.result.addquerydel).toEqual("")
    });
  
  
  });
  
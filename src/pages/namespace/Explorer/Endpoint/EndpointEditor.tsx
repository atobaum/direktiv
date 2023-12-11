import { FC, useState } from "react";
import { addEndpointHeader, useEndpointFormSchema } from "./utils";

import Button from "~/design/Button";
import { Card } from "~/design/Card";
import EndpointPreview from "./EndpointPreview";
import { JSONSchemaForm } from "~/design/JSONschemaForm";
import { Save } from "lucide-react";
import { ScrollArea } from "~/design/ScrollArea";
import { stringify } from "json-to-pretty-yaml";
import { useNodeContent } from "~/api/tree/query/node";
import { useTranslation } from "react-i18next";
import { useUpdateWorkflow } from "~/api/tree/mutate/updateWorkflow";
import yamljs from "js-yaml";

export type NodeContentType = ReturnType<typeof useNodeContent>["data"];

const EndpointEditor: FC<{
  data: NonNullable<NodeContentType>;
  path: string;
}> = ({ data, path }) => {
  const { t } = useTranslation();
  const workflowData = atob(data.revision?.source ?? "");

  const [hasUnsavedChanges, setHasUnsavedChanges] = useState(false);
  const [endpointConfigJson, setEndpointConfigJson] = useState(() => {
    let json;
    try {
      json = yamljs.load(workflowData);
    } catch (e) {
      json = null;
    }

    return json as Record<string, unknown>;
  });

  const { mutate: updateEndpoint, isLoading } = useUpdateWorkflow({
    onSuccess: () => {
      setHasUnsavedChanges(false);
    },
  });

  const endpointFormSchema = useEndpointFormSchema();

  const onSaveClicked = () => {
    const toSave = stringify(endpointConfigJson);
    if (toSave) {
      updateEndpoint({
        path,
        fileContent: toSave,
      });
    }
  };

  return (
    <div className="relative flex grow flex-col space-y-4 p-5">
      <Card className="flex grow flex-col p-4">
        <div className="grow">
          <ScrollArea className="flex grow p-4">
            <JSONSchemaForm
              formData={endpointConfigJson}
              onChange={(e) => {
                if (e.formData) {
                  const formDataWithHeader = addEndpointHeader(e.formData);
                  setHasUnsavedChanges(true);
                  setEndpointConfigJson(formDataWithHeader);
                }
              }}
              /**
               * omitExtraData is important when a the plugin selector is used
               * to change a plugin. This will wipe all the data from the previous
               * plugin.
               */
              omitExtraData={true}
              schema={endpointFormSchema}
            />
          </ScrollArea>
        </div>
        <div className="flex justify-end gap-2 pt-2 text-sm text-gray-8 dark:text-gray-dark-8">
          {hasUnsavedChanges && (
            <span className="text-center">
              {t("pages.explorer.workflow.editor.unsavedNote")}
            </span>
          )}
        </div>
      </Card>
      <div className="flex flex-col justify-end gap-4 sm:flex-row sm:items-center">
        <EndpointPreview fileContent={stringify(endpointConfigJson)} />
        <Button variant="outline" disabled={isLoading} onClick={onSaveClicked}>
          <Save />
          {t("pages.explorer.gateway.editor.saveBtn")}
        </Button>
      </div>
    </div>
  );
};

export default EndpointEditor;

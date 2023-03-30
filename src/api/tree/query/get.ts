import { apiFactory, defaultKeys } from "../../utils";
import { forceSlashIfPath, sortFoldersFirst } from "../utils";

import type { QueryFunctionContext } from "@tanstack/react-query";
import { TreeListSchema } from "../schema";
import { treeKeys } from "../";
import { useApiKey } from "../../../util/store/apiKey";
import { useNamespace } from "../../../util/store/namespace";
import { useQuery } from "@tanstack/react-query";
import { useToast } from "../../../design/Toast";

const getTree = apiFactory({
  pathFn: ({ namespace, path }: { namespace: string; path?: string }) =>
    `/api/namespaces/${namespace}/tree${forceSlashIfPath(path)}`,
  method: "GET",
  schema: TreeListSchema,
});

const fetchTree = async ({
  queryKey: [{ apiKey, namespace, path }],
}: QueryFunctionContext<ReturnType<(typeof treeKeys)["all"]>>) =>
  getTree({
    apiKey: apiKey,
    params: undefined,
    pathParams: {
      namespace,
      path,
    },
  });

export const useListDirectory = ({
  path,
}: {
  path?: string;
} = {}) => {
  const apiKey = useApiKey();
  const namespace = useNamespace();
  const { toast } = useToast();

  if (!namespace) {
    throw new Error("namespace is undefined");
  }

  return useQuery({
    queryKey: treeKeys.all(apiKey ?? defaultKeys.apiKey, namespace, path ?? ""),
    queryFn: fetchTree,
    select(data) {
      if (data?.children?.results) {
        return {
          ...data,
          children: {
            ...data.children,
            results: data.children.results.sort(sortFoldersFirst),
          },
        };
      }
      return data;
    },
    enabled: !!namespace,
    onError: () => {
      toast({
        title: "An error occurred",
        description: "could not fetch directory content 😢",
        variant: "error",
      });
    },
  });
};

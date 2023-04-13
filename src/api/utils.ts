import { z } from "zod";

const getAuthHeader = (apiKey: string) => ({
  "direktiv-token": apiKey,
});

/**
 * atm params must alway be defined. I tried to make TS infer the property
 * with
 *
 * type ReturnT<TParams> = {
 *   apiKey: string;
 * } & (TParams extends undefined ? object : { params: Partial<TParams> });
 *
 * but it didn't work. I also tried
 *
 * type ReturnT<TParams> = {
 *   apiKey: string;
 *   params?: TParams;
 * };
 *
 * but this would have the downside that params is always optional. And we would
 * lose typesafety when some api enpoints have required params
 *
 */
type ApiParams<TParams, TPathParams> = {
  apiKey?: string;
  params: TParams extends undefined ? undefined : TParams;
  pathParams: TPathParams;
};

export const apiFactory =
  <TSchema, TParams, TPathParams>({
    // the path to the api endpoint
    pathFn: path,
    // the http method that should be used for the request
    method,
    // the zod schema that the response should be parsed against. This will give
    // us not only the typesafety of the response, it also validates the response
    // at runtime. Runtime validation is important to catch unexpected responses
    // fromt the api very early in the application lifecycle and give us confidence
    // about the Typescript types. It comes with the downside that the app is more
    // likely to show errors to the user instead of trying to handle them (which
    // does not scale very well when the complexity of an app grows and leads to
    // even worse user experience).
    schema,
  }: {
    pathFn: (pathParams: TPathParams) => string;
    method: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
    schema: z.ZodSchema<TSchema>;
  }): (({
    apiKey,
    params,
    pathParams,
  }: ApiParams<TParams, TPathParams>) => Promise<TSchema>) =>
  async ({ apiKey, params, pathParams }): Promise<TSchema> => {
    const res = await fetch(path(pathParams), {
      method,
      headers: {
        ...(apiKey ? getAuthHeader(apiKey) : {}),
      },
      ...(params
        ? {
            body: typeof params === "string" ? params : JSON.stringify(params),
          }
        : {}),
    });
    if (res.ok) {
      // if we can not evaluate the response, we have null as the default
      let parsedResponse = null;
      const textResult = await res.text();
      try {
        // try to parse the response as json
        parsedResponse = JSON.parse(textResult);
      } catch (e) {
        // We use the text response if its not an empt string
        if (textResult !== "") parsedResponse = textResult;
      }
      try {
        return schema.parse(parsedResponse);
      } catch (error) {
        process.env.NODE_ENV !== "test" && console.error(error);
        return Promise.reject(
          `could not format response for ${method} ${path(pathParams)}`
        );
      }
    }
    return Promise.reject(
      `error ${res.status} for ${method} ${path(pathParams)}`
    );
  };

export const defaultKeys = {
  apiKey: "no-api-key",
};

import "cross-fetch/polyfill";

import {
  afterAll,
  afterEach,
  beforeAll,
  describe,
  expect,
  test,
  vi,
} from "vitest";
import { renderHook, waitFor } from "@testing-library/react";

import { UseQueryWrapper } from "../../../test/utils";
import { apiFactory } from "../utils";
import { rest } from "msw";
import { setupServer } from "msw/node";
import { useQuery } from "@tanstack/react-query";
import { z } from "zod";

const API_KEY = "THIS-IS-MY-API-KEY";
const apiEndpoint = "http://localhost/my-api";
const apiEndpoint404 = "http://localhost/404";
const apiEndpointWithDynamicSegment = "http://localhost/this-is-dynamic/my-api";

const testApi = setupServer(
  rest.get(apiEndpoint, (req, res, ctx) =>
    req?.headers?.get("direktiv-token") === API_KEY
      ? res(
          ctx.json({
            response: "this works",
          })
        )
      : res(ctx.status(401))
  ),
  rest.get(apiEndpointWithDynamicSegment, (req, res, ctx) =>
    req?.headers?.get("direktiv-token") === API_KEY
      ? res(
          ctx.json({
            response: "dynamic segment this works",
          })
        )
      : res(ctx.status(401))
  ),
  rest.get(apiEndpoint404, (_req, res, ctx) => res(ctx.status(404)))
);

beforeAll(() => {
  testApi.listen({ onUnhandledRequest: "error" });
});

afterAll(() => testApi.close());

afterEach(() => {
  vi.clearAllMocks();
  testApi.resetHandlers();
});

const getMyApi = apiFactory({
  pathFn: () => apiEndpoint,
  method: "GET",
  schema: z.object({
    response: z.string(),
  }),
});

const getMyApiWrongSchema = apiFactory({
  pathFn: () => apiEndpoint,
  method: "GET",
  schema: z.object({
    response: z.number(), // this will fail, since the repsonse is a string
  }),
});

const api404 = apiFactory({
  pathFn: () => apiEndpoint404,
  method: "GET",
  schema: z.object({
    response: z.string(),
  }),
});

const getMyWithDynamicSegment = apiFactory({
  pathFn: ({ segment }: { segment: string }) =>
    `http://localhost/${segment}/my-api`,
  method: "GET",
  schema: z.object({
    response: z.string(),
  }),
});

describe("processApiResponse", () => {
  test("api response and schema gets validated succesfully", async () => {
    const useCallWithApiKey = () =>
      useQuery({
        queryKey: ["getmyapikey", API_KEY],
        queryFn: () =>
          getMyApi({
            apiKey: API_KEY,
            params: undefined,
            pathParams: undefined,
          }),
      });

    const { result } = renderHook(() => useCallWithApiKey(), {
      wrapper: UseQueryWrapper,
    });
    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
      expect(result.current.data?.response).toBe("this works");
    });
  });

  test("unauthenticated response", async () => {
    const useCallWithOutApiKey = (onError: (err: unknown) => void) =>
      useQuery({
        queryKey: ["getmyapikey", "wrong-api-key"],
        queryFn: () =>
          getMyApi({
            apiKey: "wrong-api-key",
            params: undefined,
            pathParams: undefined,
          }),
        onError,
      });

    const errorMock = vi.fn();
    const { result } = renderHook(() => useCallWithOutApiKey(errorMock), {
      wrapper: UseQueryWrapper,
    });
    await waitFor(() => {
      expect(result.current.isSuccess).toBe(false);
      expect(result.current.status).toBe("error");
      expect(errorMock.mock.calls[0][0]).toMatchInlineSnapshot(
        '"error 401 for GET http://localhost/my-api"'
      );
    });
  });

  test("response fails schema validation", async () => {
    const useCallWithApiKey = (onError: (err: unknown) => void) =>
      useQuery({
        queryKey: ["getmyapikey", API_KEY],
        queryFn: () =>
          getMyApiWrongSchema({
            apiKey: API_KEY,
            params: undefined,
            pathParams: undefined,
          }),
        onError,
      });

    const errorMock = vi.fn();
    const { result } = renderHook(() => useCallWithApiKey(errorMock), {
      wrapper: UseQueryWrapper,
    });
    await waitFor(() => {
      expect(result.current.isSuccess).toBe(false);
      expect(result.current.status).toBe("error");
      expect(errorMock.mock.calls[0][0]).toMatchInlineSnapshot(
        '"could not format response for GET http://localhost/my-api"'
      );
    });
  });

  test("api path does not exist", async () => {
    const useCallWithApiKey = (onError: (err: unknown) => void) =>
      useQuery({
        queryKey: ["getmyapikey", API_KEY],
        queryFn: () =>
          api404({ apiKey: API_KEY, params: undefined, pathParams: undefined }),
        onError,
      });

    const errorMock = vi.fn();
    const { result } = renderHook(() => useCallWithApiKey(errorMock), {
      wrapper: UseQueryWrapper,
    });
    await waitFor(() => {
      expect(result.current.isSuccess).toBe(false);
      expect(result.current.status).toBe("error");
      expect(errorMock.mock.calls[0][0]).toMatchInlineSnapshot(
        '"error 404 for GET http://localhost/404"'
      );
    });
  });

  test("api with dynamic segment", async () => {
    const useCallWithApiKey = (pathParams: { segment: string }) =>
      useQuery({
        queryKey: ["getmyapikey", API_KEY, pathParams],
        queryFn: () =>
          getMyWithDynamicSegment({
            apiKey: API_KEY,
            params: null,
            pathParams: pathParams,
          }),
      });

    const { result } = renderHook(
      () => useCallWithApiKey({ segment: "this-is-dynamic" }),
      {
        wrapper: UseQueryWrapper,
      }
    );
    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
      expect(result.current.data?.response).toBe("dynamic segment this works");
    });
  });
});

import * as React from "react";
import {
  CloseEventSource,
  HandleError,
  ExtractQueryString,
  apiKeyHeaders,
} from "../util";
import fetch from "isomorphic-fetch";
import { EventSourcePolyfill } from "event-source-polyfill";

/* 
    useNamespaceServiceRevision takes
    - url
    - namespace
    - service
    - revision
    - apikey
*/
export const useDirektivNamespaceServiceRevision = (
  url,
  namespace,
  service,
  revision,
  apikey
) => {
  const [revisionDetails, setRevisionDetails] = React.useState(null);
  const [pods, setPods] = React.useState([]);
  const [err, setErr] = React.useState(null);
  const podSource = React.useRef(null);
  const revisionSource = React.useRef(null);

  const podsRef = React.useRef(pods);

  React.useEffect(() => {
    if (podSource.current === null) {
      const listener = new EventSourcePolyfill(
        `${url}functions/namespaces/${namespace}/function/${service}/revisions/${revision}/pods`,
        {
          headers: apiKeyHeaders(apikey),
        }
      );

      listener.onerror = (e) => {
        if (e.status === 404) {
          setErr(e.statusText);
        } else if (e.status === 403) {
          setErr("permission denied");
        } else {
          try {
            const json = JSON.parse(e.data);
            setErr(json.Message);
          } catch (e) {
            // TODO
          }
        }
      };

      async function readData(e) {
        const podz = podsRef.current;

        if (e.data === "") {
          return;
        }
        const json = JSON.parse(e.data);

        switch (json.event) {
          case "DELETED":
            for (var i = 0; i < pods.length; i++) {
              if (podz[i].name === json.pod.name) {
                podz.splice(i, 1);
                podsRef.current = pods;
                break;
              }
            }
            break;
          case "MODIFIED":
            for (i = 0; i < podz.length; i++) {
              if (podz[i].name === json.pod.name) {
                podz[i] = json.pod;
                podsRef.current = podz;
                break;
              }
            }
            break;
          default:
            let found = false;
            for (i = 0; i < podz.length; i++) {
              if (podz[i].name === json.pod.name) {
                found = true;
                break;
              }
            }
            if (!found) {
              podz.push(json.pod);
              podsRef.current = pods;
            }
        }
        setPods(JSON.parse(JSON.stringify(podsRef.current)));
      }
      listener.onmessage = (e) => readData(e);
      podSource.current = listener;
    }
  }, [apikey, namespace, pods, revision, service, url]);

  React.useEffect(() => {
    if (revisionSource.current === null) {
      // setup event listener
      const listener = new EventSourcePolyfill(
        `${url}functions/namespaces/${namespace}/function/${service}/revisions/${revision}`,
        {
          headers: apiKeyHeaders(apikey),
        }
      );

      listener.onerror = (e) => {
        if (e.status === 404) {
          setErr(e.statusText);
        } else if (e.status === 403) {
          setErr("permission denied");
        } else {
          try {
            const json = JSON.parse(e.data);
            setErr(json.Message);
          } catch (e) {
            // TODO handle error/error handling
          }
        }
      };

      async function readData(e) {
        if (e.data === "") {
          return;
        }
        const json = JSON.parse(e.data);
        if (json.event === "ADDED" || json.event === "MODIFIED") {
          setRevisionDetails(json.revision);
        }
        // if (json.event === "DELETED") {
        //     history.goBack()
        // }
      }

      listener.onmessage = (e) => readData(e);
      revisionSource.current = listener;
    }
  }, [apikey, url, namespace, service, revision]);

  React.useEffect(() => {
    return () => {
      CloseEventSource(revisionSource.current);
      CloseEventSource(podSource.current);
    };
  }, []);

  return {
    revisionDetails,
    pods,
    err,
  };
};
/* 
    useNamespaceService takes
    - url
    - namespace
    - service
    - navigate (react router object to navigate backwards)
    - apikey
*/
export const useDirektivNamespaceService = (
  url,
  namespace,
  service,
  navigate,
  apikey
) => {
  const [revisions, setRevisions] = React.useState(null);
  const [fn, setFn] = React.useState(null);
  const [traffic, setTraffic] = React.useState(null);
  const [config, setConfig] = React.useState(null);
  const revisionsRef = React.useRef(revisions ? revisions : []);

  const [err, setErr] = React.useState(null);

  const trafficSource = React.useRef(null);
  const eventSource = React.useRef(null);

  React.useEffect(() => {
    if (trafficSource.current === null) {
      // setup event listener
      const listener = new EventSourcePolyfill(
        `${url}functions/namespaces/${namespace}/function/${service}`,
        {
          headers: apiKeyHeaders(apikey),
        }
      );

      listener.onerror = (e) => {
        if (e.status === 404) {
          setErr(e.statusText);
        } else if (e.status === 403) {
          setErr("permission denied");
        } else {
          try {
            const json = JSON.parse(e.data);
            setErr(json.Message);
          } catch (e) {
            // TODO
          }
        }
      };

      async function readData(e) {
        if (e.data === "") {
          return;
        }
        const json = JSON.parse(e.data);

        if (json.event === "MODIFIED" || json.event === "ADDED") {
          setFn(JSON.parse(JSON.stringify(json.function)));
          setTraffic(JSON.parse(JSON.stringify(json.traffic)));
        }
      }

      listener.onmessage = (e) => readData(e);

      trafficSource.current = listener;
    }
  }, [fn, apikey, url, namespace, service]);

  React.useEffect(() => {
    if (eventSource.current === null) {
      // setup event listener
      const listener = new EventSourcePolyfill(
        `${url}functions/namespaces/${namespace}/function/${service}/revisions`,
        {
          headers: apiKeyHeaders(apikey),
        }
      );

      listener.onerror = (e) => {
        if (e.status === 404) {
          setErr(e.statusText);
        } else if (e.status === 403) {
          setErr("permission denied");
        } else {
          try {
            const json = JSON.parse(e.data);
            setErr(json.Message);
          } catch (e) {
            // TODO
          }
        }
      };

      async function readData(e) {
        const revs = revisionsRef.current;
        if (e.data === "") {
          return;
        }
        const json = JSON.parse(e.data);
        switch (json.event) {
          case "DELETED":
            for (var i = 0; i < revs.length; i++) {
              if (revs[i].name === json.revision.name) {
                revs.splice(i, 1);
                revisionsRef.current = revs;
                break;
              }
            }
            if (revs.length === 0) {
              navigate(-1);
            }
            break;
          case "MODIFIED":
            for (i = 0; i < revs.length; i++) {
              if (revs[i].name === json.revision.name) {
                revs[i] = json.revision;
                revisionsRef.current = revs;
                break;
              }
            }
            break;
          default:
            let found = false;
            for (i = 0; i < revs.length; i++) {
              if (revs[i].name === json.revision.name) {
                found = true;
                break;
              }
            }
            if (!found) {
              revs.push(json.revision);
              revisionsRef.current = revs;
            }
        }
        revisionsRef.current.sort(function (a, b) {
          return parseInt(a.generation) < parseInt(b.generation) ? 1 : -1;
        });
        setRevisions(JSON.parse(JSON.stringify(revisionsRef.current)));
      }

      listener.onmessage = (e) => readData(e);
      eventSource.current = listener;
    }
  }, [revisions, apikey, url, namespace, service, navigate]);

  React.useEffect(() => {
    return () => {
      CloseEventSource(eventSource.current);
      CloseEventSource(trafficSource.current);
    };
  }, []);

  async function getNamespaceServiceConfig(...queryParameters) {
    const resp = await fetch(
      `${url}functions/namespaces/${namespace}/function/${service}${ExtractQueryString(
        false,
        ...queryParameters
      )}`,
      {
        headers: apiKeyHeaders(apikey),
        method: "GET",
      }
    );
    if (resp.ok) {
      const json = await resp.json();
      setConfig(json.config);
      return json.config;
    } else {
      throw new Error(
        await HandleError("get namespace service", resp, "getService")
      );
    }
  }

  async function createNamespaceServiceRevision(
    image,
    minScale,
    size,
    cmd,
    traffic,
    ...queryParameters
  ) {
    const resp = await fetch(
      `${url}functions/namespaces/${namespace}/function/${service}${ExtractQueryString(
        false,
        ...queryParameters
      )}`,
      {
        headers: apiKeyHeaders(apikey),
        method: "POST",
        body: JSON.stringify({
          trafficPercent: traffic,
          cmd,
          image,
          minScale,
          size,
        }),
      }
    );
    if (!resp.ok) {
      throw new Error(
        await HandleError("create namespace service revision", resp)
      );
    }
  }

  async function deleteNamespaceServiceRevision(rev, ...queryParameters) {
    const resp = await fetch(
      `${url}functions/namespaces/${namespace}/function/${service}/revisions/${rev}${ExtractQueryString(
        false,
        ...queryParameters
      )}`,
      {
        method: "DELETE",
        headers: apiKeyHeaders(apikey),
      }
    );
    if (!resp.ok) {
      throw new Error(
        await HandleError(
          "delete namespace service revision",
          resp,
          "deleteRevision"
        )
      );
    }
  }

  async function setNamespaceServiceRevisionTraffic(
    rev1,
    rev1value,
    rev2,
    rev2value,
    ...queryParameters
  ) {
    const trafficarr = [];
    if (rev1 !== "") {
      trafficarr.push({
        revision: rev1,
        percent: rev1value,
      });
    }
    if (rev2 !== "") {
      trafficarr.push({
        revision: rev2,
        percent: rev2value,
      });
    }
    const resp = await fetch(
      `${url}functions/namespaces/${namespace}/function/${service}${ExtractQueryString(
        false,
        ...queryParameters
      )}`,
      {
        method: "PATCH",
        body: JSON.stringify({ values: trafficarr }),
        headers: apiKeyHeaders(apikey),
      }
    );
    if (!resp.ok) {
      throw new Error(
        await HandleError(
          "update traffic namespace service",
          resp,
          "updateTraffic"
        )
      );
    }
  }

  return {
    revisions,
    fn,
    config,
    traffic,
    err,
    createNamespaceServiceRevision,
    deleteNamespaceServiceRevision,
    getNamespaceServiceConfig,
    setNamespaceServiceRevisionTraffic,
  };
};
/*
    useNamespaceServices is a react hook 
    takes:
      - url to direktiv api http://x/api/
      - stream to use sse or a normal fetch
      - namespace to use for the api
      - apikey to provide authentication of an apikey
*/
export const useDirektivNamespaceServices = (
  url,
  stream,
  namespace,
  apikey
) => {
  const [data, setData] = React.useState(null);
  const [config, setConfig] = React.useState(null);
  const functionsRef = React.useRef(data ? data : []);
  const [err, setErr] = React.useState(null);
  const eventSource = React.useRef(null);

  const getNamespaceServices = React.useCallback(
    async (...queryParameters) => {
      const resp = await fetch(
        `${url}functions/namespaces/${namespace}${ExtractQueryString(
          false,
          ...queryParameters
        )}`,
        {
          headers: apiKeyHeaders(apikey),
          method: "GET",
        }
      );
      if (resp.ok) {
        const json = await resp.json();
        setData(json.functions);
        return json.functions;
      } else {
        throw new Error(
          await HandleError("get namespace service", resp, "listServices")
        );
      }
    },
    [apikey, namespace, url]
  );

  React.useEffect(() => {
    if (stream) {
      if (eventSource.current === null) {
        // setup event listener
        const listener = new EventSourcePolyfill(
          `${url}functions/namespaces/${namespace}`,
          {
            headers: apiKeyHeaders(apikey),
          }
        );

        listener.onerror = (e) => {
          if (e.status === 404) {
            setErr(e.statusText);
          } else if (e.status === 403) {
            setErr("permission denied");
          } else {
            try {
              const json = JSON.parse(e.data);
              setErr(json.Message);
            } catch (e) {
              // TODO
            }
          }
        };

        async function readData(e) {
          const funcs = functionsRef.current;
          if (e.data === "") {
            return;
          }
          const json = JSON.parse(e.data);
          switch (json.event) {
            case "DELETED":
              for (var i = 0; i < funcs.length; i++) {
                if (funcs[i].serviceName === json.function.serviceName) {
                  funcs.splice(i, 1);
                  functionsRef.current = funcs;
                  break;
                }
              }
              break;
            case "MODIFIED":
              for (i = 0; i < funcs.length; i++) {
                if (funcs[i].serviceName === json.function.serviceName) {
                  funcs[i] = json.function;
                  functionsRef.current = funcs;
                  break;
                }
              }
              break;
            default:
              let found = false;
              for (i = 0; i < funcs.length; i++) {
                if (funcs[i].serviceName === json.function.serviceName) {
                  found = true;
                  break;
                }
              }
              if (!found) {
                funcs.push(json.function);
                functionsRef.current = funcs;
              }
          }
          setData(JSON.parse(JSON.stringify(functionsRef.current)));
        }

        listener.onmessage = (e) => readData(e);
        eventSource.current = listener;
      }
    } else {
      if (data === null) {
        getNamespaceServices();
      }
    }
  }, [data, apikey, stream, url, namespace, getNamespaceServices]);

  React.useEffect(() => {
    return () => {
      CloseEventSource(eventSource.current);
    };
  }, []);

  async function getNamespaceConfig(...queryParameters) {
    const resp = await fetch(
      `${url}functions/namespaces/${namespace}${ExtractQueryString(
        false,
        ...queryParameters
      )}`,
      {
        headers: apiKeyHeaders(apikey),
        method: "GET",
      }
    );
    if (resp.ok) {
      const json = await resp.json();
      setConfig(json.config);
      return json.config;
    } else {
      throw new Error(
        await HandleError("get namespace service", resp, "listServices")
      );
    }
  }

  async function createNamespaceService(
    name,
    image,
    minScale,
    size,
    cmd,
    ...queryParameters
  ) {
    const resp = await fetch(
      `${url}functions/namespaces/${namespace}${ExtractQueryString(
        false,
        ...queryParameters
      )}`,
      {
        headers: apiKeyHeaders(apikey),
        method: "POST",
        body: JSON.stringify({
          cmd,
          image,
          minScale,
          name,
          size,
        }),
      }
    );
    if (!resp.ok) {
      throw new Error(await HandleError("create namespace service", resp));
    }
  }

  async function deleteNamespaceService(name, ...queryParameters) {
    const resp = await fetch(
      `${url}functions/namespaces/${namespace}/function/${name}${ExtractQueryString(
        false,
        ...queryParameters
      )}`,
      {
        headers: apiKeyHeaders(apikey),
        method: "DELETE",
      }
    );
    if (!resp.ok) {
      throw new Error(
        await HandleError("delete namespace service", resp, "deleteService")
      );
    }
  }

  return {
    data,
    err,
    config,
    getNamespaceConfig,
    getNamespaceServices,
    createNamespaceService,
    deleteNamespaceService,
  };
};

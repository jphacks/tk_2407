import type { AspidaClient, BasicHeaders } from 'aspida';
import type { Methods as Methods_18qsrps } from './health';

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? '' : baseURL).replace(/\/$/, '');
  const PATH0 = '/health';
  const GET = 'GET';

  return {
    health: {
      /**
       * Returns the health status of the application in JSON format.
       * @returns OK - Service is healthy
       */
      get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods_18qsrps['get']['resBody'], BasicHeaders, Methods_18qsrps['get']['status']>(prefix, PATH0, GET, option).json(),
      /**
       * Returns the health status of the application in JSON format.
       * @returns OK - Service is healthy
       */
      $get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods_18qsrps['get']['resBody'], BasicHeaders, Methods_18qsrps['get']['status']>(prefix, PATH0, GET, option).json().then(r => r.body),
      $path: () => `${prefix}${PATH0}`,
    },
  };
};

export type ApiInstance = ReturnType<typeof api>;
export default api;

import type { AspidaClient, BasicHeaders } from 'aspida';
import type { Methods as Methods_18qsrps } from './health';

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? '' : baseURL).replace(/\/$/, '');
  const PATH0 = '/health';
  const GET = 'GET';

  return {
    health: {
      /**
       * health check
       * @returns OK
       */
      get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods_18qsrps['get']['resBody'], BasicHeaders, Methods_18qsrps['get']['status']>(prefix, PATH0, GET, option).json(),
      /**
       * health check
       * @returns OK
       */
      $get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods_18qsrps['get']['resBody'], BasicHeaders, Methods_18qsrps['get']['status']>(prefix, PATH0, GET, option).json().then(r => r.body),
      $path: () => `${prefix}${PATH0}`,
    },
  };
};

export type ApiInstance = ReturnType<typeof api>;
export default api;

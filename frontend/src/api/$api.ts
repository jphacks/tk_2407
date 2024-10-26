import type { AspidaClient, BasicHeaders } from 'aspida';
import type { Methods as Methods_18qsrps } from './health';
import type { Methods as Methods_gbpz4m } from './messages';

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? '' : baseURL).replace(/\/$/, '');
  const PATH0 = '/health';
  const PATH1 = '/messages';
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
    messages: {
      /**
       * get message list
       * @returns OK
       */
      get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods_gbpz4m['get']['resBody'], BasicHeaders, Methods_gbpz4m['get']['status']>(prefix, PATH1, GET, option).json(),
      /**
       * get message list
       * @returns OK
       */
      $get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods_gbpz4m['get']['resBody'], BasicHeaders, Methods_gbpz4m['get']['status']>(prefix, PATH1, GET, option).json().then(r => r.body),
      $path: () => `${prefix}${PATH1}`,
    },
  };
};

export type ApiInstance = ReturnType<typeof api>;
export default api;

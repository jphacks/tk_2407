import type { AspidaClient, BasicHeaders } from 'aspida';
import type { Methods as Methods_4r2bqf } from './api/v1/messages';
import type { Methods as Methods_18qsrps } from './health';

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? '' : baseURL).replace(/\/$/, '');
  const PATH0 = '/api/v1/messages';
  const PATH1 = '/health';
  const GET = 'GET';

  return {
    api: {
      v1: {
        messages: {
          /**
           * get message list
           * @returns OK
           */
          get: (option?: { config?: T | undefined } | undefined) =>
            fetch<Methods_4r2bqf['get']['resBody'], BasicHeaders, Methods_4r2bqf['get']['status']>(prefix, PATH0, GET, option).json(),
          /**
           * get message list
           * @returns OK
           */
          $get: (option?: { config?: T | undefined } | undefined) =>
            fetch<Methods_4r2bqf['get']['resBody'], BasicHeaders, Methods_4r2bqf['get']['status']>(prefix, PATH0, GET, option).json().then(r => r.body),
          $path: () => `${prefix}${PATH0}`,
        },
      },
    },
    health: {
      /**
       * health check
       * @returns OK
       */
      get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods_18qsrps['get']['resBody'], BasicHeaders, Methods_18qsrps['get']['status']>(prefix, PATH1, GET, option).json(),
      /**
       * health check
       * @returns OK
       */
      $get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods_18qsrps['get']['resBody'], BasicHeaders, Methods_18qsrps['get']['status']>(prefix, PATH1, GET, option).json().then(r => r.body),
      $path: () => `${prefix}${PATH1}`,
    },
  };
};

export type ApiInstance = ReturnType<typeof api>;
export default api;

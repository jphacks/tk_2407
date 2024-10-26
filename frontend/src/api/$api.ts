import type { AspidaClient, BasicHeaders } from 'aspida';
import type { Methods as Methods_1hjgtff } from './api/v1/messages/_location_id@string';
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
          _location_id: (val3: string) => {
            const prefix3 = `${PATH0}/${val3}`;

            return {
              /**
               * get message list
               * @returns OK
               */
              get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods_1hjgtff['get']['resBody'], BasicHeaders, Methods_1hjgtff['get']['status']>(prefix, prefix3, GET, option).json(),
              /**
               * get message list
               * @returns OK
               */
              $get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods_1hjgtff['get']['resBody'], BasicHeaders, Methods_1hjgtff['get']['status']>(prefix, prefix3, GET, option).json().then(r => r.body),
              $path: () => `${prefix}${prefix3}`,
            };
          },
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

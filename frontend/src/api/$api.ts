import type { AspidaClient, BasicHeaders } from 'aspida';
import { dataToURLString } from 'aspida';
import type { Methods as Methods_6l8fl5 } from './api/v1/health';
import type { Methods as Methods_1tsk0v8 } from './api/v1/login';
import type { Methods as Methods_1ibm08e } from './api/v1/message';
import type { Methods as Methods_bjat18 } from './api/v1/messages/_locationId@string';
import type { Methods as Methods_3ai8p3 } from './api/v1/signup';
import type { Methods as Methods_p6a2jm } from './api/v1/spot/_spotId@string/photo';
import type { Methods as Methods_qb0a98 } from './api/v1/spots';
import type { Methods as Methods_6lsccp } from './api/v1/user/_userId@string';
import type { Methods as Methods_15rl5t7 } from './api/v1/user/me';

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? '' : baseURL).replace(/\/$/, '');
  const PATH0 = '/api/v1/health';
  const PATH1 = '/api/v1/login';
  const PATH2 = '/api/v1/message';
  const PATH3 = '/api/v1/messages';
  const PATH4 = '/api/v1/signup';
  const PATH5 = '/api/v1/spot';
  const PATH6 = '/photo';
  const PATH7 = '/api/v1/spots';
  const PATH8 = '/api/v1/user';
  const PATH9 = '/api/v1/user/me';
  const GET = 'GET';
  const POST = 'POST';

  return {
    api: {
      v1: {
        health: {
          /**
           * Health check
           * @returns OK
           */
          get: (option?: { config?: T | undefined } | undefined) =>
            fetch<Methods_6l8fl5['get']['resBody'], BasicHeaders, Methods_6l8fl5['get']['status']>(prefix, PATH0, GET, option).json(),
          /**
           * Health check
           * @returns OK
           */
          $get: (option?: { config?: T | undefined } | undefined) =>
            fetch<Methods_6l8fl5['get']['resBody'], BasicHeaders, Methods_6l8fl5['get']['status']>(prefix, PATH0, GET, option).json().then(r => r.body),
          $path: () => `${prefix}${PATH0}`,
        },
        login: {
          /**
           * Login
           * @returns Created
           */
          post: (option: { body: Methods_1tsk0v8['post']['reqBody'], config?: T | undefined }) =>
            fetch<Methods_1tsk0v8['post']['resBody'], BasicHeaders, Methods_1tsk0v8['post']['status']>(prefix, PATH1, POST, option).json(),
          /**
           * Login
           * @returns Created
           */
          $post: (option: { body: Methods_1tsk0v8['post']['reqBody'], config?: T | undefined }) =>
            fetch<Methods_1tsk0v8['post']['resBody'], BasicHeaders, Methods_1tsk0v8['post']['status']>(prefix, PATH1, POST, option).json().then(r => r.body),
          $path: () => `${prefix}${PATH1}`,
        },
        message: {
          /**
           * Post a message
           * @returns Created
           */
          post: (option: { body: Methods_1ibm08e['post']['reqBody'], config?: T | undefined }) =>
            fetch<Methods_1ibm08e['post']['resBody'], BasicHeaders, Methods_1ibm08e['post']['status']>(prefix, PATH2, POST, option).json(),
          /**
           * Post a message
           * @returns Created
           */
          $post: (option: { body: Methods_1ibm08e['post']['reqBody'], config?: T | undefined }) =>
            fetch<Methods_1ibm08e['post']['resBody'], BasicHeaders, Methods_1ibm08e['post']['status']>(prefix, PATH2, POST, option).json().then(r => r.body),
          $path: () => `${prefix}${PATH2}`,
        },
        messages: {
          _locationId: (val3: string) => {
            const prefix3 = `${PATH3}/${val3}`;

            return {
              /**
               * Get message list
               * @returns OK
               */
              get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods_bjat18['get']['resBody'], BasicHeaders, Methods_bjat18['get']['status']>(prefix, prefix3, GET, option).json(),
              /**
               * Get message list
               * @returns OK
               */
              $get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods_bjat18['get']['resBody'], BasicHeaders, Methods_bjat18['get']['status']>(prefix, prefix3, GET, option).json().then(r => r.body),
              $path: () => `${prefix}${prefix3}`,
            };
          },
        },
        signup: {
          /**
           * Signup
           * @returns Created
           */
          post: (option: { body: Methods_3ai8p3['post']['reqBody'], config?: T | undefined }) =>
            fetch<Methods_3ai8p3['post']['resBody'], BasicHeaders, Methods_3ai8p3['post']['status']>(prefix, PATH4, POST, option).json(),
          /**
           * Signup
           * @returns Created
           */
          $post: (option: { body: Methods_3ai8p3['post']['reqBody'], config?: T | undefined }) =>
            fetch<Methods_3ai8p3['post']['resBody'], BasicHeaders, Methods_3ai8p3['post']['status']>(prefix, PATH4, POST, option).json().then(r => r.body),
          $path: () => `${prefix}${PATH4}`,
        },
        spot: {
          _spotId: (val3: string) => {
            const prefix3 = `${PATH5}/${val3}`;

            return {
              photo: {
                /**
                 * Retrieve a list of locations based on longitude, latitude
                 * @returns OK
                 */
                get: (option?: { config?: T | undefined } | undefined) =>
                  fetch<Methods_p6a2jm['get']['resBody'], BasicHeaders, Methods_p6a2jm['get']['status']>(prefix, `${prefix3}${PATH6}`, GET, option).blob(),
                /**
                 * Retrieve a list of locations based on longitude, latitude
                 * @returns OK
                 */
                $get: (option?: { config?: T | undefined } | undefined) =>
                  fetch<Methods_p6a2jm['get']['resBody'], BasicHeaders, Methods_p6a2jm['get']['status']>(prefix, `${prefix3}${PATH6}`, GET, option).blob().then(r => r.body),
                $path: () => `${prefix}${prefix3}${PATH6}`,
              },
            };
          },
        },
        spots: {
          /**
           * Retrieve a list of locations based on longitude, latitude
           * @returns OK
           */
          get: (option: { query: Methods_qb0a98['get']['query'], config?: T | undefined }) =>
            fetch<Methods_qb0a98['get']['resBody'], BasicHeaders, Methods_qb0a98['get']['status']>(prefix, PATH7, GET, option).json(),
          /**
           * Retrieve a list of locations based on longitude, latitude
           * @returns OK
           */
          $get: (option: { query: Methods_qb0a98['get']['query'], config?: T | undefined }) =>
            fetch<Methods_qb0a98['get']['resBody'], BasicHeaders, Methods_qb0a98['get']['status']>(prefix, PATH7, GET, option).json().then(r => r.body),
          $path: (option?: { method?: 'get' | undefined; query: Methods_qb0a98['get']['query'] } | undefined) =>
            `${prefix}${PATH7}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`,
        },
        user: {
          _userId: (val3: string) => {
            const prefix3 = `${PATH8}/${val3}`;

            return {
              /**
               * Get By ID
               * @returns OK
               */
              get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods_6lsccp['get']['resBody'], BasicHeaders, Methods_6lsccp['get']['status']>(prefix, prefix3, GET, option).json(),
              /**
               * Get By ID
               * @returns OK
               */
              $get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods_6lsccp['get']['resBody'], BasicHeaders, Methods_6lsccp['get']['status']>(prefix, prefix3, GET, option).json().then(r => r.body),
              $path: () => `${prefix}${prefix3}`,
            };
          },
          me: {
            /**
             * Get current user
             * @returns OK
             */
            get: (option?: { config?: T | undefined } | undefined) =>
              fetch<Methods_15rl5t7['get']['resBody'], BasicHeaders, Methods_15rl5t7['get']['status']>(prefix, PATH9, GET, option).json(),
            /**
             * Get current user
             * @returns OK
             */
            $get: (option?: { config?: T | undefined } | undefined) =>
              fetch<Methods_15rl5t7['get']['resBody'], BasicHeaders, Methods_15rl5t7['get']['status']>(prefix, PATH9, GET, option).json().then(r => r.body),
            $path: () => `${prefix}${PATH9}`,
          },
        },
      },
    },
  };
};

export type ApiInstance = ReturnType<typeof api>;
export default api;

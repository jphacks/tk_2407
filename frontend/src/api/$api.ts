import type { AspidaClient, BasicHeaders } from 'aspida';
import { dataToURLString } from 'aspida';
import type { Methods as Methods_6l8fl5 } from './api/v1/health';
import type { Methods as Methods_1tsk0v8 } from './api/v1/login';
import type { Methods as Methods_1ibm08e } from './api/v1/message';
import type { Methods as Methods_13u71ai } from './api/v1/message/stamp';
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
  const PATH3 = '/api/v1/message/stamp';
  const PATH4 = '/api/v1/messages';
  const PATH5 = '/api/v1/signup';
  const PATH6 = '/api/v1/spot';
  const PATH7 = '/photo';
  const PATH8 = '/api/v1/spots';
  const PATH9 = '/api/v1/user';
  const PATH10 = '/api/v1/user/me';
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
          stamp: {
            /**
             * Post a stamp
             * @returns Created
             */
            post: (option: { body: Methods_13u71ai['post']['reqBody'], config?: T | undefined }) =>
              fetch<Methods_13u71ai['post']['resBody'], BasicHeaders, Methods_13u71ai['post']['status']>(prefix, PATH3, POST, option).json(),
            /**
             * Post a stamp
             * @returns Created
             */
            $post: (option: { body: Methods_13u71ai['post']['reqBody'], config?: T | undefined }) =>
              fetch<Methods_13u71ai['post']['resBody'], BasicHeaders, Methods_13u71ai['post']['status']>(prefix, PATH3, POST, option).json().then(r => r.body),
            $path: () => `${prefix}${PATH3}`,
          },
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
            const prefix3 = `${PATH4}/${val3}`;

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
            fetch<Methods_3ai8p3['post']['resBody'], BasicHeaders, Methods_3ai8p3['post']['status']>(prefix, PATH5, POST, option).json(),
          /**
           * Signup
           * @returns Created
           */
          $post: (option: { body: Methods_3ai8p3['post']['reqBody'], config?: T | undefined }) =>
            fetch<Methods_3ai8p3['post']['resBody'], BasicHeaders, Methods_3ai8p3['post']['status']>(prefix, PATH5, POST, option).json().then(r => r.body),
          $path: () => `${prefix}${PATH5}`,
        },
        spot: {
          _spotId: (val3: string) => {
            const prefix3 = `${PATH6}/${val3}`;

            return {
              photo: {
                /**
                 * Retrieve a list of locations based on longitude, latitude
                 * @returns OK
                 */
                get: (option?: { config?: T | undefined } | undefined) =>
                  fetch<Methods_p6a2jm['get']['resBody'], BasicHeaders, Methods_p6a2jm['get']['status']>(prefix, `${prefix3}${PATH7}`, GET, option).blob(),
                /**
                 * Retrieve a list of locations based on longitude, latitude
                 * @returns OK
                 */
                $get: (option?: { config?: T | undefined } | undefined) =>
                  fetch<Methods_p6a2jm['get']['resBody'], BasicHeaders, Methods_p6a2jm['get']['status']>(prefix, `${prefix3}${PATH7}`, GET, option).blob().then(r => r.body),
                $path: () => `${prefix}${prefix3}${PATH7}`,
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
            fetch<Methods_qb0a98['get']['resBody'], BasicHeaders, Methods_qb0a98['get']['status']>(prefix, PATH8, GET, option).json(),
          /**
           * Retrieve a list of locations based on longitude, latitude
           * @returns OK
           */
          $get: (option: { query: Methods_qb0a98['get']['query'], config?: T | undefined }) =>
            fetch<Methods_qb0a98['get']['resBody'], BasicHeaders, Methods_qb0a98['get']['status']>(prefix, PATH8, GET, option).json().then(r => r.body),
          $path: (option?: { method?: 'get' | undefined; query: Methods_qb0a98['get']['query'] } | undefined) =>
            `${prefix}${PATH8}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`,
        },
        user: {
          _userId: (val3: string) => {
            const prefix3 = `${PATH9}/${val3}`;

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
              fetch<Methods_15rl5t7['get']['resBody'], BasicHeaders, Methods_15rl5t7['get']['status']>(prefix, PATH10, GET, option).json(),
            /**
             * Get current user
             * @returns OK
             */
            $get: (option?: { config?: T | undefined } | undefined) =>
              fetch<Methods_15rl5t7['get']['resBody'], BasicHeaders, Methods_15rl5t7['get']['status']>(prefix, PATH10, GET, option).json().then(r => r.body),
            $path: () => `${prefix}${PATH10}`,
          },
        },
      },
    },
  };
};

export type ApiInstance = ReturnType<typeof api>;
export default api;

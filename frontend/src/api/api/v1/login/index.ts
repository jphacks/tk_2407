/* eslint-disable */
import type { DefineMethods } from 'aspida';
import type * as Types from '../../../@types';

export type Methods = DefineMethods<{
  /** Login */
  post: {
    status: 201;
    /** Created */
    resBody: Types.SuccessLoginRes;

    reqBody: {
      email: string;
      password: string;
    };
  };
}>;

/* eslint-disable */
import type { DefineMethods } from 'aspida';
import type * as Types from '../../../@types';

export type Methods = DefineMethods<{
  /** Signup */
  post: {
    status: 201;
    /** Created */
    resBody: Types.SuccessSignupRes;

    reqBody: {
      email: string;
      username: string;
      password: string;
    };
  };
}>;

/* eslint-disable */
import type { DefineMethods } from 'aspida';
import type * as Types from '../../../../@types';

export type Methods = DefineMethods<{
  /** Get current user */
  get: {
    status: 200;
    /** OK */
    resBody: Types.SuccessUserRes;
  };
}>;

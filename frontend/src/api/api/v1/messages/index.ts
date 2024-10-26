/* eslint-disable */
import type { DefineMethods } from 'aspida';
import type * as Types from '../../../@types';

export type Methods = DefineMethods<{
  /** get message list */
  get: {
    status: 200;
    /** OK */
    resBody: Types.SuccessMessageRes;
  };
}>;

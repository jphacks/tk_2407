/* eslint-disable */
import type { DefineMethods } from 'aspida';
import type * as Types from '../../../@types';

export type Methods = DefineMethods<{
  /** health check */
  get: {
    status: 200;
    /** OK */
    resBody: Types.HealthRes;
  };
}>;

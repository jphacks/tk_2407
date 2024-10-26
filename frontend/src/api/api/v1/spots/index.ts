/* eslint-disable */
import type { DefineMethods } from 'aspida';
import type * as Types from '../../../@types';

export type Methods = DefineMethods<{
  /** Retrieve a list of locations based on longitude, latitude */
  get: {
    query: {
      /** The longitude of the location. */
      longitude: number;
      /** The latitude of the location. */
      latitude: number;
    };

    status: 200;
    /** OK */
    resBody: Types.SuccessLocationRes;
  };
}>;

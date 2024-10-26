/* eslint-disable */
import type { DefineMethods } from 'aspida';

export type Methods = DefineMethods<{
  /** Retrieve a list of locations based on longitude, latitude */
  get: {
    status: 200;
    /** OK */
    resBody: Blob;
  };
}>;

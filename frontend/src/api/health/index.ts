/* eslint-disable */
import type { DefineMethods } from 'aspida';

export type Methods = DefineMethods<{
  /** Returns the health status of the application in JSON format. */
  get: {
    status: 200;

    /** OK - Service is healthy */
    resBody: {
      status?: string | undefined;
      timestamp?: string | undefined;
    };
  };
}>;

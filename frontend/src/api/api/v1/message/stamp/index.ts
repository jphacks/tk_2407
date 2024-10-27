/* eslint-disable */
import type { DefineMethods } from 'aspida';
import type * as Types from '../../../../@types';

export type Methods = DefineMethods<{
  /** Post a stamp */
  post: {
    status: 201;
    /** Created */
    resBody: Types.SuccessStampRes;

    reqBody: {
      messageId: string;
      userId: string;
      stampType: string;
    };
  };
}>;

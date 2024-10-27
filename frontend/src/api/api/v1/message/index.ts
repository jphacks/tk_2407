/* eslint-disable */
import type { DefineMethods } from 'aspida';
import type * as Types from '../../../@types';

export type Methods = DefineMethods<{
  /** Post a message */
  post: {
    status: 201;
    /** Created */
    resBody: Types.SuccessMessageCreateRes;

    reqBody: {
      spotId: string;
      userid: string;
      photoUrl: string;
      content: string;
    };
  };
}>;

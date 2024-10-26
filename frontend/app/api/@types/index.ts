/* eslint-disable */
export type HealthRes = {
  message: string;
}

export type SuccessMessageRes = {
  messages?: {
    /** Message ID */
    id?: number | undefined;
    /** Author of the message */
    author?: string | undefined;
    /** Message content */
    content?: string | undefined;

    /** Message stamps */
    stamps?: {
      [key: string]: {
        /** Stamp count */
        count?: number | undefined;
        /** Check if the user has stamped */
        is_reacted?: boolean | undefined;
      };
    } | undefined;
  }[] | undefined;
}

export type ErrorRes = {
  message: string;
}

export type UnauthorizedError = ErrorRes

export type BadRequestError = ErrorRes

export type NotFoundError = ErrorRes

export type InternalError = ErrorRes

/* eslint-disable */
export type HealthRes = {
  message: string;
}

export type SuccessMessageRes = {
  messages: {
    id: string;
    author_name: string;
    content: string;

    stamps?: {
      [key: string]: {
        count: number;
        is_reacted: boolean;
      };
    } | null | undefined;
  }[];
}

export type ErrorRes = {
  message: string;
}

export type UnauthorizedError = ErrorRes

export type BadRequestError = ErrorRes

export type NotFoundError = ErrorRes

export type InternalError = ErrorRes

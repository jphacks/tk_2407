/* eslint-disable */
export type HealthRes = {
  message: string;
}

export type SuccessRes = {
  message: string;
}

export type ErrorRes = {
  message: string;
}

export type SuccessResponse = SuccessRes

export type UnauthorizedError = ErrorRes

export type BadRequestError = ErrorRes

export type NotFoundError = ErrorRes

export type InternalError = ErrorRes

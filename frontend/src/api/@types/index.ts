/* eslint-disable */
export type HealthRes = {
  message: string;
}

export type SuccessStampRes = {
  message: string;
}

export type SuccessMessageRes = {
  messages: {
    id: string;
    author_name: string;
    content: string;

    stamps?: {
      [key: string]: {
        type: string;
        count: number;
        is_reacted: boolean;
      };
    } | null | undefined;
  }[];
}

export type SuccessLocationRes = {
  spots?: Spot[] | null | undefined;
}

export type SuccessUserRes = {
  userId: string;
  username: string;
  email: string;
}

export type SuccessLoginRes = {
  userId: string;
}

export type SuccessSignupRes = {
  userId: string;
}

export type ErrorRes = {
  message: string;
}

export type SuccessMessageCreateRes = {
  messageId: string;
  spotId: string;
  userId: string;
  content: string;
  photoUrl: string;
}

export type Spot = {
  spot_id: string;
  google_map_id: string;
  google_map_place_id: string;
  name: string;
  latitude: number;
  longitude: number;
  address: string;
  types: string[];
  photo_url: string;
}

export type UnauthorizedError = ErrorRes

export type BadRequestError = ErrorRes

export type NotFoundError = ErrorRes

export type InternalError = ErrorRes

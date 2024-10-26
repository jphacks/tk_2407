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

export type SuccessLocationRes = {
  spots?: {
    spot_id: string;
    google_map_place_id: string;
    name: string;
    description: string;
    photo_url: string;
    latitude: number;
    longitude: number;
  }[] | null | undefined;
}

export type SuccessUserRes = {
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

export type UnauthorizedError = ErrorRes

export type BadRequestError = ErrorRes

export type NotFoundError = ErrorRes

export type InternalError = ErrorRes

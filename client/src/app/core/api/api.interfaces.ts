export type SignInRequest = {
  email: string;
  password: string;
};

export type SignInUser = {
  id: string;
  email: string;
};

export type SignInResponse = RefreshTokenResponse & {
  user: SignInUser;
};

export type SignInConfirmResponse = {
  sent: boolean;
};

export type SignUpRequest = {
  email: string;
  password: string;
  name: string;
};

export type SignUpResponse = {
  id: string;
};

export type SignOutRequest = {
  refreshToken: string;
};

export type SignOutResponse = {
  ok: boolean;
};

export type RefreshTokenRequest = {
  refreshToken: string;
};

export type RefreshTokenResponse = {
  accessToken: string;
  accessTokenExpiresIn: number;
  refreshToken: string;
  refreshTokenExpiresIn: number;
};

export type ConfirmEmailRequest = {
  token: string;
};
export type ConfirmEmailResponse = SignInResponse;

export type ResendConfirmEmailRequest = {
  email: string;
};

export type ResendConfirmEmailResponse = {
  sent: boolean;
};

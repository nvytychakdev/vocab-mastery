import { HttpClient, HttpContext } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Api } from '@core/api/api';
import { Profile } from './auth-profile.interface';
import {
  ConfirmEmailRequest,
  ConfirmEmailResponse,
  RefreshTokenRequest,
  RefreshTokenResponse,
  ResendConfirmEmailRequest,
  ResendConfirmEmailResponse,
  SignInConfirmResponse,
  SignInRequest,
  SignInResponse,
  SignOutRequest,
  SignOutResponse,
  SignUpRequest,
  SignUpResponse,
} from './auth.interface';
import { IS_AUTHORIZED_REQUEST } from './auth.model';

enum AuthEndpoint {
  Profile = 'api/v1/auth/profile',
  SignIn = 'api/v1/auth/sign-in',
  SignOut = 'api/v1/auth/sign-out',
  SignUp = 'api/v1/auth/sign-up',
  RefreshToken = 'api/v1/auth/refresh-token',
  ConfirmEmail = 'api/v1/auth/confirm-email',
  ResendConfirmEmail = 'api/v1/auth/resend-confirm-email',
  OAuthGoogle = 'api/v1/auth/oauth/google',
}

const IsAuthorizedContext = new HttpContext().set(IS_AUTHORIZED_REQUEST, true);

@Injectable({
  providedIn: 'root',
})
export class AuthApi {
  private readonly http = inject(HttpClient);
  private readonly api = inject(Api);

  private getApiUrl(endpoint: AuthEndpoint, params?: Record<string, string>) {
    return this.api.getUrl(endpoint, params);
  }

  getGooglePopupUrl() {
    return this.getApiUrl(AuthEndpoint.OAuthGoogle);
  }

  getAuthProfile() {
    const url = this.getApiUrl(AuthEndpoint.Profile);
    return this.http.get<Profile>(url, { context: IsAuthorizedContext });
  }

  signIn(request: SignInRequest) {
    const url = this.getApiUrl(AuthEndpoint.SignIn);
    return this.http.post<SignInResponse | SignInConfirmResponse>(url, request);
  }

  signUp(request: SignUpRequest) {
    const url = this.getApiUrl(AuthEndpoint.SignUp);
    return this.http.post<SignUpResponse | SignInConfirmResponse>(url, request);
  }

  signOut(request: SignOutRequest) {
    const url = this.getApiUrl(AuthEndpoint.SignOut);
    return this.http.post<SignOutResponse>(url, request, { context: IsAuthorizedContext });
  }

  refreshToken(request: RefreshTokenRequest) {
    const url = this.getApiUrl(AuthEndpoint.RefreshToken);
    return this.http.post<RefreshTokenResponse>(url, request);
  }

  confirmEmail(request: ConfirmEmailRequest) {
    const url = this.getApiUrl(AuthEndpoint.ConfirmEmail);
    return this.http.post<ConfirmEmailResponse>(url, request);
  }

  resendConfirmEmail(request: ResendConfirmEmailRequest) {
    const url = this.getApiUrl(AuthEndpoint.ResendConfirmEmail);
    return this.http.post<ResendConfirmEmailResponse>(url, request);
  }
}

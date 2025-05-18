import { HttpClient, HttpContext } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { IS_AUTHORIZED_REQUEST } from '../auth/auth.interfaces';
import { User } from '../interfaces/user.interface';
import {
  RefreshTokenRequest,
  RefreshTokenResponse,
  SignInRequest,
  SignInResponse,
  SignOutRequest,
  SignOutResponse,
  SignUpRequest,
  SignUpResponse,
} from './api.interfaces';

enum ApiEndpoint {
  Profile = 'api/v1/auth/profile',
  SignIn = 'api/v1/auth/sign-in',
  SignOut = 'api/v1/auth/sign-out',
  SignUp = 'api/v1/auth/sign-up',
  RefreshToken = 'api/v1/auth/refresh-token',
}

const IsAuthorizedContext = new HttpContext().set(IS_AUTHORIZED_REQUEST, true);

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private readonly http = inject(HttpClient);

  private getApiUrl(endpoint: ApiEndpoint, params?: Record<string, string>) {
    return `${environment.hostUrl}/${endpoint}`.trim();
  }

  getAuthProfile() {
    const url = this.getApiUrl(ApiEndpoint.Profile);
    return this.http.get<User>(url, { context: IsAuthorizedContext });
  }

  signIn(request: SignInRequest) {
    const url = this.getApiUrl(ApiEndpoint.SignIn);
    return this.http.post<SignInResponse>(url, request);
  }

  signUp(request: SignUpRequest) {
    const url = this.getApiUrl(ApiEndpoint.SignUp);
    return this.http.post<SignUpResponse>(url, request);
  }

  signOut(request: SignOutRequest) {
    const url = this.getApiUrl(ApiEndpoint.SignOut);
    return this.http.post<SignOutResponse>(url, request, { context: IsAuthorizedContext });
  }

  refreshToken(request: RefreshTokenRequest) {
    const url = this.getApiUrl(ApiEndpoint.RefreshToken);
    return this.http.post<RefreshTokenResponse>(url, request);
  }
}

import { inject, Injectable } from '@angular/core';
import { AuthApi } from '@domain/auth/auth.api';
import { filter, fromEvent, Observable, of, tap } from 'rxjs';
import { environment } from '../../../environments/environment';
import { AuthStorage } from '../../core/storage/auth-storage.service';
import { AuthTokens } from '../../domains/auth/auth-tokens';
import { SignInConfirmResponse, SignOutResponse, SignUpRequest } from '../../domains/auth/auth.interface';

export const isConfirmResponse = <T extends object>(
  response: T | SignInConfirmResponse
): response is SignInConfirmResponse => {
  return 'sent' in response;
};

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly authApi = inject(AuthApi);
  private readonly storage = inject(AuthStorage);
  private readonly tokens = inject(AuthTokens);

  signIn(email: string, password: string) {
    return this.authApi.signIn({ email, password }).pipe(
      tap(res => {
        if (isConfirmResponse(res)) return;
        this.tokens.setTokens(res);
      })
    );
  }

  signInWithGoogle() {
    window.open(this.authApi.getGooglePopupUrl(), '_blank', 'width=500,height=600');

    return fromEvent(window, 'message').pipe(
      filter(event => {
        if (!(event instanceof MessageEvent)) return false;
        return event.origin === environment.hostUrl;
      }),
      tap(event => {
        if (!(event instanceof MessageEvent)) return;
        if (isConfirmResponse(event.data)) return;
        this.tokens.setTokens(event.data);
      })
    );
  }

  refreshToken(refreshToken: string) {
    return this.authApi.refreshToken({ refreshToken }).pipe(tap(res => this.tokens.setTokens(res)));
  }

  signUp(request: SignUpRequest) {
    return this.authApi.signUp(request);
  }

  signOut() {
    const refreshToken = this.storage.getRefreshToken()?.jwtToken;
    const signOut$: Observable<SignOutResponse | null> = refreshToken
      ? this.authApi.signOut({ refreshToken })
      : of(null);
    return signOut$.pipe(tap(() => this.tokens.removeTokens()));
  }

  getAuthProfile() {
    return this.authApi.getAuthProfile();
  }

  isAuthorized(): boolean {
    const accessToken = this.storage.getAccessToken();

    if (!accessToken) return false;
    if (accessToken.isExpired()) return false;

    return true;
  }

  isAuthenticated() {
    const refreshToken = this.storage.getRefreshToken();

    if (!refreshToken) return false;
    if (refreshToken.isExpired()) return false;

    return true;
  }

  confirmEmail(token: string) {
    return this.authApi.confirmEmail({ token }).pipe(tap(res => this.tokens.setTokens(res)));
  }

  resendConfirmEmail(email: string) {
    return this.authApi.resendConfirmEmail({ email });
  }
}

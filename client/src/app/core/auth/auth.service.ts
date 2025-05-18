import { inject, Injectable } from '@angular/core';
import { Observable, of, tap } from 'rxjs';
import { SignOutResponse, SignUpRequest } from '../api/api.interfaces';
import { ApiService } from '../api/api.service';
import { AuthStorageService } from './auth-storage.service';
import { AuthTokensService } from './auth-tokens.service';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly api = inject(ApiService);
  private readonly storage = inject(AuthStorageService);
  private readonly tokens = inject(AuthTokensService);

  signIn(email: string, password: string) {
    return this.api.signIn({ email, password }).pipe(tap(res => this.tokens.setTokens(res)));
  }

  refreshToken(refreshToken: string) {
    return this.api.refreshToken({ refreshToken }).pipe(tap(res => this.tokens.setTokens(res)));
  }

  signUp(request: SignUpRequest) {
    return this.api.signUp(request);
  }

  signOut() {
    const refreshToken = this.storage.getRefreshToken()?.jwtToken;
    const signOut$: Observable<SignOutResponse | null> = refreshToken ? this.api.signOut({ refreshToken }) : of(null);
    return signOut$.pipe(tap(() => this.tokens.removeTokens()));
  }

  getAuthProfile() {
    return this.api.getAuthProfile();
  }

  isAuthenticated(): boolean {
    const accessToken = this.storage.getAccessToken();
    const refreshToken = this.storage.getRefreshToken();

    if (!accessToken || !refreshToken) return false;
    if (accessToken.isExpired() || refreshToken.isExpired()) return false;

    return true;
  }
}

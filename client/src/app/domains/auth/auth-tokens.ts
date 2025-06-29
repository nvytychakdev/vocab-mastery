import { inject, Injectable } from '@angular/core';
import { Token } from '../../core/models/token.model';
import { AuthStorage } from '../../core/storage/auth-storage.service';
import { RefreshTokenResponse } from './auth.interface';

@Injectable({
  providedIn: 'root',
})
export class AuthTokens {
  private readonly storage = inject(AuthStorage);

  getTokens() {
    return {
      accessToken: this.storage.getAccessToken(),
      refreshToken: this.storage.getRefreshToken(),
    };
  }

  setTokens(tokens: RefreshTokenResponse) {
    const accessTokenExpiresAt = new Date(Date.now() + tokens.accessTokenExpiresIn * 1000).toISOString();
    const accessToken = new Token(tokens.accessToken, accessTokenExpiresAt);
    this.storage.setAccessToken(accessToken);

    const refreshTokenExpiresAt = new Date(Date.now() + tokens.refreshTokenExpiresIn * 1000).toISOString();
    const refreshToken = new Token(tokens.refreshToken, refreshTokenExpiresAt);
    this.storage.setRefreshToken(refreshToken);
  }

  removeTokens() {
    this.storage.removeAccessToken();
    this.storage.removeRefreshToken();
  }
}

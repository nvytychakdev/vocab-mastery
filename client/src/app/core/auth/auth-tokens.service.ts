import { inject, Injectable } from '@angular/core';
import { RefreshTokenResponse } from '../api/api.interfaces';
import { Token } from '../models/token.model';
import { AuthStorageService } from './auth-storage.service';

@Injectable({
  providedIn: 'root',
})
export class AuthTokensService {
  private readonly storage = inject(AuthStorageService);

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

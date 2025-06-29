import { Injectable } from '@angular/core';
import { Token } from '../models/token.model';

const storagePrefix = 'vm';
enum StorageKey {
  RefreshToken = 'refreshToken',
  RefreshTokenExpiresAt = 'refreshTokenExpiresIn',
  AccessToken = 'accessToken',
  AccessTokenExpiresAt = 'accessTokenExpiresIn',
}

@Injectable({
  providedIn: 'root',
})
export class AuthStorage {
  private getStorageKey(key: StorageKey) {
    return `${storagePrefix}_${key}`;
  }

  setRefreshToken(refreshToken: Token) {
    localStorage.setItem(this.getStorageKey(StorageKey.RefreshToken), refreshToken.jwtToken);
    localStorage.setItem(this.getStorageKey(StorageKey.RefreshTokenExpiresAt), refreshToken.expiresAt);
  }

  getRefreshToken(): Token | null {
    const refreshToken = localStorage.getItem(this.getStorageKey(StorageKey.RefreshToken));
    const refreshTokenExpiresAt = localStorage.getItem(this.getStorageKey(StorageKey.RefreshTokenExpiresAt));
    if (!refreshToken || !refreshTokenExpiresAt) return null;
    return new Token(refreshToken, refreshTokenExpiresAt);
  }

  removeRefreshToken() {
    localStorage.removeItem(this.getStorageKey(StorageKey.RefreshToken));
    localStorage.removeItem(this.getStorageKey(StorageKey.RefreshTokenExpiresAt));
  }

  setAccessToken(accessToken: Token) {
    localStorage.setItem(this.getStorageKey(StorageKey.AccessToken), accessToken.jwtToken);
    localStorage.setItem(this.getStorageKey(StorageKey.AccessTokenExpiresAt), accessToken.expiresAt);
  }

  getAccessToken() {
    const accesstToken = localStorage.getItem(this.getStorageKey(StorageKey.AccessToken));
    const accessTokenExpiresAt = localStorage.getItem(this.getStorageKey(StorageKey.AccessTokenExpiresAt));
    if (!accesstToken || !accessTokenExpiresAt) return null;
    return new Token(accesstToken, accessTokenExpiresAt);
  }

  removeAccessToken() {
    localStorage.removeItem(this.getStorageKey(StorageKey.AccessToken));
    localStorage.removeItem(this.getStorageKey(StorageKey.AccessTokenExpiresAt));
  }
}

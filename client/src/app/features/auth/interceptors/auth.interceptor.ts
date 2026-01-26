import type { HttpInterceptorFn, HttpRequest } from '@angular/common/http';
import { inject } from '@angular/core';
import { IS_AUTHORIZED_REQUEST } from '@core/models/authorized.model';
import { AuthTokens } from '@domain/auth/auth-tokens';
import { RefreshTokenResponse } from '@domain/auth/auth.interface';
import { finalize, Observable, switchMap } from 'rxjs';
import { AuthService } from '../auth.service';

let refreshToken$: Observable<RefreshTokenResponse> | null = null;

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  if (!req.context.get(IS_AUTHORIZED_REQUEST)) {
    return next(req);
  }

  const auth = inject(AuthService);
  const tokens = inject(AuthTokens);

  const { accessToken, refreshToken } = tokens.getTokens();

  if (!accessToken || !refreshToken) throw new Error('No access or refresh tokens');
  if (refreshToken.isExpired()) throw new Error('Refresh token expired');

  if (accessToken.isExpired()) {
    if (!refreshToken$) {
      refreshToken$ = auth.refreshToken(refreshToken.jwtToken);
    }

    return refreshToken$.pipe(
      switchMap(res => next(authReq(req, res.accessToken))),
      finalize(() => (refreshToken$ = null))
    );
  }

  return next(authReq(req, accessToken.jwtToken));
};

const authReq = (req: HttpRequest<unknown>, token: string) => {
  return req.clone({
    headers: req.headers.append('Authorization', `Bearer ${token}`),
  });
};

import type { HttpInterceptorFn, HttpRequest } from '@angular/common/http';
import { inject } from '@angular/core';
import { AuthTokens } from '@domain/auth/auth-tokens';
import { IS_AUTHORIZED_REQUEST } from '@domain/auth/auth.model';
import { switchMap } from 'rxjs';
import { AuthService } from '../auth.service';

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
    return auth.refreshToken(refreshToken.jwtToken).pipe(
      switchMap(res => {
        return next(authReq(req, res.accessToken));
      })
    );
  }

  return next(authReq(req, accessToken.jwtToken));
};

const authReq = (req: HttpRequest<unknown>, token: string) => {
  return req.clone({
    headers: req.headers.append('Authorization', `Bearer ${token}`),
  });
};

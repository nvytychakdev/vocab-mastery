import type { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { switchMap } from 'rxjs';
import { AuthTokensService } from '../auth-tokens.service';
import { IS_AUTHORIZED_REQUEST } from '../auth.interfaces';
import { AuthService } from '../auth.service';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  if (!req.context.get(IS_AUTHORIZED_REQUEST)) {
    return next(req);
  }

  const auth = inject(AuthService);
  const tokens = inject(AuthTokensService);

  const { accessToken, refreshToken } = tokens.getTokens();

  if (!accessToken || !refreshToken) throw new Error('No access or refresh tokens');
  if (refreshToken.isExpired()) throw new Error('Refresh token expired');

  if (accessToken.isExpired()) {
    return auth.refreshToken(refreshToken.jwtToken).pipe(switchMap(() => next(req)));
  }

  return next(req);
};

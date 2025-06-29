import { inject } from '@angular/core';
import { Router, type CanActivateFn } from '@angular/router';
import { AuthTokens } from '../../../domains/auth/auth-tokens';
import { AuthService } from '../auth.service';

export const authGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  return auth.isAuthenticated();
};

export const authRedirectGuards = (options?: { redirectAuth: string; redirectUnauth: string }) => {
  const redirectIfAuthenticated = () => {
    const router = inject(Router);
    const auth = inject(AuthService);
    if (!options?.redirectAuth) return true;
    return auth.isAuthorized() ? router.createUrlTree([options.redirectAuth]) : true;
  };

  const redirectIfUnauthenticated = () => {
    const router = inject(Router);
    const auth = inject(AuthService);
    const tokens = inject(AuthTokens);
    if (!options?.redirectUnauth) return true;
    if (!auth.isAuthorized()) {
      const { refreshToken } = tokens.getTokens();
      if (!refreshToken || !auth.isAuthenticated()) {
        return router.createUrlTree([options.redirectUnauth]);
      }
      return auth.refreshToken(refreshToken.jwtToken);
    }

    return true;
  };

  return {
    redirectIfAuthenticated,
    redirectIfUnauthenticated,
  };
};

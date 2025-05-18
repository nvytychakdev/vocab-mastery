import { inject } from '@angular/core';
import { Router, type CanActivateFn } from '@angular/router';
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
    return auth.isAuthenticated() ? router.createUrlTree([options.redirectAuth]) : true;
  };

  const redirectIfUnauthenticated = () => {
    const router = inject(Router);
    const auth = inject(AuthService);
    if (!options?.redirectUnauth) return true;
    return auth.isAuthenticated() ? true : router.createUrlTree([options.redirectUnauth]);
  };

  return {
    redirectIfAuthenticated,
    redirectIfUnauthenticated,
  };
};

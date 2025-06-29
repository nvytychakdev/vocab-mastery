import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { Observable, tap } from 'rxjs';
import { AuthProfile } from '../../domains/auth/auth-profile';
import { User } from '../user/user.interface';
import { AuthService } from './auth.service';

export const authProfileResolve: ResolveFn<Observable<User>> = (
  route: ActivatedRouteSnapshot,
  state: RouterStateSnapshot
) => {
  const auth = inject(AuthService);
  const authProfile = inject(AuthProfile);
  return auth.getAuthProfile().pipe(tap(profile => authProfile.setProfile(profile)));
};

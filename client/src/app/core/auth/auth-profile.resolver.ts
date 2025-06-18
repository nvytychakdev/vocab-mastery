import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { Observable, tap } from 'rxjs';
import { User } from '../interfaces/user.interface';
import { AuthProfileService } from './auth-profile.service';
import { AuthService } from './auth.service';

export const authProfileResolve: ResolveFn<Observable<User>> = (
  route: ActivatedRouteSnapshot,
  state: RouterStateSnapshot
) => {
  const auth = inject(AuthService);
  const authProfile = inject(AuthProfileService);
  return auth.getAuthProfile().pipe(tap(profile => authProfile.setProfile(profile)));
};

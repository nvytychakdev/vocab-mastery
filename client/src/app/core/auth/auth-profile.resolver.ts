import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';
import { User } from '../interfaces/user.interface';
import { AuthService } from './auth.service';

export const authProfileResolve: ResolveFn<Observable<User>> = (route: ActivatedRouteSnapshot, state: RouterStateSnapshot) => {
  const auth = inject(AuthService);
  return auth.getAuthProfile();
};

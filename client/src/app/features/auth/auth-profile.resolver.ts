import { inject } from '@angular/core';
import { ResolveFn } from '@angular/router';
import { AuthProfileState } from '@domain/auth/auth-profile-state';
import { Profile } from '@domain/auth/auth-profile.interface';
import { Observable, tap } from 'rxjs';
import { AuthService } from './auth.service';

export const authProfileResolve: ResolveFn<Observable<Profile>> = () => {
  const auth = inject(AuthService);
  const authProfile = inject(AuthProfileState);
  return auth.getAuthProfile().pipe(tap(profile => authProfile.setProfile(profile)));
};

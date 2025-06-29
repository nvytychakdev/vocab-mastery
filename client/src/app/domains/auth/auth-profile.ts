import { Injectable, signal } from '@angular/core';
import { Profile } from './auth-profile.interface';

@Injectable({
  providedIn: 'root',
})
export class AuthProfile {
  private readonly _profile = signal<Profile | undefined>(undefined);
  public profile = this._profile.asReadonly();

  setProfile(profile: Profile) {
    this._profile.set(profile);
  }
}

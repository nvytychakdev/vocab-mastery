import { Injectable, signal } from '@angular/core';
import { User } from '../interfaces/user.interface';

@Injectable({
  providedIn: 'root',
})
export class AuthProfileService {
  private readonly _profile = signal<User | undefined>(undefined);
  public profile = this._profile.asReadonly();

  setProfile(profile: User) {
    this._profile.set(profile);
  }
}

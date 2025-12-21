import { Injectable, signal } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';

@Injectable()
export class ToggleButtonState {
  private readonly _activeToggle = signal<string | number | null>(null);
  readonly activeToggle = this._activeToggle.asReadonly();
  readonly activeToggleChanges$ = toObservable(this.activeToggle);

  setActiveToggle(value: string | number) {
    this._activeToggle.set(value === this.activeToggle() ? null : value);
  }
}

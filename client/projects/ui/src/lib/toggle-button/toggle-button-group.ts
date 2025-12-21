import { ChangeDetectionStrategy, Component, DestroyRef, inject, model, OnInit } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { ToggleButtonState } from './toggle-button-state';

@Component({
  selector: 'vm-toggle-button-group',
  imports: [],
  template: `<div class="vm-toggle-button-group">
    <ng-content select="vm-toggle-button" />
  </div>`,
  providers: [ToggleButtonState],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ToggleButtonGroup implements OnInit {
  private readonly destroyRef = inject(DestroyRef);
  readonly toggleState = inject(ToggleButtonState);

  readonly value = model<string | number | undefined>(undefined);

  ngOnInit() {
    const value = this.value();
    if (value) this.toggleState.setActiveToggle(value);

    this.toggleState.activeToggleChanges$.pipe(takeUntilDestroyed(this.destroyRef)).subscribe(value => {
      this.value.set(value || undefined);
    });
  }
}

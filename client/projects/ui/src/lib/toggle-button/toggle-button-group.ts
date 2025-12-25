import { ChangeDetectionStrategy, Component, InjectionToken, model } from '@angular/core';
import { FormValueControl } from '@angular/forms/signals';

export const TOGGLE_BUTTON_GROUP = new InjectionToken<ToggleButtonGroup>('TOGGLE_BUTTON_GROUP');

@Component({
  selector: 'vm-toggle-button-group',
  imports: [],
  template: `<div class="vm-toggle-button-group">
    <ng-content select="vm-toggle-button" />
  </div>`,
  providers: [{ provide: TOGGLE_BUTTON_GROUP, useExisting: ToggleButtonGroup }],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ToggleButtonGroup implements FormValueControl<string> {
  readonly value = model<string>('');
}

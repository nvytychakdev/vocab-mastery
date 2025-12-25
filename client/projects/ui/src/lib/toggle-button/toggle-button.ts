import { ChangeDetectionStrategy, Component, computed, inject, input } from '@angular/core';
import { TOGGLE_BUTTON_GROUP } from './toggle-button-group';

@Component({
  selector: 'vm-toggle-button',
  imports: [],
  template: `
    <button
      tabindex="0"
      class="vm-toggle-button"
      [class.vm-toggle-button-active]="isSelected()"
      (click)="onButtonSelect()"
      (keypress)="onButtonSelect()">
      <ng-content />
    </button>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ToggleButton {
  private readonly toggleState = inject(TOGGLE_BUTTON_GROUP, { optional: true });

  readonly value = input.required<string>();
  readonly isSelected = computed(() => this.toggleState?.value() === this.value());

  onButtonSelect() {
    const value = this.value();
    const currentValue = this.toggleState?.value();

    if (currentValue && currentValue === value) {
      this.toggleState?.value.set('');
      return;
    }

    if (value) {
      this.toggleState?.value.set(value);
      return;
    }
  }
}

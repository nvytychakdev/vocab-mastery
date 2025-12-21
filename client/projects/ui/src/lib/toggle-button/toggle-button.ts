import { ChangeDetectionStrategy, Component, computed, inject, input } from '@angular/core';
import { ToggleButtonState } from './toggle-button-state';

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
  private readonly toggleState = inject(ToggleButtonState);
  readonly value = input.required<string | number>();
  readonly isSelected = computed(() => this.toggleState.activeToggle() === this.value());

  onButtonSelect() {
    this.toggleState.setActiveToggle(this.value());
  }
}

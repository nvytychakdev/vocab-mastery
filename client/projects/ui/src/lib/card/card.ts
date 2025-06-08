import { ChangeDetectionStrategy, Component } from '@angular/core';

@Component({
  selector: 'vm-card',
  imports: [],
  template: `
    <div class="vm-card-title">
      <ng-content select="[title]"></ng-content>
    </div>

    <div class="vm-card-value">
      <ng-content select="[value]"></ng-content>
    </div>

    <div class="vm-class-action">
      <ng-content select="[action]"></ng-content>
    </div>
  `,
  host: {
    class: 'vm-card',
  },
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Card {}

import { ChangeDetectionStrategy, Component, HostBinding, input } from '@angular/core';

type CardType = 'stat' | 'info';

const CardTypes: Record<CardType, string> = {
  info: 'vm-card-info',
  stat: 'vm-card-stat',
};

@Component({
  selector: 'vm-card',
  imports: [],
  template: `
    <div class="vm-card-content">
      <div class="vm-card-title">
        <ng-content select="[title]"></ng-content>
      </div>

      <div class="vm-card-value">
        <ng-content select="[value]"></ng-content>
      </div>
    </div>

    <div class="vm-card-action">
      <ng-content select="[action]"></ng-content>
    </div>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Card {
  readonly type = input<CardType>('stat');

  @HostBinding('class')
  get className() {
    return `vm-card ${CardTypes[this.type()]}`;
  }
}

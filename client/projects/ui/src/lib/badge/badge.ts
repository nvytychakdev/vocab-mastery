import { ChangeDetectionStrategy, Component, HostBinding, input } from '@angular/core';

export type BadgeColor = 'gray' | 'green' | 'blue' | 'yellow' | 'purple' | 'lime' | 'sky' | 'pink' | 'orange' | 'red';

const BadgeColors: Record<BadgeColor, string> = {
  green: 'vm-badge-green',
  lime: 'vm-badge-lime',
  blue: 'vm-badge-blue',
  sky: 'vm-badge-sky',
  purple: 'vm-badge-purple',
  pink: 'vm-badge-pink',
  yellow: 'vm-badge-yellow',
  orange: 'vm-badge-orange',
  red: 'vm-badge-red',
  gray: 'vm-badge-gray',
};

@Component({
  selector: 'vm-badge',
  imports: [],
  template: '<ng-content />',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Badge {
  readonly color = input<BadgeColor>('gray');

  @HostBinding('class') get className() {
    return `vm-badge ${this.classColor}`;
  }

  get classColor() {
    return BadgeColors[this.color()];
  }
}

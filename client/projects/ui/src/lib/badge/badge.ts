import { ChangeDetectionStrategy, Component, HostBinding, input } from '@angular/core';

type BadgeColor = 'gray' | 'green' | 'blue' | 'yellow' | 'purple';

const BadgeColors: Record<BadgeColor, string> = {
  green: 'vm-badge-green',
  blue: 'vm-badge-blue',
  purple: 'vm-badge-purple',
  yellow: 'vm-badge-yellow',
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

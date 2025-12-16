import { ChangeDetectionStrategy, Component, HostBinding, input } from '@angular/core';

type BadgeColor = 'gray' | 'green' | 'blue' | 'yellow' | 'purple';

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
    switch (this.color()) {
      case 'green':
        return 'bg-badge-green/50 text-badge-green-subtle';
      case 'blue':
        return 'bg-badge-blue/50 text-badge-blue-subtle';
      case 'purple':
        return 'bg-badge-purple/50 text-badge-purple-subtle';
      case 'yellow':
        return `bg-badge-yellow/50 text-badge-yellow-subtle`;

      default:
        return 'bg-badge-zinc/50 text-badge-zinc-subtle';
    }
  }
}

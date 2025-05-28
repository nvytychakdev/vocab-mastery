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
        return 'bg-green-900/50 text-green-500';
      case 'blue':
        return 'bg-blue-900/50 text-blue-400';
      case 'purple':
        return 'bg-purple-900/50 text-purple-500';
      case 'yellow':
        return `bg-yellow-900/50 text-yellow-500`;

      default:
        return 'bg-gray-900/50 text-gray-300';
    }
  }
}

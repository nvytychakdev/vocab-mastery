import { ChangeDetectionStrategy, Component } from '@angular/core';

@Component({
  selector: 'vm-divider',
  imports: [],
  template: '<div class="bg-surface-muted h-[1px]"></div>',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Divider {}

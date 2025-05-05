import { Directive, HostBinding, input } from '@angular/core';

@Directive({
  selector: '[vmButton]',
})
export class ButtonDirective {
  readonly appearance = input<'button' | 'link' | 'icon'>('button');

  @HostBinding('class') get className() {
    switch (this.appearance()) {
      case 'link':
        return `vm-button-link`;
      case 'icon':
        return `vm-button vm-button-icon`;
      default:
        return `vm-button`;
    }
  }
}

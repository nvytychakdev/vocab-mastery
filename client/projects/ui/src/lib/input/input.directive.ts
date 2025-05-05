import { Directive, HostBinding } from '@angular/core';

@Directive({
  selector: '[vmInput]',
})
export class InputDirective {
  @HostBinding('class') className = 'vm-input';
}

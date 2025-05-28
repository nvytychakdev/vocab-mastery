import { Directive, HostBinding } from '@angular/core';

@Directive({
  selector: '[vmInput]',
})
export class Input {
  @HostBinding('class') className = 'vm-input';
}

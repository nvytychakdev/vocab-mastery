import { Directive, HostListener, inject } from '@angular/core';
import { MENU } from './menu';

@Directive({
  selector: '[vmMenuItem]',
  host: {
    class: 'vm-menu-item',
  },
})
export class MenuItem {
  private readonly menu = inject(MENU);

  @HostListener('click')
  onClick() {
    this.menu.closeMenu();
  }
}

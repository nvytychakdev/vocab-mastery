import { Component, InjectionToken, input, signal, TemplateRef, viewChild } from '@angular/core';
import { Subject } from 'rxjs';

export const MENU = new InjectionToken<Menu>('VM_MENU');

@Component({
  selector: 'vm-menu',
  template: `
    <ng-template>
      <div class="vm-menu vm-menu-left-top" [class]="className()">
        <div class="vm-menu-content">
          <ng-content></ng-content>
        </div>
      </div>
    </ng-template>
  `,
  host: {
    class: 'contents',
  },
  exportAs: 'vmMenu',
  providers: [{ provide: MENU, useExisting: Menu }],
})
export class Menu {
  readonly templateRef = viewChild.required(TemplateRef);

  readonly className = input<string>();

  private readonly _close = new Subject<void>();
  private readonly _open = new Subject<void>();
  readonly close = this._close.asObservable();
  readonly open = this._open.asObservable();

  readonly isOpen = signal(false);

  openMenu() {
    this._open.next();
  }

  closeMenu() {
    this._close.next();
  }
}

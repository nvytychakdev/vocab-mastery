import { ChangeDetectionStrategy, Component, computed, HostBinding, input, signal } from '@angular/core';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { lucideX } from '@ng-icons/lucide';
import { Observable, Subject } from 'rxjs';

export interface BaseToast {
  onRemove$: Observable<void>;
  onClose$: Observable<void>;
  isRemoving: boolean;
  show(): void;
  hide(): void;
}

export type ToastType = 'error' | 'success' | 'warn' | 'default';

@Component({
  selector: 'vm-toast-card',
  imports: [NgIcon],
  providers: [provideIcons({ lucideX })],
  template: `
    <div class="flex gap-2 overflow-hidden">
      <div class="vm-toast-content">
        <div class="vm-toast-content-bar" [class]="barClass()"></div>

        <div class="vm-toast-content-data">
          <div class="text-sm font-semibold">{{ title() }}</div>
          @if (description(); as description) {
            <div class="text-xs text-gray-400">{{ description }}</div>
          }
        </div>

        <div class="vm-toast-content-close" (click)="close()">
          <ng-icon name="lucideX" />
        </div>
      </div>
    </div>
  `,

  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ToastCard implements BaseToast {
  readonly title = input.required<string>();
  readonly description = input<string>();
  readonly type = input<ToastType>('default');

  @HostBinding('class') get className() {
    return this.vmClass();
  }

  readonly vmClass = computed(() => `vm-toast ${this.actionClass()}`);
  readonly actionClass = signal<string>('');
  readonly barClass = computed(() => {
    switch (this.type()) {
      case 'success':
        return `bg-green-600`;
      case 'error':
        return `bg-red-800`;
      case 'warn':
        return `bg-amber-600`;

      default:
        return `transparent`;
    }
  });

  isRemoving = false;
  private readonly _onRemove$ = new Subject<void>();
  private readonly _onClose$ = new Subject<void>();
  readonly onRemove$ = this._onRemove$.asObservable();
  readonly onClose$ = this._onClose$.asObservable();

  show() {
    this.actionClass.set('vm-toast-enter');
    setTimeout(() => this.actionClass.set(''), 500);
  }

  hide() {
    this.isRemoving = true;
    this.actionClass.set('vm-toast-leave vm-toast-hidden');
    setTimeout(() => this._onRemove$.next(), 490);
  }

  close() {
    if (this.isRemoving) return;
    this._onClose$.next();
  }
}

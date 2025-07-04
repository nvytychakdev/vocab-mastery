import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { lucidePlus } from '@ng-icons/lucide';
import { Button, Card, Menu, MenuItem, MenuTrigger, ToastService } from '@vm/ui';

@Component({
  selector: 'app-home',
  imports: [Button, NgIcon, MenuTrigger, Menu, MenuItem, Card],
  providers: [provideIcons({ lucidePlus })],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HomeComponent {
  private readonly toast = inject(ToastService);

  openDefault() {
    this.toast.info('Title', 'Test description');
  }

  openShort() {
    this.toast.info('Title');
  }

  openSuccess() {
    this.toast.success('Title', 'Test description');
  }

  openWarn() {
    this.toast.warn('Title', 'Test description');
  }

  openError() {
    this.toast.error('Title', 'Test description');
  }
}

import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { ButtonDirective, ToastService } from '@vm/ui';

@Component({
  selector: 'app-home',
  imports: [ButtonDirective],
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
